package user

import (
	model "petstore/internal/models/user"
)

func ConvertUserDTOtoEntity(dto *model.UserDTO) *model.UserEntity {
	return &model.UserEntity{

		Id: dto.Id,

		Username: dto.Username,

		FirstName: dto.FirstName,

		LastName: dto.LastName,

		Email: dto.Email,

		Password: dto.Password,

		Phone: dto.Phone,

		UserStatus: dto.UserStatus,
	}
}

func ConvertUserEntityToDTO(entity *model.UserEntity) *model.UserDTO {
	return &model.UserDTO{

		Id: entity.Id,

		Username: entity.Username,

		FirstName: entity.FirstName,

		LastName: entity.LastName,

		Email: entity.Email,

		Password: entity.Password,

		Phone: entity.Phone,

		UserStatus: entity.UserStatus,
	}
}
