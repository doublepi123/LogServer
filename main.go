package main

import (
	"LogServer/dao"
	"LogServer/util"
)

func main() {
	db := &util.Database{}
	db.Init()
	cookieDao := &dao.CookieDao{}
	cookieDao.Init(db)
	logDao := &dao.LogDao{}
	logDao.Init(db)

}
