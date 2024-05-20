package generated

import (
	"context"
	petRepository "petstore/internal/database/repositories/pet"
	"petstore/internal/models/apiResponse"
	"petstore/internal/models/pet"
)

type ServiceImpl struct {
	repo petRepository.Repository
}

func NewService(repo petRepository.Repository) *ServiceImpl {
	return &ServiceImpl{
		repo: repo,
	}
}

// UpdatePet - Summary: Update an existing pet. Description: Update an existing pet by Id
func (s *ServiceImpl) UpdatePet(
	ctx context.Context,
	petDTO *pet.PetDTO,
) (*pet.PetDTO, error) {
	dto, err := s.repo.UpdatePet(
		ctx, petDTO,
	)
	if err != nil {
		return nil, err
	}
	return dto, nil

}

// AddPet - Summary: Add a new pet to the store. Description: Add a new pet to the store
func (s *ServiceImpl) AddPet(
	ctx context.Context,
	petDTO *pet.PetDTO,
) (*pet.PetDTO, error) {
	dto, err := s.repo.AddPet(
		ctx, petDTO,
	)
	if err != nil {
		return nil, err
	}
	return dto, nil

}

// FindPetsByStatus - Summary: Finds Pets by status. Description: Multiple status values can be provided with comma separated strings
func (s *ServiceImpl) FindPetsByStatus(
	ctx context.Context,
	status string, // Required: false, Description: Status values that need to be considered for filter
) ([]*pet.PetDTO, error) {
	panic("not implemented")
}

// FindPetsByTags - Summary: Finds Pets by tags. Description: Multiple tags can be provided with comma separated strings. Use tag1, tag2, tag3 for testing.
func (s *ServiceImpl) FindPetsByTags(
	ctx context.Context,
	tags []string, // Required: false, Description: Tags to filter by
) ([]*pet.PetDTO, error) {
	panic("not implemented")
}

// GetPetById - Summary: Find pet by ID. Description: Returns a single pet
func (s *ServiceImpl) GetPetById(
	ctx context.Context,
	petId int64, // Required: true, Description: ID of pet to return
) (*pet.PetDTO, error) {
	dto, err := s.repo.GetPetById(
		ctx, petId,
	)
	if err != nil {
		return nil, err
	}
	return dto, nil

}

// UpdatePetWithForm - Summary: Updates a pet in the store with form data. Description:
func (s *ServiceImpl) UpdatePetWithForm(
	ctx context.Context,
	petId int64, // Required: true, Description: ID of pet that needs to be updated
	name string, // Required: true, Description: Name of pet that needs to be updated
	status string, // Required: true, Description: Status of pet that needs to be updated
) error {
	panic("not implemented")
}

// DeletePet - Summary: Deletes a pet. Description: delete a pet
func (s *ServiceImpl) DeletePet(
	ctx context.Context,
	petId int64, // Required: true, Description: Pet id to delete
) error {
	if err := s.repo.DeletePet(
		ctx, petId,
	); err != nil {
		return err
	}
	return nil

}

// UploadFile - Summary: uploads an image. Description:
func (s *ServiceImpl) UploadFile(
	ctx context.Context,
	petId int64, // Required: true, Description: ID of pet to update
	additionalMetadata string, // Required: false, Description: Additional Metadata
) (*apiResponse.ApiResponseDTO, error) {
	panic("not implemented")
}
