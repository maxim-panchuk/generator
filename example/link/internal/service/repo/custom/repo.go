package custom

import (
	repoRepository "example/link/internal/database/repositories/repo"
	"example/link/internal/service/repo/generated"
)

type ServiceImpl struct {
	*generated.ServiceImpl
}

func NewService(repo repoRepository.Repository) *ServiceImpl {
	return &ServiceImpl{
		ServiceImpl: generated.NewService(repo),
	}
}
