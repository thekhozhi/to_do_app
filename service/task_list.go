package service

import (
	"context"
	"fmt"
	"to_do_app/api/models"
	"to_do_app/storage"
)

type taskListService struct {
	storage storage.IStorage
}

func NewTaskListService(storage storage.IStorage) taskListService {
	return taskListService{
		storage: storage,
	}
}

func (t taskListService) Create(ctx context.Context, req models.CreateTaskList) (models.TaskList, error) {
	taskList, err := t.storage.TaskList().Create(ctx,req)
	if err != nil{
		fmt.Println("Error in service layer, while creating taskList!", err.Error())
		return models.TaskList{}, err
	}

	return taskList, nil
}

func (t taskListService) Get(ctx context.Context, req models.PrimaryKey) (models.TaskList, error) {
	taskList, err := t.storage.TaskList().Get(ctx, req)
	if err != nil{
		fmt.Println("Error in service layer, while getting taskList!", err.Error())
		return models.TaskList{}, err
	}

	return taskList, nil
}

func (t taskListService) Update(ctx context.Context, req models.UpdateTaskList) (models.UpdateResponse, error) {
	taskList, err := t.storage.TaskList().Update(ctx, req)
	if err != nil{
		fmt.Println("Error in service layer, while updating taskList!", err.Error())
		return models.UpdateResponse{},err
	}

	return taskList, nil
}

func (t taskListService) Delete(ctx context.Context, req models.PrimaryKey) (error) {
	err := t.storage.Task().DeleteByTaskListID(ctx, req)
	if err != nil{
		fmt.Println("Error in service layer, while deleting tasks by task list id!", err.Error())
		return err
	}

	err = t.storage.TaskList().Delete(ctx, req)
	if err != nil{
		fmt.Println("Error in service layer, while deleting taskList!", err.Error())
		return err
	}
	
	return nil
}