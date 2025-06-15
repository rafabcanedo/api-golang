package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rafabcanedo/api-golang/src/controller"
)

func InitRoutes(
	r *gin.RouterGroup,
	userController controller.UserControllerInterface) {
	
	r.GET("/users", userController.FindAllUsers)
	r.GET("/getUserById/:userId", userController.FindUserById)
	r.GET("/getUserByEmail/:userEmail", userController.FindUserByEmail)
	r.POST("/users", userController.CreateUser)
	r.PUT("/users/:userId", userController.UpdateUser)
	r.DELETE("/users/:userId", userController.DeleteUser)
}