package lib

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

var tokenKey = []byte("secret")

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func GenerateJWT(user User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username":   user.Username,
		"exp":        time.Now().Add(10 * time.Minute),
		"authorized": true,
	})

	tokenString, err := token.SignedString(tokenKey)
	if err != nil {
		return "", err
	}
	return tokenString, err
}

func ValidateJWT(tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("validation error: %v", token.Header["alg"])
		}
		return tokenKey, nil
	})

	if err != nil {
		return false, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if claims["exp"].(time.Time).After(time.Now()) {
			return true, nil
		}
		return false, fmt.Errorf("token expired: %v", claims["exp"])
	}
	return false, nil
}
