package appret

import (
	"backend/proto/pb"

	"github.com/sirupsen/logrus"

	"github.com/golang/protobuf/proto"
)

type PoolObj interface {
	Put()
}

func MakeErrRet(code int64, msg string) *pb.TAppRet {
	ret := pb.NewTAppRet()
	ret.Msg = msg
	ret.Code = code
	ret.Data = nil
	return ret
}

func MakeSuccessRet(code int64, resp interface{}) *pb.TAppRet {
	r := pb.NewTAppRet()
	logrus.Infof("Response: %v\n", resp)
	if rpObj, ok := resp.(PoolObj); ok {
		defer rpObj.Put()
	}

	r.Code = code
	pMsg, ok := resp.(proto.Message)
	if !ok {
		r.Msg = "resp Type error"
		return r
	}
	d, err := proto.Marshal(pMsg)
	if err != nil {
		r.Msg = err.Error()
		return r
	}
	r.Data = d
	return nil
}
