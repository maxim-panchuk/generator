package generated

import (
	userMapper "petstore/internal/mapper/user"
	"petstore/internal/models/user"

	"context"
	"gorm.io/gorm"
)

type RepositoryImpl struct {
	Db *gorm.DB
}

func NewRepository(db *gorm.DB) *RepositoryImpl {
	return &RepositoryImpl{db}
}

// CreateUser - Summary: Create user. Description: This can only be done by the logged in user.
func (r *RepositoryImpl) CreateUser(
	ctx context.Context,
	userDTO *user.UserDTO,
) (*user.UserDTO, error) {

	entity := userMapper.ConvertUserDTOtoEntity(userDTO)
	if err := r.Db.Table("user_dtos").WithContext(ctx).Create(entity).Error; err != nil {
		return nil, err
	}
	return userMapper.ConvertUserEntityToDTO(entity), nil

}

// CreateUsersWithListInput - Summary: Creates list of users with given input array. Description: Creates list of users with given input array
func (r *RepositoryImpl) CreateUsersWithListInput(
	ctx context.Context,
	userDTO []*user.UserDTO, // Creates list of users with given input array

) (*user.UserDTO, error) {

	panic("not implemented")

}

// LoginUser - Summary: Logs user into the system. Description:
func (r *RepositoryImpl) LoginUser(
	ctx context.Context,
	username string, // Required: false, Description: The user name for login
	password string, // Required: false, Description: The password for login in clear text
) error {

	panic("not implemented")

}

// LogoutUser - Summary: Logs out current logged in user session. Description:
func (r *RepositoryImpl) LogoutUser(
	ctx context.Context,
) error {

	panic("not implemented")

}

// GetUserByName - Summary: Get user by user name. Description:
func (r *RepositoryImpl) GetUserByName(
	ctx context.Context,
	username string, // Required: true, Description: The name that needs to be fetched. Use user1 for testing.
) (*user.UserDTO, error) {

	panic("not implemented")

}

// UpdateUser - Summary: Update user. Description: This can only be done by the logged in user.
func (r *RepositoryImpl) UpdateUser(
	ctx context.Context,
	userDTO *user.UserDTO,
	username string, // Required: true, Description: name that need to be deleted
) error {

	panic("not implemented")

}

// DeleteUser - Summary: Delete user. Description: This can only be done by the logged in user.
func (r *RepositoryImpl) DeleteUser(
	ctx context.Context,
	username string, // Required: true, Description: The name that needs to be deleted
) error {

	panic("not implemented")

}
