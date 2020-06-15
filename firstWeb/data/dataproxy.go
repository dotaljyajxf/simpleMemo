package data

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"unicode"

	"reflect"
)

var Db *sql.DB

func Exec(ctx context.Context, sql string, args ...interface{}) (sql.Result, error) {
	return Db.ExecContext(ctx, sql, args)
}

func QueryContext(ctx context.Context, resp interface{}, sql string, args ...interface{}) error {
	rows, err := Db.QueryContext(ctx, sql, args)
	if err != nil {
		return err
	}
	return query(rows, resp)
}

func Query(resp interface{}, sql string, args ...interface{}) error {
	rows, err := Db.Query(sql, args)
	if err != nil {
		return err
	}
	return query(rows, resp)
}

func fillFieldAddr(columnNames []string, val reflect.Value) []interface{} {
	typ := val.Type()

	retAddr := make([]interface{}, len(columnNames))
	fieldNum := val.NumField()
	for _, name := range columnNames {
		for i := 0; i < fieldNum; i++ {
			if !val.Field(i).CanSet() {
				continue
			}

			if tag, ok := typ.Field(i).Tag.Lookup("sql"); ok {
				if strings.ToLower(tag) == strings.ToLower(name) {
					retAddr = append(retAddr, val.Field(i).Addr().Interface())
					break
				}
			} else {
				if camelToUnderscore(typ.Field(i).Name) == name {
					retAddr = append(retAddr, val.Field(i).Addr().Interface())
					break
				}
			}
		}
	}
	return retAddr
}

func parseRow(rows *sql.Rows, columnNames []string, val reflect.Value) error {
	for rows.Next() {
		scan := fillFieldAddr(columnNames, val)
		return rows.Scan(scan)
	}
	return rows.Err()
}

func parseRows(rows *sql.Rows, columnNames []string, val reflect.Value) error {
	typ := val.Type()
	t := typ.Elem()
	results := reflect.MakeSlice(t, 0, 0)

	isPtr := false
	if t.Elem().Kind() == reflect.Struct {
		t = t.Elem() // struct
	} else if t.Elem().Kind() == reflect.Ptr && !val.Elem().IsNil() && t.Elem().Elem().Kind() == reflect.Struct {
		isPtr = true
		t = t.Elem().Elem() // struct
	} else {
		return fmt.Errorf("scan data invalid(%v,%v)", t.Elem().Kind(), t.Elem().Elem().Kind())
	}

	for rows.Next() {
		row := reflect.New(t).Elem()
		scans := fillFieldAddr(columnNames, row)
		err := rows.Scan(scans...)
		if err != nil {
			return err
		}

		if isPtr {
			results = reflect.Append(results, row.Addr())
			continue
		}
		results = reflect.Append(results, row)
	}

	val.Elem().Set(results)
	return rows.Err()
}

func query(rows *sql.Rows, resp interface{}) error {
	columnNames, err := rows.Columns()
	if err != nil {
		return err
	}
	val := reflect.ValueOf(resp)
	typ := val.Type()

	if typ.Kind() != reflect.Slice && typ.Elem().Kind() == reflect.Struct {
		return parseRow(rows, columnNames, val)
	} else {
		return parseRows(rows, columnNames, val)
	}
}

// 驼峰转下划线
func camelToUnderscore(name string) string {
	buf := make([]rune, 0, len(name)+4)
	var preIsUpper bool
	for i, r := range name {
		if unicode.IsUpper(r) {
			if i != 0 && !preIsUpper {
				buf = append(buf, '_')
			}
			buf = append(buf, unicode.ToLower(r))
		} else {
			buf = append(buf, r)
		}
		preIsUpper = unicode.IsUpper(r)
	}
	return string(buf)
}
