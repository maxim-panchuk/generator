package custom

import (
	metadataService "example/uspto/internal/service/metadata"
	"example/uspto/internal/transport/http/in/metadata/generated"
)

type ControllerImpl struct {
	*generated.ControllerImpl
}

func NewController(service metadataService.Service) *ControllerImpl {
	return &ControllerImpl{
		ControllerImpl: generated.NewController(service),
	}
}
