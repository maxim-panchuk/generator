package custom

import (
	"gorm.io/gorm"
	"petstore/internal/database/repositories/user/generated"
)

type RepositoryImpl struct {
	*generated.RepositoryImpl
}

func NewRepository(db *gorm.DB) *RepositoryImpl {
	return &RepositoryImpl{
		RepositoryImpl: generated.NewRepository(db),
	}
}
