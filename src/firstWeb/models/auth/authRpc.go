package auth

import (
	"firstWeb/proto/pb"
)

func GetAuthInfo(arg *pb.TGetAuthArg, ret *pb.TAuthInfo) error {

	pb.NewTAuthInfo()
	ret.SetAge(12)
	ret.SetName(*arg.Name)
	ret.SetSex(1)
	return nil
}
