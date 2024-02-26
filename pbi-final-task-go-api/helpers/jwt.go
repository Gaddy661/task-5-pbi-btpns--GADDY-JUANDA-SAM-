package helpers

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func createToken(ID uint) *jwt.Token {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
}

func CreateTokenString(ID uint) (string, error) {
	token := createToken(ID)
	return token.SignedString([]byte(os.Getenv("SECRET")))
}

func ParseTokenString(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRET")), nil
	})
}
