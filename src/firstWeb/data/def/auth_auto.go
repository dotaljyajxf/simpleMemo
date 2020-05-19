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
package def

import (
	"firstWeb/data"
	"sync"
)

var Authpool = sync.Pool{New: func() interface{} {
	return new(Auth)
}}

func NewAuth() *Auth {
	ret := Authpool.Get().(*Auth)
	*ret = Auth{}
	return ret
}

func (auth *Auth) GetName() string {
	return auth.Name
}

func (auth *Auth) SetName(name string) {
	auth.Name = name
}

func (auth *Auth) GetPassWord() string {
	return auth.PassWord
}

func (auth *Auth) SetPassWord(passWord string) {
	auth.PassWord = passWord
}

func FindAuthObj(name string) *Auth {
	auth := NewAuth()
	auth.Name = name
	data.Db.Find(auth).Find(auth)
	return auth
}
