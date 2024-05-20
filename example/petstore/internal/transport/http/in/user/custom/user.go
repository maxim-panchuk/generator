package custom

import (
	userService "petstore/internal/service/user"
	"petstore/internal/transport/http/in/user/generated"
)

type ControllerImpl struct {
	*generated.ControllerImpl
}

func NewController(service userService.Service) *ControllerImpl {
	return &ControllerImpl{
		ControllerImpl: generated.NewController(service),
	}
}
