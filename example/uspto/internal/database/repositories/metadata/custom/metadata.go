package custom

import (
	"example/uspto/internal/database/repositories/metadata/generated"
	"gorm.io/gorm"
)

type RepositoryImpl struct {
	*generated.RepositoryImpl
}

func NewRepository(db *gorm.DB) *RepositoryImpl {
	return &RepositoryImpl{
		RepositoryImpl: generated.NewRepository(db),
	}
}
