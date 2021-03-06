package app

import (
	"net/http"

	"github.com/miguelhun/go-microservices/mvc/controllers"
)

func StartApplication() {
	http.HandleFunc("/users", controllers.GetUser)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
