package generated

import (
	"gorm.io/gorm/clause"

	petMapper "petstore/internal/mapper/pet"
	"petstore/internal/models/pet"

	"context"
	"gorm.io/gorm"
	"petstore/internal/models/apiResponse"
)

type RepositoryImpl struct {
	Db *gorm.DB
}

func NewRepository(db *gorm.DB) *RepositoryImpl {
	return &RepositoryImpl{db}
}

// UpdatePet - Summary: Update an existing pet. Description: Update an existing pet by Id
func (r *RepositoryImpl) UpdatePet(
	ctx context.Context,
	petDTO *pet.PetDTO,
) (*pet.PetDTO, error) {

	entity := petMapper.ConvertPetDTOtoEntity(petDTO)
	db := r.Db.Table("pet_dtos").WithContext(ctx).Session(&gorm.Session{FullSaveAssociations: true}).Updates(entity)
	if err := db.Error; err != nil {
		return nil, err
	}
	return petMapper.ConvertPetEntityToDTO(entity), nil

}

// AddPet - Summary: Add a new pet to the store. Description: Add a new pet to the store
func (r *RepositoryImpl) AddPet(
	ctx context.Context,
	petDTO *pet.PetDTO,
) (*pet.PetDTO, error) {

	entity := petMapper.ConvertPetDTOtoEntity(petDTO)
	if err := r.Db.Table("pet_dtos").WithContext(ctx).Create(entity).Error; err != nil {
		return nil, err
	}
	return petMapper.ConvertPetEntityToDTO(entity), nil

}

// FindPetsByStatus - Summary: Finds Pets by status. Description: Multiple status values can be provided with comma separated strings
func (r *RepositoryImpl) FindPetsByStatus(
	ctx context.Context,
	status string, // Required: false, Description: Status values that need to be considered for filter
) ([]*pet.PetDTO, error) {

	panic("not implemented")

}

// FindPetsByTags - Summary: Finds Pets by tags. Description: Multiple tags can be provided with comma separated strings. Use tag1, tag2, tag3 for testing.
func (r *RepositoryImpl) FindPetsByTags(
	ctx context.Context,
	tags []string, // Required: false, Description: Tags to filter by
) ([]*pet.PetDTO, error) {

	panic("not implemented")

}

// GetPetById - Summary: Find pet by ID. Description: Returns a single pet
func (r *RepositoryImpl) GetPetById(
	ctx context.Context,
	petId int64, // Required: true, Description: ID of pet to return
) (*pet.PetDTO, error) {

	var entity *pet.PetEntity
	if err := r.Db.Table("pet_dtos").WithContext(ctx).Preload(clause.Associations).First(&entity, petId).Error; err != nil {
		return nil, err
	}
	return petMapper.ConvertPetEntityToDTO(entity), nil

}

// UpdatePetWithForm - Summary: Updates a pet in the store with form data. Description:
func (r *RepositoryImpl) UpdatePetWithForm(
	ctx context.Context,
	petId int64, // Required: true, Description: ID of pet that needs to be updated
	name string, // Required: true, Description: Name of pet that needs to be updated
	status string, // Required: true, Description: Status of pet that needs to be updated
) error {

	panic("not implemented")

}

// DeletePet - Summary: Deletes a pet. Description: delete a pet
func (r *RepositoryImpl) DeletePet(
	ctx context.Context,
	petId int64, // Required: true, Description: Pet id to delete
) error {

	tx := r.Db.Begin()
	if err := tx.Table("pet_dtos").WithContext(ctx).Delete(&pet.PetEntity{}, petId).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil

}

// UploadFile - Summary: uploads an image. Description:
func (r *RepositoryImpl) UploadFile(
	ctx context.Context,
	petId int64, // Required: true, Description: ID of pet to update
	additionalMetadata string, // Required: false, Description: Additional Metadata
) (*apiResponse.ApiResponseDTO, error) {

	panic("not implemented")

}
