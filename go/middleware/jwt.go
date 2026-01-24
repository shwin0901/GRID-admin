package middleware

import (
	"strings"
	"time"

	"go-admin/config"
	"go-admin/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// Claims represents JWT claims
type Claims struct {
	UserID   uint     `json:"user_id"`
	Username string   `json:"username"`
	Roles    []string `json:"roles"`
	jwt.RegisteredClaims
}

// GenerateToken generates a JWT token
func GenerateToken(userID uint, username string, roles []string) (string, error) {
	claims := Claims{
		UserID:   userID,
		Username: username,
		Roles:    roles,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // Expires in 24 hours
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "grid-admin",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.AppConfig.JWTSecret))
}

// ParseToken parses a JWT token
func ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.AppConfig.JWTSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, jwt.ErrSignatureInvalid
}

// JWTAuth is a middleware for JWT authentication
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the token from the header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			utils.Unauthorized(c, "Please log in first")
			c.Abort()
			return
		}

		// Check the Bearer prefix
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			utils.Unauthorized(c, "Token format error")
			c.Abort()
			return
		}

		// Parse the token
		claims, err := ParseToken(parts[1])
		if err != nil {
			utils.Unauthorized(c, "Token is invalid or expired")
			c.Abort()
			return
		}

		// Store user information in the context
		c.Set("userID", claims.UserID)
		c.Set("username", claims.Username)
		c.Set("roles", claims.Roles)
		c.Next()
	}
}
