package routers

import (
	"backend/api"
	"backend/conf"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CommonRouter(r *gin.Engine) *gin.Engine {

	//api.Login(r)
	r.Use(LocalRecover())
	r.POST("/Login", AfterHook(api.Login))
	r.POST("/Register", AfterHook(api.Register))

	LocalStatic("/static", conf.Config.StaticPath, r)

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusFound, gin.H{"message": "notFound"})
	})
	return r
}
