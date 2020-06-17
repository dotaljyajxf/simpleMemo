package util

import (
	"backend/proto/pb"
	"errors"

	"github.com/golang/protobuf/proto"
)

type PoolObj interface {
	Put()
}

func MakeErrRet(r *pb.TAppRet, code int32, msg string) error {
	r.Msg = msg
	r.Code = code
	r.Data = nil
	return nil
}

func MakeSuccessRet(r *pb.TAppRet, code int32, resp interface{}) error {
	if rpObj, ok := resp.(PoolObj); ok {
		defer rpObj.Put()
	}

	r.Code = code
	pMsg, ok := resp.(proto.Message)
	if !ok {
		return errors.New("resp Type error")
	}
	d, err := proto.Marshal(pMsg)
	if err != nil {
		return err
	}
	r.Data = d
	return nil
}
