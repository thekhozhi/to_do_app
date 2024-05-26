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

type taskListRepo struct {
	DB *pgxpool.Pool
}

func NewTaskListRepo(db *pgxpool.Pool) storage.ITaskListStorage {
	return &taskListRepo{
		DB: db,
	}
}

func (t *taskListRepo) Create(ctx context.Context, createTaskList models.CreateTaskList) (models.TaskList, error) {
	taskList := models.TaskList{}
	uid := uuid.New()

	query := `INSERT INTO task_lists (id, title, description, user_id)
			values ($1, $2, $3, $4)
		returning id, title, description, user_id, created_at::text `

	err := t.DB.QueryRow(ctx, query, 
		uid,
		createTaskList.Title,
		createTaskList.Description,
		createTaskList.UserID,
 	).Scan(
		&taskList.ID,
		&taskList.Title,
		&taskList.Description,
		&taskList.UserID,
 		&taskList.CreatedAt,
	)
	if err != nil{
		fmt.Println("error while creating taskList!", err.Error())
		return models.TaskList{},err
	}

	return taskList, nil
}

func (t *taskListRepo) Get(ctx context.Context, pKey models.PrimaryKey) (models.TaskList, error) {
	taskList := models.TaskList{}

	query := `SELECT id, title, description, user_id, created_at::text, updated_at::text 
			FROM task_lists where id = $1 `

	err := t.DB.QueryRow(ctx,query, pKey.ID).Scan(
		&taskList.ID,
		&taskList.Title,
		&taskList.Description,
		&taskList.UserID,
		&taskList.CreatedAt,
		&taskList.UpdatedAt,
	)

	if err != nil{
		fmt.Println("error while getting taskList!", err.Error())
		return models.TaskList{}, err
	}

	return taskList, nil
}

func (t *taskListRepo) Update(ctx context.Context, updTaskList models.UpdateTaskList) (models.UpdateResponse, error) {
	var (
		params = make(map[string]interface{})
		filter = ""
		query = `UPDATE task_lists set `
		taskList = models.UpdateResponse{}
	)

	params["id"] = updTaskList.ID

	if updTaskList.Title != ""{
		params["title"] = updTaskList.Title
		filter += " title = @title, "
	}

	if updTaskList.Description != ""{
		params["description"] = updTaskList.Description
		filter += " description = @description, "
	}

	query += filter + ` updated_at = now() WHERE id = @id 
		returning id, updated_at::text `
	
	fullQuery, args := helper.ReplaceQueryParams(query, params)

	err := t.DB.QueryRow(ctx, fullQuery, args...).Scan(
		&taskList.ID,
		&taskList.UpdatedAt,
	)
	fmt.Println(fullQuery)
	if err != nil{
		fmt.Println("error while updating task_lists!", err.Error())
		return models.UpdateResponse{}, err
	}

	return taskList, nil
}

func (t *taskListRepo) Delete(ctx context.Context, pKey models.PrimaryKey) (error) {
	query := `DELETE FROM task_lists where id = $1 `

	_, err := t.DB.Exec(ctx, query, pKey.ID)
	if err != nil{
		fmt.Println("error while deleting taskList!", err.Error())
		return err
	}

	return nil
}

func (t *taskListRepo) GetTaskListIDByUserID(ctx context.Context, pKey models.PrimaryKey) ([]string, error) {
	sliceID := []string{}
	
	query := `SELECT id from task_lists where user_id = $1 `

	rows, err := t.DB.Query(ctx, query, pKey.ID)
	if err != nil{
		fmt.Println("error while query task list id query rows!", err.Error())
		return []string{},err
	}

	for rows.Next(){
		id := ""

		err := rows.Scan(
			&id,
		)
		if err != nil{
			fmt.Println("error while scannig task list ids!", err.Error())
			return []string{},err
		}

		sliceID = append(sliceID, id)
	}

	return sliceID, nil
}

func (t *taskListRepo) DeleteByUserID(ctx context.Context, pKey models.PrimaryKey) (error) {
	query := `DELETE from task_lists where user_id = $1 `

	_, err := t.DB.Exec(ctx,query, pKey.ID)
	if err != nil{
		fmt.Println("error while deleting task lists by user id!", err.Error())
		return err
	}

	return nil
}