package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"

	rest_err "github.com/gustavoaurelio/go-crud/src/configuration"
	"github.com/gustavoaurelio/go-crud/src/controller/model/request"
)

func CreateUser(c *gin.Context) {

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := rest_err.NewBadRequestError(
			fmt.Sprintf("There are some incorrect fields!, error=%s \n", err.Error))

		c.JSON(restErr.Code, restErr)
		return
	}

	fmt.Println(userRequest)
}
