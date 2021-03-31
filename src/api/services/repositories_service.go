package services

import (
	"net/http"
	"sync"

	"github.com/miguelhun/go-microservices/src/api/config"
	"github.com/miguelhun/go-microservices/src/api/domain/github"
	"github.com/miguelhun/go-microservices/src/api/domain/repositories"
	"github.com/miguelhun/go-microservices/src/api/log"
	"github.com/miguelhun/go-microservices/src/api/providers/github_provider"
	"github.com/miguelhun/go-microservices/src/api/utils/errors"
)

type repoService struct{}

type reposServiceInterface interface {
	CreateRepo(repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError)
	CreateRepos([]repositories.CreateRepoRequest) (repositories.CreateReposResponse, errors.ApiError)
}

var (
	RepositoryService reposServiceInterface
)

func init() {
	RepositoryService = &repoService{}
}

func (rs *repoService) CreateRepo(input repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError) {
	if err := input.Validate(); err != nil {
		return nil, err
	}

	request := github.CreateRepoRequest{
		Name:        input.Name,
		Description: input.Description,
		Private:     false,
	}

	response, err := github_provider.CreateRepo(config.GetGithubAccessToken(), request)
	if err != nil {
		log.Error("response from external api", err, "status:error")
		return nil, errors.NewApiError(err.StatusCode, err.Message)
	}

	result := repositories.CreateRepoResponse{
		Id:    response.Id,
		Owner: response.Owner.Login,
		Name:  response.Name,
	}

	return &result, nil
}

func (rs *repoService) CreateRepos(request []repositories.CreateRepoRequest) (repositories.CreateReposResponse, errors.ApiError) {
	input := make(chan repositories.CreateRepositoriesResult)
	output := make(chan repositories.CreateReposResponse)
	defer close(output)

	var wg sync.WaitGroup
	go rs.handleRepoResults(&wg, input, output)

	for _, current := range request {
		wg.Add(1)
		go rs.createRepoConcurrent(current, input)
	}

	wg.Wait()
	close(input)

	result := <-output

	successfulResponses := 0
	for _, current := range result.Results {
		if current.Response != nil {
			successfulResponses++
		}
	}

	if successfulResponses == 0 {
		result.StatusCode = result.Results[0].Error.Status()
	} else if successfulResponses == len(result.Results) {
		result.StatusCode = http.StatusCreated
	} else {
		result.StatusCode = http.StatusPartialContent
	}

	return result, nil
}

func (rs *repoService) handleRepoResults(wg *sync.WaitGroup, input chan repositories.CreateRepositoriesResult, output chan repositories.CreateReposResponse) {
	var results repositories.CreateReposResponse

	for incomingRequest := range input {
		repoResult := repositories.CreateRepositoriesResult{
			Response: incomingRequest.Response,
			Error:    incomingRequest.Error,
		}
		results.Results = append(results.Results, repoResult)
		wg.Done()
	}
	output <- results
}

func (rs *repoService) createRepoConcurrent(input repositories.CreateRepoRequest, output chan repositories.CreateRepositoriesResult) {
	if err := input.Validate(); err != nil {
		output <- repositories.CreateRepositoriesResult{Error: err}
		return
	}

	result, err := rs.CreateRepo(input)
	if err != nil {
		output <- repositories.CreateRepositoriesResult{Error: err}
		return
	}
	output <- repositories.CreateRepositoriesResult{Response: result}

}
