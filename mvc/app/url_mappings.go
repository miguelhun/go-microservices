package app

import "github.com/miguelhun/go-microservices/mvc/controllers"

func mapUrls() {
	router.GET("/users/:user_id", controllers.GetUser)
}
