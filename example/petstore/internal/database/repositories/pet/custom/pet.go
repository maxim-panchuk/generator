package custom

import (
	"gorm.io/gorm"
	"petstore/internal/database/repositories/pet/generated"
)

type RepositoryImpl struct {
	*generated.RepositoryImpl
}

func NewRepository(db *gorm.DB) *RepositoryImpl {
	return &RepositoryImpl{
		RepositoryImpl: generated.NewRepository(db),
	}
}
