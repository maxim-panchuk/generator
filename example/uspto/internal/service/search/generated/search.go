package generated

import (
	"context"
	searchRepository "example/uspto/internal/database/repositories/search"
)

type ServiceImpl struct {
	repo searchRepository.Repository
}

func NewService(repo searchRepository.Repository) *ServiceImpl {
	return &ServiceImpl{
		repo: repo,
	}
}

// PerformSearch - Summary: Provides search capability for the data set with the given search criteria.. Description: This API is based on Solr/Lucene Search. The data is indexed using SOLR. This GET API returns the list of all the searchable field names that are in the Solr Index. Please see the 'fields' attribute which returns an array of field names. Each field or a combination of fields can be searched using the Solr/Lucene Syntax. Please refer https://lucene.apache.org/core/3_6_2/queryparsersyntax.html#Overview for the query syntax. List of field names that are searchable can be determined using above GET api.
func (s *ServiceImpl) PerformSearch(
	ctx context.Context,
	version string, // Required: true, Description: Version of the dataset.
	dataset string, // Required: true, Description: Name of the dataset. In this case, the default value is oa_citations
) error {
	panic("not implemented")
}
