package users

import (
	"fmt"
	"net/http"

	"github.com/ZerepL/bookstore_users-api/domain/users"
	"github.com/ZerepL/bookstore_users-api/services"
	internalErrors "github.com/ZerepL/bookstore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user users.User
	fmt.Println(user)

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := internalErrors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, restErr := services.CreateUser(user)
	if restErr != nil {
		c.JSON(restErr.Status, restErr)
		return
	}

	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}

func FindUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}
