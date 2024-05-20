package user

import (
	"context"
	userRepository "example/link/internal/database/repositories/user"
	"example/link/internal/models/user"
	"example/link/internal/service/user/custom"
)

type Service interface {
	//
	GetUserByName(
		ctx context.Context,
		username string, // Required: true, Description:
	) (*user.userDTO, error)
}

func NewService(repo userRepository.Repository) *custom.ServiceImpl {
	return custom.NewService(repo)
}
