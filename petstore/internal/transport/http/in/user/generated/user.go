package generated

import (
	"encoding/json"
	"petstore/internal/models/user"
	userService "petstore/internal/service/user"

	"context"
	"net/http"
)

type ControllerImpl struct {
	service userService.Service
}

func NewController(service userService.Service) *ControllerImpl {
	return &ControllerImpl{service: service}
}

// CreateUser godoc
// @Summary  "Create user"
// @Description  "This can only be done by the logged in user."
// @Tags user
// @Accept json
// @Param User body user.UserDTO true "This can only be done by the logged in user."
// @Success  default    {object}  user.UserDTO  "successful operation"
// @Router /user [post]
// @Security BearerAuth
func (c *ControllerImpl) CreateUser(
	w http.ResponseWriter, r *http.Request) {

	// TODO генерация, если слайс схем
	var userDTO *user.UserDTO
	err := json.NewDecoder(r.Body).Decode(&userDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid input"))
		return
	}

	dto, err := c.service.CreateUser(
		context.Background(), userDTO,
	)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	b, err := json.Marshal(dto)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
	return

}

// CreateUsersWithListInput godoc
// @Summary  "Creates list of users with given input array"
// @Description  "Creates list of users with given input array"
// @Tags user
// @Accept json
// @Param User body user.UserDTO true "Creates list of users with given input array"
// @Success  200    {object}  user.UserDTO  "Successful operation"
// @Success  default   "successful operation"
// @Router /user/createWithList [post]
// @Security BearerAuth
func (c *ControllerImpl) CreateUsersWithListInput(
	w http.ResponseWriter, r *http.Request) {
	panic("not implemented")

}

// LoginUser godoc
// @Summary  "Logs user into the system"
// @Description  "default description"
// @Tags user
// @Accept json
// @Param username query string false "The user name for login"
// @Param password query string false "The password for login in clear text"
// @Success  200   "successful operation"
// @Failure  400   "Invalid username/password supplied"
// @Router /user/login [get]
// @Security BearerAuth
func (c *ControllerImpl) LoginUser(
	w http.ResponseWriter, r *http.Request) {
	panic("not implemented")

}

// LogoutUser godoc
// @Summary  "Logs out current logged in user session"
// @Description  "default description"
// @Tags user
// @Accept json
// @Success  default   "successful operation"
// @Router /user/logout [get]
// @Security BearerAuth
func (c *ControllerImpl) LogoutUser(
	w http.ResponseWriter, r *http.Request) {
	panic("not implemented")

}

// GetUserByName godoc
// @Summary  "Get user by user name"
// @Description  "default description"
// @Tags user
// @Accept json
// @Param username path string true "The name that needs to be fetched. Use user1 for testing. "
// @Success  200    {object}  user.UserDTO  "successful operation"
// @Failure  400   "Invalid username supplied"
// @Failure  404   "User not found"
// @Router /user/{username} [get]
// @Security BearerAuth
func (c *ControllerImpl) GetUserByName(
	w http.ResponseWriter, r *http.Request) {
	panic("not implemented")

}

// UpdateUser godoc
// @Summary  "Update user"
// @Description  "This can only be done by the logged in user."
// @Tags user
// @Accept json
// @Param User body user.UserDTO true "This can only be done by the logged in user."
// @Param username path string true "name that need to be deleted"
// @Success  default   "successful operation"
// @Router /user/{username} [put]
// @Security BearerAuth
func (c *ControllerImpl) UpdateUser(
	w http.ResponseWriter, r *http.Request) {
	panic("not implemented")

}

// DeleteUser godoc
// @Summary  "Delete user"
// @Description  "This can only be done by the logged in user."
// @Tags user
// @Accept json
// @Param username path string true "The name that needs to be deleted"
// @Failure  400   "Invalid username supplied"
// @Failure  404   "User not found"
// @Router /user/{username} [delete]
// @Security BearerAuth
func (c *ControllerImpl) DeleteUser(
	w http.ResponseWriter, r *http.Request) {
	panic("not implemented")

}
