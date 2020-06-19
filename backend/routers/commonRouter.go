package routers

import (
	"backend/api"
	"backend/conf"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CommonRouter(r *gin.Engine) *gin.Engine {

	//api.Login(r)
	r.POST("/Login", AfterHook(api.Login))
	r.POST("/Regist", AfterHook(api.Regist))
	r.POST("/doRpc", AfterHook(api.DoRpc))

	LocalStatic("/static", conf.Config.StaticPath, r)

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusFound, gin.H{"message": "notFound"})
	})
	return r
}
