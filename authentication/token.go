package authentication

import (
	"api/src/config"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func MakeToken(userID uint64) (string, error) {
	//
	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	permissions["expiration"] = time.Now().Add(time.Hour * 6).Unix()
	permissions["UserId"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)
	return token.SignedString([]byte(config.SecretKey)) // secret
}

func ValidateToken(r *http.Request) error {
	//
	strToken := extractToken(r)
	token, err := jwt.Parse(strToken, returnVerificationKey)
	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("token inválido")
}

func extractToken(r *http.Request) string {
	//
	token := r.Header.Get("Authorization")

	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}

	return ""
}

func returnVerificationKey(token *jwt.Token) (interface{}, error) {
	//
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("método de assinatura inválido! %v", token.Header["alg"])
	}

	return config.SecretKey, nil
}
