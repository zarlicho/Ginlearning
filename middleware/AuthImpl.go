package middleware

import (
	// "net/http"

	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type MiddlewareImpl struct{}

func NewMiddleware() Middlewares {
	return &MiddlewareImpl{}
}

func (m *MiddlewareImpl) Midware(c *gin.Context) {
	// Your middleware logic goes here

	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "login error",
		})
	}
	fmt.Printf("Nilai Cookie: %s\n", tokenString)

	// Parse token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("asc22332afg0061729940qqr"), nil
	})

	if err != nil {
		panic("token not found!")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			// Token expired
			return
		}

		// Call the actual handler if the token is valid
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

}
