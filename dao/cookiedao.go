package dao

import (
	"LogServer/util"
	"time"
)

const ValidTimeForCookie = time.Minute * 15

type CookieDao struct {
	dao
}

func (dao CookieDao) GetUsername(userid string) (string, error) {
	username, err := dao.db.Redis.Get(userid).Result()
	if err != nil {
		return "", err
	}
	return username, err
}
func (dao CookieDao) generateUserid() string {
	return util.GetPWD(string(time.Now().Unix()))
}

func (dao CookieDao) SetCookie(username string) (string, error) {
	userid := dao.generateUserid() + util.GetPWD(username)
	err := dao.db.Redis.Set(userid, username, ValidTimeForCookie).Err()
	return userid, err
}

func (dao CookieDao) UpdateCookie(userid string) error {
	username, err := dao.GetUsername(userid)
	if err != nil {
		return err
	}
	return dao.db.Redis.Set(userid, username, ValidTimeForCookie).Err()
}
