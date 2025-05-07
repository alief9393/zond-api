package api

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"

	"zond-api/internal/api/dto"
	"zond-api/internal/api/handler"
	aaRepo "zond-api/internal/api/repository/account_abstraction"
	addrRepo "zond-api/internal/api/repository/address"
	beaconDepositRepo "zond-api/internal/api/repository/beacon_deposit"
	beaconWithdrawalRepo "zond-api/internal/api/repository/beacon_withdrawal"
	blobRepo "zond-api/internal/api/repository/blob"
	blkRepo "zond-api/internal/api/repository/block"
	chainRepo "zond-api/internal/api/repository/chain"
	forkRepo "zond-api/internal/api/repository/fork"
	reorgRepo "zond-api/internal/api/repository/reorg"
	searchRepo "zond-api/internal/api/repository/search"
	txRepo "zond-api/internal/api/repository/transaction"
	valRepo "zond-api/internal/api/repository/validator"

	"zond-api/internal/api/service"

	_ "zond-api/docs"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/time/rate"
)

type rateLimiter struct {
	limiters map[string]*rate.Limiter
	mu       sync.Mutex
	db       *pgxpool.Pool
}

func newRateLimiter() *rateLimiter {
	return &rateLimiter{
		limiters: make(map[string]*rate.Limiter),
	}
}

func (rl *rateLimiter) getLimiter(username string, isPaid bool) *rate.Limiter {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	if limiter, exists := rl.limiters[username]; exists {
		return limiter
	}

	var requestsPerMinute int
	err := rl.db.QueryRow(context.Background(), `
		SELECT requests_per_minute 
		FROM rate_limits WHERE is_paid = $1`, isPaid).
		Scan(&requestsPerMinute)
	if err != nil {
		// Fallback to default
		if isPaid {
			requestsPerMinute = 100
		} else {
			requestsPerMinute = 10
		}
	}

	limit := rate.Every(time.Minute / time.Duration(requestsPerMinute))
	limiter := rate.NewLimiter(limit, 1)
	rl.limiters[username] = limiter
	return limiter
}

var globalRateLimiter = newRateLimiter()

func SetupRouter(db *pgxpool.Pool, jwtSecret string) *gin.Engine {
	blockRepo := blkRepo.NewBlockRepoPG(db)
	txRepository := txRepo.NewTransactionRepoPG(db)
	addrRepository := addrRepo.NewAddressRepoPG(db)
	forkRepository := forkRepo.NewForkRepoPG(db)
	chainRepository := chainRepo.NewChainRepoPG(db)
	reorgRepository := reorgRepo.NewReorgRepoPG(db)
	validatorRepository := valRepo.NewValidatorRepoPG(db)
	searchRepository := searchRepo.NewSearchRepoPG(db)
	beaconDepositRepo := beaconDepositRepo.NewBeaconDepositRepoPG(db)
	searchRepo := searchRepo.NewSearchRepoPG(db)
	bcWithdrawalRepo := beaconWithdrawalRepo.NewBeaconWithdrawalRepoPG(db)
	blobRepo := blobRepo.NewBlobRepoPG(db)
	aaRepo := aaRepo.NewAccountAbstractionRepoPG(db)

	blockService := service.NewBlockService(blockRepo)
	txService := service.NewTransactionService(txRepository)
	addrService := service.NewAddressService(addrRepository)
	forkService := service.NewForkService(forkRepository)
	chainService := service.NewChainService(chainRepository)
	reorgService := service.NewReorgService(reorgRepository)
	validatorService := service.NewValidatorService(validatorRepository)
	beaconDepositService := service.NewBeaconDepositService(beaconDepositRepo)
	searchService := service.NewSearchService(searchRepo)
	bcWithdrawalSvc := service.NewBeaconWithdrawalService(bcWithdrawalRepo)
	blobSvc := service.NewBlobService(blobRepo)
	aaSvc := service.NewAccountAbstractionService(aaRepo)
	_ = searchRepository

	blockHandler := handler.NewBlockHandler(blockService)
	txHandler := handler.NewTransactionHandler(txService)
	addrHandler := handler.NewAddressHandler(addrService)
	forkHandler := handler.NewForkHandler(forkService)
	chainHandler := handler.NewChainHandler(chainService)
	reorgHandler := handler.NewReorgHandler(reorgService)
	validatorHandler := handler.NewValidatorHandler(validatorService)
	searchHandler := handler.NewSearchHandler(searchService)
	beaconDepositHandler := handler.NewBeaconDepositHandler(beaconDepositService)
	beaconWithdrawalHandler := handler.NewBeaconWithdrawalHandler(bcWithdrawalSvc)
	blobHandler := handler.NewBlobHandler(blobSvc)
	aaHandler := handler.NewAccountAbstractionHandler(aaSvc)

	r := gin.Default()
	r.POST("/login", loginHandler(db))
	r.POST("/register", registerHandler(db))

	api := r.Group("/api")
	api.Use(jwtMiddleware(jwtSecret))

	searchGroup := api.Group("/search")
	{
		searchGroup.GET("/suggestions", searchHandler.GetSuggestions)
	}

	blocks := api.Group("/blocks")
	{
		blocks.GET("/latest", blockHandler.GetLatestBlocks)
		blocks.GET("/:block_number", blockHandler.GetBlockByNumber)
		blocks.GET("/hash/:hash", blockHandler.GetBlockByHash)
	}

	transactions := api.Group("/transactions")
	{
		transactions.GET("/latest", txHandler.GetLatestTransactionsWithFilter)
		transactions.GET("/:tx_hash", txHandler.GetTransactionByHash)
		transactions.GET("/block/:block_number", txHandler.GetTransactionsByBlockNumber)
		transactions.GET("/metrics", txHandler.GetTransactionMetrics)
		transactions.GET("/contract", txHandler.GetContractTransactions)
		transactions.GET("/stats/daily", txHandler.GetDailyTransactionStats)
		transactions.GET("/stats/tps", txHandler.GetTPSStats)
		transactions.GET("/stats/fee/daily", txHandler.GetDailyFeeStats)
	}

	addresses := api.Group("/addresses")
	{
		addresses.GET("/:address/balance", addrHandler.GetAddressBalance)
		addresses.GET("/:address/transactions", addrHandler.GetAddressTransactions)
		addresses.GET("/:address/details", addrHandler.GetAddressDetails)
	}

	forks := api.Group("/forks")
	{
		forks.GET("", forkHandler.GetForks)
	}

	chain := api.Group("/chain")
	{
		chain.GET("", chainHandler.GetChainInfo)
	}

	reorgs := api.Group("/reorgs")
	{
		reorgs.GET("", reorgHandler.GetReorgs)
	}

	validators := api.Group("/validators")
	{
		validators.GET("", validatorHandler.GetValidators)
		validators.GET("/:validatorId", validatorHandler.GetValidatorByID)
	}

	beacon := api.Group("/beacon-deposits")
	{
		beacon.GET("/", beaconDepositHandler.GetBeaconDeposits)
	}

	beaconWithdrawals := api.Group("/beacon-withdrawals")
	{
		beaconWithdrawals.GET("/", beaconWithdrawalHandler.GetBeaconWithdrawals)
	}

	blob := api.Group("/blob")
	{
		blob.GET("/", blobHandler.GetBlobs)
	}

	aaGroup := r.Group("/api/account-abstraction")
	{
		aaGroup.GET("", aaHandler.GetAccountAbstraction)
		aaGroup.GET("/bundles", aaHandler.GetOnlyBundleTransactions)
		aaGroup.GET("/aa", aaHandler.GetOnlyAATransactions)
	}

	api.GET("/premium", premiumHandler)

	return r
}

func loginHandler(db *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req dto.LoginRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		var username, password string
		var isPaid bool
		err := db.QueryRow(context.Background(), `
			SELECT username, password, is_paid 
			FROM users WHERE username = $1`, req.Username).
			Scan(&username, &password, &isPaid)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			return
		}

		if err := bcrypt.CompareHashAndPassword([]byte(password), []byte(req.Password)); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			return
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"username": username,
			"is_paid":  isPaid,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})
		tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
			return
		}

		c.JSON(http.StatusOK, dto.LoginResponse{Token: tokenString, IsPaid: isPaid})
	}
}

func registerHandler(db *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req dto.LoginRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
			return
		}

		_, err = db.Exec(context.Background(), `
			INSERT INTO users (username, password, is_paid) 
			VALUES ($1, $2, $3)`, req.Username, string(hashedPassword), req.IsPaid)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "User created"})
	}
}

func jwtMiddleware(jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		const bearerPrefix = "Bearer "
		if len(authHeader) < len(bearerPrefix) || authHeader[:len(bearerPrefix)] != bearerPrefix {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			c.Abort()
			return
		}

		tokenString := authHeader[len(bearerPrefix):]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(jwtSecret), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			username, _ := claims["username"].(string)
			isPaid, _ := claims["is_paid"].(bool)
			c.Set("username", username)
			c.Set("is_paid", isPaid)

			// Apply rate limiting
			limiter := globalRateLimiter.getLimiter(username, isPaid)
			if !limiter.Allow() {
				c.JSON(http.StatusTooManyRequests, gin.H{"error": "Rate limit exceeded"})
				c.Abort()
				return
			}
		}

		c.Next()
	}
}

func premiumHandler(c *gin.Context) {
	isPaid, exists := c.Get("is_paid")
	if !exists || !isPaid.(bool) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Paid account required"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Premium data for paid users"})
}
