package main

import (
	"context"
	"firstWeb/conf"
	"firstWeb/data"
	"firstWeb/routers"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	flag.Parse()

	confPath := flag.String("c", "../conf/app.ini", "set confFile path")

	conf.Init(*confPath)
	data.Init()

	if conf.Config.HTTPPort < 1 || conf.Config.HTTPPort > 65535 {
		logrus.Fatal("server port must be a number between 1 and 65535")
	}

	gin.DefaultWriter = logrus.StandardLogger().Out

	r := gin.Default()

	gin.SetMode(conf.Config.RunMode)

	r.LoadHTMLGlob("./views/*")

	r = routers.InitRouter(r)
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", conf.Config.HTTPPort),
		Handler:        r,
		ReadTimeout:    conf.Config.ReadTimeout,
		WriteTimeout:   conf.Config.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := s.ListenAndServe(); err != nil {
			logrus.Infof("Listen: %s", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	logrus.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		logrus.Infof("Server Shutdown:%s", err)
	}
	data.CloseDB()
	logrus.Info("Server exiting")
}
