package custom

import (
	searchService "example/uspto/internal/service/search"
	"example/uspto/internal/transport/http/in/search/generated"
)

type ControllerImpl struct {
	*generated.ControllerImpl
}

func NewController(service searchService.Service) *ControllerImpl {
	return &ControllerImpl{
		ControllerImpl: generated.NewController(service),
	}
}
