package custom

import (
	petRepository "petstore/internal/database/repositories/pet"
	"petstore/internal/service/pet/generated"
)

type ServiceImpl struct {
	*generated.ServiceImpl
}

func NewService(repo petRepository.Repository) *ServiceImpl {
	return &ServiceImpl{
		ServiceImpl: generated.NewService(repo),
	}
}
