package order

import (
	model "petstore/internal/models/order"
)

func ConvertOrderDTOtoEntity(dto *model.OrderDTO) *model.OrderEntity {
	return &model.OrderEntity{

		Id: dto.Id,

		PetId: dto.PetId,

		Quantity: dto.Quantity,

		ShipDate: dto.ShipDate,

		Status: dto.Status,

		Complete: dto.Complete,
	}
}

func ConvertOrderEntityToDTO(entity *model.OrderEntity) *model.OrderDTO {
	return &model.OrderDTO{

		Id: entity.Id,

		PetId: entity.PetId,

		Quantity: entity.Quantity,

		ShipDate: entity.ShipDate,

		Status: entity.Status,

		Complete: entity.Complete,
	}
}
