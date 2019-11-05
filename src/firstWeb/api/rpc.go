package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type RpcType struct {
	Method string `form:"method"  binding:"required"`
	Args   []byte `form:"args"    binding:"required"`
}

func DoMethod(router *gin.RouterGroup) {
	router.POST("/doRpc", func(c *gin.Context) {
		var call RpcType
		if err := c.ShouldBindJSON(&call); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	})
}
