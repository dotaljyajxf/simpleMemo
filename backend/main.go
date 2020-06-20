package main

import (
	"backend/conf"
	"backend/routers"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {

	conf.Init()
	//data.Init()

	if conf.Config.HTTPPort < 1 || conf.Config.HTTPPort > 65535 {
		logrus.Fatal("server port must be a number between 1 and 65535")
	}

	gin.DefaultWriter = logrus.StandardLogger().Out

	r := gin.Default()

	store := cookie.NewStore([]byte(conf.Config.JwtSecret))
	r.Use(sessions.Sessions("mysession", store))

	gin.SetMode(conf.Config.RunMode)

	//r.LoadHTMLGlob(conf.Config.StaticPath + "/views/*")
	r.LoadHTMLFiles(conf.Config.StaticPath + "/index.html")

	r = routers.CommonRouter(r)
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

	logrus.Info("Shutdown Server Start...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		logrus.Infof("Server Shutdown:%s", err)
	}
	//data.CloseDB()
	logrus.Info("Server exiting")
}
