package util

import (
	"fmt"
	"github.com/go-redis/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

type Database struct {
	DB    *gorm.DB
	Redis *redis.Client
}

func (db *Database) Init() {
	var err error
	db.DB, err = gorm.Open(mysql.Open(os.Getenv("DSN")), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}
	db.Redis = redis.NewClient(&redis.Options{
		Addr:     "localhost:" + os.Getenv("redis_port"),
		Password: os.Getenv("redis_pass"), // no password set
		DB:       0,                       // use default DB
	})

}
