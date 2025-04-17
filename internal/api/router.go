package api

import (
	"fmt"
	"net/http"
	"time"

	"zond-api/internal/api/handler"
	"zond-api/internal/api/repository"
	"zond-api/internal/api/service"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func SetupRouter(db *pgxpool.Pool, jwtSecret string) *gin.Engine {
	blockRepo := repository.NewBlockRepoPG(db)
	txRepo := repository.NewTransactionRepoPG(db)
	addrRepo := repository.NewAddressRepoPG(db)
	forkRepo := repository.NewForkRepoPG(db)
	chainRepo := repository.NewChainRepoPG(db)
	reorgRepo := repository.NewReorgRepoPG(db)
	validatorRepo := repository.NewValidatorRepoPG(db)
	blockService := service.NewBlockService(blockRepo)
	txService := service.NewTransactionService(txRepo)
	addrService := service.NewAddressService(addrRepo)
	forkService := service.NewForkService(forkRepo)
	chainService := service.NewChainService(chainRepo)
	reorgService := service.NewReorgService(reorgRepo)
	validatorService := service.NewValidatorService(validatorRepo)
	blockHandler := handler.NewBlockHandler(blockService)
	txHandler := handler.NewTransactionHandler(txService)
	addrHandler := handler.NewAddressHandler(addrService)
	forkHandler := handler.NewForkHandler(forkService)
	chainHandler := handler.NewChainHandler(chainService)
	reorgHandler := handler.NewReorgHandler(reorgService)
	validatorHandler := handler.NewValidatorHandler(validatorService)
	r := gin.Default()
	r.POST("/api/login", loginHandler(jwtSecret))
	api := r.Group("/api")
	api.Use(jwtMiddleware(jwtSecret))
	blocks := api.Group("/blocks")
	{
		blocks.GET("/latest", blockHandler.GetLatestBlocks)
		blocks.GET("/:block_number", blockHandler.GetBlockByNumber)
	}

	transactions := api.Group("/transactions")
	{
		transactions.GET("/latest", txHandler.GetLatestTransactions)
		transactions.GET("/:tx_hash", txHandler.GetTransactionByHash)
	}

	addresses := api.Group("/addresses")
	{
		addresses.GET("/:address/balance", addrHandler.GetAddressBalance)
		addresses.GET("/:address/transactions", addrHandler.GetAddressTransactions)
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
	}

	return r
}

func loginHandler(jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}
		if user.Username != "admin" || user.Password != "password123" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			return
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"username": user.Username,
			"role":     "admin",
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})
		tokenString, err := token.SignedString([]byte(jwtSecret))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": tokenString})
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
			c.Set("username", claims["username"])
			c.Set("role", claims["role"])
		}

		c.Next()
	}
}
