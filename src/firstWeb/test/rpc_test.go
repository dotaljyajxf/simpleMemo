package test

import (
	"encoding/json"
	"errors"
	"firstWeb/data/table"
	"firstWeb/proto/pb"
	"fmt"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
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

	request, _ := http.NewRequest("POST", registUrl, strings.NewReader("account=liujianyong&password=liujy594148&name=lalal"))
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
	p := &struct {
		authObj table.Auth
		token   string
		message string
	}{}
	fmt.Println(body)
	err = json.Unmarshal(body, p)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("%v", p.message)
	fmt.Printf("%v", p.token)
	fmt.Printf("%v", p.authObj)

}

func TestLogin(t *testing.T) {
	registUrl := "http://106.12.16.96:8000/Login"

	request, _ := http.NewRequest("POST", registUrl, strings.NewReader("account=liujianyong&password=liujy594148"))
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

func TestRpc(t *testing.T) {

	arg := pb.NewTAuthLoginArg()
	arg.SetAccount("haha")

	msg, err := proto.Marshal(arg)

	postValue := Call{
		Method: "auth.GetAuthInfo",
		Args:   msg,
	}

	postString, err := json.Marshal(&postValue)
	if err != nil {
		fmt.Println(err)
	}

	request, err := http.NewRequest("POST", urlRoot, strings.NewReader(string(postString)))
	if err != nil {
		// handle error
	}

	request.Header.Set("content-type", "application/json")
	client := &http.Client{}

	resp, err := client.Do(request)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	ret := pb.NewTAuthInfo()

	err = decodePb(body, ret)
	fmt.Println(ret)
}

func Benchmark_Add(b *testing.B) {
	var n int
	arg := pb.NewTAuthLoginArg()
	arg.SetAccount("haha")

	msg, err := proto.Marshal(arg)

	postValue := Call{
		Method: "auth.GetAuthInfo",
		Args:   msg,
	}

	postString, err := json.Marshal(&postValue)
	if err != nil {
		fmt.Println(err)
	}

	sPostString := string(postString)
	client := &http.Client{}

	for i := 0; i < b.N; i++ {

		request, err := http.NewRequest("POST", urlRoot, strings.NewReader(sPostString))
		if err != nil {
			// handle error
		}

		request.Header.Set("content-type", "application/json")
		_, err = client.Do(request)

		n++
	}
}
