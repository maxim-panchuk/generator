package category

import (
	model "petstore/internal/models/category"
)

func ConvertCategoryDTOtoEntity(dto *model.CategoryDTO) *model.CategoryEntity {
	return &model.CategoryEntity{

		Id: dto.Id,

		Name: dto.Name,
	}
}

func ConvertCategoryEntityToDTO(entity *model.CategoryEntity) *model.CategoryDTO {
	return &model.CategoryDTO{

		Id: entity.Id,

		Name: entity.Name,
	}
}
