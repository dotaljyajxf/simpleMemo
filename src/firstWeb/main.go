package main

import (
	"context"
	"firstWeb/conf"
	"firstWeb/data"
	"firstWeb/routers"
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	conf.Init()
	data.Init()

	if conf.Config.HTTPPort < 1 || conf.Config.HTTPPort > 65535 {
		logrus.Fatal("server port must be a number between 1 and 65535")
	}

	r := routers.InitRouter()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", conf.Config.HTTPPort),
		Handler:        r,
		ReadTimeout:    conf.Config.ReadTimeout,
		WriteTimeout:   conf.Config.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := s.ListenAndServe(); err != nil {
			logrus.Infof("Listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	logrus.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		logrus.Fatalf("Server Shutdown:%s\n", err)
	}
	data.CloseDB()
	logrus.Info("Server exiting")
}
