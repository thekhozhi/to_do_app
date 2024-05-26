package service

import (
	"context"
	"errors"
	"fmt"
	"to_do_app/api/models"
	"to_do_app/pkg"
	"to_do_app/pkg/security"
	"to_do_app/storage"
)

type userService struct {
	storage storage.IStorage
}

func NewUserService(storage storage.IStorage) userService {
	return userService{
		storage: storage,
	}
}

func (u userService) Create(ctx context.Context, req models.CreateUser) (models.User, error) {
	if len(req.Email) < 13{
		fmt.Println("Error in service layer, email is too short!", errors.New(pkg.EmailErr))
		return models.User{}, errors.New(pkg.EmailErr)
	} 

	if len(req.Password) < 8{
		fmt.Println("Error in service layer, password is too short!", errors.New(pkg.PasswordErr))
		return  models.User{}, errors.New(pkg.PasswordErr)
	}

	if len(req.UserName) < 5{
		fmt.Println("Error in service layer, user name is too short!", errors.New(pkg.UserNameErr))
		return  models.User{}, errors.New(pkg.UserNameErr)
	}

	hashedPassword, err := security.HashPassword(req.Password)
	if err != nil{
		fmt.Println("Error in service layer, while hashin password!", err.Error())
		return models.User{}, err
	}

	req.Password = hashedPassword

	user, err := u.storage.User().Create(ctx,req)
	if err != nil{
		fmt.Println("Error in service layer, while creating user!", err.Error())
		return models.User{}, err
	}

	return user, nil
}

func (u userService) Get(ctx context.Context, req models.PrimaryKey) (models.User, error) {
	user, err := u.storage.User().Get(ctx, req)
	if err != nil{
		fmt.Println("Error in service layer, while getting user!", err.Error())
		return models.User{}, err
	}

	return user, nil
}

func (u userService) Update(ctx context.Context, req models.UpdateUser) (models.UpdateResponse, error) {
	if len(req.Email) > 0 && len(req.Email) < 13{
		fmt.Println("Error in service layer, email is too short!", errors.New(pkg.EmailErr))
		return models.UpdateResponse{}, errors.New(pkg.EmailErr)
	}

	if len(req.UserName) > 0 && len(req.UserName) < 5{
		fmt.Println("Error in service layer, user name is too short!", errors.New(pkg.UserNameErr))
		return models.UpdateResponse{}, errors.New(pkg.UserNameErr)
	}

	user, err := u.storage.User().Update(ctx, req)
	if err != nil{
		fmt.Println("Error in service layer, while updating user!", err.Error())
		return models.UpdateResponse{},err
	}

	return user, nil
}

func (u userService) Delete(ctx context.Context, req models.PrimaryKey) (error) {
	err := u.storage.Label().DeleteByUserID(ctx, req)
	if err != nil{
		fmt.Println("Error in service layer, while deleting label by user id!", err.Error())
		return err
	}

	taskListIDs, err := u.storage.TaskList().GetTaskListIDByUserID(ctx, req)
	if err != nil{
		fmt.Println("Error in service layer. while getting task list ids!", err.Error())
		return err
	}

	err = u.storage.Task().DeleteByTaskListIDs(ctx, taskListIDs)
	if err != nil{
		fmt.Println("Error in service layer, while deleting tasks!", err.Error())
		return err
	}

	err = u.storage.TaskList().DeleteByUserID(ctx, req)
	if err != nil{
		fmt.Println("Error in service layer, while deleting task lists!", err.Error())
		return err
	}
	
	err = u.storage.User().Delete(ctx, req)
	if err != nil{
		fmt.Println("Error in service layer, while deleting user!", err.Error())
		return err
	}
	
	return nil
}

func (u userService) ChangePassword(ctx context.Context, req models.UpdateUserPassword) (error) {
	oldHashedPassword, err := u.storage.User().GetPassword(ctx, models.PrimaryKey{ID: req.ID})
	if err != nil{
		fmt.Println("Error in service layer, while getting user hashed password!", err.Error())
		return err
	}

	err = security.CompareHashAndPassword(oldHashedPassword,req.OldPassword)
	if err != nil{
		fmt.Println("Error in service layer, given old password is wrong!", err.Error())
		return err
	}

	if len(req.NewPassword) < 8{
		fmt.Println("Error in service layer, password is too short!", errors.New(pkg.PasswordErr))
		return errors.New(pkg.PasswordErr)
	}

	hashedPassword, err := security.HashPassword(req.NewPassword)
	if err != nil{
		fmt.Println("Error in service layer, while hashing new password!", err.Error())
		return err
	}

	req.NewPassword = hashedPassword

	err = u.storage.User().ChangePassword(ctx, req)
	if err != nil{
		fmt.Println("Error in service layer, while changing password!", err.Error())
		return err
	}

	return nil
}