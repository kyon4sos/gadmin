package dto

type UserDto struct {
	UserName string `form:"username" json:"username" binding:"required"`
	PassWord string `form:"password" json:"password" binding:"required"`
}

