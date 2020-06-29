package test

import (
	"net/http"
	"strings"
	"testing"
)

func BenchmarkSend(b *testing.B) {
	registUrl := "http://127.0.0.1:8000/Login"
	//构建请求对象
	//req, _ := http.NewRequest("POST", registUrl, strings.NewReader("account=dotaljyajxf&password=liujy594148"))
	//req.Header.Set("content-type", "application/x-www-form-urlencoded")
	//req, _ := http.NewRequest("POST", "/req_url", io.Reader(bytes.NewReader(testData)))
	//req.Header.Set("X-Forwarded-For" , "10.78.48.10")

	b.ReportAllocs()
	client := &http.Client{}
	for i := 0; i < b.N; i++ {
		//req.Body = ioutil.NopCloser(io.Reader(bytes.NewReader(testData)))
		//handles包代表http handler的包
		//Handlers.Send代表handles包内部的真正逻辑函数
		req, _ := http.NewRequest("POST", registUrl, strings.NewReader("account=dotaljyajxf&password=liujy594148"))
		req.Header.Set("content-type", "application/x-www-form-urlencoded")
		_, err := client.Do(req)
		if err != nil {
			b.Errorf("do error:%s\n", err.Error())
		}
	}
}

//func send(c *http.Client,req *http.Request) error{
//	_, err := c.Do(req)
//	return err
//}
