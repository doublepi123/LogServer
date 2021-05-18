package dao

import "LogServer/entity"

type LogDao struct {
	dao
}

func (dao LogDao) GetItems() (*[]entity.ItemReturn, error) {
	var items []entity.ItemReturn
	err := dao.db.DB.Model(&entity.Item{}).Find(&items).Error
	return &items, err
}

func (dao LogDao) GetLog(level string) (*[]entity.LogEntity, error) {
	var logs []entity.LogEntity
	err := dao.db.DB.Raw("SELECT * FROM " + level).Scan(&logs).Error
	return &logs, err
}
