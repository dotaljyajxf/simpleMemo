package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type Auth struct {
	gorm.Model
}

func CheckAuth(name string, password string) bool {
	fmt.Println(name, password)
	return true
}
