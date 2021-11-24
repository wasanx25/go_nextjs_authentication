package main

import (
	"github.com/pelletier/go-toml"
	"github.com/wasanx25/go_nextjs_authentication/backend/config"
	"github.com/wasanx25/go_nextjs_authentication/backend/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	conf := config.Config{}
	tree, err := toml.LoadFile("config.toml")
	if err != nil {
		panic(err)
	}

	err = tree.Unmarshal(&conf)
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(mysql.Open(conf.DatabaseDSN), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&model.User{}, &model.Post{}, &model.Follow{})
	if err != nil {
		panic(err)
	}
}
