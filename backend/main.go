package main

import (
	"backend/api"
	"backend/conf"
	"backend/data"
	"backend/data/cache"
	"backend/routers"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	_ "net/http/pprof"
)

func main() {

	conf.Init()
	data.InitDataManager()
	cache.InitKvCache()

	if conf.Config.HTTPPort < 1 || conf.Config.HTTPPort > 65535 {
		logrus.Fatal("server port must be a number between 1 and 65535")
	}

	gin.DefaultWriter = logrus.StandardLogger().Out

	r := gin.Default()

	gin.SetMode(conf.Config.RunMode)

	//r.LoadHTMLGlob(conf.Config.StaticPath + "/views/*")
	r.LoadHTMLFiles(conf.Config.StaticPath + "/index.html")

	r = routers.CommonRouter(r)
	r = api.InitRouter(r)
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", conf.Config.HTTPPort),
		Handler:        r,
		ReadTimeout:    time.Second * conf.Config.ReadTimeout,
		WriteTimeout:   time.Second * conf.Config.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := s.ListenAndServe(); err != nil {
			logrus.Infof("Listen: %s", err)
		}
	}()

	go func() {
		err := http.ListenAndServe(":9909", nil)
		if err != nil {
			panic(err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	logrus.Info("Shutdown Server Start...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		logrus.Infof("Server Shutdown:%s", err)
	}
	data.CloseDB()
	logrus.Info("Server exiting")
}
