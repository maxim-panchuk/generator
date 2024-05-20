package generated

import (
	"context"
	storeRepository "petstore/internal/database/repositories/store"
	"petstore/internal/models/order"
)

type ServiceImpl struct {
	repo storeRepository.Repository
}

func NewService(repo storeRepository.Repository) *ServiceImpl {
	return &ServiceImpl{
		repo: repo,
	}
}

// GetInventory - Summary: Returns pet inventories by status. Description: Returns a map of status codes to quantities
func (s *ServiceImpl) GetInventory(
	ctx context.Context,
) error {
	panic("not implemented")
}

// PlaceOrder - Summary: Place an order for a pet. Description: Place a new order in the store
func (s *ServiceImpl) PlaceOrder(
	ctx context.Context,
	orderDTO *order.OrderDTO,
) (*order.OrderDTO, error) {
	dto, err := s.repo.PlaceOrder(
		ctx, orderDTO,
	)
	if err != nil {
		return nil, err
	}
	return dto, nil

}

// GetOrderById - Summary: Find purchase order by ID. Description: For valid response try integer IDs with value <= 5 or > 10. Other values will generate exceptions.
func (s *ServiceImpl) GetOrderById(
	ctx context.Context,
	orderId int64, // Required: true, Description: ID of order that needs to be fetched
) (*order.OrderDTO, error) {
	dto, err := s.repo.GetOrderById(
		ctx, orderId,
	)
	if err != nil {
		return nil, err
	}
	return dto, nil

}

// DeleteOrder - Summary: Delete purchase order by ID. Description: For valid response try integer IDs with value < 1000. Anything above 1000 or nonintegers will generate API errors
func (s *ServiceImpl) DeleteOrder(
	ctx context.Context,
	orderId int64, // Required: true, Description: ID of the order that needs to be deleted
) error {
	if err := s.repo.DeleteOrder(
		ctx, orderId,
	); err != nil {
		return err
	}
	return nil

}
