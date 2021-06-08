package dao

import (
	"LogServer/entity"
	"fmt"
	"strings"
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

func (dao LogDao) CountPic() (int64, error) {
	var count int64
	err := dao.db.DB.Model(&entity.PictureEntity{}).Count(&count).Error
	return count, err
}

func (dao LogDao) SelectPic(from int, to int) (*[]entity.PictureEntity, error) {
	var pics []entity.PictureEntity
	err := dao.db.DB.Raw("SELECT * FROM pitcute_entities LIMIT " + fmt.Sprint(from) + " , " + fmt.Sprint(to-from)).Scan(&pics).Error
	return &pics, err
}

func (dao LogDao) GetRecentInfo() (*[]entity.LogReturn, error) {
	t, err := dao.GetItems()
	if err != nil {
		return nil, err
	}
	logs := *t
	l := 0
	for i := range logs {
		if strings.Contains(strings.ToUpper(logs[i].Name), "INFO") {
			l = i
		}
	}
	var ans []entity.LogReturn
	err = dao.db.DB.Raw("SELECT * FROM " + logs[l].Name + " order by id desc limit 10").Find(&ans).Error
	if err != nil {
		return nil, err
	}
	return &ans, err
}

func (dao LogDao) GetRecentError() (*[]entity.LogReturn, error) {
	t, err := dao.GetItems()
	if err != nil {
		return nil, err
	}
	logs := *t
	l := 0
	for i := range logs {
		if strings.Contains(strings.ToUpper(logs[i].Name), "ERROR") {
			l = i
		}
	}
	var ans []entity.LogReturn
	err = dao.db.DB.Raw("SELECT * FROM " + logs[l].Name + " order by id desc limit 10").Find(&ans).Error
	if err != nil {
		return nil, err
	}
	return &ans, err
}

func (dao LogDao) Select(name string, from int, to int) (*[]entity.LogReturn, error) {
	var logs []entity.LogReturn
	fmt.Println(name, from, to)
	err := dao.db.DB.Raw("SELECT * FROM " + name + " LIMIT " + fmt.Sprint(from) + "," + fmt.Sprint(to-from)).Scan(&logs).Error
	return &logs, err
}

func (dao LogDao) Count(name string) (int, error) {
	var logs []entity.LogReturn
	err := dao.db.DB.Raw("SELECT * FROM " + name).Scan(&logs).Error
	count := len(logs)
	return count, err
}
