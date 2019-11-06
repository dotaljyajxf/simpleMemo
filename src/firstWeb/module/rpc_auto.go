//DO NOT MODIFY GENERATED CODE !!!
//DO NOT MODIFY GENERATED CODE !!!
//DO NOT MODIFY GENERATED CODE !!!

package module
import (
	"firstWeb/proto/pb"
	"github.com/golang/protobuf/proto"
	"github.com/sirupsen/logrus"
	"errors"
    "firstWeb/module/auth"
	
)


var gRpcMethodMap = map[string]func(args []byte) (interface{} ,error) {
		"auth.GetAuthInfo": proxy_auth_GetAuthInfo,
		
}


func proxy_auth_GetAuthInfo(args []byte) (interface{} ,error) {
	
	request := pb.NewTGetAuthArg()
	defer request.Put()

	response := pb.NewTAuthInfo()
	defer response.Put()
	
	err := decodePb(args, request)
	if err != nil {
		return nil,	errors.New("decodeArgs error")

	}
	
	err = auth.GetAuthInfo(request,response)
	return response,err
}


func decodePb(msg []byte  ,obj interface{}) error {
	pMsg,ok := obj.(proto.Message)
	if !ok {
		return errors.New("type error")
	}

	return proto.Unmarshal(msg,pMsg)
}

func DoRpcMethod(method string,arg []byte) (interface{} ,error){
	if f ,ok := gRpcMethodMap[method];ok {
		return f(arg)
	}
	return nil,errors.New("unknow method")
}
