package auth

import (
	"fmt"
)

func GetAuthInfo(name string, password string) bool {
	fmt.Println(name, password)
	return true
}
