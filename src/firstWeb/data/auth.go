package data

import (
	"github.com/jinzhu/gorm"
	"sync"
)

type Auth struct {
	gorm.Model
	Name     string `gorm:"unique_index;"`
	PassWord string `gorm:"type:varchar(32)"`
}

var dummyAuth = Auth{}

var pool = sync.Pool{New: func() interface{} {
	return new(Auth)
}}

func NewAuth() *Auth {
	ret := pool.Get().(*Auth)
	*ret = dummyAuth
	return ret
}
