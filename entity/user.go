package entity

import "gorm.io/gorm"

type UserEntity struct {
	gorm.Model
	Username string
	Password string
}
