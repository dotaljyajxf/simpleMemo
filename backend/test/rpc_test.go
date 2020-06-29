package test

import (
	"backend/proto/pb"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/golang/protobuf/proto"
)

var urlRoot = "http://106.12.16.96:8000/doRpc"

type Call struct {
	Method string `json:"method"`
	Args   []byte `json:"args"`
}

func decodePb(msg []byte, obj interface{}) error {
	pMsg, ok := obj.(proto.Message)
	if !ok {
		return errors.New("type error")
	}

	return proto.Unmarshal(msg, pMsg)
}

func TestRegist(t *testing.T) {
	registUrl := "http://106.12.16.96:8000/Regist"

	request, _ := http.NewRequest("POST", registUrl, strings.NewReader("account=ljy&password=liujy594148&name=lalal"))
	request.Header.Set("content-type", "application/x-www-form-urlencoded")

	client := &http.Client{}

	resp, err := client.Do(request)
	if err != nil {
		fmt.Printf("%v", err.Error())
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	ret := pb.NewTAuthInfo()
	err = proto.Unmarshal(body, ret)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("%v", ret.String())
}

func TestLogin(t *testing.T) {
	registUrl := "http://127.0.0.1:8000/Login"

	request, _ := http.NewRequest("POST", registUrl, strings.NewReader("account=dotaljyajxf&password=liujy594148"))
	request.Header.Set("content-type", "application/x-www-form-urlencoded")

	client := &http.Client{}

	resp, err := client.Do(request)
	if err != nil {
		fmt.Printf("%v", err.Error())
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	ret := pb.NewTAppRet()
	err = proto.Unmarshal(body, ret)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("%v\n", ret.String())
}

func TestLoginN(t *testing.T) {
	registUrl := "http://127.0.0.1:8000/Login"

	//request, _ := http.NewRequest("POST", registUrl, strings.NewReader("account=dotaljyajxf&password=liujy594148"))
	//request.Header.Set("content-type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	for i := 0; i < 10; i++ {
		//req.Body = ioutil.NopCloser(io.Reader(bytes.NewReader(testData)))
		//handles包代表http handler的包
		//Handlers.Send代表handles包内部的真正逻辑函数
		request, _ := http.NewRequest("POST", registUrl, strings.NewReader("account=dotaljyajxf&password=liujy594148"))
		request.Header.Set("content-type", "application/x-www-form-urlencoded")
		_, err := client.Do(request)
		if err != nil {
			t.Errorf("do error:%s\n", err.Error())
		}
	}
}

func TestRpc(t *testing.T) {
	//
	//arg := pb.NewTAuthLoginArg()
	//arg.SetAccount("haha")
	//
	//msg, err := proto.Marshal(arg)
	//
	//postValue := Call{
	//	Method: "auth.GetAuthInfo",
	//	Args:   msg,
	//}
	//
	//postString, err := json.Marshal(&postValue)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//request, err := http.NewRequest("POST", urlRoot, strings.NewReader(string(postString)))
	//if err != nil {
	//	// handle error
	//}
	//
	//request.Header.Set("content-type", "application/json")
	//client := &http.Client{}
	//
	//resp, err := client.Do(request)
	//defer resp.Body.Close()
	//
	//body, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	// handle error
	//}
	//
	//ret := pb.NewTAuthInfo()
	//
	//err = decodePb(body, ret)
	//fmt.Println(ret)
}

func Benchmark_Add(b *testing.B) {
	//var n int
	//arg := pb.NewTAuthLoginArg()
	//arg.SetAccount("haha")
	//
	//msg, err := proto.Marshal(arg)
	//
	//postValue := Call{
	//	Method: "auth.GetAuthInfo",
	//	Args:   msg,
	//}
	//
	//postString, err := json.Marshal(&postValue)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//sPostString := string(postString)
	//client := &http.Client{}
	//
	//for i := 0; i < b.N; i++ {
	//
	//	request, err := http.NewRequest("POST", urlRoot, strings.NewReader(sPostString))
	//	if err != nil {
	//		// handle error
	//	}
	//
	//	request.Header.Set("content-type", "application/json")
	//	_, err = client.Do(request)
	//
	//	n++
	//}
}
