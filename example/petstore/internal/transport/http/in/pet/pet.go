package pet

import (
	"net/http"
	petService "petstore/internal/service/pet"
	"petstore/internal/transport/http/in/pet/custom"
)

type Controller interface {
	// UpdatePet - Summary: Update an existing pet. Description: Update an existing pet by Id
	UpdatePet(
		w http.ResponseWriter, r *http.Request)
	// AddPet - Summary: Add a new pet to the store. Description: Add a new pet to the store
	AddPet(
		w http.ResponseWriter, r *http.Request)
	// FindPetsByStatus - Summary: Finds Pets by status. Description: Multiple status values can be provided with comma separated strings
	FindPetsByStatus(
		w http.ResponseWriter, r *http.Request)
	// FindPetsByTags - Summary: Finds Pets by tags. Description: Multiple tags can be provided with comma separated strings. Use tag1, tag2, tag3 for testing.
	FindPetsByTags(
		w http.ResponseWriter, r *http.Request)
	// GetPetById - Summary: Find pet by ID. Description: Returns a single pet
	GetPetById(
		w http.ResponseWriter, r *http.Request)
	// UpdatePetWithForm - Summary: Updates a pet in the store with form data. Description:
	UpdatePetWithForm(
		w http.ResponseWriter, r *http.Request)
	// DeletePet - Summary: Deletes a pet. Description: delete a pet
	DeletePet(
		w http.ResponseWriter, r *http.Request)
	// UploadFile - Summary: uploads an image. Description:
	UploadFile(
		w http.ResponseWriter, r *http.Request)
}

func NewController(service petService.Service) *custom.ControllerImpl {
	return custom.NewController(service)
}
