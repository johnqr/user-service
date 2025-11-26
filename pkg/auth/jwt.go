package auth

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var ErrInvalidToken = errors.New("invalid token")

func secret() []byte {
	if s := os.Getenv("JWT_SECRET"); s != "" { return []byte(s) }
	return []byte("changeme")
}

func GenerateToken(userID string) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"sub": userID, "exp": time.Now().Add(24 * time.Hour).Unix()})
	return t.SignedString(secret())
}

func ParseToken(token string) (string, error) {
	t, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) { return secret(), nil })
	if err != nil || !t.Valid { return "", ErrInvalidToken }
	if claims, ok := t.Claims.(jwt.MapClaims); ok {
		if sub, ok := claims["sub"].(string); ok { return sub, nil }
	}
	return "", ErrInvalidToken
}
