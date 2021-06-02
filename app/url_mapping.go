package app

import (
	"github.com/frech87/bookstore_users-api/controllers/ping"
	"github.com/frech87/bookstore_users-api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.GET("/users/:user_id", users.GetUser)
	router.POST("/users", users.CreateUser)
	//router.GET("/users/search", controllers.SearchUser)
}
