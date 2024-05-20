package user

import (
	"net/http"
	userService "petstore/internal/service/user"
	"petstore/internal/transport/http/in/user/custom"
)

type Controller interface {
	// CreateUser - Summary: Create user. Description: This can only be done by the logged in user.
	CreateUser(
		w http.ResponseWriter, r *http.Request)
	// CreateUsersWithListInput - Summary: Creates list of users with given input array. Description: Creates list of users with given input array
	CreateUsersWithListInput(
		w http.ResponseWriter, r *http.Request)
	// LoginUser - Summary: Logs user into the system. Description:
	LoginUser(
		w http.ResponseWriter, r *http.Request)
	// LogoutUser - Summary: Logs out current logged in user session. Description:
	LogoutUser(
		w http.ResponseWriter, r *http.Request)
	// GetUserByName - Summary: Get user by user name. Description:
	GetUserByName(
		w http.ResponseWriter, r *http.Request)
	// UpdateUser - Summary: Update user. Description: This can only be done by the logged in user.
	UpdateUser(
		w http.ResponseWriter, r *http.Request)
	// DeleteUser - Summary: Delete user. Description: This can only be done by the logged in user.
	DeleteUser(
		w http.ResponseWriter, r *http.Request)
}

func NewController(service userService.Service) *custom.ControllerImpl {
	return custom.NewController(service)
}
