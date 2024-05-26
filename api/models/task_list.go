package models

type TaskList struct {
	ID		    string`json:"id"`
	Title       string`json:"title"`
    Description string`json:"description"`
	UserID      string`json:"user_id"`
	CreatedAt   string`json:"created_at"`
	UpdatedAt   string`json:"updated_at"`
}

type CreateTaskList struct {
	Title       string`json:"title"`
    Description string`json:"description"`
	UserID      string`json:"user_id"`
}

type UpdateTaskList struct {
	ID		    string`json:"-"`
	Title       string`json:"title"`
    Description string`json:"description"`
}

 