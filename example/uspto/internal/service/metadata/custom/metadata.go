package custom

import (
	metadataRepository "example/uspto/internal/database/repositories/metadata"
	"example/uspto/internal/service/metadata/generated"
)

type ServiceImpl struct {
	*generated.ServiceImpl
}

func NewService(repo metadataRepository.Repository) *ServiceImpl {
	return &ServiceImpl{
		ServiceImpl: generated.NewService(repo),
	}
}
