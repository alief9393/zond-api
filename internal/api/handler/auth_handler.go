package handler

import (
	"context"
	"os"
	"time"
	"zond-api/internal/api/dto"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	DB *pgx.Conn
}

// Login godoc
// @Summary      User login
// @Description  Authenticate user and return JWT token if credentials are valid
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        request  body      dto.LoginRequest  true  "Login credentials"
// @Success      200      {object}  dto.LoginResponse
// @Failure      400      {object}  map[string]string "Invalid request format"
// @Failure      401      {object}  map[string]string "Invalid credentials"
// @Failure      500      {object}  map[string]string "Failed to generate token"
// @Router       /auth/login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	var password, username string
	var isPaid bool
	err := h.DB.QueryRow(context.Background(), `
		SELECT username, password, is_paid 
		FROM users WHERE username = $1`, req.Username).
		Scan(&username, &password, &isPaid)
	if err != nil {
		c.JSON(401, gin.H{"error": "Invalid credentials"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(password), []byte(req.Password)); err != nil {
		c.JSON(401, gin.H{"error": "Invalid credentials"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"is_paid":  isPaid,
		"exp":      time.Now().Add(time.Hour * 1).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(200, dto.LoginResponse{Token: tokenString, IsPaid: isPaid})
}

// Register godoc
// @Summary      Register new user
// @Description  Create a new user with username and password
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        request  body      dto.LoginRequest  true  "User registration data"
// @Success      201      {object}  map[string]string "User created successfully"
// @Failure      400      {object}  map[string]string "Username already exists or invalid request"
// @Failure      500      {object}  map[string]string "Failed to hash password"
// @Router       /auth/register [post]
func (h *AuthHandler) Register(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to hash password"})
		return
	}

	_, err = h.DB.Exec(context.Background(), `
		INSERT INTO users (username, password, is_paid) 
		VALUES ($1, $2, $3)`, req.Username, string(hashedPassword), false)
	if err != nil {
		c.JSON(400, gin.H{"error": "Username already exists"})
		return
	}

	c.JSON(201, gin.H{"message": "User created"})
}
