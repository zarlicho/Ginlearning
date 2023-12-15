package security

import (
	"errors"
	"fmt"
	"ginlearning/model/domain"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func PassEncrypt(pass domain.RegisterData) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass.Passwords), 20)
	if err != nil {
		return "", errors.New("failed to create hash")
	}
	return string(hash), nil
}

func ClaimsJwt(logData domain.LoginData, userData string) (string, error) {
	err := bcrypt.CompareHashAndPassword([]byte(logData.Passwords), []byte(userData))
	if err != nil {
		return "", errors.New("password salah")
	}

	// Membuat token JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": logData.Id,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	tokenString, err := token.SignedString([]byte("asc22332afg0061729940qqr"))
	if err != nil {
		return "", errors.New("gagal membuat token JWT")
	}
	fmt.Println(tokenString)
	return tokenString, nil
}
