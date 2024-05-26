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

type taskRepo struct {
	DB *pgxpool.Pool
}

func NewTaskRepo(db *pgxpool.Pool) storage.ITaskStorage {
	return &taskRepo{
		DB: db,
	}
}

func (t *taskRepo) Create(ctx context.Context, createTask models.CreateTask) (models.Task, error) {
	task := models.Task{}
	uid := uuid.New()

	query := `INSERT INTO tasks (id, title, description, due_date, task_list_id)
			values ($1, $2, $3, $4, $5)
		returning id, title, description, due_date::text, task_list_id, created_at::text `

	err := t.DB.QueryRow(ctx, query, 
		uid,
		createTask.Title,
		createTask.Description,
		createTask.DueDate,
 		createTask.TaskListID,
	).Scan(
		&task.ID,
		&task.Title,
		&task.Description,
		&task.DueDate,
 		&task.TaskListID,
		&task.CreatedAt,
	)
	if err != nil{
		fmt.Println("error while creating task!", err.Error())
		return models.Task{},err
	}

	return task, nil
}

func (t *taskRepo) Get(ctx context.Context, pKey models.PrimaryKey) (models.Task, error) {
	task := models.Task{}

	query := `SELECT id, title, description, due_date::text, task_list_id, created_at::text, updated_at::text 
			FROM tasks where id = $1 `

	err := t.DB.QueryRow(ctx,query, pKey.ID).Scan(
		&task.ID,
		&task.Title,
		&task.Description,
		&task.DueDate,
 		&task.TaskListID,
		&task.CreatedAt,
		&task.UpdatedAt,
	)

	if err != nil{
		fmt.Println("error while getting task!", err.Error())
		return models.Task{}, err
	}

	return task, nil
}

func (t *taskRepo) Update(ctx context.Context, updTask models.UpdateTask) (models.UpdateResponse, error) {
	var (
		params = make(map[string]interface{})
		filter = ""
		query = `UPDATE tasks set `
		task = models.UpdateResponse{}
	)

	params["id"] = updTask.ID

	if updTask.Title != ""{
		params["title"] = updTask.Title
		filter += " title = @title, "
	}

	if updTask.Description != ""{
		params["description"] = updTask.Description
		filter += " description = @description, "
	}

	if updTask.DueDate != ""{
		params["due_date"] = updTask.DueDate
		filter += " due_date = @due_date, "
	}

	if updTask.TaskListID != ""{
		params["task_list_id"] = updTask.TaskListID
		filter += " task_list_id = @task_list_id, "
	}

	query += filter + ` updated_at = now() WHERE id = @id 
		returning id, updated_at::text `
	
	fullQuery, args := helper.ReplaceQueryParams(query, params)

	err := t.DB.QueryRow(ctx, fullQuery, args...).Scan(
		&task.ID,
		&task.UpdatedAt,
	)
	fmt.Println(fullQuery)
	if err != nil{
		fmt.Println("error while updating tasks!", err.Error())
		return models.UpdateResponse{}, err
	}

	return task, nil
}

func (t *taskRepo) Delete(ctx context.Context, pKey models.PrimaryKey) (error) {
	query := `DELETE FROM tasks where id = $1 `

	_, err := t.DB.Exec(ctx, query, pKey.ID)
	if err != nil{
		fmt.Println("error while deleting task!", err.Error())
		return err
	}

	return nil
}

func (t *taskRepo) DeleteByTaskListID(ctx context.Context, pKey models.PrimaryKey) (error) {
	query := `DELETE FROM tasks where task_list_id = $1 `

	_, err := t.DB.Exec(ctx, query, pKey.ID)
	if err != nil{
		fmt.Println("error while deletin task by task_list id!",err.Error())
		return err
	}

	return nil
}

func (t *taskRepo) DeleteByTaskListIDs(ctx context.Context, taskListIDs []string) (error) {
	query := `DELETE FROM tasks where task_list_id = $1 `

	for _, id := range taskListIDs{
		_, err := t.DB.Exec(ctx,query,id)
		if err != nil{
			fmt.Println("error while deleting tasks by task list ids!", err.Error())
			return err
		}
	}

	return nil
}