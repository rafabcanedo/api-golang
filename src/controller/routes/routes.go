package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rafabcanedo/api-golang/src/controller"
)

func InitRoutes(r *gin.RouterGroup) {
	
	r.GET("/users", controller.FindUsers)
	r.GET("/users/:userId", controller.FindUserById)
	r.GET("/users/:userEmail", controller.FindUserByEmail)
	r.POST("/users", controller.CreateUser)
	r.PUT("/users/:userId", controller.UpdateUser)
	r.DELETE("/users/:userId", controller.DeletUser)
}