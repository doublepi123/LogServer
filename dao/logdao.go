package dao

type LogDao struct {
	dao
}

func (dao LogDao) GetItems() (string, error) {
	items, err := dao.db.Redis.Get("items").Result()
	return items, err
}
