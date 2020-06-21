package table

import (
	"encoding/json"
	"fmt"
	"sync"
)

var aTAuthPool = &sync.Pool{New: func() interface{} {
	return new(TAuth)
}}

func NewTAuth() *TAuth {
	ret := aTAuthPool.Get().(*TAuth)
	*ret = TAuth{}
	return ret
}

func (this *TAuth) Put() {
	*this = TAuth{}
	aTAuthPool.Put(this)
}
func (this *TAuth) GetStringKey() string {
	return fmt.Sprintf("auth#%v", this.Uid)
}

func (this *TAuth) Decode(v []byte) error {
	return json.Unmarshal(v, this)
}

func (this *TAuth) Encode() []byte {
	b, _ := json.Marshal(this)
	return b
}

func (this *TAuth) UpdateSql() (string, []interface{}) {
	sql := "update auth set  `nick_name` = ? and `account` = ? and `mail` = ? and `pass_word` = ? and `phone_num` = ? where `uid` = ?"
	return sql, []interface{}{this.NickName, this.Account, this.Mail, this.PassWord, this.PhoneNum, this.Uid}
}

func (this *TAuth) InsertSql() (string, []interface{}) {
	sql := "insert into auth(`nick_name`,`account`,`mail`,`pass_word`,`phone_num`) values(?,?,?,?,?)"
	return sql, []interface{}{this.NickName, this.Account, this.Mail, this.PassWord, this.PhoneNum}
}

func (this *TAuth) TableName() string {
	return "auth"
}

func (this *TAuth) SelectStr() string {
	return "`uid`,`nick_name`,`account`,`mail`,`pass_word`,`phone_num`,`create_at`,`update_at`"
}
func (this *TAuth) SelectByAccountSql() string {
	return "select `uid`,`nick_name`,`account`,`mail`,`pass_word`,`phone_num`,`create_at`,`update_at` from auth where `account` = ?"
}
func (this *TAuth) SelectByCreateAtSql() string {
	return "select `uid`,`nick_name`,`account`,`mail`,`pass_word`,`phone_num`,`create_at`,`update_at` from auth where `create_at` = ?"
}
func (this *TAuth) SelectByMailSql() string {
	return "select `uid`,`nick_name`,`account`,`mail`,`pass_word`,`phone_num`,`create_at`,`update_at` from auth where `mail` = ?"
}
func (this *TAuth) SelectByNickNameSql() string {
	return "select `uid`,`nick_name`,`account`,`mail`,`pass_word`,`phone_num`,`create_at`,`update_at` from auth where `nick_name` = ?"
}
func (this *TAuth) SelectByUidSql() string {
	return "select `uid`,`nick_name`,`account`,`mail`,`pass_word`,`phone_num`,`create_at`,`update_at` from auth where `uid` = ?"
}
func (this *TAuth) SelectByUpdateAtSql() string {
	return "select `uid`,`nick_name`,`account`,`mail`,`pass_word`,`phone_num`,`create_at`,`update_at` from auth where `update_at` = ?"
}
