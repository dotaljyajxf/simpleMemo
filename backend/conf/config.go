package conf

import (
	"flag"
	"log"
	"os"
	"time"

	"github.com/go-ini/ini"
	"github.com/sirupsen/logrus"
)

var Config = &ConfigIni{}

type ConfigIni struct {
	JwtSecret          string        `ini:"JWT_SECRET"`
	HTTPPort           int           `ini:"HTTP_PORT"`
	ReadTimeout        time.Duration `ini:"READ_TIMEOUT"`
	WriteTimeout       time.Duration `ini:"WRITE_TIMEOUT"`
	DBUser             string        `ini:"USER"`
	DBPassWord         string        `ini:"PASSWORD"`
	MasterHost         string        `ini:"MASTER_HOST"`
	SlaveHost          string        `ini:"SLAVE_HOST"`
	DBName             string        `ini:"DB_NAME"`
	LogPath            string        `ini:"LOG_PATH"`
	LogLevel           int           `ini:"LOG_LEVEL"`
	KvRedisHost        string        `ini:"KV_REDIS_HOST"`
	KvRedisPassWd      string        `ini:"KV_REDIS_PW"`
	KvRedisDB          int           `ini:"KV_DB_NUMBER" `
	KvRedisMaxIdel     int           `ini:"KV_MAX_IDEL"`
	KvRedisMaxActive   int           `ini:"KV_MAX_ACTIVE"`
	KvRedisIdelTimeout int           `ini:"KV_IDEL_TIMEOUT"`
	ConfPath           string        `ini:"-"`
	ViewPath           string        `ini:"-"`
	StaticPath         string        `ini:"-"`
	RunMode            string        `ini:"-"`
}

func handleCmdFlag() {
	flag.StringVar(&Config.ConfPath, "c", "./conf/app.ini", "set confFile path")
	flag.StringVar(&Config.ViewPath, "v", "./views", "set view path")
	flag.StringVar(&Config.StaticPath, "s", "./static", "set static path")
	flag.StringVar(&Config.RunMode, "m", "debug", "run mode")
	flag.Parse()
}

func Init() {
	handleCmdFlag()
	Cfg, err := ini.Load(Config.ConfPath)
	if err != nil {
		log.Fatal(2, "Fail to parse 'conf/app.ini': %v", err)
	}

	//TODO mode 来自命令行
	mode := Cfg.Section("").Key(Config.RunMode).String()
	if mode == "debug" {
		err = Cfg.Section(mode).MapTo(Config)
		if err != nil {
			log.Fatal(2, "Fail to map config : %v", err)
			return
		}
	}

	//logrus.SetFormatter(&logrus.TextFormatter{
	//	DisableColors:   false,
	//	FullTimestamp:   true,
	//	TimestampFormat: "2006-01-02 15:04:05.999999999 -0700 MST",
	//})
	logrus.SetFormatter(&MyFormatter{})

	logrus.SetReportCaller(true)
	log.Println(Config)

	logFile, err := os.OpenFile(Config.LogPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(2, " open LogFile error : %s,%v", Config.LogPath, err)
		return
	}

	logrus.SetOutput(logFile)
	logrus.SetLevel(logrus.Level(Config.LogLevel))
}
