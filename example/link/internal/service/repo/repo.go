package repo

import (
	"context"
	repoRepository "example/link/internal/database/repositories/repo"
	"example/link/internal/models/pullrequest"
	"example/link/internal/models/repository"
	"example/link/internal/service/repo/custom"
)

type Service interface {
	//
	GetRepositoriesByOwner(
		ctx context.Context,
		username string, // Required: true, Description:
	) ([]*repository.repositoryDTO, error)
	//
	GetRepository(
		ctx context.Context,
		username string, // Required: true, Description:
		slug string, // Required: true, Description:
	) (*repository.repositoryDTO, error)
	//
	GetPullRequestsByRepository(
		ctx context.Context,
		username string, // Required: true, Description:
		slug string, // Required: true, Description:
		state string, // Required: true, Description:
	) ([]*pullrequest.pullrequestDTO, error)
	//
	GetPullRequestsById(
		ctx context.Context,
		username string, // Required: true, Description:
		slug string, // Required: true, Description:
		pid string, // Required: true, Description:
	) (*pullrequest.pullrequestDTO, error)
	//
	MergePullRequest(
		ctx context.Context,
		username string, // Required: true, Description:
		slug string, // Required: true, Description:
		pid string, // Required: true, Description:
	) error
}

func NewService(repo repoRepository.Repository) *custom.ServiceImpl {
	return custom.NewService(repo)
}
