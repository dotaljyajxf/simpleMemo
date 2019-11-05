package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
)

type RpcType struct {
	Method string `form:"method"  binding:"required"`
	Args   []byte `form:"args"    binding:"required"`
}

var poolRpcCall = sync.Pool{
	New: func() interface{} {
		return new(RpcType)
	},
}
var dummyRpcType RpcType

func NewRpcCall() *RpcType {
	obj := poolRpcCall.Get().(*RpcType)
	return obj
}

func (r *RpcType) Put() {
	*r = dummyRpcType
	poolRpcCall.Put(r)
}

//auth.GetInfo  XXXXXNNJMH
func DoRpc(router *gin.RouterGroup) {
	router.POST("/doRpc", func(c *gin.Context) {
		var call = NewRpcCall()
		if err := c.ShouldBindJSON(call); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	})
}
