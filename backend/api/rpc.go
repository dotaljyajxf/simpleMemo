package api

import (
	"backend/model"
	"backend/proto/pb"
	"backend/util"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
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
func DoRpc(c *gin.Context, ret *pb.TAppRet) error {
	var call = NewRpcCall()
	if err := c.ShouldBindJSON(call); err != nil {
		return util.MakeErrRet(ret, http.StatusBadRequest, err.Error())
	}
	logrus.Infof("recive  method: %v", call)
	resp, err := model.DoRpcMethod(call.Method, call.Args)
	if err != nil {
		return util.MakeErrRet(ret, http.StatusBadRequest, err.Error())
	}
	respPoolObj, ok := resp.(util.PoolObj)
	if !ok {
		return util.MakeErrRet(ret, http.StatusBadRequest, "respNotPoolObj")
	}

	logrus.Infof("response method: %s ret: %v", call.Method, ret)
	call.Put()

	return util.MakeSuccessRet(ret, http.StatusBadRequest, respPoolObj)
}
