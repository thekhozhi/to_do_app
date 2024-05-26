package handler

import (
	"context"
	"errors"
	"net/http"
	"to_do_app/api/models"
	"to_do_app/pkg"

	"github.com/gin-gonic/gin"
)

// CreateTaskList godoc
// @Security ApiKeyAuth
// @Router       /task_list [POST]
// @Summary      Creates a new task list
// @Description  create a new task list
// @Tags         task_list
// @Accept       json
// @Produce      json
// @Param        task_list body models.CreateTaskList false "task_list"
// @Success      201  {object}  models.TaskList
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) CreateTaskList(c *gin.Context) {
	auth, err := h.GetAuthInfoFromToken(c)
	if err != nil{
		handleResponse(c, "Error Unauthorized!",http.StatusUnauthorized, pkg.ErrUnauthorized.Error())
	}

	createTaskList := models.CreateTaskList{}

	err = c.ShouldBindJSON(&createTaskList)
	if err != nil{
		handleResponse(c,"Error in handler, while reading body from client!",http.StatusBadRequest, err.Error())
		return
	}

	if auth.UserID != createTaskList.UserID{
		handleResponse(c, "given user id is wrong!", http.StatusBadRequest, errors.New("given user id is wrong").Error())
		return
	}
	
	taskList, err := h.services.TaskList().Create(context.Background(),createTaskList)
	if err != nil{
		handleResponse(c, "Error in handler, while crating task_list!",http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "TaskList created!", http.StatusCreated, taskList)
}

// GetTaskList godoc
// @Security ApiKeyAuth
// @Router       /task_list/{id} [GET]
// @Summary      Get task list by id
// @Description  get task list by id
// @Tags         task_list
// @Accept       json
// @Produce      json
// @Param        id path string true "task_list_id"
// @Success      201  {object}  models.TaskList
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetTaskList(c *gin.Context)  {
	auth, err := h.GetAuthInfoFromToken(c)
	if err != nil{
		handleResponse(c, "Error Unauthorized!",http.StatusUnauthorized, pkg.ErrUnauthorized.Error())
	}

	id := c.Param("id")

	if id == ""{
		handleResponse(c, pkg.IdErr, http.StatusBadRequest, errors.New(pkg.IdErr).Error())
		return
	}

	taskList, err := h.services.TaskList().Get(context.Background(),models.PrimaryKey{ID: id})
	if err != nil{
		handleResponse(c, "Error in handler, while getting task_list!", http.StatusInternalServerError, err.Error())
		return
	}

	if auth.UserID != taskList.UserID{
		handleResponse(c, "given id is wrong!", http.StatusBadRequest, errors.New("given id is wrong").Error())
		return
	}

	handleResponse(c, "", http.StatusOK, taskList)
}

// UpdateTaskList godoc
// @Security ApiKeyAuth
// @Router       /task_list/{id} [PUT]
// @Summary      Update task list
// @Description  update task list
// @Tags         task_list
// @Accept       json
// @Produce      json
// @Param        id path string true "task_list_id"
// @Param        task_list body models.UpdateTaskList false "task_list"
// @Success      201  {object}  models.UpdateResponse
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) UpdateTaskList(c *gin.Context) {
	auth, err := h.GetAuthInfoFromToken(c)
	if err != nil{
		handleResponse(c, "Error Unauthorized!",http.StatusUnauthorized, pkg.ErrUnauthorized.Error())
	}

	updTaskList := models.UpdateTaskList{}

	id := c.Param("id")
	if id == ""{
		handleResponse(c, pkg.IdErr, http.StatusBadRequest, errors.New(pkg.IdErr).Error())
		return
	}

	err = c.ShouldBindJSON(&updTaskList)
	if err != nil{
		handleResponse(c, "Error in handler, while reading body from client!", http.StatusBadRequest, err.Error())
		return
	}

	getTaskList, err := h.services.TaskList().Get(context.Background(), models.PrimaryKey{ID: id})
	if err != nil{
		handleResponse(c, "Error in handler, while getting task list!", http.StatusBadRequest, err.Error())
		return
	}

	if auth.UserID != getTaskList.UserID{
		handleResponse(c, "Error in handler, id is wrong!", http.StatusBadRequest, errors.New("id is wrong").Error())
		return
	}

	updTaskList.ID = id

	taskList, err := h.services.TaskList().Update(context.Background(), updTaskList)
	if err != nil{
		handleResponse(c, "Error in handler, while updating task_list!", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "Successfully Updated!", http.StatusOK, taskList)
}

// DeleteTaskList godoc
// @Security ApiKeyAuth
// @Router       /task_list/{id} [DELETE]
// @Summary      Delete task list
// @Description  delete task list
// @Tags         task_list
// @Accept       json
// @Produce      json
// @Param        id path string true "task_list_id"
// @Success      201  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) DeleteTaskList(c *gin.Context) {
	auth, err := h.GetAuthInfoFromToken(c)
	if err != nil{
		handleResponse(c, "Error Unauthorized!",http.StatusUnauthorized, pkg.ErrUnauthorized.Error())
	}

	id := c.Param("id")
	if id == ""{
		handleResponse(c, pkg.IdErr, http.StatusBadRequest, errors.New(pkg.IdErr).Error())
		return
	}

	getTaskList, err := h.services.TaskList().Get(context.Background(), models.PrimaryKey{ID: id})
	if err != nil{
		handleResponse(c, "Error in handler, while getting task list!", http.StatusBadRequest, err.Error())
		return
	}

	if auth.UserID != getTaskList.UserID{
		handleResponse(c, "Error in handler, id is wrong!", http.StatusBadRequest, errors.New("id is wrong").Error())
		return
	}

	err = h.services.TaskList().Delete(context.Background(), models.PrimaryKey{ID: id})
	if err != nil{
		handleResponse(c, "Error in handler, while deleting task_list!", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "TaskList succesfully deleted!", http.StatusOK, "TaskList succesfully deleted!")
}