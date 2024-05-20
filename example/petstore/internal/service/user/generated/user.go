package generated

import (
	"context"
	userRepository "petstore/internal/database/repositories/user"
	"petstore/internal/models/user"
)

type ServiceImpl struct {
	repo userRepository.Repository
}

func NewService(repo userRepository.Repository) *ServiceImpl {
	return &ServiceImpl{
		repo: repo,
	}
}

// CreateUser - Summary: Create user. Description: This can only be done by the logged in user.
func (s *ServiceImpl) CreateUser(
	ctx context.Context,
	userDTO *user.UserDTO,
) (*user.UserDTO, error) {
	dto, err := s.repo.CreateUser(
		ctx, userDTO,
	)
	if err != nil {
		return nil, err
	}
	return dto, nil

}

// CreateUsersWithListInput - Summary: Creates list of users with given input array. Description: Creates list of users with given input array
func (s *ServiceImpl) CreateUsersWithListInput(
	ctx context.Context,
	userDTO []*user.UserDTO, // Creates list of users with given input array

) (*user.UserDTO, error) {
	panic("not implemented")
}

// LoginUser - Summary: Logs user into the system. Description:
func (s *ServiceImpl) LoginUser(
	ctx context.Context,
	username string, // Required: false, Description: The user name for login
	password string, // Required: false, Description: The password for login in clear text
) error {
	panic("not implemented")
}

// LogoutUser - Summary: Logs out current logged in user session. Description:
func (s *ServiceImpl) LogoutUser(
	ctx context.Context,
) error {
	panic("not implemented")
}

// GetUserByName - Summary: Get user by user name. Description:
func (s *ServiceImpl) GetUserByName(
	ctx context.Context,
	username string, // Required: true, Description: The name that needs to be fetched. Use user1 for testing.
) (*user.UserDTO, error) {
	panic("not implemented")
}

// UpdateUser - Summary: Update user. Description: This can only be done by the logged in user.
func (s *ServiceImpl) UpdateUser(
	ctx context.Context,
	userDTO *user.UserDTO,
	username string, // Required: true, Description: name that need to be deleted
) error {
	panic("not implemented")
}

// DeleteUser - Summary: Delete user. Description: This can only be done by the logged in user.
func (s *ServiceImpl) DeleteUser(
	ctx context.Context,
	username string, // Required: true, Description: The name that needs to be deleted
) error {
	panic("not implemented")
}
