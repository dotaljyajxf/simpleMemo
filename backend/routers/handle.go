package routers

import (
	"backend/proto/pb"
	"fmt"
	"runtime"
	"strings"

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
				logrus.Error("panic", recv, "stack", strings.Join(stack, " "))
				//SetHttpStatus(c, http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}

func StackInfo() []string {
	var pc [8]uintptr
	sep := "/app/"
	data := make([]string, 0, 8)
	n := runtime.Callers(5, pc[:])
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
		data = append(data, fmt.Sprintf("(%s:%d)", file, line))
	}
	return data
}
