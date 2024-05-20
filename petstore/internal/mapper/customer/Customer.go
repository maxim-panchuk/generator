package customer

import (
	model "petstore/internal/models/customer"

	addressMapper "petstore/internal/mapper/address"

	addressModel "petstore/internal/models/address"
)

func ConvertCustomerDTOtoEntity(dto *model.CustomerDTO) *model.CustomerEntity {
	return &model.CustomerEntity{

		Id: dto.Id,

		Username: dto.Username,

		Address: func() []*addressModel.AddressEntity {
			slice := make([]*addressModel.AddressEntity, 0)
			for _, el := range dto.Address {
				slice = append(slice, addressMapper.ConvertAddressDTOtoEntity(el))
			}
			return slice
		}(),
	}
}

func ConvertCustomerEntityToDTO(entity *model.CustomerEntity) *model.CustomerDTO {
	return &model.CustomerDTO{

		Id: entity.Id,

		Username: entity.Username,

		Address: func() []*addressModel.AddressDTO {
			slice := make([]*addressModel.AddressDTO, 0)
			for _, el := range entity.Address {
				slice = append(slice, addressMapper.ConvertAddressEntityToDTO(el))
			}
			return slice
		}(),
	}
}
