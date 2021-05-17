package dao

import (
	"crypto/sha256"
	"fmt"
	"math/rand"
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
	return fmt.Sprint(sha256.Sum256([]byte(time.Now().String() + fmt.Sprint(rand.Uint64()))))
}

func (dao CookieDao) SetCookie(username string) error {
	userid := time.UTC.String() + dao.generateUserid() + fmt.Sprint(sha256.Sum256([]byte(username)))
	err := dao.db.Redis.Set(userid, username, ValidTimeForCookie).Err()
	return err
}

func (dao CookieDao) UpdateCookie(userid string) error {
	username, err := dao.GetUsername(userid)
	if err != nil {
		return err
	}
	return dao.db.Redis.Set(userid, username, ValidTimeForCookie).Err()
}
