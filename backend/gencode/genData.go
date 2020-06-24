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
{{- if .IsOriginTb}}
	"encoding/json"
	"fmt"
{{- end}}
	"sync"
)

var a{{.ModuleName}}Pool = &sync.Pool{New: func() interface{} {
	return new({{.ModuleName}})
}}

func New{{.ModuleName}}() *{{.ModuleName}} {
	ret := a{{.ModuleName}}Pool.Get().(*{{.ModuleName}})
	*ret = {{.ModuleName}}{}
	return ret
}

func (this *{{.ModuleName}}) Put() {
	*this = {{.ModuleName}}{}
	a{{.ModuleName}}Pool.Put(this)
}

{{- if .IsOriginTb}}
func (this *{{.ModuleName}}) GetStringKey() string {
	return fmt.Sprintf("{{.TableName}}#{{range $index,$val := .KeyFields -}}
		{{- if $index -}}
			#%v
		{{- else -}}
			%v
		{{- end -}}
	{{- end }}",{{range $index,$val := .KeyFields -}}
		{{- if $index -}}
			,this.{{index $.SqlName2Field $val}}
		{{- else -}}
			this.{{index $.SqlName2Field $val}}
		{{- end -}}
	{{- end }})
}

func (this *{{.ModuleName}}) Decode(v []byte) error {
	return json.Unmarshal(v, this)
}

func (this *{{.ModuleName}}) Encode() []byte {
	b, _ := json.Marshal(this)
	return b
}

func (this *{{.ModuleName}}) UpdateSql() (string, []interface{}) {
	sql := "update {{$.TableName}} set {{range $index,$val := .UpdateFields -}}
		{{- if $index -}}
			{{print " "}}and {{dot}}{{.}}{{dot}} = ?
		{{- else -}}
			{{print " "}}{{dot}}{{.}}{{dot}} = ?
		{{- end -}}
	{{- end }} where {{range $index,$val := .KeyFields -}}
		{{- if $index -}}
			{{print " "}}and {{dot}}{{.}}{{dot}} = ?
		{{- else -}}
			{{dot}}{{.}}{{dot}} = ?
		{{- end -}}
	{{- end }}"
	return sql, []interface{}{ {{range $index,$val := .UpdateFields -}}
		{{- if $index -}}
			,this.{{index $.SqlName2Field $val}}
		{{- else -}}
			this.{{index $.SqlName2Field $val}}
		{{- end -}}
	{{- end }} {{range $index,$val := .KeyFields -}} 
			 ,this.{{index $.SqlName2Field $val}}
	{{- end -}} }
}

func (this *{{.ModuleName}}) InsertSql() (string, []interface{}) {
	sql := "insert into {{$.TableName}}({{range $index,$val := .InsertFields -}}
		{{- if $index -}}
			 ,{{dot}}{{.}}{{dot}}
		{{- else -}}
			{{dot}}{{.}}{{dot}}
		{{- end -}}
	{{- end }}) values({{range $index,$val := .InsertFields -}} 
		{{- if $index -}}
			 ,?
		{{- else -}}
			?
		{{- end -}} 
	{{- end -}})"
	return sql, []interface{}{ {{range $index,$val := .InsertFields -}} 
		{{- if $index -}}
			 ,this.{{index $.SqlName2Field $val}}
		{{- else -}}
			this.{{index $.SqlName2Field $val}}
		{{- end -}} 
	{{- end -}} }
}
{{- end}}

func (this *{{.ModuleName}}) TableName() string {
	return "{{.TableName}}"
}

func (this *{{.ModuleName}}) SelectStr() string {
	return "{{range $index,$val := .SelectFields -}}
		{{- if $index -}}
			,{{dot}}{{.}}{{dot}}
		{{- else -}}
			{{dot}}{{.}}{{dot}}
		{{- end -}}
	{{- end }}"
}

{{- range $name,$keys := .IndexKeys}}
func (this *{{$.ModuleName}}) {{$name}}Sql() string {
	return "select {{range $index,$val := $.SelectFields -}}
		{{- if $index -}}
			,{{dot}}{{.}}{{dot}}
		{{- else -}}
			{{dot}}{{.}}{{dot}}
		{{- end -}}
	{{- end }} from {{$.TableName}} where {{range $index,$val := $keys -}}
		{{- if $index -}}
			{{print " "}}and {{dot}}{{.}}{{dot}} = ?
		{{- else -}}
			{{dot}}{{.}}{{dot}} = ?
		{{- end -}}
	{{- end }}"
}
{{- end}}

{{- range $name,$keys := .SliceKeys}}
func (this []*{{$.ModuleName}}) {{$name}}Sql() string {
	return "select {{range $index,$val := $.SelectFields -}}
		{{- if $index -}}
			,{{dot}}{{.}}{{dot}}
		{{- else -}}
			{{dot}}{{.}}{{dot}}
		{{- end -}}
	{{- end }} from {{$.TableName}} where {{range $index,$val := $keys -}}
		{{- if $index -}}
			{{print " "}}and {{dot}}{{.}}{{dot}} = ?
		{{- else -}}
			{{dot}}{{.}}{{dot}} = ?
		{{- end -}}
	{{- end }}"
}
{{- end}}

`

type TableModule struct {
	IsOriginTb    bool
	ModuleName    string
	TableName     string
	Fields        []FieldsType
	SelectFields  []string
	InsertFields  []string
	UpdateFields  []string
	KeyFields     []string
	SqlName2Field map[string]string

	IndexKeys map[string][]string

	IsAutoIncrement bool
	SliceKeys       map[string][]string
}

type FieldsType struct {
	Name string
	Type string
}

func (tInfo *TableModule) handleComment(doc string) map[string]bool {
	tInfo.IsOriginTb = false
	tmpIndexMap := make(map[string]bool)
	tInfo.IsAutoIncrement = false

	lines := strings.Split(doc, "\n")
	for _, line := range lines {

		if strings.Contains(line, "AUTO_INCREMENT") {
			tInfo.IsAutoIncrement = true
		}

		if strings.Contains(line, "CREATE TABLE") {
			tInfo.IsOriginTb = true
			r := strings.Split(line, " ")
			tInfo.TableName = strings.Trim(r[2], "`")
			continue
		}
		if strings.Contains(line, "TABLE_NAME") {
			r := strings.Split(line, " ")
			tInfo.TableName = r[len(r)-1]
			continue
		}
		if strings.Contains(line, "KEY") {
			r := strings.Split(line, " ")
			c := strings.Split(r[len(r)-1], "`")
			c = c[1 : len(c)-1]
			if len(c) <= 0 {
				continue
			}
			f := ""
			isPk := false
			if len(c) > 0 {
				for _, field := range c {
					if field == "," {
						continue
					}
					f += field + " "
					if _, ok := tmpIndexMap[f]; !ok {
						tmpIndexMap[f] = false
					}

					if strings.Contains(line, "PRIMARY") {
						tInfo.KeyFields = append(tInfo.KeyFields, field)
						isPk = true
					}
				}
			} else {
				f = c[0] + " "
				if _, ok := tmpIndexMap[f]; !ok {
					tmpIndexMap[f] = false
				}
				if strings.Contains(line, "PRIMARY") {
					tInfo.KeyFields = append(tInfo.KeyFields, c[0])
					isPk = true
				}
			}
			if strings.Contains(line, "UNIQUE") || isPk {
				//唯一索引或者主键索引  区别于 -->普通索引，批量查找
				tmpIndexMap[f] = true
			}

		}
	}
	return tmpIndexMap
}

func (tb *TableModule) IsPK(key string) bool {
	for _, n := range tb.KeyFields {
		if n == key {
			return true
		}
	}
	return false
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

		if len(gd.Specs) > 1 {
			continue
		}

		sp, ok := gd.Specs[0].(*ast.TypeSpec)
		if !ok {
			continue
		}

		tmpIndexMap := tb.handleComment(pf.Comments[len(pf.Comments)-1].Text())

		tb.ModuleName = sp.Name.Name
		tb.SqlName2Field = make(map[string]string)
		tb.IndexKeys = make(map[string][]string)

		st, ok := sp.Type.(*ast.StructType)
		if !ok {
			fmt.Printf("single type not struct")
			continue
		}
		for _, fl := range st.Fields.List {
			//_, ok := fl.Type.(*ast.Ident)
			//if !ok {
			//	continue
			//}
			tag := fl.Tag.Value
			parts := strings.Split(tag, "\"")
			sqlFieldName := parts[1]
			tb.SqlName2Field[sqlFieldName] = fl.Names[0].Name
			//sqlFieldName := ""
			//if strings.Contains(tag, "primary_key") {
			//	parts := strings.Split(tag, "\"")
			//	firstPart := strings.Split(parts[1], ",")
			//
			//	sqlFieldName = firstPart[0]
			//	tb.SqlName2Field[sqlFieldName] = fl.Names[0].Name
			//} else {
			//	parts := strings.Split(tag, "\"")
			//	sqlFieldName = parts[1]
			//	tb.SqlName2Field[sqlFieldName] = fl.Names[0].Name
			//}
			isTimeField := strings.Contains(sqlFieldName, "create_at") || strings.Contains(sqlFieldName, "update_at")
			isPk := tb.IsPK(sqlFieldName)
			isAutoField := isPk && tb.IsAutoIncrement

			if isTimeField || isAutoField {
				tb.SelectFields = append(tb.SelectFields, sqlFieldName)
				continue
			} else if isPk {
				tb.SelectFields = append(tb.SelectFields, sqlFieldName)
				tb.InsertFields = append(tb.InsertFields, sqlFieldName)
				continue
			} else {
				tb.InsertFields = append(tb.InsertFields, sqlFieldName)
				tb.UpdateFields = append(tb.UpdateFields, sqlFieldName)
				tb.SelectFields = append(tb.SelectFields, sqlFieldName)
			}
		}

		for key, isUniqueOrPk := range tmpIndexMap {
			indexName := "SelectBy"
			v := strings.Split(key, " ")
			v = v[:len(v)-1]
			for _, one := range v {
				indexName += tb.SqlName2Field[one]
			}
			if !isUniqueOrPk {
				tb.SliceKeys[indexName] = v
			} else {
				tb.IndexKeys[indexName] = v
			}
		}
	}
}

//func (m *Modules) genMap(dataDirPath string) {
//	funcMap := template.FuncMap{
//		"dec": func(i int) int {
//			return i - 1
//		},
//	}
//	t := template.New("templateMap")
//	t = t.Funcs(funcMap)
//	t, err := t.Parse(iMapTpl)
//	if err != nil {
//		fmt.Println(err.Error())
//		return
//	}
//
//	fpMap, err := os.OpenFile(dataDirPath+"/map_auto.go",
//		os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0644)
//	if err != nil {
//		fmt.Println("create file error : %s", err.Error())
//		return
//	}
//	err = t.Execute(fpMap, m)
//	if err != nil {
//		fmt.Println("genMap err : ", err.Error())
//		return
//	}
//	fpMap.Close()
//}

func genTableFile() {
	path := os.Getenv("HOME")
	if path == "" {
		fmt.Println("can not get GOPATH")
		return
	}

	dataDirPath := path + "/LittleCai/backend/data/table"

	fd, err := ioutil.ReadDir(dataDirPath)
	if err != nil {
		fmt.Println("read dir error : %s", err.Error())
		return
	}

	funcMap := template.FuncMap{
		"dot": func() string {
			return "`"
		},
	}
	t := template.New("template")
	t = t.Funcs(funcMap)
	t, err = t.Parse(iTableTpl)
	if err != nil {
		fmt.Println(err.Error())
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

		fileNameNoExt := file.Name()[:len(file.Name())-3]
		//fmt.Printf("%v", *tb)
		fpAuto, err := os.OpenFile(dataDirPath+"/"+fileNameNoExt+"_auto.go",
			os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0644)
		if err != nil {
			fmt.Println("create file error : %s", err.Error())
			return
		}

		err = t.Execute(fpAuto, tb)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fpAuto.Close()
	}
}

func main() {
	genTableFile()
	fmt.Println("done ")
}
