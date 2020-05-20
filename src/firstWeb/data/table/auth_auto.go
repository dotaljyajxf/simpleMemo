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
package table

import (
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

func (auth *Auth) Release() {
	*auth = Auth{}
	Authpool.Put(auth)
}

func (auth *Auth) TableName() string {
	return "auth"
}

func (auth *Auth) GetUid() uint64 {
	return auth.Uid
}

func (auth *Auth) SetNickName(nickName string) {
	auth.NickName = nickName
}

func (auth *Auth) GetAccount() string {
	return auth.Account
}

func (auth *Auth) SetAccount(account string) {
	auth.Account = account
}

func (auth *Auth) GetPassWord() string {
	return auth.PassWord
}

func (auth *Auth) SetPassWord(passWord string) {
	auth.PassWord = passWord
}

func (auth *Auth) GetMail() string {
	return auth.PassWord
}

func (auth *Auth) SetMail(mail string) {
	auth.Mail = mail
}

func (auth *Auth) GetPhoneNum() string {
	return auth.PhoneNum
}

func (auth *Auth) SetPhoneNum(phoneNum string) {
	auth.PhoneNum = phoneNum
}
