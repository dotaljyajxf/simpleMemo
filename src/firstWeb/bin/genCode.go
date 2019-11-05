package main

import (
	"fmt"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
)

const iCodeTpl = `//DO NOT MODIFY GENERATED CODE !!!
//DO NOT MODIFY GENERATED CODE !!!
//DO NOT MODIFY GENERATED CODE !!!
import (
	"sync"
	"fmt"
	
	"firstWeb/proto/pb"
	{{range $module, $fn := .Modules}}"{{$fn}}"
	{{end}}
)
func init() {
	gRpcMethodMap = map[string]func(method string, args []byte) interface{} ,error {
		{{range $method :=.Methods}}"{{$method.Module}}.{{$method.Method}}": proxy_{{$method.Module}}_{{$method.Method}},
		{{end}}
	}
}

{{range $method := .Methods}}
func proxy_{{$method.Module}}_{{$method.Method}}(method string, args []byte) interface{} ,error {
	
	request := new_{{$method.Request}}()
	defer g_{{$method.Request}}_pool.Put(request)
	
	err := codec.Decode(args, request)
	if err != nil {
		return nil,	errors.New("decodeArgs error")

	}
	
	ret := {{$method.Module}}.{{$method.Method}}(request)
	return ret,nil
}

`

type ModuleInfo struct {
	Types   []string
	Modules map[string]string
	Methods []RpcMethod
}

type RpcMethod struct {
	Request   string
	Response  string
	Module    string
	Method    string
	ParamType []string
	Comment   string
}

type ImplInfo struct {
	Struct  string
	Module  string
	Package string
	Path    string
}

func check(wd string, fi os.FileInfo, moduleInfo *ModuleInfo, typeMap map[string]bool) {

	fs := new(token.FileSet)

	pkgs, err := parser.ParseDir(fs, wd+"/../module/"+fi.Name(), nil, parser.ParseComments)
	if err != nil {
		fmt.Println("parseDir Error:%s", err.Error())
		return
	}

	for _, pkg := range pkgs {
		for fn, f := range pkg.Files {
			fmt.Println(fn)
			fmt.Println(f)
		}
	}
}

func genModuleInfo(wd string) *ModuleInfo {
	moduleInfo := new(ModuleInfo)
	moduleInfo.Modules = make(map[string]string)

	typeMap := make(map[string]bool)

	dir, err := ioutil.ReadDir(wd + "/../module/")
	if err != nil {
		return nil
	}
	for _, fi := range dir {
		if !fi.IsDir() {
			continue
		}
		switch fi.Name() {
		case "errs":
			continue
		case "global":
			continue
		case "handler":
			continue
		case "hook":
			continue
		case "public":
			continue
		case "web":
			continue
		case "modidata":
			continue
		}
		check(wd, fi, moduleInfo, typeMap)
	}

	for typeName, _ := range typeMap {
		moduleInfo.Types = append(moduleInfo.Types, typeName)
	}
	return moduleInfo
}

func GenGoFile() error {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("get work directory failed:%s", err.Error())
		return err
	}

	_ = genModuleInfo(wd)
	//funcMap := template.FuncMap{
	//	"dec": func(i int) int {
	//		return i - 1
	//	},
	//}
	//t := template.New("template")
	//t = t.Funcs(funcMap)
	//t, err = t.Parse(iCodeTpl)
	//if err != nil {
	//	return err
	//}
	//
	//file, err := os.OpenFile(wd+"/../rpc_auto.go", os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0644)
	//if err != nil {
	//	return err
	//}
	//defer file.Close()
	//err = t.Execute(file, param)
	//if err != nil {
	//	return err
	//}
	//
	////	fmtFile(wd+"/../rpc_auto.go")
	//
	//t = template.New("module")
	//t = t.Funcs(funcMap)
	//t, err = t.Parse(iImplTpl)
	//if err != nil {
	//	return err
	//}
	//for _, info := range param.ImplMap {
	//	path := wd + "/../../../" + info.Path + "/" + info.Package + "/" + info.Module + "_pool_auto.go"
	//
	//	file, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0644)
	//	if err != nil {
	//		return err
	//	}
	//
	//	defer file.Close()
	//	err = t.Execute(file, info)
	//	if err != nil {
	//		return err
	//	}
	//
	//	//		fmtFile(path)
	//}

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
