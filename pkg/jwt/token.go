package jwt

import (
	"errors"
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func GenerateJWT(m map[string]interface{}, tokenExpireTime time.Duration, tokenSecretKey string) (string, error){

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	for key, value := range m {
		claims[key] = value
	}

	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(tokenExpireTime).Unix()

	tokenString, err := token.SignedString([]byte(tokenSecretKey))
	if err != nil{
		fmt.Println("Error while getting token signed string!", err.Error())
		return "", err
	}

	return tokenString, nil
}

func ExtractClaims(tokenString string, tokenSecretKey string) (jwt.MapClaims, error){
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(tokenSecretKey), nil
	})

	if err != nil{
		fmt.Println("Error while parsing to token from token string!", err.Error())
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !(ok && token.Valid) {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}