package generated

import (
	"context"
	repoRepository "example/link/internal/database/repositories/repo"
	"example/link/internal/models/pullrequest"
	"example/link/internal/models/repository"
)

type ServiceImpl struct {
	repo repoRepository.Repository
}

func NewService(repo repoRepository.Repository) *ServiceImpl {
	return &ServiceImpl{
		repo: repo,
	}
}

//
func (s *ServiceImpl) GetRepositoriesByOwner(
	ctx context.Context,
	username string, // Required: true, Description:
) ([]*repository.repositoryDTO, error) {
	panic("not implemented")
}

//
func (s *ServiceImpl) GetRepository(
	ctx context.Context,
	username string, // Required: true, Description:
	slug string, // Required: true, Description:
) (*repository.repositoryDTO, error) {
	panic("not implemented")
}

//
func (s *ServiceImpl) GetPullRequestsByRepository(
	ctx context.Context,
	username string, // Required: true, Description:
	slug string, // Required: true, Description:
	state string, // Required: true, Description:
) ([]*pullrequest.pullrequestDTO, error) {
	panic("not implemented")
}

//
func (s *ServiceImpl) GetPullRequestsById(
	ctx context.Context,
	username string, // Required: true, Description:
	slug string, // Required: true, Description:
	pid string, // Required: true, Description:
) (*pullrequest.pullrequestDTO, error) {
	panic("not implemented")
}

//
func (s *ServiceImpl) MergePullRequest(
	ctx context.Context,
	username string, // Required: true, Description:
	slug string, // Required: true, Description:
	pid string, // Required: true, Description:
) error {
	panic("not implemented")
}
