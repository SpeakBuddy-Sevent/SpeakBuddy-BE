package utils

import (
	"os"
	"time"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(id uint, role string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": id,
		"role":    role,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JwtSecret())
}

func JwtSecret() []byte {
    return []byte(os.Getenv("JWT_SECRET"))
}
