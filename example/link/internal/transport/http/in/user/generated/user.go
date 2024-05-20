package generated

import (
	userService "example/link/internal/service/user"
	"net/http"
)

type ControllerImpl struct {
	service userService.Service
}

func NewController(service userService.Service) *ControllerImpl {
	return &ControllerImpl{service: service}
}

// GetUserByName godoc
// @Summary  "default summary"
// @Description  "default description"
// @Tags user
// @Accept json
// @Param username path string true ""
// @Success  200    {object}  user.userDTO  "The User"
// @Router /2.0/users/{username} [get]
// @Security BearerAuth
func (c *ControllerImpl) GetUserByName(
	w http.ResponseWriter, r *http.Request) {
	panic("not implemented")

}
