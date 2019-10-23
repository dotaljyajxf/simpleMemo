package routers

import (
	"firstWeb/api"
	"firstWeb/conf"
	"firstWeb/util"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(conf.Config.RunMode)

	api.GetAuth(r)

	groupRouter := r.Group("/api/v1")
	groupRouter.Use(util.JWT())
	{
		api.GetArticles(groupRouter)
	}

	return r
}
