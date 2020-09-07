package routers

import (
	"backend/proto/pb"
	"fmt"
	"net/http"
	"runtime"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type HandleFunc func(c *gin.Context) *pb.TAppRet

func AfterHook(f HandleFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		ret := f(c)
		defer ret.Put()
		if len(ret.Msg) != 0 {
			logrus.Error("server error : %s", ret.Msg)
			c.ProtoBuf(http.StatusServiceUnavailable, ret)
		}

		c.ProtoBuf(http.StatusOK, ret)
	}
}

func LocalRecover() gin.HandlerFunc {
	return func(c *gin.Context) {
		//处理panic 未知的错误
		defer func() {
			if r := recover(); r != nil {
				var recv error
				switch r := r.(type) {
				case error:
					recv = r
				default:
					recv = fmt.Errorf("%v", r)
				}
				stack := StackInfo()
				logrus.Errorf("panic: %v, stack:\n %v", recv, strings.Join(stack, " "))
			}
		}()
		c.Next()
	}
}

func StackInfo() []string {
	var pc [8]uintptr
	sep := "backend/"
	data := make([]string, 0, 10)
	n := runtime.Callers(0, pc[:])
	for _, pc := range pc[:n] {
		fn := runtime.FuncForPC(pc)
		if fn == nil {
			continue
		}
		file, line := fn.FileLine(pc)
		if !strings.Contains(file, sep) {
			continue
		}
		ret := strings.Split(file, sep)
		file = ret[1]
		//name := fn.Name()
		data = append(data, fmt.Sprintf("(%s:%d)\n", file, line))
	}
	return data
}
