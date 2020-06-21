package data

import (
	"backend/conf"
	"backend/data/cache"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
)

func InitDataManager() {
	masterDns := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		conf.Config.DBUser, conf.Config.DBPassWord, conf.Config.MasterHost, conf.Config.DBName)
	var err error
	Manager.Master, err = sql.Open("mysql", masterDns)
	if err != nil {
		logrus.Fatalf("open master db error src %s,%s\n", masterDns, err)
		return
	}
	Manager.Slave = Manager.Master
	if conf.Config.MasterHost != conf.Config.SlaveHost {
		slaveDns := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
			conf.Config.DBUser, conf.Config.DBPassWord, conf.Config.SlaveHost, conf.Config.DBName)
		Manager.Master, err = sql.Open("mysql", slaveDns)
		if err != nil {
			logrus.Fatalf("open slave db error src %s,%s\n", masterDns, err)
			return
		}
	}

	if conf.Config.CacheUse == 1 {
		Manager.Cache = cache.InitDbCache()
	} else {
		Manager.Cache = &cache.DbCache{}
	}

	if err = Manager.Master.Ping(); err != nil {
		panic(err)
	}
	if err = Manager.Slave.Ping(); err != nil {
		panic(err)
	}

}

//user:password@tcp(localhost)/dbname?charset=utf8&parseTime=True&loc=Local
//
//var Db *gorm.DB
//
//func Init() {
//	dbSrc := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
//		conf.Config.DBUser, conf.Config.DBPassWord, conf.Config.DBHost, conf.Config.DBName)
//
//	var err error
//	Db, err = gorm.Open("mysql", dbSrc)
//	if err != nil {
//		logrus.Fatalf("open db error src %s,%s\n", dbSrc, err)
//		return
//	}
//
//	//gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
//	//	return conf.Config.TablePrfix + defaultTableName
//	//}
//
//	Db.SingularTable(true)
//	Db.DB().SetMaxIdleConns(10)
//	Db.DB().SetMaxOpenConns(100)
//
//	Db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(table.DbMap...)
//
//}
//
//func CloseDB() {
//	if err := Db.Close(); err != nil {
//		logrus.Errorf("could not close database connection: %s", err)
//	} else {
//		logrus.Info("closed database connection")
//	}
//}
