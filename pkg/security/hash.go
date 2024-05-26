package security

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil{
		fmt.Println("Error while hashing the password!", err.Error())
		return "", nil
	}

	return string(hashedPassword), nil
}

func CompareHashAndPassword(hashedPassword, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword),[]byte(password))
	if err != nil{
		fmt.Println("Error, password is not correct!", err.Error())
		return err
	}

	return nil
}