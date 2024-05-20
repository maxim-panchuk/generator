package pet

import (
	"petstore/internal/models/category"

	"petstore/internal/models/tag"
)

type Status string

const (
	AVAILABLE Status = "available"
	PENDING   Status = "pending"
	SOLD      Status = "sold"
)

type PetDTO struct {
	Id        int64                 `json:"id"`
	Name      string                `json:"name"`
	Category  *category.CategoryDTO `json:"category"`
	PhotoUrls []string              `json:"photoUrls"`
	Tags      []*tag.TagDTO         `json:"tags"`
	Status    string                `json:"status"` // pet status in the store

}

type PetEntity struct {
	Id        int64                      `gorm:"type:integer"`
	Name      string                     `gorm:"type:varchar"`
	Category  []*category.CategoryEntity `gorm:"many2many:pet_category"`
	PhotoUrls []string                   `gorm:"type:varchar"`
	Tags      []*tag.TagEntity           `gorm:"many2many:pet_tag"`
	Status    string                     `gorm:"type:varchar"`
}

func (PetEntity) TableName() string {
	return "public.pet_dtos"
}
