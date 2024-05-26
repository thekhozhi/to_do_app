package models

type User struct {
	ID		  string`json:"id"`
	UserName  string`json:"user_name"`
	Email     string`json:"email"`
 	CreatedAt string`json:"created_at"`
	UpdatedAt string`json:"updated_at"`
}

type CreateUser struct {
	UserName  string`json:"user_name"`
	Email     string`json:"email"`
	Password  string`json:"password"`
}

type UpdateUser struct {
	ID	      string`json:"-"`
	UserName  string`json:"user_name"`
	Email     string`json:"email"`
}

type UpdateUserPassword struct {
	ID          string`json:"-"`
	OldPassword string`json:"old_password"`
	NewPassword string`json:"new_password"`
}

type UserCredentials struct {
	ID 			   string`json:"id"`
	HashedPassword string`json:"hashed_password"`
}
