package generated

import (
	metadataService "example/uspto/internal/service/metadata"
	"net/http"
)

type ControllerImpl struct {
	service metadataService.Service
}

func NewController(service metadataService.Service) *ControllerImpl {
	return &ControllerImpl{service: service}
}

// ListDataSets godoc
// @Summary  "List available data sets"
// @Description  "default description"
// @Tags metadata
// @Accept json
// @Success  200    {object}  dataSetList.dataSetListDTO  "Returns a list of data sets"
// @Router / [get]
// @Security BearerAuth
func (c *ControllerImpl) ListDataSets(
	w http.ResponseWriter, r *http.Request) {
	panic("not implemented")

}

// ListSearchableFields godoc
// @Summary  "Provides the general information about the API and the list of fields that can be used to query the dataset."
// @Description  "This GET API returns the list of all the searchable field names that are in the oa_citations. Please see the 'fields' attribute which returns an array of field names. Each field or a combination of fields can be searched using the syntax options shown below."
// @Tags metadata
// @Accept json
// @Param dataset path string true "Name of the dataset."
// @Param version path string true "Version of the dataset."
// @Success  200   "The dataset API for the given version is found and it is accessible to consume."
// @Failure  404   "The combination of dataset name and version is not found in the system or it is not published yet to be consumed by public."
// @Router /{dataset}/{version}/fields [get]
// @Security BearerAuth
func (c *ControllerImpl) ListSearchableFields(
	w http.ResponseWriter, r *http.Request) {
	panic("not implemented")

}
