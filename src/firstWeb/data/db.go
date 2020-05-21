package data

import (
	"firstWeb/conf"
	"firstWeb/data/table"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
)

//user:password@tcp(localhost)/dbname?charset=utf8&parseTime=True&loc=Local

var Db *gorm.DB

func Init() {
	dbSrc := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		conf.Config.DBUser, conf.Config.DBPassWord, conf.Config.DBHost, conf.Config.DBName)

	var err error
	Db, err = gorm.Open("mysql", dbSrc)
	if err != nil {
		logrus.Fatalf("open db error src %s,%s\n", dbSrc, err)
		return
	}

	//gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
	//	return conf.Config.TablePrfix + defaultTableName
	//}

	Db.SingularTable(true)
	Db.DB().SetMaxIdleConns(10)
	Db.DB().SetMaxOpenConns(100)

	Db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(table.DbMap...)
}

func CloseDB() {
	if err := Db.Close(); err != nil {
		logrus.Errorf("could not close database connection: %s", err)
	} else {
		logrus.Info("closed database connection")
	}
}
