package main

import (
	"log"
	"test/database"
	"test/routes"

	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatal("fatal error config file: %w", err)
	}
	db := database.InitDb()
	rdb := database.InitRdb()
	defer func() {
		//db.DropTable(&model.User{})
		db.Close()
		rdb.Close()
	}()
	engine := routes.InitRouter()
	if err := engine.Run(":3000"); err != nil {
		log.Fatal("service failed", err)
	}
}
func initConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	return err
}
