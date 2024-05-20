package custom

import (
	storeService "petstore/internal/service/store"
	"petstore/internal/transport/http/in/store/generated"
)

type ControllerImpl struct {
	*generated.ControllerImpl
}

func NewController(service storeService.Service) *ControllerImpl {
	return &ControllerImpl{
		ControllerImpl: generated.NewController(service),
	}
}
