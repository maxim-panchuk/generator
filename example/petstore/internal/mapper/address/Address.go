package address

import (
	model "petstore/internal/models/address"
)

func ConvertAddressDTOtoEntity(dto *model.AddressDTO) *model.AddressEntity {
	return &model.AddressEntity{

		Id: dto.Id,

		Street: dto.Street,

		City: dto.City,

		State: dto.State,

		Zip: dto.Zip,
	}
}

func ConvertAddressEntityToDTO(entity *model.AddressEntity) *model.AddressDTO {
	return &model.AddressDTO{

		Id: entity.Id,

		Street: entity.Street,

		City: entity.City,

		State: entity.State,

		Zip: entity.Zip,
	}
}
