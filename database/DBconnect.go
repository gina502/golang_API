package database

import (
	"log"

	"github.com/go-redis/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DBconnect *gorm.DB
var err error
var rdb *redis.Client

func DD() {
	dsn := "root:root@tcp(127.0.0.1:8889)/Golang_GinTest?charset=utf8mb4&parseTime=True&loc=Local"
	DBconnect, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
}

func RR() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	_, err = rdb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}
