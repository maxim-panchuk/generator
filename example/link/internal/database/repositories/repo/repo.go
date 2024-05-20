package repo

import (
	"context"
	"example/link/internal/database/repositories/repo/custom"
	"example/link/internal/models/pullrequest"
	"example/link/internal/models/repository"
	"gorm.io/gorm"
)

type Repository interface {
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

func NewRepository(db *gorm.DB) *custom.RepositoryImpl {
	return custom.NewRepository(db)
}
