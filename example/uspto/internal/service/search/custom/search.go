package custom

import (
	searchRepository "example/uspto/internal/database/repositories/search"
	"example/uspto/internal/service/search/generated"
)

type ServiceImpl struct {
	*generated.ServiceImpl
}

func NewService(repo searchRepository.Repository) *ServiceImpl {
	return &ServiceImpl{
		ServiceImpl: generated.NewService(repo),
	}
}
