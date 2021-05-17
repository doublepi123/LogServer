package dao

import "LogServer/util"

type dao struct {
	db *util.Database
}

func (dao *dao) Init(db *util.Database) {
	dao.db = db
}
