package storage

import (
	"context"
	"to_do_app/api/models"
)

type IStorage interface {
	Close()
	User() IUserStorage
	Task() ITaskStorage
	TaskList() ITaskListStorage
	Label() ILabelStorage
}

type IUserStorage interface {
	Create(context.Context, models.CreateUser) (models.User, error)
	Get(context.Context, models.PrimaryKey) (models.User, error)
 	Update(context.Context, models.UpdateUser) (models.UpdateResponse, error)
	Delete(context.Context, models.PrimaryKey) (error)
	GetPassword(context.Context, models.PrimaryKey) (string, error)
	ChangePassword(context.Context, models.UpdateUserPassword) (error)
	GetUserCredentialsByEmail(context.Context, models.UserLoginRequest) (models.UserCredentials, error)
}

type ITaskStorage interface {
	Create(context.Context, models.CreateTask) (models.Task, error)
	Get(context.Context, models.PrimaryKey) (models.Task, error)
 	Update(context.Context, models.UpdateTask) (models.UpdateResponse, error)
	Delete(context.Context, models.PrimaryKey) (error)
	DeleteByTaskListID(context.Context, models.PrimaryKey) (error)
	DeleteByTaskListIDs(context.Context, []string) (error)
}

type ITaskListStorage interface {
	Create(context.Context, models.CreateTaskList) (models.TaskList, error)
	Get(context.Context, models.PrimaryKey) (models.TaskList, error)
 	Update(context.Context, models.UpdateTaskList) (models.UpdateResponse, error)
	Delete(context.Context, models.PrimaryKey) (error)
	GetTaskListIDByUserID(context.Context, models.PrimaryKey) ([]string, error)
	DeleteByUserID(context.Context, models.PrimaryKey) (error)
}

type ILabelStorage interface {
	Create(context.Context, models.CreateLabel) (models.Label, error)
	Get(context.Context, models.PrimaryKey) (models.Label, error)
 	Update(context.Context, models.UpdateLabel) (models.UpdateResponse, error)
	Delete(context.Context, models.PrimaryKey) (error)
	DeleteByUserID(context.Context, models.PrimaryKey) (error)
}

