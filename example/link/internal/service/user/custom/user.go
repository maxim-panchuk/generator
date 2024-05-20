package custom

import (
	userRepository "example/link/internal/database/repositories/user"
	"example/link/internal/service/user/generated"
)

type ServiceImpl struct {
	*generated.ServiceImpl
}

func NewService(repo userRepository.Repository) *ServiceImpl {
	return &ServiceImpl{
		ServiceImpl: generated.NewService(repo),
	}
}
