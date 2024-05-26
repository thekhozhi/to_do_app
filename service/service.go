package service

import "to_do_app/storage"

type IServiceManager interface {
	User() userService
	Task() taskService
	TaskList() taskListService
	Label() labelService
	Auth() authService
}

type Service struct {
	userService userService
	taskService taskService
	taskListService taskListService
	labelService labelService
	authService authService
}

func New(storage storage.IStorage) IServiceManager {
	service := Service{}

	service.userService = NewUserService(storage)
	service.taskService = NewTaskService(storage)
	service.taskListService = NewTaskListService(storage)
	service.labelService = NewLabelService(storage)
	service.authService = NewAuthService(storage)

	return service
}

func (s Service) User() userService {
	return s.userService
}

func (s Service) Task() taskService {
	return s.taskService
}

func (s Service) TaskList() taskListService {
	return s.taskListService
}

func (s Service) Label() labelService {
	return s.labelService
}

func (s Service) Auth() authService {
	return s.authService
}