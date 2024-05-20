package user

import (
	"context"
	"gorm.io/gorm"
	"petstore/internal/database/repositories/user/custom"
	"petstore/internal/models/user"
)

type Repository interface {
	// CreateUser - Summary: Create user. Description: This can only be done by the logged in user.
	CreateUser(
		ctx context.Context,
		userDTO *user.UserDTO,
	) (*user.UserDTO, error)
	// CreateUsersWithListInput - Summary: Creates list of users with given input array. Description: Creates list of users with given input array
	CreateUsersWithListInput(
		ctx context.Context,
		userDTO []*user.UserDTO, // Creates list of users with given input array

	) (*user.UserDTO, error)
	// LoginUser - Summary: Logs user into the system. Description:
	LoginUser(
		ctx context.Context,
		username string, // Required: false, Description: The user name for login
		password string, // Required: false, Description: The password for login in clear text
	) error
	// LogoutUser - Summary: Logs out current logged in user session. Description:
	LogoutUser(
		ctx context.Context,
	) error
	// GetUserByName - Summary: Get user by user name. Description:
	GetUserByName(
		ctx context.Context,
		username string, // Required: true, Description: The name that needs to be fetched. Use user1 for testing.
	) (*user.UserDTO, error)
	// UpdateUser - Summary: Update user. Description: This can only be done by the logged in user.
	UpdateUser(
		ctx context.Context,
		userDTO *user.UserDTO,
		username string, // Required: true, Description: name that need to be deleted
	) error
	// DeleteUser - Summary: Delete user. Description: This can only be done by the logged in user.
	DeleteUser(
		ctx context.Context,
		username string, // Required: true, Description: The name that needs to be deleted
	) error
}

func NewRepository(db *gorm.DB) *custom.RepositoryImpl {
	return custom.NewRepository(db)
}
