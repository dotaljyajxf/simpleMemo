package table

import (
	"sync"
)

var aTAuthSessionPool = &sync.Pool{New: func() interface{} {
	return new(TAuthSession)
}}

func NewTAuthSession() *TAuthSession {
	ret := aTAuthSessionPool.Get().(*TAuthSession)
	*ret = TAuthSession{}
	return ret
}

func (this *TAuthSession) Put() {
	*this = TAuthSession{}
	aTAuthSessionPool.Put(this)
}

func (this *TAuthSession) TableName() string {
	return "auth"
}

func (this *TAuthSession) SelectSql() (string, []interface{}) {
	sql := "select `uid`,`nick_name`,`account`,`mail`,`phone_num` from auth where `uid` = ?"
	return sql, []interface{}{this.Uid}
}

func (this *TAuthSession) FieldsStr() string {
	return "`uid`,`nick_name`,`account`,`mail`,`phone_num`"
}
func (this *TAuthSession) SelectByAccountSql() string {
	return "select `uid`,`nick_name`,`account`,`mail`,`phone_num` from auth where `account` = ?"
}
func (this *TAuthSession) SelectByMailSql() string {
	return "select `uid`,`nick_name`,`account`,`mail`,`phone_num` from auth where `mail` = ?"
}
func (this *TAuthSession) SelectByNickNameSql() string {
	return "select `uid`,`nick_name`,`account`,`mail`,`phone_num` from auth where `nick_name` = ?"
}
func (this *TAuthSession) SelectByUidSql() string {
	return "select `uid`,`nick_name`,`account`,`mail`,`phone_num` from auth where `uid` = ?"
}
