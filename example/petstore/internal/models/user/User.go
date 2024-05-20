package user

type UserDTO struct {
	Id         int64  `json:"id"`
	Username   string `json:"username"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Phone      string `json:"phone"`
	UserStatus int32  `json:"userStatus"` // User Status

}

type UserEntity struct {
	Id         int64  `gorm:"type:integer"`
	Username   string `gorm:"type:varchar"`
	FirstName  string `gorm:"type:varchar"`
	LastName   string `gorm:"type:varchar"`
	Email      string `gorm:"type:varchar"`
	Password   string `gorm:"type:varchar"`
	Phone      string `gorm:"type:varchar"`
	UserStatus int32  `gorm:"type:integer"`
}

func (UserEntity) TableName() string {
	return "public.user_dtos"
}
