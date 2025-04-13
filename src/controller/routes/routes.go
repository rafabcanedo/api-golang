package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rafabcanedo/api-golang/src/controller"
)

func InitRoutes(r *gin.RouterGroup) {
	
	r.GET("/users", controller.FindUsers)
	r.GET("/getUserById/:userId", controller.FindUserById)
	r.GET("/getUserByEmail/:userEmail", controller.FindUserByEmail)
	r.POST("/users", controller.CreateUser)
	r.PUT("/users/:userId", controller.UpdateUser)
	r.DELETE("/users/:userId", controller.DeletUser)
}