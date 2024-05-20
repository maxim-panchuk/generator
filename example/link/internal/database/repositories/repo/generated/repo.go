package generated

import (
	"context"
	"example/link/internal/models/pullrequest"
	"example/link/internal/models/repository"
	"gorm.io/gorm"
)

type RepositoryImpl struct {
	Db *gorm.DB
}

func NewRepository(db *gorm.DB) *RepositoryImpl {
	return &RepositoryImpl{db}
}

//
func (r *RepositoryImpl) GetRepositoriesByOwner(
	ctx context.Context,
	username string, // Required: true, Description:
) ([]*repository.repositoryDTO, error) {

	panic("not implemented")

}

//
func (r *RepositoryImpl) GetRepository(
	ctx context.Context,
	username string, // Required: true, Description:
	slug string, // Required: true, Description:
) (*repository.repositoryDTO, error) {

	panic("not implemented")

}

//
func (r *RepositoryImpl) GetPullRequestsByRepository(
	ctx context.Context,
	username string, // Required: true, Description:
	slug string, // Required: true, Description:
	state string, // Required: true, Description:
) ([]*pullrequest.pullrequestDTO, error) {

	panic("not implemented")

}

//
func (r *RepositoryImpl) GetPullRequestsById(
	ctx context.Context,
	username string, // Required: true, Description:
	slug string, // Required: true, Description:
	pid string, // Required: true, Description:
) (*pullrequest.pullrequestDTO, error) {

	panic("not implemented")

}

//
func (r *RepositoryImpl) MergePullRequest(
	ctx context.Context,
	username string, // Required: true, Description:
	slug string, // Required: true, Description:
	pid string, // Required: true, Description:
) error {

	panic("not implemented")

}
