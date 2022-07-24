package database

import (
	"fmt"
	"log"
	"test/model"

	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
)

var DB *gorm.DB
var Rdb *redis.Client

func InitDb() *gorm.DB {
	host := viper.GetString("postgresql.host")
	port := viper.GetString("postgresql.port")
	passwd := viper.GetString("postgresql.password")
	user := viper.GetString("postgresql.user")
	dbname := viper.GetString("postgresql.dbname")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host,
		user,
		passwd,
		dbname,
		port,
	)
	//log.Fatal(dsn)

	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		panic("connect postgres failed   " + err.Error())
	}
	DB = db
	if !db.HasTable(&model.User{}) {
		db.CreateTable(&model.User{})
	}
	if !db.HasTable(&model.Post{}) {
		db.CreateTable(&model.Post{})
	}
	if !db.HasTable(&model.Comment{}) {
		db.CreateTable(&model.Comment{})
	}
	if !db.HasTable(&model.Nick{}) {
		db.CreateTable(&model.Nick{})
	}
	log.Println("connect postgres successfully")
	return db
}
func GetDB() *gorm.DB {
	return DB
}
func InitRdb() *redis.Client {
	host := viper.GetString("redis.host")
	rdb := redis.NewClient(&redis.Options{
		Addr: host,
	})

	_, err := rdb.Ping().Result()
	if err != nil {
		panic("connect redis failed  " + err.Error())
	}
	Rdb = rdb
	log.Println("redis connect successfully")
	return rdb
}
func GetRdb() *redis.Client {
	return Rdb
}
