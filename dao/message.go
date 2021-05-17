package dao

type MessageDao struct {
	dao
}

func (dao MessageDao) LogAdd(level string, message string) {
	dao.db.Redis.LPush(level, message, 0)
}
