package generated

import (
	"context"
	"example/link/internal/models/user"
	"gorm.io/gorm"
)

type RepositoryImpl struct {
	Db *gorm.DB
}

func NewRepository(db *gorm.DB) *RepositoryImpl {
	return &RepositoryImpl{db}
}

//
func (r *RepositoryImpl) GetUserByName(
	ctx context.Context,
	username string, // Required: true, Description:
) (*user.userDTO, error) {

	panic("not implemented")

}
