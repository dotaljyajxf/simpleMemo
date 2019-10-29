package conf

import (
	"flag"
	"github.com/go-ini/ini"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"time"
)

var Config = &ConfigIni{}

type ConfigIni struct {
	RunMode      string        `ini:"RUN_MODE"`
	PageSize     int           `ini:"PAGE_SIZE"`
	JwtSecret    string        `ini:"JWT_SECRET"`
	HTTPPort     int           `ini:"HTTP_PORT"`
	ReadTimeout  time.Duration `ini:"READ_TIMEOUT"`
	WriteTimeout time.Duration `ini:"WRITE_TIMEOUT"`
	DBUser       string        `ini:"USER"`
	DBPassWord   string        `ini:"PASSWORD"`
	DBHost       string        `ini:"HOST"`
	DBName       string        `ini:"DB_NAME"`
	TablePrfix   string        `ini:"TABLE_PREFIX"`
	LogPath      string        `ini:"LOG_PATH"`
	LogLevel     int           `ini:"LOG_LEVEL"`
	ConfPath     string        `ini:"-"`
	ViewPath     string        `ini:"-"`
}

func handleCmdFlag() {
	flag.Parse()
	flag.StringVar(&Config.ConfPath, "c", "./conf/app.ini", "set confFile path")
	flag.StringVar(&Config.ViewPath, "v", "../views/", "set view path")
}

func Init() {
	handleCmdFlag()
	Cfg, err := ini.Load(Config.ConfPath)
	if err != nil {
		log.Fatal(2, "Fail to parse 'conf/app.ini': %v", err)
	}

	mode := Cfg.Section("").Key("RUN_MODE").String()
	if mode == "debug" {
		err = Cfg.Section(mode).MapTo(Config)
		if err != nil {
			log.Fatal(2, "Fail to map config : %v", err)
			return
		}
	}

	Config.ReadTimeout = Config.ReadTimeout * time.Second
	Config.WriteTimeout = Config.ReadTimeout * time.Second

	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})

	log.Println(Config)

	logFile, err := os.OpenFile(Config.LogPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(2, " open LogFile error : %s,%v", Config.LogPath, err)
		return
	}

	logrus.SetOutput(logFile)
	logrus.SetLevel(logrus.Level(Config.LogLevel))
}
