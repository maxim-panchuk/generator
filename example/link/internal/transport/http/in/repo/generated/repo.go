package generated

import (
	repoService "example/link/internal/service/repo"
	"net/http"
)

type ControllerImpl struct {
	service repoService.Service
}

func NewController(service repoService.Service) *ControllerImpl {
	return &ControllerImpl{service: service}
}

// GetRepositoriesByOwner godoc
// @Summary  "default summary"
// @Description  "default description"
// @Tags repo
// @Accept json
// @Param username path string true ""
// @Success  200  {array}    repository.repositoryDTO  "repositories owned by the supplied user"
// @Router /2.0/repositories/{username} [get]
// @Security BearerAuth
func (c *ControllerImpl) GetRepositoriesByOwner(
	w http.ResponseWriter, r *http.Request) {
	panic("not implemented")

}

// GetRepository godoc
// @Summary  "default summary"
// @Description  "default description"
// @Tags repo
// @Accept json
// @Param username path string true ""
// @Param slug path string true ""
// @Success  200    {object}  repository.repositoryDTO  "The repository"
// @Router /2.0/repositories/{username}/{slug} [get]
// @Security BearerAuth
func (c *ControllerImpl) GetRepository(
	w http.ResponseWriter, r *http.Request) {
	panic("not implemented")

}

// GetPullRequestsByRepository godoc
// @Summary  "default summary"
// @Description  "default description"
// @Tags repo
// @Accept json
// @Param username path string true ""
// @Param slug path string true ""
// @Param state query string true ""
// @Success  200  {array}    pullrequest.pullrequestDTO  "an array of pull request objects"
// @Router /2.0/repositories/{username}/{slug}/pullrequests [get]
// @Security BearerAuth
func (c *ControllerImpl) GetPullRequestsByRepository(
	w http.ResponseWriter, r *http.Request) {
	panic("not implemented")

}

// GetPullRequestsById godoc
// @Summary  "default summary"
// @Description  "default description"
// @Tags repo
// @Accept json
// @Param username path string true ""
// @Param slug path string true ""
// @Param pid path string true ""
// @Success  200    {object}  pullrequest.pullrequestDTO  "a pull request object"
// @Router /2.0/repositories/{username}/{slug}/pullrequests/{pid} [get]
// @Security BearerAuth
func (c *ControllerImpl) GetPullRequestsById(
	w http.ResponseWriter, r *http.Request) {
	panic("not implemented")

}

// MergePullRequest godoc
// @Summary  "default summary"
// @Description  "default description"
// @Tags repo
// @Accept json
// @Param username path string true ""
// @Param slug path string true ""
// @Param pid path string true ""
// @Failure  204   "the PR was successfully merged"
// @Router /2.0/repositories/{username}/{slug}/pullrequests/{pid}/merge [post]
// @Security BearerAuth
func (c *ControllerImpl) MergePullRequest(
	w http.ResponseWriter, r *http.Request) {
	panic("not implemented")

}
