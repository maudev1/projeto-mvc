package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/maudev1/meu-primeiro-crud/src/controller"
)

func InitRoutes(r *gin.RouterGroup) {

	r.GET("/findUser/:userId", controller.FindUserById)
	r.POST("/createUser", controller.CreateUser)
	r.DELETE("/deleteUser/:userId", controller.DeleteUser)
	r.PUT("/updateUser/:userId", controller.DeleteUser)

}
