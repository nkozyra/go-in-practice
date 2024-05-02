package main

import (
	"log"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

var SIGNING_KEY = []byte("this-value-should-be-secret")

type claim struct {
	jwt.RegisteredClaims
}

func generateClaim() (string, error) {
	claims := claim{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Subject:   "nobody@example.com",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(SIGNING_KEY)
	if err != nil {
		return "", err
	}
	return ss, nil
}

func main() {
	if signed, err := generateClaim(); err != nil {
		panic(err)
	} else {
		token, err := jwt.Parse(signed, func(token *jwt.Token) (interface{}, error) {
			return SIGNING_KEY, nil
		})
		if err != nil {
			panic(err)
		}
		if validatedClaims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			log.Println(validatedClaims["sub"])
		} else {
			panic("error getting claims")
		}
	}
}
