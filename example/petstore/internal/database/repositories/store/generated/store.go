package generated

import (
	"gorm.io/gorm/clause"

	orderMapper "petstore/internal/mapper/order"
	"petstore/internal/models/order"

	"context"
	"gorm.io/gorm"
)

type RepositoryImpl struct {
	Db *gorm.DB
}

func NewRepository(db *gorm.DB) *RepositoryImpl {
	return &RepositoryImpl{db}
}

// GetInventory - Summary: Returns pet inventories by status. Description: Returns a map of status codes to quantities
func (r *RepositoryImpl) GetInventory(
	ctx context.Context,
) error {

	panic("not implemented")

}

// PlaceOrder - Summary: Place an order for a pet. Description: Place a new order in the store
func (r *RepositoryImpl) PlaceOrder(
	ctx context.Context,
	orderDTO *order.OrderDTO,
) (*order.OrderDTO, error) {

	entity := orderMapper.ConvertOrderDTOtoEntity(orderDTO)
	if err := r.Db.Table("order_dtos").WithContext(ctx).Create(entity).Error; err != nil {
		return nil, err
	}
	return orderMapper.ConvertOrderEntityToDTO(entity), nil

}

// GetOrderById - Summary: Find purchase order by ID. Description: For valid response try integer IDs with value <= 5 or > 10. Other values will generate exceptions.
func (r *RepositoryImpl) GetOrderById(
	ctx context.Context,
	orderId int64, // Required: true, Description: ID of order that needs to be fetched
) (*order.OrderDTO, error) {

	var entity *order.OrderEntity
	if err := r.Db.Table("order_dtos").WithContext(ctx).Preload(clause.Associations).First(&entity, orderId).Error; err != nil {
		return nil, err
	}
	return orderMapper.ConvertOrderEntityToDTO(entity), nil

}

// DeleteOrder - Summary: Delete purchase order by ID. Description: For valid response try integer IDs with value < 1000. Anything above 1000 or nonintegers will generate API errors
func (r *RepositoryImpl) DeleteOrder(
	ctx context.Context,
	orderId int64, // Required: true, Description: ID of the order that needs to be deleted
) error {

	tx := r.Db.Begin()
	if err := tx.Table("order_dtos").WithContext(ctx).Delete(&order.OrderEntity{}, orderId).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil

}
