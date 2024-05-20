package tag

import (
	model "petstore/internal/models/tag"
)

func ConvertTagDTOtoEntity(dto *model.TagDTO) *model.TagEntity {
	return &model.TagEntity{

		Id: dto.Id,

		Name: dto.Name,
	}
}

func ConvertTagEntityToDTO(entity *model.TagEntity) *model.TagDTO {
	return &model.TagDTO{

		Id: entity.Id,

		Name: entity.Name,
	}
}
