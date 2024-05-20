package generated

import (
	"encoding/json"
	"petstore/internal/models/pet"
	petService "petstore/internal/service/pet"

	"github.com/gorilla/mux"
	"strconv"

	"context"
	"net/http"
)

type ControllerImpl struct {
	service petService.Service
}

func NewController(service petService.Service) *ControllerImpl {
	return &ControllerImpl{service: service}
}

// UpdatePet godoc
// @Summary  "Update an existing pet"
// @Description  "Update an existing pet by Id"
// @Tags pet
// @Accept json
// @Param Pet body pet.PetDTO true "Update an existing pet by Id"
// @Success  200    {object}  pet.PetDTO  "Successful operation"
// @Failure  400   "Invalid ID supplied"
// @Failure  404   "Pet not found"
// @Failure  422   "Validation exception"
// @Router /pet [put]
// @Security BearerAuth
func (c *ControllerImpl) UpdatePet(
	w http.ResponseWriter, r *http.Request) {

	// TODO генерация, если слайс схем
	var petDTO *pet.PetDTO
	err := json.NewDecoder(r.Body).Decode(&petDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid input"))
		return
	}

	dto, err := c.service.UpdatePet(
		context.Background(), petDTO,
	)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	b, err := json.Marshal(dto)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
	return

}

// AddPet godoc
// @Summary  "Add a new pet to the store"
// @Description  "Add a new pet to the store"
// @Tags pet
// @Accept json
// @Param Pet body pet.PetDTO true "Add a new pet to the store"
// @Success  200    {object}  pet.PetDTO  "Successful operation"
// @Failure  400   "Invalid input"
// @Failure  422   "Validation exception"
// @Router /pet [post]
// @Security BearerAuth
func (c *ControllerImpl) AddPet(
	w http.ResponseWriter, r *http.Request) {

	// TODO генерация, если слайс схем
	var petDTO *pet.PetDTO
	err := json.NewDecoder(r.Body).Decode(&petDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid input"))
		return
	}

	dto, err := c.service.AddPet(
		context.Background(), petDTO,
	)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	b, err := json.Marshal(dto)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
	return

}

// FindPetsByStatus godoc
// @Summary  "Finds Pets by status"
// @Description  "Multiple status values can be provided with comma separated strings"
// @Tags pet
// @Accept json
// @Param status query string false "Status values that need to be considered for filter"
// @Success  200  {array}    pet.PetDTO  "successful operation"
// @Failure  400   "Invalid status value"
// @Router /pet/findByStatus [get]
// @Security BearerAuth
func (c *ControllerImpl) FindPetsByStatus(
	w http.ResponseWriter, r *http.Request) {
	panic("not implemented")

}

// FindPetsByTags godoc
// @Summary  "Finds Pets by tags"
// @Description  "Multiple tags can be provided with comma separated strings. Use tag1, tag2, tag3 for testing."
// @Tags pet
// @Accept json
// @Param tags query string false "Tags to filter by"
// @Success  200  {array}    pet.PetDTO  "successful operation"
// @Failure  400   "Invalid tag value"
// @Router /pet/findByTags [get]
// @Security BearerAuth
func (c *ControllerImpl) FindPetsByTags(
	w http.ResponseWriter, r *http.Request) {
	panic("not implemented")

}

// GetPetById godoc
// @Summary  "Find pet by ID"
// @Description  "Returns a single pet"
// @Tags pet
// @Accept json
// @Param petId path integer true "ID of pet to return"
// @Success  200    {object}  pet.PetDTO  "successful operation"
// @Failure  400   "Invalid ID supplied"
// @Failure  404   "Pet not found"
// @Router /pet/{petId} [get]
// @Security BearerAuth
func (c *ControllerImpl) GetPetById(
	w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	petIdIn, ok := vars["petId"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid ID supplied"))
		return

	}

	petIdInt, err := strconv.Atoi(petIdIn)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("bad user input"))
		return
	}

	petId := int64(petIdInt)

	dto, err := c.service.GetPetById(
		context.Background(), petId,
	)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	b, err := json.Marshal(dto)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
	return

}

// UpdatePetWithForm godoc
// @Summary  "Updates a pet in the store with form data"
// @Description  "default description"
// @Tags pet
// @Accept json
// @Param petId path integer true "ID of pet that needs to be updated"
// @Param name query string true "Name of pet that needs to be updated"
// @Param status query string true "Status of pet that needs to be updated"
// @Failure  400   "Invalid input"
// @Router /pet/{petId} [post]
// @Security BearerAuth
func (c *ControllerImpl) UpdatePetWithForm(
	w http.ResponseWriter, r *http.Request) {
	panic("not implemented")

}

// DeletePet godoc
// @Summary  "Deletes a pet"
// @Description  "delete a pet"
// @Tags pet
// @Accept json
// @Param petId path integer true "Pet id to delete"
// @Failure  400   "Invalid pet value"
// @Router /pet/{petId} [delete]
// @Security BearerAuth
func (c *ControllerImpl) DeletePet(
	w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	petIdIn, ok := vars["petId"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid pet value"))
		return

	}

	petIdInt, err := strconv.Atoi(petIdIn)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("bad user input"))
		return
	}

	petId := int64(petIdInt)

	if err := c.service.DeletePet(
		context.Background(), petId,
	); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	} else {
		w.WriteHeader(http.StatusOK)
		return
	}

}

// UploadFile godoc
// @Summary  "uploads an image"
// @Description  "default description"
// @Tags pet
// @Accept json
// @Param petId path integer true "ID of pet to update"
// @Param additionalMetadata query string false "Additional Metadata"
// @Success  200    {object}  apiResponse.ApiResponseDTO  "successful operation"
// @Router /pet/{petId}/uploadImage [post]
// @Security BearerAuth
func (c *ControllerImpl) UploadFile(
	w http.ResponseWriter, r *http.Request) {
	panic("not implemented")

}
