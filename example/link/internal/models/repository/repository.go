package repository

import (
	"example/link/internal/models/user"
)

type repositoryDTO struct {
	Slug  string        `json:"slug"`
	Owner *user.userDTO `json:"owner"`
}
