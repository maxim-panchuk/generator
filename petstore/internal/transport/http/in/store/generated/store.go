package generated

import (
	"encoding/json"
	"petstore/internal/models/order"
	storeService "petstore/internal/service/store"

	"github.com/gorilla/mux"
	"strconv"

	"context"
	"net/http"
)

type ControllerImpl struct {
	service storeService.Service
}

func NewController(service storeService.Service) *ControllerImpl {
	return &ControllerImpl{service: service}
}

// GetInventory godoc
// @Summary  "Returns pet inventories by status"
// @Description  "Returns a map of status codes to quantities"
// @Tags store
// @Accept json
// @Success  200   "successful operation"
// @Router /store/inventory [get]
// @Security BearerAuth
func (c *ControllerImpl) GetInventory(
	w http.ResponseWriter, r *http.Request) {
	panic("not implemented")

}

// PlaceOrder godoc
// @Summary  "Place an order for a pet"
// @Description  "Place a new order in the store"
// @Tags store
// @Accept json
// @Param Order body order.OrderDTO true "Place a new order in the store"
// @Success  200    {object}  order.OrderDTO  "successful operation"
// @Failure  400   "Invalid input"
// @Failure  422   "Validation exception"
// @Router /store/order [post]
// @Security BearerAuth
func (c *ControllerImpl) PlaceOrder(
	w http.ResponseWriter, r *http.Request) {

	// TODO генерация, если слайс схем
	var orderDTO *order.OrderDTO
	err := json.NewDecoder(r.Body).Decode(&orderDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid input"))
		return
	}

	dto, err := c.service.PlaceOrder(
		context.Background(), orderDTO,
	)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	b, err := json.Marshal(dto)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
	return

}

// GetOrderById godoc
// @Summary  "Find purchase order by ID"
// @Description  "For valid response try integer IDs with value <= 5 or > 10. Other values will generate exceptions."
// @Tags store
// @Accept json
// @Param orderId path integer true "ID of order that needs to be fetched"
// @Success  200    {object}  order.OrderDTO  "successful operation"
// @Failure  400   "Invalid ID supplied"
// @Failure  404   "Order not found"
// @Router /store/order/{orderId} [get]
// @Security BearerAuth
func (c *ControllerImpl) GetOrderById(
	w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	orderIdIn, ok := vars["orderId"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid ID supplied"))
		return

	}

	orderIdInt, err := strconv.Atoi(orderIdIn)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("bad user input"))
		return
	}

	orderId := int64(orderIdInt)

	dto, err := c.service.GetOrderById(
		context.Background(), orderId,
	)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	b, err := json.Marshal(dto)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
	return

}

// DeleteOrder godoc
// @Summary  "Delete purchase order by ID"
// @Description  "For valid response try integer IDs with value < 1000. Anything above 1000 or nonintegers will generate API errors"
// @Tags store
// @Accept json
// @Param orderId path integer true "ID of the order that needs to be deleted"
// @Failure  400   "Invalid ID supplied"
// @Failure  404   "Order not found"
// @Router /store/order/{orderId} [delete]
// @Security BearerAuth
func (c *ControllerImpl) DeleteOrder(
	w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	orderIdIn, ok := vars["orderId"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid ID supplied"))
		return

	}

	orderIdInt, err := strconv.Atoi(orderIdIn)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("bad user input"))
		return
	}

	orderId := int64(orderIdInt)

	if err := c.service.DeleteOrder(
		context.Background(), orderId,
	); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	} else {
		w.WriteHeader(http.StatusOK)
		return
	}

}
