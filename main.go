// @title           Bookstore Users API
// @version         1.0
// @description     API to handle and store differents users.
// @termsOfService  http://swagger.io/terms/

// @host      localhost:8080
// @BasePath  /users

// @securityDefinitions.basic  BasicAuth
package main

import (
	"github.com/ZerepL/bookstore_users-api/app"
)

func main() {
	app.StartApplication()
}
