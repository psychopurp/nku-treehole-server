package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"nku-treehole-server/config"
)

var (
	db *gorm.DB
)

type DBConfig struct {
	UserName string
	Password string
	Addr     string
	Name     string
}

func InitDB() {
	dbConf := &DBConfig{
		UserName: fmt.Sprint(config.Conf.Get("db.username")),
		Password: fmt.Sprint(config.Conf.Get("db.password")),
		Addr:     fmt.Sprint(config.Conf.Get("db.addr")),
		Name:     fmt.Sprint(config.Conf.Get("db.name")),
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbConf.UserName, dbConf.Password, dbConf.Addr, dbConf.Name)
	_db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = _db
	fmt.Println("success")
}

func GetDBConn() *gorm.DB {
	return db
}
