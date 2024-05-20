package user

import (
	userService "example/link/internal/service/user"
	"example/link/internal/transport/http/in/user/custom"
	"net/http"
)

type Controller interface {
	//
	GetUserByName(
		w http.ResponseWriter, r *http.Request)
}

func NewController(service userService.Service) *custom.ControllerImpl {
	return custom.NewController(service)
}
