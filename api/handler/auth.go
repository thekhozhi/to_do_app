package handler

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"to_do_app/api/models"
	"to_do_app/pkg"
	"to_do_app/pkg/jwt"

	"github.com/gin-gonic/gin"
 	"github.com/spf13/cast"
)

// UserLogin godoc
// @Router       /user_login [POST]
// @Summary      User Login
// @Description  User Login
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        login body models.UserLoginRequest false "login"
// @Success      201  {object}  models.UserLoginResponse
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) UserLogin(c *gin.Context)  {
	loginReq := models.UserLoginRequest{}
	loginResp := models.UserLoginResponse{}

	err := c.ShouldBindJSON(&loginReq)
	if err != nil{
		handleResponse(c, "Error in handler, while reading body from client!", http.StatusBadRequest, err.Error())
		return
	}

	userID, err := h.services.Auth().UserLogin(context.Background(), loginReq)
	if err != nil{
		handleResponse(c,"Given email or password is wrong!", http.StatusBadRequest, errors.New("given email or password is wrong").Error())
		return
	}

	claims := map[string]interface{}{
		"user_id": userID,
	}

	accessToken, err := jwt.GenerateJWT(claims, pkg.AccessExpireTime, string(pkg.SignKey))
	if err != nil{
		handleResponse(c, "Error in handler, while generating access token JWT!", http.StatusInternalServerError, err.Error())
		return
	}

	refreshToken, err := jwt.GenerateJWT(claims, pkg.RefreshExpireTime, string(pkg.SignKey))
	if err != nil{
		handleResponse(c, "Error in handler, while generating refresh token JWT!", http.StatusInternalServerError, err.Error())
		return
	}

	loginResp.AccessToken = accessToken
	loginResp.RefreshToken = refreshToken

	handleResponse(c, "Tokens created!", http.StatusCreated, loginResp)
}

func (h Handler) GetAuthInfoFromToken(c *gin.Context) (models.AuthInfo, error) {
	accessToken := c.GetHeader("Authorization")
	if accessToken == ""{
		return models.AuthInfo{}, pkg.ErrUnauthorized
	}

	claims, err := jwt.ExtractClaims(accessToken, string(pkg.SignKey))
	if err != nil{
		fmt.Println("Error in handlers, while extracting claims (user_id from token)!", err.Error())
		return models.AuthInfo{}, err
	}

	userID := cast.ToString(claims["user_id"])

	return models.AuthInfo{
		UserID: userID,
	}, nil
}