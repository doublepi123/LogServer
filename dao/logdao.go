package dao

import (
	"LogServer/entity"
	"fmt"
)

type LogDao struct {
	dao
}

func (dao LogDao) GetItems() (*[]entity.ItemReturn, error) {
	var items []entity.ItemReturn
	err := dao.db.DB.Model(&entity.Item{}).Find(&items).Error
	return &items, err
}

func (dao LogDao) GetLog(level string) (*[]entity.LogReturn, error) {
	var logs []entity.LogReturn
	err := dao.db.DB.Raw("SELECT * FROM " + level).Scan(&logs).Error
	return &logs, err
}

func (dao LogDao) Select(name string, from int, to int) (*[]entity.LogReturn, error) {
	var logs []entity.LogReturn
	fmt.Println(name, from, to)
	err := dao.db.DB.Raw("SELECT * FROM " + name + " LIMIT " + fmt.Sprint(from) + "," + fmt.Sprint(to)).Scan(&logs).Error
	return &logs, err
}

func (dao LogDao) Count(name string) (int, error) {
	var logs []entity.LogReturn
	err := dao.db.DB.Raw("SELECT * FROM " + name).Scan(&logs).Error
	count := len(logs)
	return count, err
}
