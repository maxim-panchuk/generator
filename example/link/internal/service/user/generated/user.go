package generated

import (
	"context"
	userRepository "example/link/internal/database/repositories/user"
	"example/link/internal/models/user"
)

type ServiceImpl struct {
	repo userRepository.Repository
}

func NewService(repo userRepository.Repository) *ServiceImpl {
	return &ServiceImpl{
		repo: repo,
	}
}

//
func (s *ServiceImpl) GetUserByName(
	ctx context.Context,
	username string, // Required: true, Description:
) (*user.userDTO, error) {
	panic("not implemented")
}
