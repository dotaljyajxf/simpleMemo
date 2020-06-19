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

const iPoolTpl = `
{{range $NAME := .TypeName }}
var {{$NAME}}Pool = sync.Pool{
	New: func() interface{} {
		return new({{$NAME}})
	},
}

func New{{$NAME}}() *{{$NAME}} {
	return {{$NAME}}Pool.Get().(*{{$NAME}})
}

func (x *{{$NAME}})Put() {
	*x = {{$NAME}}{}
	{{$NAME}}Pool.Put(x)
}
{{end}}
`

type PbInfo struct {
	TypeName []string
}

func (tb *PbInfo) makePbType(dir string, fileName string) {
	tk := token.NewFileSet()

	pf, err := parser.ParseFile(tk, dir+"/"+fileName, nil, parser.ParseComments)
	if err != nil {
		fmt.Println("ParseFile failed err : %s", err.Error())
		return
	}

	for _, decl := range pf.Decls {
		gd, ok := decl.(*ast.GenDecl)
		if !ok {
			continue
		}

		if len(gd.Specs) <= 0 {
			continue
		}

		for _, oneSpec := range gd.Specs {
			sp, ok := oneSpec.(*ast.TypeSpec)
			if !ok {
				continue
			}
			st, ok := sp.Type.(*ast.StructType)
			if !ok {
				fmt.Printf("type not struct")
				continue
			}
			if len(st.Fields.List) < 3 {
				fmt.Printf("struct %s field less\n", sp.Name.Name)
				continue
			}

			tb.TypeName = append(tb.TypeName, sp.Name.Name)
		}
	}
}

func GenPbPool() error {
	path := os.Getenv("HOME")
	if path == "" {
		fmt.Println("can not get GOPATH")
		return nil
	}

	pbDir := path + "/LittleCai/backend/proto/pb"
	fd, err := ioutil.ReadDir(pbDir)
	if err != nil {
		fmt.Println("read dir error : %s", err.Error())
		return err
	}

	funcMap := template.FuncMap{
		"dot": func() string {
			return "`"
		},
	}
	t := template.New("template")
	t = t.Funcs(funcMap)
	t, err = t.Parse(iPoolTpl)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	for _, file := range fd {
		if file.IsDir() {
			continue
		}

		if !strings.HasSuffix(file.Name(), ".pb.go") {
			continue
		}

		pbInfos := new(PbInfo)
		pbInfos.makePbType(pbDir, file.Name())

		//fmt.Printf("%v", *tb)
		fpAuto, err := os.OpenFile(pbDir+"/"+file.Name(),
			os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println("open Pbfile error : %s", err.Error())
			return err
		}

		err = t.Execute(fpAuto, pbInfos)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
		fpAuto.Close()
	}
	return nil
}

func main() {
	err := GenPbPool()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("done")
	}
}
