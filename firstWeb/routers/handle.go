package routers

import (
	"firstWeb/proto/pb"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type HandleFunc func(c *gin.Context, ret *pb.TAppRet) error

func AfterHook(f HandleFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		ret := pb.NewTAppRet()
		err := f(c, ret)
		if err != nil {
			logrus.Error("server error : %s", err.Error())
		}
		c.ProtoBuf(int(ret.Code), ret)
		ret.Put()
	}
}
