package custom

import (
	petService "petstore/internal/service/pet"
	"petstore/internal/transport/http/in/pet/generated"
)

type ControllerImpl struct {
	*generated.ControllerImpl
}

func NewController(service petService.Service) *ControllerImpl {
	return &ControllerImpl{
		ControllerImpl: generated.NewController(service),
	}
}
