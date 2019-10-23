package routers

import (
	"firstWeb/api"
	"firstWeb/conf"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(conf.Config.RunMode)

	groupRouter := r.Group("/api/v1")
	{
		api.GetArticles(groupRouter)
	}

	return r
}
