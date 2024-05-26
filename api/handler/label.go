package handler

import (
	"context"
	"errors"
	"net/http"
	"to_do_app/api/models"
	"to_do_app/pkg"

	"github.com/gin-gonic/gin"
)

// CreateLabel godoc
// @Security ApiKeyAuth
// @Router       /label [POST]
// @Summary      Creates a new label
// @Description  create a new label
// @Tags         label
// @Accept       json
// @Produce      json
// @Param        label body models.CreateLabel false "label"
// @Success      201  {object}  models.Label
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) CreateLabel(c *gin.Context) {
	auth, err := h.GetAuthInfoFromToken(c)
	if err != nil{
		handleResponse(c, "Error Unauthorized!",http.StatusUnauthorized, pkg.ErrUnauthorized.Error())
	}

	createLabel := models.CreateLabel{}

	err = c.ShouldBindJSON(&createLabel)
	if err != nil{
		handleResponse(c,"Error in handler, while reading body from client!",http.StatusBadRequest, err.Error())
		return
	}

	if auth.UserID != createLabel.UserID{
		handleResponse(c, "Error in handler, user id is wrong!", http.StatusBadRequest, errors.New("user id is wrong").Error())
		return
	}
	
	label, err := h.services.Label().Create(context.Background(),createLabel)
	if err != nil{
		handleResponse(c, "Error in handler, while crating label!",http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "Label created!", http.StatusCreated, label)
}

// GetLabel godoc
// @Security ApiKeyAuth
// @Router       /label/{id} [GET]
// @Summary      Get label by id
// @Description  get label by id
// @Tags         label
// @Accept       json
// @Produce      json
// @Param        id path string true "label_id"
// @Success      201  {object}  models.Label
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) GetLabel(c *gin.Context)  {
	auth, err := h.GetAuthInfoFromToken(c)
	if err != nil{
		handleResponse(c, "Error Unauthorized!",http.StatusUnauthorized, pkg.ErrUnauthorized.Error())
	}

	id := c.Param("id")

	if id == ""{
		handleResponse(c, pkg.IdErr, http.StatusBadRequest, errors.New(pkg.IdErr).Error())
		return
	}

	label, err := h.services.Label().Get(context.Background(),models.PrimaryKey{ID: id})
	if err != nil{
		handleResponse(c, "Error in handler, while getting label!", http.StatusInternalServerError, err.Error())
		return
	}

	if auth.UserID != label.UserID{
		handleResponse(c, "Error in handler, id is wrong!", http.StatusBadRequest, errors.New(" id is wrong").Error())
		return
	}

	handleResponse(c, "", http.StatusOK, label)
}

// UpdateLabel godoc
// @Security ApiKeyAuth
// @Router       /label/{id} [PUT]
// @Summary      Update label
// @Description  update label
// @Tags         label
// @Accept       json
// @Produce      json
// @Param        id path string true "label_id"
// @Param        label body models.UpdateLabel false "label"
// @Success      201  {object}  models.UpdateResponse
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) UpdateLabel(c *gin.Context) {
	auth, err := h.GetAuthInfoFromToken(c)
	if err != nil{
		handleResponse(c, "Error Unauthorized!",http.StatusUnauthorized, pkg.ErrUnauthorized.Error())
	}

	updLabel := models.UpdateLabel{}

	id := c.Param("id")
	if id == ""{
		handleResponse(c, pkg.IdErr, http.StatusBadRequest, errors.New(pkg.IdErr).Error())
		return
	}

	getLabel, err := h.services.Label().Get(context.Background(), models.PrimaryKey{ID: id})
	if err != nil{
		handleResponse(c, "Error in handler, while getting label while updating!", http.StatusInternalServerError,err.Error())
		return
	}

	if auth.UserID != getLabel.UserID{
		handleResponse(c, "Error in handler, label id is wrong!", http.StatusBadRequest,errors.New("label id is wrong").Error())
		return
	}

	err = c.ShouldBindJSON(&updLabel)
	if err != nil{
		handleResponse(c, "Error in handler, while reading body from client!", http.StatusBadRequest, err.Error())
		return
	}

	updLabel.ID = id

	label, err := h.services.Label().Update(context.Background(), updLabel)
	if err != nil{
		handleResponse(c, "Error in handler, while updating label!", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "Successfully Updated!", http.StatusOK, label)
}

// DeleteLabel godoc
// @Security ApiKeyAuth
// @Router       /label/{id} [DELETE]
// @Summary      Delete label
// @Description  delete label
// @Tags         label
// @Accept       json
// @Produce      json
// @Param        id path string true "label_id"
// @Success      201  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) DeleteLabel(c *gin.Context) {
	auth, err := h.GetAuthInfoFromToken(c)
	if err != nil{
		handleResponse(c, "Error Unauthorized!",http.StatusUnauthorized, pkg.ErrUnauthorized.Error())
	}

	id := c.Param("id")
	if id == ""{
		handleResponse(c, pkg.IdErr, http.StatusBadRequest, errors.New(pkg.IdErr).Error())
		return
	}

	getLabel, err := h.services.Label().Get(context.Background(), models.PrimaryKey{ID: id})
	if err != nil{
		handleResponse(c, "Error in handler, while getting label while updating!", http.StatusInternalServerError,err.Error())
		return
	}

	if auth.UserID != getLabel.UserID{
		handleResponse(c, "Error in handler, label id is wrong!", http.StatusBadRequest,errors.New("label id is wrong").Error())
		return
	}

	err = h.services.Label().Delete(context.Background(), models.PrimaryKey{ID: id})
	if err != nil{
		handleResponse(c, "Error in handler, while deleting label!", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "Label succesfully deleted!", http.StatusOK, "Label succesfully deleted!")
}