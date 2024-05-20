package custom

import (
	userService "example/link/internal/service/user"
	"example/link/internal/transport/http/in/user/generated"
)

type ControllerImpl struct {
	*generated.ControllerImpl
}

func NewController(service userService.Service) *ControllerImpl {
	return &ControllerImpl{
		ControllerImpl: generated.NewController(service),
	}
}
