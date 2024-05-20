package customer

import (
	"petstore/internal/models/address"
)

type CustomerDTO struct {
	Id       int64                 `json:"id"`
	Username string                `json:"username"`
	Address  []*address.AddressDTO `json:"address"`
}

type CustomerEntity struct {
	Id       int64                    `gorm:"type:integer"`
	Username string                   `gorm:"type:varchar"`
	Address  []*address.AddressEntity `gorm:"many2many:customer_address"`
}

func (CustomerEntity) TableName() string {
	return "public.customer_dtos"
}
