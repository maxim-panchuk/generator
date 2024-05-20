package repo

import (
	repoService "example/link/internal/service/repo"
	"example/link/internal/transport/http/in/repo/custom"
	"net/http"
)

type Controller interface {
	//
	GetRepositoriesByOwner(
		w http.ResponseWriter, r *http.Request)
	//
	GetRepository(
		w http.ResponseWriter, r *http.Request)
	//
	GetPullRequestsByRepository(
		w http.ResponseWriter, r *http.Request)
	//
	GetPullRequestsById(
		w http.ResponseWriter, r *http.Request)
	//
	MergePullRequest(
		w http.ResponseWriter, r *http.Request)
}

func NewController(service repoService.Service) *custom.ControllerImpl {
	return custom.NewController(service)
}
