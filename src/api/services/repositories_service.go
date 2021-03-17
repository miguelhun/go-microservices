package services

import (
	"strings"

	"github.com/miguelhun/go-microservices/src/api/config"
	"github.com/miguelhun/go-microservices/src/api/domain/github"
	"github.com/miguelhun/go-microservices/src/api/domain/repositories"
	"github.com/miguelhun/go-microservices/src/api/providers/github_provider"
	"github.com/miguelhun/go-microservices/src/api/utils/errors"
)

type repoService struct{}

type reposServiceInterface interface {
	CreateRepo(repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError)
}

var (
	RepositoryService reposServiceInterface
)

func init() {
	RepositoryService = &repoService{}
}

func (rs *repoService) CreateRepo(input repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError) {
	input.Name = strings.TrimSpace(input.Name)
	if input.Name == "" {
		return nil, errors.NewBadRequestError("invalid repo name")
	}

	request := github.CreateRepoRequest{
		Name:        input.Name,
		Description: input.Description,
		Private:     false,
	}

	response, err := github_provider.CreateRepo(config.GetGithubAccessToken(), request)
	if err != nil {
		return nil, errors.NewApiError(err.StatusCode, err.Message)
	}

	result := repositories.CreateRepoResponse{
		Id:    response.Id,
		Owner: response.Owner.Login,
		Name:  response.Name,
	}

	return &result, nil
}
