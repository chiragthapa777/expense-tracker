package utils

import (
	"time"

	"github.com/chiragthapa777/expense-tracker-api/internal/config"
	"github.com/chiragthapa777/expense-tracker-api/internal/logger"
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

func GenerateJWT(userID string) (string, error) {
	log := logger.GetLogger()
	cfg := config.GetConfig()
	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(cfg.JWTSecret))
	if err != nil {
		log.Errorf("Failed to generate JWT: %v", err)
		return "", err
	}
	return signedToken, nil
}

func ValidateJWT(tokenStr string) (*Claims, error) {
	log := logger.GetLogger()
	cfg := config.GetConfig()
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(cfg.JWTSecret), nil
	})
	if err != nil {
		log.Errorf("Failed to validate JWT: %v", err)
		return nil, err
	}
	if !token.Valid {
		log.Error("Invalid JWT token")
		return nil, jwt.ErrSignatureInvalid
	}
	return claims, nil
}
