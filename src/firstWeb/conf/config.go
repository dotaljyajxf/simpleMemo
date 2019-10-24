package conf

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

var Config = &ConfigIni{}

type ConfigIni struct {
	RunMode      string        `ini:RUN_MODE`
	PageSize     int           `ini:PAGE_SIZE`
	JwtSecret    string        `ini:JWT_SECRET`
	HTTPPort     int           `ini:HTTP_PORT`
	ReadTimeout  time.Duration `ini:READ_TIMEOUT`
	WriteTimeout time.Duration `ini:WRITE_TIMEOUT`
	DBUser       string        `ini:USER`
	DBPassWord   string        `ini:PASSWORD`
	DBHost       string        `ini:HOST`
	DBName       string        `ini:NAME`
	TablePrfix   string        `ini:TABLE_PREFIX`
}

func init() {
	Cfg, err := ini.Load("conf/app.ini")
	if err != nil {
		log.Fatal(2, "Fail to parse 'conf/app.ini': %v", err)
	}

	mode := Cfg.Section("").Key("RUN_MODE").String()

	if mode == "debug" {
		err = Cfg.Section(mode).MapTo(Config)
		if err != nil {
			log.Fatal(2, "Fail to map config : %v", err)
		}
	}

	Config.ReadTimeout = Config.ReadTimeout * time.Second
	Config.WriteTimeout = Config.ReadTimeout * time.Second
}
