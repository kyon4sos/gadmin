package model

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"neko/src/main/dao"
	"time"
)

type User struct {
	ID         uint
	Username   string    `json:"username"`
	Password   string    `json:"password"`
	Name       string    `json:"name"`
	Email      *string   `json:"email"`
	Age        uint8     `json:"age"`
	appid      string    `json:"appid"`
	Openid     string    `json:"openid"`
	Unionid    string    `json:"unionid"`
	SessionKey string    `json:"session_key"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

var cachePrefix = "user"

func (user *User) GetCachePrefix() string {
	return "user"
}
func (user *User) SetCachePrefix(name string) {
	cachePrefix = name
}
func (user *User) FindUserByOpenId(opendId string) (*User, error) {
	err := dao.Db().Where("openid = ?", opendId).First(&user)
	if err.Error == nil {
		return user, nil
	}
	if errors.Is(err.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return nil, err.Error
}
func (user *User) CreateUser(userinfo *User) (*gorm.DB, error) {
	db := dao.Db().Create(userinfo)
	if db.Error != nil {
		fmt.Printf("db error %s\n", db.Error)
		return db, db.Error
	}
	return db, nil
}
