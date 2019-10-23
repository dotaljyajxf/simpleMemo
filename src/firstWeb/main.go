package main

import (
	"firstWeb/conf"
	"firstWeb/routers"
	"fmt"
	"net/http"
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

	s.ListenAndServe()
}
