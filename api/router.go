package api

import (
	"to_do_app/api/handler"
	"to_do_app/service"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "to_do_app/api/docs"
)

// New ...
// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func New(services service.IServiceManager) *gin.Engine{
	h := handler.New(services)
	
	r := gin.New()
	
	{

		// Users
		r.POST("/user", h.CreateUser)
		r.GET("/user/:id", h.GetUser)
		r.PUT("/user/:id", h.UpdateUser)
		r.DELETE("/user/:id", h.DeleteUser)
		r.PATCH("/user/:id", h.ChangeUserPassword)

		// Tasks
		r.POST("/task", h.CreateTask)
		r.GET("/task/:id", h.GetTask)
		r.PUT("/task/:id", h.UpdateTask)
		r.DELETE("/task/:id", h.DeleteTask)

		// Task Lists
		r.POST("/task_list", h.CreateTaskList)
		r.GET("/task_list/:id", h.GetTaskList)
		r.PUT("/task_list/:id", h.UpdateTaskList)
		r.DELETE("/task_list/:id", h.DeleteTaskList)

		// Labels
		r.POST("/label", h.CreateLabel)
		r.GET("/label/:id", h.GetLabel)
		r.PUT("/label/:id", h.UpdateLabel)
		r.DELETE("/label/:id", h.DeleteLabel)

		// Auth
		r.POST("/user_login", h.UserLogin)

		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
	
	return r
}