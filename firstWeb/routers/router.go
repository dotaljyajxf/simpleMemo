package routers

import (
	"firstWeb/api"
	"firstWeb/conf"
	"firstWeb/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter(r *gin.Engine) *gin.Engine {

	api.Regist(r)
	api.Login(r)
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
		api.DoRpc(r)
		api.GetArticles(groupRouter)
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusFound, gin.H{"message": "notFound"})
	})
	return r
}
