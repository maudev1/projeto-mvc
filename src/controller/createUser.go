package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/maudev1/meu-primeiro-crud/src/configuration/rest_err"
	"github.com/maudev1/meu-primeiro-crud/src/controller/model/request"
)

func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	if err := c.ShouldBindBodyWithJSON(&userRequest); err != nil {
		restErr := rest_err.NewBadRequestError(fmt.Sprintf("There are some incorrect field, erro=%s\n", err.Error()))

		c.JSON(restErr.Code, restErr)

		return
	}

	fmt.Println(userRequest)

	//aula 6

}
