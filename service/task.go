package service

import (
	"context"
	"fmt"
	"to_do_app/api/models"
	"to_do_app/storage"
)

type taskService struct {
	storage storage.IStorage
}

func NewTaskService(storage storage.IStorage) taskService {
	return taskService{
		storage: storage,
	}
}

func (t taskService) Create(ctx context.Context, req models.CreateTask) (models.Task, error) {
	task, err := t.storage.Task().Create(ctx,req)
	if err != nil{
		fmt.Println("Error in service layer, while creating task!", err.Error())
		return models.Task{}, err
	}

	return task, nil
}

func (t taskService) Get(ctx context.Context, req models.PrimaryKey) (models.Task, error) {
	task, err := t.storage.Task().Get(ctx, req)
	if err != nil{
		fmt.Println("Error in service layer, while getting task!", err.Error())
		return models.Task{}, err
	}

	return task, nil
}

func (t taskService) Update(ctx context.Context, req models.UpdateTask) (models.UpdateResponse, error) {
	task, err := t.storage.Task().Update(ctx, req)
	if err != nil{
		fmt.Println("Error in service layer, while updating task!", err.Error())
		return models.UpdateResponse{},err
	}

	return task, nil
}

func (t taskService) Delete(ctx context.Context, req models.PrimaryKey) (error) {
	err := t.storage.Task().Delete(ctx, req)
	if err != nil{
		fmt.Println("Error in service layer, while deleting task!", err.Error())
		return err
	}
	
	return nil
}