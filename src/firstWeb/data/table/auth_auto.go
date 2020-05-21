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

func (auth *Auth) SetUid(aUid uint64) {
	auth.Uid = aUid
}

func (auth *Auth) GetNickName() string {
	return auth.NickName
}

func (auth *Auth) SetNickName(aNickName string) {
	auth.NickName = aNickName
}

func (auth *Auth) GetAccount() string {
	return auth.Account
}

func (auth *Auth) SetAccount(aAccount string) {
	auth.Account = aAccount
}

func (auth *Auth) GetMail() string {
	return auth.Mail
}

func (auth *Auth) SetMail(aMail string) {
	auth.Mail = aMail
}

func (auth *Auth) GetPassWord() string {
	return auth.PassWord
}

func (auth *Auth) SetPassWord(aPassWord string) {
	auth.PassWord = aPassWord
}

func (auth *Auth) GetPhoneNum() string {
	return auth.PhoneNum
}

func (auth *Auth) SetPhoneNum(aPhoneNum string) {
	auth.PhoneNum = aPhoneNum
}

func (auth *Auth) GetCreateTime() int64 {
	return auth.CreateTime
}

func (auth *Auth) SetCreateTime(aCreateTime int64) {
	auth.CreateTime = aCreateTime
}
