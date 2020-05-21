/**********************************************************************************************************************
 *
 * Copyright (c) 2010 babeltime.com, Inc. All Rights Reserved
 * $
 *
 **********************************************************************************************************************/

/**
 * @file $
 * @author $(liujianyong@babeltime.com)
 * @date $
 * @version $
 * @brief
 *
 **/
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

const iTableTpl = `
package table

import (
	"sync"
)

var {{.ModuleName}}pool = sync.Pool{New: func() interface{} {
	return new({{.ModuleName}})
}}

func New{{.ModuleName}}() *{{.ModuleName}} {
	ret := {{.ModuleName}}pool.Get().(*{{.ModuleName}})
	*ret = {{.ModuleName}}{}
	return ret
}

func ({{.FileNameNoExt}} *{{.ModuleName}}) Release() {
	*a{{.ModuleName}} = {{.ModuleName}}{}
	Authpool.Put(a{{.ModuleName}})
}

func ({{.FileNameNoExt}} *{{.ModuleName}}) TableName() string {
	return "{{.FileNameNoExt}}"
}

{{range $filed := .Fields}}
func ({{.FileNameNoExt}} *{{.ModuleName}}) Get{{$field.Name}}() {{$field.Type}} {
	return a{{.ModuleName}}.{{$field.Name}}
}

func ({{.FileNameNoExt}} *{{.ModuleName}}) SetNickName(a{{$field.Type}} {{$field.Type}}) {
	{{.FileNameNoExt}}.NickName = a{{$field.Type}}
}

{{end}}
`

type TableModule struct {
	ModuleName    string
	FileNameNoExt string
	Fields        []*FieldsType
}

type FieldsType struct {
	Name string
	Type string
}

func (tb *TableModule) makeFileStruct(dir string, fileName string) {
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

		if len(gd.Specs) > 0 {
			continue
		}

		sp, ok := gd.Specs[0].(*ast.TypeSpec)
		if !ok {
			continue
		}

		tb.ModuleName = sp.Name.Name
		tb.FileNameNoExt = fileName[:len(fileName)-3]
		tb.Fields = make([]*FieldsType, 0)

		st, ok := sp.Type.(*ast.StructType)
		if !ok {
			fmt.Printf("single type not struct")
			continue
		}

		for _, fl := range st.Fields.List {
			fident, ok := fl.Type.(*ast.Ident)
			if ok {
				tb.Fields = append(tb.Fields, &FieldsType{fl.Names[0].Name, fident.Name})
			}
		}

	}
}

func genTableFile() {
	path := os.Getenv("GOPATH")
	if path == "" {
		fmt.Println("can not get GOPATH")
		return
	}

	dataDirPath := path + "/src/firstWeb/data/table"

	fd, err := ioutil.ReadDir(dataDirPath)
	if err != nil {
		fmt.Println("read dir error : %s", err.Error())
		return
	}

	funcMap := template.FuncMap{
		"dec": func(i int) int {
			return i - 1
		},
	}
	t := template.New("tableTpl")
	t = t.Funcs(funcMap)
	t, err = t.Parse(iTableTpl)
	if err != nil {
		return
	}

	for _, file := range fd {
		if file.IsDir() {
			continue
		}

		if strings.Contains(file.Name(), "_auto") {
			continue
		}

		tb := new(TableModule)
		tb.makeFileStruct(dataDirPath, file.Name())

		fmt.Printf("%v", *tb)
		//fpAuto ,err := os.OpenFile(dataDirPath +"/"+ tb.FileNameNoExt + "_auto.go",
		//	os.O_CREATE|os.O_TRUNC|os.O_RDWR,0644)
		//if err != nil {
		//	fmt.Println("create file error : %s",err.Error())
		//	return
		//}
		//
		//t.Execute(fpAuto,tb)
		//fpAuto.Close()
	}
}

func main() {
	genTableFile()
}
