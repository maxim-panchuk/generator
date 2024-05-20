package custom

import (
	userRepository "petstore/internal/database/repositories/user"
	"petstore/internal/service/user/generated"
)

type ServiceImpl struct {
	*generated.ServiceImpl
}

func NewService(repo userRepository.Repository) *ServiceImpl {
	return &ServiceImpl{
		ServiceImpl: generated.NewService(repo),
	}
}
