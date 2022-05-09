package app

import (
	"github.com/ZerepL/bookstore_users-api/controllers/ping"
	"github.com/ZerepL/bookstore_users-api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.GET("/users/:userd_id", users.GetUser)
	router.POST("/users", users.CreateUser)
}
