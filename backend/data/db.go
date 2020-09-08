package data

import (
	"backend/conf"
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

	if err = Manager.Master.Ping(); err != nil {
		logrus.Fatalf("ping master db err %s: %s\n", masterDns, err)
		return
	}
	if err = Manager.Slave.Ping(); err != nil {
		logrus.Fatalf("ping slave db err %s: %s\n", masterDns, err)
		return
	}

}

func CloseDB() {
	if err := Manager.Master.Close(); err != nil {
		logrus.Errorf("could not close master database connection: %s", err)
	} else {
		logrus.Info("closed master database connection")
	}
	if Manager.Slave != Manager.Master {
		if err := Manager.Slave.Close(); err != nil {
			logrus.Errorf("could not close slave database connection: %s", err)
		} else {
			logrus.Info("closed slave database connection")
		}
	}

	if conf.Config.CacheUse == 1 {
		Manager.Cache.Close()
	}
}
