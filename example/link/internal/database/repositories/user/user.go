package user

import (
	"context"
	"example/link/internal/database/repositories/user/custom"
	"example/link/internal/models/user"
	"gorm.io/gorm"
)

type Repository interface {
	//
	GetUserByName(
		ctx context.Context,
		username string, // Required: true, Description:
	) (*user.userDTO, error)
}

func NewRepository(db *gorm.DB) *custom.RepositoryImpl {
	return custom.NewRepository(db)
}
