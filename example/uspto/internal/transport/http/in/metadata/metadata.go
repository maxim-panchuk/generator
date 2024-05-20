package metadata

import (
	metadataService "example/uspto/internal/service/metadata"
	"example/uspto/internal/transport/http/in/metadata/custom"
	"net/http"
)

type Controller interface {
	// ListDataSets - Summary: List available data sets. Description:
	ListDataSets(
		w http.ResponseWriter, r *http.Request)
	// ListSearchableFields - Summary: Provides the general information about the API and the list of fields that can be used to query the dataset.. Description: This GET API returns the list of all the searchable field names that are in the oa_citations. Please see the 'fields' attribute which returns an array of field names. Each field or a combination of fields can be searched using the syntax options shown below.
	ListSearchableFields(
		w http.ResponseWriter, r *http.Request)
}

func NewController(service metadataService.Service) *custom.ControllerImpl {
	return custom.NewController(service)
}
