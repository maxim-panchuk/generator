package custom

import (
	repoService "example/link/internal/service/repo"
	"example/link/internal/transport/http/in/repo/generated"
)

type ControllerImpl struct {
	*generated.ControllerImpl
}

func NewController(service repoService.Service) *ControllerImpl {
	return &ControllerImpl{
		ControllerImpl: generated.NewController(service),
	}
}
