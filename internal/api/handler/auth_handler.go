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
