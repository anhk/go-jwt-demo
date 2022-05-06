package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

var JWTKEY = []byte("HelloWorld")

func JwtSign(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"expire":   time.Now().Add(time.Hour * 2).Unix(),
	})
	return token.SignedString(JWTKEY)
}

func JwtVerify(tokenStr string) (string, bool) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) { return JWTKEY, nil })
	if err != nil {
		return "", false
	}
	if token.Valid {
		fmt.Println(token.Claims)
		claim := token.Claims.(jwt.MapClaims)
		u := claim["username"].(string)
		return u, true
	}
	return "", false
}

func init() {
	if key := os.Getenv("JWTKEY"); key != "" {
		JWTKEY = []byte(key)
	}
}
