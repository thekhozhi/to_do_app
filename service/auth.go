package service

import (
	"context"
	"fmt"
	"to_do_app/api/models"
	"to_do_app/pkg/security"
	"to_do_app/storage"
)

type authService struct {
	storage storage.IStorage
}

func NewAuthService(storage storage.IStorage) authService {
	return authService{
		storage: storage,
	}
}

func (a authService) UserLogin(ctx context.Context, req models.UserLoginRequest) (string, error) {
	user, err := a.storage.User().GetUserCredentialsByEmail(ctx, req)
	if err != nil{
		fmt.Println("Error in service layer, given email is wrong!", err.Error())
		return "", err
	}

	err = security.CompareHashAndPassword(user.HashedPassword, req.Password)
	if err != nil{
		fmt.Println("Error in service layer, given password is wrong!", err.Error())
		return "", err
	}

	return user.ID, nil
}

