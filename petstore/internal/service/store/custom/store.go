package custom

import (
	storeRepository "petstore/internal/database/repositories/store"
	"petstore/internal/service/store/generated"
)

type ServiceImpl struct {
	*generated.ServiceImpl
}

func NewService(repo storeRepository.Repository) *ServiceImpl {
	return &ServiceImpl{
		ServiceImpl: generated.NewService(repo),
	}
}
