package dao

import (
	"LogServer/entity"
	"LogServer/util"
)

type UserDao struct {
	dao
}

func (dao UserDao) Check(username string, password string) bool {
	if password == "" {
		return false
	}
	user := &entity.UserEntity{}
	dao.db.DB.Where("username = ? ", username).First(user)
	return util.CmpPWD(user.Password, password)
}

func (dao UserDao) Add(username string, password string) error {
	dao.db.DB.AutoMigrate(&entity.UserEntity{})
	user := &entity.UserEntity{
		Username: username,
		Password: util.GetPWD(password),
	}
	return dao.db.DB.Create(user).Error
}
