package main

import (
	"LogServer/dao"
	"LogServer/service"
	"LogServer/util"
)

func main() {
	db := &util.Database{}
	db.Init()
	cookieDao := &dao.CookieDao{}
	cookieDao.Init(db)
	logDao := &dao.LogDao{}
	logDao.Init(db)
	userdao := &dao.UserDao{}
	userdao.Init(db)
	messagedao := &dao.MessageDao{}
	messagedao.Init(db)
	server := service.LogServer{}
	server.Init(logDao, cookieDao, userdao, messagedao)
	server.ListenAndServer()
	util.PauseForRun()
}
