package generated

import (
	searchService "example/uspto/internal/service/search"
	"net/http"
)

type ControllerImpl struct {
	service searchService.Service
}

func NewController(service searchService.Service) *ControllerImpl {
	return &ControllerImpl{service: service}
}

// PerformSearch godoc
// @Summary  "Provides search capability for the data set with the given search criteria."
// @Description  "This API is based on Solr/Lucene Search. The data is indexed using SOLR. This GET API returns the list of all the searchable field names that are in the Solr Index. Please see the 'fields' attribute which returns an array of field names. Each field or a combination of fields can be searched using the Solr/Lucene Syntax. Please refer https://lucene.apache.org/core/3_6_2/queryparsersyntax.html#Overview for the query syntax. List of field names that are searchable can be determined using above GET api."
// @Tags search
// @Accept json
// @Param version path string true "Version of the dataset."
// @Param dataset path string true "Name of the dataset. In this case, the default value is oa_citations"
// @Success  200  {array}   "successful operation"
// @Failure  404   "No matching record found for the given criteria."
// @Router /{dataset}/{version}/records [post]
// @Security BearerAuth
func (c *ControllerImpl) PerformSearch(
	w http.ResponseWriter, r *http.Request) {
	panic("not implemented")

}
