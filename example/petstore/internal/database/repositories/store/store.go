package store

import (
	"context"
	"gorm.io/gorm"
	"petstore/internal/database/repositories/store/custom"
	"petstore/internal/models/order"
)

type Repository interface {
	// GetInventory - Summary: Returns pet inventories by status. Description: Returns a map of status codes to quantities
	GetInventory(
		ctx context.Context,
	) error
	// PlaceOrder - Summary: Place an order for a pet. Description: Place a new order in the store
	PlaceOrder(
		ctx context.Context,
		orderDTO *order.OrderDTO,
	) (*order.OrderDTO, error)
	// GetOrderById - Summary: Find purchase order by ID. Description: For valid response try integer IDs with value <= 5 or > 10. Other values will generate exceptions.
	GetOrderById(
		ctx context.Context,
		orderId int64, // Required: true, Description: ID of order that needs to be fetched
	) (*order.OrderDTO, error)
	// DeleteOrder - Summary: Delete purchase order by ID. Description: For valid response try integer IDs with value < 1000. Anything above 1000 or nonintegers will generate API errors
	DeleteOrder(
		ctx context.Context,
		orderId int64, // Required: true, Description: ID of the order that needs to be deleted
	) error
}

func NewRepository(db *gorm.DB) *custom.RepositoryImpl {
	return custom.NewRepository(db)
}
