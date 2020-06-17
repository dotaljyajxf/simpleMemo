package routers

import (
	"backend/api"
	"backend/conf"
	"backend/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) *gin.Engine {

	//api.Login(r)
	r.POST("/Login", AfterHook(api.Login))
	r.POST("/Regist", AfterHook(api.Regist))
	r.POST("/doRpc", AfterHook(api.DoRpc))

	// Favicon
	//r.StaticFile("/favicon.ico", conf.HttpFaviconsPath()+"/favicon.ico")

	// Static assets like js and css files
	//r.Static("/static", conf.Config.StaticPath)
	LocalStatic("/static", conf.Config.StaticPath, r)

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	groupRouter := r.Group("/api/v1")
	groupRouter.Use(util.JWT())
	{
		groupRouter.POST("/memoList", AfterHook(api.GetMemo))
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusFound, gin.H{"message": "notFound"})
	})
	return r
}
