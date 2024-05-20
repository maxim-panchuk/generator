package generated

import (
	"context"
	"example/uspto/internal/models/dataSetList"
	"gorm.io/gorm"
)

type RepositoryImpl struct {
	Db *gorm.DB
}

func NewRepository(db *gorm.DB) *RepositoryImpl {
	return &RepositoryImpl{db}
}

// ListDataSets - Summary: List available data sets. Description:
func (r *RepositoryImpl) ListDataSets(
	ctx context.Context,
) (*dataSetList.dataSetListDTO, error) {

	panic("not implemented")

}

// ListSearchableFields - Summary: Provides the general information about the API and the list of fields that can be used to query the dataset.. Description: This GET API returns the list of all the searchable field names that are in the oa_citations. Please see the 'fields' attribute which returns an array of field names. Each field or a combination of fields can be searched using the syntax options shown below.
func (r *RepositoryImpl) ListSearchableFields(
	ctx context.Context,
	dataset string, // Required: true, Description: Name of the dataset.
	version string, // Required: true, Description: Version of the dataset.
) error {

	panic("not implemented")

}
