package pkg

import (
	"errors"
	"time"
)

var(
	SignKey = []byte("asd@#lskd2!aw32k34242WSASdsk32")


	IdErr = "id column is empty, please fill the id column"
	ErrUnauthorized = errors.New("Unauthorized")
	EmailErr = "email is too short, the minimum size is 13"
	PasswordErr = "password is too short, the minimum size is 8"
	UserNameErr = "user name is too short, the minimum size is 5"
)

const (
	AccessExpireTime = time.Minute * 20
	RefreshExpireTime = time.Hour * 24
)