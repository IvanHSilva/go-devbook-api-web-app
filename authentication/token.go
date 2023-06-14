package authentication

import (
	"api/src/config"
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