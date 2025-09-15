package controller

import (
	"github.com/gin-gonic/gin"

	rest_err "github.com/gustavoaurelio/go-crud/src/configuration"
)

func CreateUser(c *gin.Context) {
	err := rest_err.NewBadRequestError("Rota errada")
	c.JSON(err.Code, err)
}
