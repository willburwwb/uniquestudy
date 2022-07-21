package model

import "github.com/jinzhu/gorm"

//user 包含名字密码和邮箱
type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
}
