package test

import (
	"encoding/json"
	"errors"
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

func TestRpc(t *testing.T) {

	arg := pb.NewTGetAuthArg()
	arg.SetName("haha")

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
	arg := pb.NewTGetAuthArg()
	arg.SetName("haha")

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
