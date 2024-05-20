package metadata

import (
	"context"
	metadataRepository "example/uspto/internal/database/repositories/metadata"
	"example/uspto/internal/models/dataSetList"
	"example/uspto/internal/service/metadata/custom"
)

type Service interface {
	// ListDataSets - Summary: List available data sets. Description:
	ListDataSets(
		ctx context.Context,
	) (*dataSetList.dataSetListDTO, error)
	// ListSearchableFields - Summary: Provides the general information about the API and the list of fields that can be used to query the dataset.. Description: This GET API returns the list of all the searchable field names that are in the oa_citations. Please see the 'fields' attribute which returns an array of field names. Each field or a combination of fields can be searched using the syntax options shown below.
	ListSearchableFields(
		ctx context.Context,
		dataset string, // Required: true, Description: Name of the dataset.
		version string, // Required: true, Description: Version of the dataset.
	) error
}

func NewService(repo metadataRepository.Repository) *custom.ServiceImpl {
	return custom.NewService(repo)
}
