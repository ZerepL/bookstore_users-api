package users

import (
	"net/http"
	"strconv"

	"github.com/ZerepL/bookstore_oauth-go/oauth"
	"github.com/ZerepL/bookstore_users-api/domain/users"
	"github.com/ZerepL/bookstore_users-api/services"
	internalErrors "github.com/ZerepL/bookstore_utils/internal_errors"
	"github.com/gin-gonic/gin"
)

func getUserId(userIdParam string) (int64, internalErrors.RestErr) {
	userId, userErr := strconv.ParseInt(userIdParam, 10, 64)
	if userErr != nil {
		return 0, internalErrors.NewBadRequestError("user id should be a number")
	}

	return userId, nil
}

// ShowAccount godoc
// @Summary      Create user
// @Description  create a new user
// @Tags         users
// @Produce      json
// @Success      200  {object} users.User
// @Failure      400  {object}  internalErrors.RestErr
// @Failure      500  {object}  internalErrors.RestErr
// @Router       /users [post]
func Create(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := internalErrors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status(), restErr)
		return
	}

	result, saveErr := services.UserService.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status(), saveErr)
		return
	}
	c.JSON(http.StatusCreated, result.Marshall(c.GetHeader("X-Public") == "true"))
}

// ShowAccount godoc
// @Summary      Get user
// @Description  get user info
// @Tags         users
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Header       all  {string}  X-Public    "true"
// @Success      200  {object}  users.User
// @Failure      400  {object}  internalErrors.RestErr
// @Failure      404  {object}  internalErrors.RestErr
// @Failure      500  {object}  internalErrors.RestErr
// @Router       /users/{id} [get]
func Get(c *gin.Context) {
	if err := oauth.AuthenticateRequest(c.Request); err != nil {
		c.JSON(err.Status(), err)
		return
	}

	userId, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status(), idErr)
		return
	}
	user, getErr := services.UserService.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status(), getErr)
		return
	}

	if oauth.GetCallerId(c.Request) == user.Id {
		c.JSON(http.StatusOK, user.Marshall(false))
		return
	}
	c.JSON(http.StatusOK, user.Marshall(oauth.IsPublic(c.Request)))
}

// ShowAccount godoc
// @Summary      Update user
// @Description  update user info
// @Tags         users
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Header       all  {string}  X-Public    "true"
// @Success      200  {object}  users.User
// @Failure      400  {object}  internalErrors.RestErr
// @Failure      404  {object}  internalErrors.RestErr
// @Failure      500  {object}  internalErrors.RestErr
// @Router       /users/{id} [put]
// @Router       /users/{id} [patch]
func Update(c *gin.Context) {
	userId, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status(), idErr)
		return
	}

	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := internalErrors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status(), restErr)
		return
	}

	user.Id = userId

	isPartial := c.Request.Method == http.MethodPatch

	result, err := services.UserService.UpdateUser(isPartial, user)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, result.Marshall(c.GetHeader("X-Public") == "true"))
}

// ShowAccount godoc
// @Summary      Delete user
// @Description  delete user from db
// @Tags         users
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200  {string}  string "status:deleted"
// @Failure      400  {object}  internalErrors.RestErr
// @Failure      404  {object}  internalErrors.RestErr
// @Failure      500  {object}  internalErrors.RestErr
// @Router       /users/{id} [delete]
func Delete(c *gin.Context) {
	userId, idErr := getUserId(c.Param("user_id"))
	if idErr != nil {
		c.JSON(idErr.Status(), idErr)
		return
	}

	if err := services.UserService.DeleteUser(userId); err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})

}

// ShowAccount godoc
// @Summary      Search user by status
// @Description  search a user based on status
// @Tags         users
// @Produce      json
// @Header       all  {string}  X-Public    "true"
// @Param        enumstring  query string false  "status" Enums(active)
// @Success      200  {array}   users.User
// @Failure      400  {object}  internalErrors.RestErr
// @Failure      404  {object}  internalErrors.RestErr
// @Failure      500  {object}  internalErrors.RestErr
// @Router       /internal/users/search [get]
func Search(c *gin.Context) {
	status := c.Query("status")

	users, err := services.UserService.SearchUser(status)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, users.Marshall(c.GetHeader("X-Public") == "true"))
}

func Login(c *gin.Context) {
	var request users.LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		restErr := internalErrors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status(), restErr)
		return
	}
	user, err := services.UserService.LoginUser(request)
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, user.Marshall(c.GetHeader("X-Public") == "true"))
}
