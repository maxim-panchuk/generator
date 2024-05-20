package order

import (
	"time"
)

type Status string

const (
	PLACED    Status = "placed"
	APPROVED  Status = "approved"
	DELIVERED Status = "delivered"
)

type OrderDTO struct {
	Id       int64     `json:"id"`
	PetId    int64     `json:"petId"`
	Quantity int32     `json:"quantity"`
	ShipDate time.Time `json:"shipDate"`
	Status   string    `json:"status"` // Order Status
	Complete bool      `json:"complete"`
}

type OrderEntity struct {
	Id       int64     `gorm:"type:integer"`
	PetId    int64     `gorm:"type:integer"`
	Quantity int32     `gorm:"type:integer"`
	ShipDate time.Time `gorm:"type:timestamp"`
	Status   string    `gorm:"type:varchar"`
	Complete bool      `gorm:"type:bool"`
}

func (OrderEntity) TableName() string {
	return "public.order_dtos"
}
