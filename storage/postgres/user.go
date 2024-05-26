package postgres

import (
	"context"
	"fmt"
	"to_do_app/api/models"
	"to_do_app/pkg/helper"
	"to_do_app/storage"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type userRepo struct {
	DB *pgxpool.Pool
}

func NewUserRepo(db *pgxpool.Pool) storage.IUserStorage {
	return &userRepo{
		DB: db,
	}
}

func (u *userRepo) Create(ctx context.Context, createUser models.CreateUser) (models.User, error) {
	user := models.User{}
	uid := uuid.New()

	query := `INSERT INTO users (id, user_name, email, password)
			values ($1, $2, $3, $4)
		returning id, user_name, email, created_at::text `

	err := u.DB.QueryRow(ctx, query, 
		uid,
		createUser.UserName,
		createUser.Email,
		createUser.Password,
	).Scan(
		&user.ID,
		&user.UserName,
		&user.Email,
		&user.CreatedAt,
	)
	if err != nil{
		fmt.Println("error while creating user!", err.Error())
		return models.User{},err
	}

	return user, nil
}

func (u *userRepo) Get(ctx context.Context, pKey models.PrimaryKey) (models.User, error) {
	user := models.User{}

	query := `SELECT id, user_name, email, created_at::text, updated_at::text 
			FROM users where id = $1 `

	err := u.DB.QueryRow(ctx,query, pKey.ID).Scan(
		&user.ID,
		&user.UserName,
		&user.Email,
 		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil{
		fmt.Println("error while getting user!", err.Error())
		return models.User{}, err
	}

	return user, nil
}

func (u *userRepo) Update(ctx context.Context, updUser models.UpdateUser) (models.UpdateResponse, error) {
	var (
		params = make(map[string]interface{})
		filter = ""
		query = `UPDATE users set `
		user = models.UpdateResponse{}
	)

	params["id"] = updUser.ID

	if updUser.UserName != ""{
		params["user_name"] = updUser.UserName
		filter += " user_name = @user_name, "
	}

	if updUser.Email != ""{
		params["email"] = updUser.Email
		filter += " email = @email, "
	}

	query += filter + ` updated_at = now() WHERE id = @id 
		returning id, updated_at::text `
	
	fullQuery, args := helper.ReplaceQueryParams(query, params)

	err := u.DB.QueryRow(ctx, fullQuery, args...).Scan(
		&user.ID,
		&user.UpdatedAt,
	)
	fmt.Println(fullQuery)
	if err != nil{
		fmt.Println("error while updating users!", err.Error())
		return models.UpdateResponse{}, err
	}

	return user, nil
}

func (u *userRepo) Delete(ctx context.Context, pKey models.PrimaryKey) (error) {
	query := `DELETE FROM users where id = $1 `

	_, err := u.DB.Exec(ctx, query, pKey.ID)
	if err != nil{
		fmt.Println("error while deleting user!", err.Error())
		return err
	}

	return nil
}

func (u *userRepo) GetPassword(ctx context.Context, pKey models.PrimaryKey) (string, error) {
	password := ""
	query := `SELECT password from users where id = $1 `

	err := u.DB.QueryRow(ctx, query, pKey.ID).Scan(
		&password,
	)
	if err != nil{
		fmt.Println("error while getting password by id!", err.Error())
		return "", err
	}

	return password, nil
}

func (u *userRepo) ChangePassword(ctx context.Context, updUser models.UpdateUserPassword) (error) {
	query := `UPDATE users set password = $1 where id = $2 `

	_, err := u.DB.Exec(ctx, query,
		updUser.NewPassword,
		updUser.ID,
	)
	if err != nil{
		fmt.Println("error while changing users password!", err.Error())
		return err
	}

	return nil
}

func (u *userRepo) GetUserCredentialsByEmail(ctx context.Context, auth models.UserLoginRequest) (models.UserCredentials, error) {
	user := models.UserCredentials{}

	query := `SELECT id, password from users where email = $1 `

	err := u.DB.QueryRow(ctx, query, auth.Email).Scan(
		&user.ID,
		&user.HashedPassword,
	)
	if err != nil{
		fmt.Println("error while getting user credentials!", err.Error())
		return models.UserCredentials{}, err
	}

	return user, nil
}