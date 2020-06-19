#!/bin/bash


dirRoot="$HOME/LittleCai"
curPwd=$(pwd)

cd $dirRoot/backend/gencode
go run genData.go
echo "genData done ..."
go run genRpc.go
echo "genRpc done ..."

if [ ! -x $GOBIN/protoc ];then
   echo "can not found protoc in $GOBIN/"
   exit
fi

if [ ! -x $GOBIN/protoc-gen-go ];then
  cd $GOPATH/pkg/mod/github.com/golang/protobuf\@v1.4.2/protoc-gen-go/ && go install

  if [ ! -x $GOBIN/protoc-gen-go ];then
    echo "install protoc-gen-go faild!"
  fi
fi

cd $dirRoot/protofile
#/Users/liujianyong/goDownload/bin/protoc --plugin=/Users/liujianyong/goDownload/bin/protoc-gen-go --go_out=../pb *.proto
$GOBIN/protoc --plugin=$GOBIN/protoc-gen-go --go_out=$dirRoot/backend/proto/pb *.proto

cd $curPwd

echo "protoc pb done..."

go run $dirRoot/backend/gencode/genPbPool.go
echo "genPbPool done ..."



