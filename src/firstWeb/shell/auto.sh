#/bin/bash


dirRoot="/home/game/LittleCai"
curPwd=`pwd`

cd $dirRoot/src/firstWeb/bin 
go run genCode.go
echo "genCode ok ..."

cd $dirRoot/src/firstWeb/proto/protofile 
protoc --plugin=../../../github.com/golang/protobuf/protoc-gen-go/protoc-gen-go --go_out=../pb *.proto 

cd $curPwd

echo "protoc pb ok..."



