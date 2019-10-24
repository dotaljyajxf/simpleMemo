package routers

import (
	"firstWeb/api"
	"firstWeb/conf"
	"firstWeb/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger(), gin.Recovery())

	gin.SetMode(conf.Config.RunMode)

	api.GetAuth(r)
	// Favicon
	//r.StaticFile("/favicon.ico", conf.HttpFaviconsPath()+"/favicon.ico")

	// Static assets like js and css files
	//r.Static("/static", conf.HttpStaticPath())

	groupRouter := r.Group("/api/v1")
	groupRouter.Use(util.JWT())
	{
		api.GetArticles(groupRouter)
	}

	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", nil)
	})
	return r
}
