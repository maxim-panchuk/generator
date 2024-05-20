package pet

import (
	"context"
	petRepository "petstore/internal/database/repositories/pet"
	"petstore/internal/models/apiResponse"
	"petstore/internal/models/pet"
	"petstore/internal/service/pet/custom"
)

type Service interface {
	// UpdatePet - Summary: Update an existing pet. Description: Update an existing pet by Id
	UpdatePet(
		ctx context.Context,
		petDTO *pet.PetDTO,
	) (*pet.PetDTO, error)
	// AddPet - Summary: Add a new pet to the store. Description: Add a new pet to the store
	AddPet(
		ctx context.Context,
		petDTO *pet.PetDTO,
	) (*pet.PetDTO, error)
	// FindPetsByStatus - Summary: Finds Pets by status. Description: Multiple status values can be provided with comma separated strings
	FindPetsByStatus(
		ctx context.Context,
		status string, // Required: false, Description: Status values that need to be considered for filter
	) ([]*pet.PetDTO, error)
	// FindPetsByTags - Summary: Finds Pets by tags. Description: Multiple tags can be provided with comma separated strings. Use tag1, tag2, tag3 for testing.
	FindPetsByTags(
		ctx context.Context,
		tags []string, // Required: false, Description: Tags to filter by
	) ([]*pet.PetDTO, error)
	// GetPetById - Summary: Find pet by ID. Description: Returns a single pet
	GetPetById(
		ctx context.Context,
		petId int64, // Required: true, Description: ID of pet to return
	) (*pet.PetDTO, error)
	// UpdatePetWithForm - Summary: Updates a pet in the store with form data. Description:
	UpdatePetWithForm(
		ctx context.Context,
		petId int64, // Required: true, Description: ID of pet that needs to be updated
		name string, // Required: true, Description: Name of pet that needs to be updated
		status string, // Required: true, Description: Status of pet that needs to be updated
	) error
	// DeletePet - Summary: Deletes a pet. Description: delete a pet
	DeletePet(
		ctx context.Context,
		petId int64, // Required: true, Description: Pet id to delete
	) error
	// UploadFile - Summary: uploads an image. Description:
	UploadFile(
		ctx context.Context,
		petId int64, // Required: true, Description: ID of pet to update
		additionalMetadata string, // Required: false, Description: Additional Metadata
	) (*apiResponse.ApiResponseDTO, error)
}

func NewService(repo petRepository.Repository) *custom.ServiceImpl {
	return custom.NewService(repo)
}
