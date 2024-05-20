package tag

import ()

type TagDTO struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type TagEntity struct {
	Id   int64  `gorm:"type:integer"`
	Name string `gorm:"type:varchar"`
}

func (TagEntity) TableName() string {
	return "public.tag_dtos"
}
