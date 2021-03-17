package app

import "github.com/miguelhun/go-microservices/src/api/controllers/repositories"

func mapUrls() {
	router.POST("/repositories", repositories.CreateRepo)
}
