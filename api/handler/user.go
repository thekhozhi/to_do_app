package handler

import (
	"context"
	"errors"
	"net/http"
	"to_do_app/api/models"
	"to_do_app/pkg"

	"github.com/gin-gonic/gin"
)

// CreateUser godoc
// @Router       /user [POST]
// @Summary      Creates a new user
// @Description  create a new user
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        user body models.CreateUser false "user"
// @Success      201  {object}  models.User
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) CreateUser(c *gin.Context) {
	createUser := models.CreateUser{}

	err := c.ShouldBindJSON(&createUser)
	if err != nil{
		handleResponse(c,"Error in handler, while reading body from client!",http.StatusBadRequest, err.Error())
		return
	}
	
	user, err := h.services.User().Create(context.Background(),createUser)
	if err != nil{
		handleResponse(c, "Error in handler, while crating user!",http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "User created!", http.StatusCreated, user)
}

// GetUser godoc
// @Security ApiKeyAuth
// @Router       /user/{id} [GET]
// @Summary      Get user by id
// @Description  get user by id
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        id path string true "user_id"
// @Success      201  {object}  models.User
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetUser(c *gin.Context)  {
	auth, err := h.GetAuthInfoFromToken(c)
	if err != nil{
		handleResponse(c, "Error Unauthorized!",http.StatusUnauthorized, pkg.ErrUnauthorized.Error())
	}

	id := c.Param("id")

	if id == ""{
		handleResponse(c, pkg.IdErr, http.StatusBadRequest, errors.New(pkg.IdErr).Error())
		return
	}

	if auth.UserID != id {
		handleResponse(c, "given id is wrong!", http.StatusBadRequest, errors.New("given id is wrong").Error())
		return
	}

	user, err := h.services.User().Get(context.Background(),models.PrimaryKey{ID: id})
	if err != nil{
		handleResponse(c, "Error in handler, while getting user!", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "", http.StatusOK, user)
}

// UpdateUser godoc
// @Security ApiKeyAuth
// @Router       /user/{id} [PUT]
// @Summary      Update user
// @Description  update user
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        id path string true "user_id"
// @Param        user body models.UpdateUser false "user"
// @Success      201  {object}  models.UpdateResponse
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) UpdateUser(c *gin.Context) {
	auth, err := h.GetAuthInfoFromToken(c)
	if err != nil{
		handleResponse(c, "Error Unauthorized!",http.StatusUnauthorized, pkg.ErrUnauthorized.Error())
	}

	updUser := models.UpdateUser{}

	id := c.Param("id")
	if id == ""{
		handleResponse(c, pkg.IdErr, http.StatusBadRequest, errors.New(pkg.IdErr).Error())
		return
	}

	if auth.UserID != id {
		handleResponse(c, "given id is wrong!", http.StatusBadRequest, errors.New("given id is wrong").Error())
		return
	}

	err = c.ShouldBindJSON(&updUser)
	if err != nil{
		handleResponse(c, "Error in handler, while reading body from client!", http.StatusBadRequest, err.Error())
		return
	}

	updUser.ID = id

	user, err := h.services.User().Update(context.Background(), updUser)
	if err != nil{
		handleResponse(c, "Error in handler, while updating user!", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "Successfully Updated!", http.StatusOK, user)
}

// DeleteUser godoc
// @Security ApiKeyAuth
// @Router       /user/{id} [DELETE]
// @Summary      Delete user
// @Description  delete user
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        id path string true "user_id"
// @Success      201  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) DeleteUser(c *gin.Context) {
	auth, err := h.GetAuthInfoFromToken(c)
	if err != nil{
		handleResponse(c, "Error Unauthorized!",http.StatusUnauthorized, pkg.ErrUnauthorized.Error())
	}

	id := c.Param("id")
	if id == ""{
		handleResponse(c, pkg.IdErr, http.StatusBadRequest, errors.New(pkg.IdErr).Error())
		return
	}

	if auth.UserID != id {
		handleResponse(c, "given id is wrong!", http.StatusBadRequest, errors.New("given id is wrong").Error())
		return
	}

	err = h.services.User().Delete(context.Background(), models.PrimaryKey{ID: id})
	if err != nil{
		handleResponse(c, "Error in handler, while deleting user!", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "User succesfully deleted!", http.StatusOK, "User succesfully deleted!")
}

// ChangeUserPassword godoc
// @Security ApiKeyAuth
// @Router       /user/{id} [PATCH]
// @Summary      Change password
// @Description  Change password
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        id path string true "user_id"
// @Param        user body models.UpdateUserPassword false "user"
// @Success      200  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) ChangeUserPassword(c *gin.Context)  {
	auth, err := h.GetAuthInfoFromToken(c)
	if err != nil{
		handleResponse(c, "Error Unauthorized!",http.StatusUnauthorized, pkg.ErrUnauthorized.Error())
	}

	id := c.Param("id")
	if id == ""{
		handleResponse(c, pkg.IdErr, http.StatusBadRequest, errors.New(pkg.IdErr).Error())
		return
	}

	if auth.UserID != id {
		handleResponse(c, "given id is wrong!", http.StatusBadRequest, errors.New("given id is wrong").Error())
		return
	}

	updUserPassword := models.UpdateUserPassword{}

	err = c.ShouldBindJSON(&updUserPassword)
	if err != nil{
		handleResponse(c, "Error while reading boy from client!", http.StatusBadRequest, err.Error())
		return
	}

	updUserPassword.ID = id

	err = h.services.User().ChangePassword(context.Background(), updUserPassword)
	if err != nil{
		handleResponse(c, "Error in handler, while changing password!", http.StatusInternalServerError, errors.New("given old password is wrong").Error())
		return
	}

	handleResponse(c, "Your password has been changed!", http.StatusOK, "Your password has been changed!")
}