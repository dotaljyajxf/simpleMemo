package routers

import (
	"backend/api"
	"backend/util"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) *gin.Engine {
	groupRouter := r.Group("/api/v1")
	groupRouter.Use(util.JWT())

	groupRouter.POST("/memoList", AfterHook(api.GetMemo))
	groupRouter.POST("/createMemo", AfterHook(api.CreateMemo))
	return r
}
