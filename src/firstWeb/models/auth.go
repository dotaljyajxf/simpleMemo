package models

import "fmt"

func CheckAuth(name string, password string) bool {
	fmt.Println(name, password)
	return true
}
