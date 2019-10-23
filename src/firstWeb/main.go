package main

import (
	"context"
	"firstWeb/conf"
	"firstWeb/routers"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
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
			log.Printf("Listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	log.Println("Server exiting")
}
