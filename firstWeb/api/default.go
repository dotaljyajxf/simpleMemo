package api

import "github.com/gin-gonic/gin"

func GetArticles(router *gin.RouterGroup) {
	router.GET("/ljy", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "HelloWorld"})
	})
}
