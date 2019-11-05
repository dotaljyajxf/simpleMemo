package auth

import (
	"firstWeb/proto/pb"
)

func GetInfo(arg *pb.TGetAuthArg) *pb.TAuthInfo {

	ret := pb.NewTAuthInfo()
	ret.SetAge(12)
	ret.SetName("LittleCai")
	ret.SetSex(1)
	return ret
}
