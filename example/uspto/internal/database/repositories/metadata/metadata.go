package metadata

import (
	"context"
	"example/uspto/internal/database/repositories/metadata/custom"
	"example/uspto/internal/models/dataSetList"
	"gorm.io/gorm"
)

type Repository interface {
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

func NewRepository(db *gorm.DB) *custom.RepositoryImpl {
	return custom.NewRepository(db)
}
