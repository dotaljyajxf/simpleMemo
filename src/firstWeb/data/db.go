package data

import (
	"firstWeb/conf"
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
)

//user:password@tcp(localhost)/dbname?charset=utf8&parseTime=True&loc=Local

var Db *gorm.DB

func init() {
	dbSrc := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		conf.Config.DBUser, conf.Config.DBPassWord, conf.Config.DBHost, conf.Config.DBName)

	Db, err := gorm.Open("mysql", dbSrc)
	if err != nil {
		log.Fatal("open db error src:%s", dbSrc)
		return
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return conf.Config.TablePrfix + defaultTableName
	}

	Db.SingularTable(true)
	Db.DB().SetMaxIdleConns(10)
	Db.DB().SetMaxOpenConns(100)
}

func CloseDB() {
	Db.Close()
}
