package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"nku-treehole-server/config"
	"nku-treehole-server/pkg/logger"
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
		UserName: config.Conf.GetString("db.username"),
		Password: config.Conf.GetString("db.password"),
		Addr:     config.Conf.GetString("db.addr"),
		Name:     config.Conf.GetString("db.name"),
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbConf.UserName, dbConf.Password, dbConf.Addr, dbConf.Name)
	_db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = _db
	logger.Infof("DB Connect success")
}

func GetDBConn() *gorm.DB {
	if config.Conf.GetBool("debug") {
		return db.Debug()
	}
	return db
}
