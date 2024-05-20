package category

import ()

type CategoryDTO struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type CategoryEntity struct {
	Id   int64  `gorm:"type:integer"`
	Name string `gorm:"type:varchar"`
}

func (CategoryEntity) TableName() string {
	return "public.category_dtos"
}
