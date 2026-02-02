package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/opaulosaa/crud-go/src/controller"
)

func InitRoutes(r *gin.RouterGroup) {
	r.GET("/getUserById/:userId", controller.FindUserByID)
	r.GET("/getUserByEmai/:userEmail", controller.FindUserByEmail)
	r.POST("/createUser/", controller.CreateUser)
	r.PUT("/updateUser/:userId", controller.UpdateUser)
	r.DELETE("/deleteUser/:userID", controller.DeleteUser)
}
