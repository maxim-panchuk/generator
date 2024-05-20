package pullrequest

import (
	"example/link/internal/models/repository"

	"example/link/internal/models/user"
)

type pullrequestDTO struct {
	Id         int64                     `json:"id"`
	Title      string                    `json:"title"`
	Repository *repository.repositoryDTO `json:"repository"`
	Author     *user.userDTO             `json:"author"`
}
