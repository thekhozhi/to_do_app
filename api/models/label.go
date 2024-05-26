package models

type Label struct {
	ID		    string`json:"id"`
	Name        string`json:"name"`
	Color       string`json:"color"`
	UserID      string`json:"user_id"`
	CreatedAt   string`json:"created_at"`
	UpdatedAt   string`json:"updated_at"`
}

type CreateLabel struct {
	Name        string`json:"name"`
	Color       string`json:"color"`
	UserID      string`json:"user_id"`
}

type UpdateLabel struct {
	ID		    string`json:"-"`
	Name        string`json:"name"`
	Color       string`json:"color"`
}