package service

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"neko/src/main/cache"
	"neko/src/main/dao"
	"neko/src/main/dto"
	"neko/src/main/model"
)

func Login(userDto dto.UserDto) (*model.User,error) {
	var user model.User
	get,err := cache.Get("neko")
	if err !=nil {
		fmt.Printf("redis 异常 :%s",err.Error())
		return nil, err
	}
	if get !=nil {
		return &user,nil
	}
	err = dao.Db().Where(&model.User{UserName: userDto.UserName}).First(&user).Error
	if errors.Is(err,gorm.ErrRecordNotFound) {
		user  = user
	}
	cache.Set("neko",nil,600)
	return &user,err
}
