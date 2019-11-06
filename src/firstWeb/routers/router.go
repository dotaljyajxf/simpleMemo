package routers

import (
	"firstWeb/api"
	"firstWeb/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter(r *gin.Engine) *gin.Engine {

	api.GetAuth(r)
	api.DoRpc(r)
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
		c.HTML(http.StatusOK, "index.html", nil)
	})
	return r
}
