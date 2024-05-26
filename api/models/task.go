package models

type Task struct {
	ID		    string`json:"id"`
	Title       string`json:"title"`
    Description string`json:"description"`
    DueDate     string`json:"due_date"`
 	TaskListID  string`json:"task_list_id"`
	CreatedAt   string`json:"created_at"`
	UpdatedAt   string`json:"updated_at"`
}

type CreateTask struct {
	Title       string`json:"title"`
    Description string`json:"description"`
    DueDate     string`json:"due_date"`
 	TaskListID  string`json:"task_list_id"`
}

type UpdateTask struct {
	ID		    string`json:"-"`
	Title       string`json:"title"`
    Description string`json:"description"`
    DueDate     string`json:"due_date"`
 	TaskListID  string`json:"task_list_id"`
}