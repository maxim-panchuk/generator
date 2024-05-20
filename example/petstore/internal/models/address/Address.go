package address

type AddressDTO struct {
	Id     int64  `json:"id"`
	Street string `json:"street"`
	City   string `json:"city"`
	State  string `json:"state"`
	Zip    string `json:"zip"`
}

type AddressEntity struct {
	Id     int64  `gorm:"type:integer"`
	Street string `gorm:"type:varchar"`
	City   string `gorm:"type:varchar"`
	State  string `gorm:"type:varchar"`
	Zip    string `gorm:"type:varchar"`
}

func (AddressEntity) TableName() string {
	return "public.address_dtos"
}
