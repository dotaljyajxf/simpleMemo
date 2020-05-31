package api

import (
	"firstWeb/module"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"sync"
)

type RpcType struct {
	Method string `form:"method" json:"method"  binding:"required" `
	Args   []byte `form:"args"   json:"args"   binding:"required" `
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
func DoRpc(router *gin.Engine) {
	router.POST("/doRpc", func(c *gin.Context) {
		var call = NewRpcCall()
		if err := c.ShouldBindJSON(call); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		logrus.Infof("recive  method: %v", call)
		ret, err := module.DoRpcMethod(call.Method, call.Args)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		logrus.Infof("response method: %s ret: %v", call.Method, ret)
		call.Put()

		c.ProtoBuf(http.StatusOK, ret)
	})
}
