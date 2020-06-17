package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
)

const iCodeTpl = `//DO NOT MODIFY GENERATED CODE !!!
//DO NOT MODIFY GENERATED CODE !!!
//DO NOT MODIFY GENERATED CODE !!!

package model
import (
	"backend/proto/pb"
	"github.com/golang/protobuf/proto"
	"errors"
)


var gRpcMethodMap = map[string]func(args []byte) (interface{} ,error) {
		{{range $method :=.Methods}}"{{$method.Module}}.{{$method.Method}}": proxy_{{$method.Module}}_{{$method.Method}},
		{{end}}
}

{{range $method := .Methods}}
func proxy_{{$method.Module}}_{{$method.Method}}(args []byte) (interface{} ,error) {
	
	request := pb.New{{$method.Request}}()
	defer request.Put()

	response := pb.New{{$method.Response}}()
	
	err := decodePb(args, request)
	if err != nil {
		return nil,	errors.New("decodeArgs error")

	}
	
	err = {{$method.Module}}.{{$method.Method}}(request,response)
	return response,err
}
{{end}}

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
`

type ModuleInfo struct {
	Methods []RpcMethod
	Modules map[string]string
}

type RpcMethod struct {
	Request  string
	Response string
	Module   string
	Method   string
}

func getParamType(param ast.Expr) (string, string) {
	st, ok := param.(*ast.StarExpr)
	if !ok {
		return "", ""
	}
	sel, ok := st.X.(*ast.SelectorExpr)
	if !ok {
		return "", ""
	}
	m, ok := sel.X.(*ast.Ident)
	if !ok {
		return "", ""
	}
	return m.Name, sel.Sel.Name
}

func check(wd string, fi os.FileInfo, moduleInfo *ModuleInfo) {

	fs := new(token.FileSet)

	pkgs, err := parser.ParseDir(fs, wd+"/../model/"+fi.Name(), nil, parser.ParseComments)
	if err != nil {
		fmt.Println("parseDir Error:%s", err.Error())
		return
	}

	for _, pkg := range pkgs {
		moduleInfo.Modules[pkg.Name] = "backend/model/" + pkg.Name
		for fn, f := range pkg.Files {

			if strings.Split(fn, ".")[1] != "go" {
				continue
			}
			start := strings.LastIndex(fn, "/")
			fs := []byte(fn)
			name := fs[start+1 : len(fn)-3]
			if !strings.Contains(string(name), "Rpc") {
				continue
			}

			for _, d := range f.Decls {
				ft, ok := d.(*ast.FuncDecl)
				if !ok {
					continue
				}

				if len(ft.Type.Params.List) != 1 {

				}
				var methodInfo RpcMethod
				x, sel := getParamType(ft.Type.Params.List[0].Type)
				if x != "pb" {

				}
				x, rel := getParamType(ft.Type.Params.List[1].Type)
				if x != "pb" {

				}
				methodInfo.Request = sel
				methodInfo.Response = rel
				methodInfo.Module = pkg.Name
				methodInfo.Method = ft.Name.Name

				moduleInfo.Methods = append(moduleInfo.Methods, methodInfo)

			}
		}
	}
}

func genModuleInfo(wd string) *ModuleInfo {
	moduleInfo := new(ModuleInfo)
	moduleInfo.Modules = make(map[string]string)

	dir, err := ioutil.ReadDir(wd + "/../model/")
	if err != nil {
		return nil
	}
	for _, fi := range dir {
		if !fi.IsDir() {
			continue
		}
		check(wd, fi, moduleInfo)
	}

	return moduleInfo
}

func GenGoFile() error {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("get work directory failed:%s", err.Error())
		return err
	}

	param := genModuleInfo(wd)
	funcMap := template.FuncMap{
		"dec": func(i int) int {
			return i - 1
		},
	}
	t := template.New("template")
	t = t.Funcs(funcMap)
	t, err = t.Parse(iCodeTpl)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(wd+"/../model/rpc_auto.go", os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	err = t.Execute(file, param)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	err := GenGoFile()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("done")
	}
}
