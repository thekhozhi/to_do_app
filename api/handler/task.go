package handler

import (
	"context"
	"errors"
	"net/http"
 	"to_do_app/api/models"
	"to_do_app/pkg"

	"github.com/gin-gonic/gin"
)

// CreateTask godoc
// @Security ApiKeyAuth
// @Router       /task [POST]
// @Summary      Creates a new task
// @Description  create a new task
// @Tags         task
// @Accept       json
// @Produce      json
// @Param        task body models.CreateTask false "task"
// @Success      201  {object}  models.Task
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) CreateTask(c *gin.Context) {
	auth, err := h.GetAuthInfoFromToken(c)
	if err != nil{
		handleResponse(c, "Error Unauthorized!",http.StatusUnauthorized, pkg.ErrUnauthorized.Error())
	}

	createTask := models.CreateTask{}

	err = c.ShouldBindJSON(&createTask)
	if err != nil{
		handleResponse(c,"Error in handler, while reading body from client!",http.StatusBadRequest, err.Error())
		return
	}

	taskList, err := h.services.TaskList().Get(context.Background(), models.PrimaryKey{ID: createTask.TaskListID})
	if err != nil{
		handleResponse(c, "Error in handler, while getting task list while creating task list", http.StatusInternalServerError, err.Error())
		return
	}

	if auth.UserID != taskList.UserID{
		handleResponse(c, "task list id is wrong!", http.StatusBadRequest, errors.New("task list id is wrong").Error())
		return
	}

	task, err := h.services.Task().Create(context.Background(),createTask)
	if err != nil{
		handleResponse(c, "Error in handler, while crating task!",http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "Task created!", http.StatusCreated, task)
}

// GetTask godoc
// @Security ApiKeyAuth
// @Router       /task/{id} [GET]
// @Summary      Get task by id
// @Description  get task by id
// @Tags         task
// @Accept       json
// @Produce      json
// @Param        id path string true "task_id"
// @Success      201  {object}  models.Task
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetTask(c *gin.Context) {
	auth, err := h.GetAuthInfoFromToken(c)
	if err != nil{
		handleResponse(c, "Error Unauthorized!",http.StatusUnauthorized, pkg.ErrUnauthorized.Error())
	}

	id := c.Param("id")

	if id == ""{
		handleResponse(c, pkg.IdErr, http.StatusBadRequest, errors.New(pkg.IdErr).Error())
		return
	}


	task, err := h.services.Task().Get(context.Background(),models.PrimaryKey{ID: id})
	if err != nil{
		handleResponse(c, "Error in handler, while getting task!", http.StatusInternalServerError, err.Error())
		return
	}

	taskList, err := h.services.TaskList().Get(context.Background(), models.PrimaryKey{ID: task.TaskListID})
	if err != nil{
		handleResponse(c, "Error in handler, while getting task list!", http.StatusInternalServerError, err.Error())
		return
	}

	if auth.UserID != taskList.UserID{
		handleResponse(c, "Error in handlers, task id is wrong!", http.StatusBadRequest, errors.New("task id is wrong").Error())
		return
	}

	handleResponse(c, "", http.StatusOK, task)
}

// UpdateTask godoc
// @Security ApiKeyAuth
// @Router       /task/{id} [PUT]
// @Summary      Update task
// @Description  update task
// @Tags         task
// @Accept       json
// @Produce      json
// @Param        id path string true "task_id"
// @Param        task body models.UpdateTask false "task"
// @Success      201  {object}  models.UpdateResponse
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) UpdateTask(c *gin.Context) {
	auth, err := h.GetAuthInfoFromToken(c)
	if err != nil{
		handleResponse(c, "Error Unauthorized!",http.StatusUnauthorized, pkg.ErrUnauthorized.Error())
	}

	updTask := models.UpdateTask{}

	id := c.Param("id")
	if id == ""{
		handleResponse(c, pkg.IdErr, http.StatusBadRequest, errors.New(pkg.IdErr).Error())
		return
	}

	err = c.ShouldBindJSON(&updTask)
	if err != nil{
		handleResponse(c, "Error in handler, while reading body from client!", http.StatusBadRequest, err.Error())
		return
	}

	taskList, err := h.services.TaskList().Get(context.Background(), models.PrimaryKey{ID: updTask.TaskListID})
	if err != nil{
		handleResponse(c, "Error in handler, while getting task list while updating task!", http.StatusInternalServerError, err.Error())
		return
	}

	if auth.UserID != taskList.UserID{
		handleResponse(c, "Error in handler, task list id is wrong!", http.StatusBadRequest, errors.New("task list id is wrong").Error())
		return
	}

	updTask.ID = id

	task, err := h.services.Task().Update(context.Background(), updTask)
	if err != nil{
		handleResponse(c, "Error in handler, while updating task!", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "Successfully Updated!", http.StatusOK, task)
}

// DeleteTask godoc
// @Security ApiKeyAuth
// @Router       /task/{id} [DELETE]
// @Summary      Delete task
// @Description  delete task
// @Tags         task
// @Accept       json
// @Produce      json
// @Param        id path string true "task_id"
// @Success      201  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) DeleteTask(c *gin.Context) {
	auth, err := h.GetAuthInfoFromToken(c)
	if err != nil{
		handleResponse(c, "Error Unauthorized!",http.StatusUnauthorized, pkg.ErrUnauthorized.Error())
	}

	id := c.Param("id")
	if id == ""{
		handleResponse(c, pkg.IdErr, http.StatusBadRequest, errors.New(pkg.IdErr).Error())
		return
	}

	task, err := h.services.Task().Get(context.Background(), models.PrimaryKey{ID: id})
	if err != nil{
		handleResponse(c, "Error in handler, while getting task while deleting task!", http.StatusInternalServerError,err.Error())
		return
	}

	taskList, err := h.services.TaskList().Get(context.Background(), models.PrimaryKey{ID: task.TaskListID})
	if err != nil{
		handleResponse(c, "Error in handler, while getting task list while updating task!", http.StatusInternalServerError, err.Error())
		return
	}

	if auth.UserID != taskList.UserID{
		handleResponse(c, "Error in handler, task id is wrong!", http.StatusBadRequest, errors.New("task id is wrong").Error())
		return
	}

	err = h.services.Task().Delete(context.Background(), models.PrimaryKey{ID: id})
	if err != nil{
		handleResponse(c, "Error in handler, while deleting task!", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "Task succesfully deleted!", http.StatusOK, "Task succesfully deleted!")
}