package store

import (
	"net/http"
	storeService "petstore/internal/service/store"
	"petstore/internal/transport/http/in/store/custom"
)

type Controller interface {
	// GetInventory - Summary: Returns pet inventories by status. Description: Returns a map of status codes to quantities
	GetInventory(
		w http.ResponseWriter, r *http.Request)
	// PlaceOrder - Summary: Place an order for a pet. Description: Place a new order in the store
	PlaceOrder(
		w http.ResponseWriter, r *http.Request)
	// GetOrderById - Summary: Find purchase order by ID. Description: For valid response try integer IDs with value <= 5 or > 10. Other values will generate exceptions.
	GetOrderById(
		w http.ResponseWriter, r *http.Request)
	// DeleteOrder - Summary: Delete purchase order by ID. Description: For valid response try integer IDs with value < 1000. Anything above 1000 or nonintegers will generate API errors
	DeleteOrder(
		w http.ResponseWriter, r *http.Request)
}

func NewController(service storeService.Service) *custom.ControllerImpl {
	return custom.NewController(service)
}
