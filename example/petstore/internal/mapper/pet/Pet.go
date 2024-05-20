package pet

import (
	model "petstore/internal/models/pet"

	categoryMapper "petstore/internal/mapper/category"

	tagMapper "petstore/internal/mapper/tag"

	categoryModel "petstore/internal/models/category"

	tagModel "petstore/internal/models/tag"
)

func ConvertPetDTOtoEntity(dto *model.PetDTO) *model.PetEntity {
	return &model.PetEntity{

		Id: dto.Id,

		Name: dto.Name,

		Category: func() []*categoryModel.CategoryEntity {
			return []*categoryModel.CategoryEntity{
				categoryMapper.ConvertCategoryDTOtoEntity(dto.Category),
			}
		}(),
		//Category: categoryMapper.ConvertCategoryDTOtoEntity(dto.Category[0]),

		PhotoUrls: dto.PhotoUrls,

		Tags: func() []*tagModel.TagEntity {
			slice := make([]*tagModel.TagEntity, 0)
			for _, el := range dto.Tags {
				slice = append(slice, tagMapper.ConvertTagDTOtoEntity(el))
			}
			return slice
		}(),

		Status: dto.Status,
	}
}

func ConvertPetEntityToDTO(entity *model.PetEntity) *model.PetDTO {
	return &model.PetDTO{

		Id: entity.Id,

		Name: entity.Name,

		Category: categoryMapper.ConvertCategoryEntityToDTO(entity.Category[0]),

		PhotoUrls: entity.PhotoUrls,

		Tags: func() []*tagModel.TagDTO {
			slice := make([]*tagModel.TagDTO, 0)
			for _, el := range entity.Tags {
				slice = append(slice, tagMapper.ConvertTagEntityToDTO(el))
			}
			return slice
		}(),

		Status: entity.Status,
	}
}
