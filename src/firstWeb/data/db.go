package data

import (
	"firstWeb/conf"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

//user:password@tcp(localhost)/dbname?charset=utf8&parseTime=True&loc=Local

var Db *gorm.DB

func Init() {
	dbSrc := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		conf.Config.DBUser, conf.Config.DBPassWord, conf.Config.DBHost, conf.Config.DBName)

	Db, err := gorm.Open("mysql", dbSrc)
	if err != nil {
		logrus.Fatalf("open db error src %s,%s\n", dbSrc, err)
		return
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return conf.Config.TablePrfix + defaultTableName
	}

	Db.SingularTable(true)
	Db.DB().SetMaxIdleConns(10)
	Db.DB().SetMaxOpenConns(100)

	Db.AutoMigrate(DbMap)
}

func CloseDB() {
	if err := Db.Close(); err != nil {
		logrus.Errorf("could not close database connection: %s", err)
	} else {
		logrus.Info("closed database connection")
	}
}
