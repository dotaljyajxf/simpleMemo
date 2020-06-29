package api

import (
	"backend/routers"
	"backend/util"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) *gin.Engine {
	groupRouter := r.Group("/api/v1")
	groupRouter.Use(util.JWT())

	groupRouter.POST("/memoList", routers.AfterHook(GetMemo))
	groupRouter.POST("/createMemo", routers.AfterHook(CreateMemo))
	return r
}
