package table

import (
	"encoding/json"
	"fmt"
	"sync"
)

var authPool = &sync.Pool{New: func() interface{} {
	return new(TAuth)
}}

func NewTAuth() *TAuth {
	ret := authPool.Get().(*TAuth)
	*ret = TAuth{}
	return ret
}

func (this *TAuth) Release() {
	*this = TAuth{}
	authPool.Put(this)
}

func (this *TAuth) GetStringKey() string {
	return fmt.Sprintf("%d", this.Uid)
}

func (this *TAuth) Decode(v []byte) error {
	return json.Unmarshal(v, this)
}

func (this *TAuth) Encode() []byte {
	b, _ := json.Marshal(this)
	return b
}

func (this *TAuth) TableName() string {
	return "auth"
}

func (this *TAuth) SelectStr() string {
	return "select `uid`,`nick_name`,`account`,`mail`,`pass_word`,`phone_num`,`create_at`,`update_at`"
}

func (this *TAuth) SelectSql() (string, []interface{}) {
	sql := "select `uid`,`nick_name`,`account`,`mail`,`pass_word`,`phone_num`,`create_at`,`update_at` from auth where uid = ?"
	return sql, []interface{}{this.Uid}
}

func (this *TAuth) InsertSql() (string, []interface{}) {
	sql := "insert into auth(`uid`,`nick_name`,`account`,`mail`,`pass_word`,`phone_num`,)  values(?,?,?,?,?,?,)"
	return sql, []interface{}{this.Uid, this.NickName, this.Account, this.Mail, this.PassWord, this.PhoneNum}
}

func (this *TAuth) UpdateSql() (string, []interface{}) {
	sql := "update auth set`uid` = ?,`nick_name` = ?,`account` = ?,`mail` = ?,`pass_word` = ?,`phone_num` = ?, where uid = ?"
	return sql, []interface{}{this.Uid}
}

func (this *TAuth) GetUid() uint64 {
	return this.Uid
}

func (this *TAuth) SetUid(aUid uint64) {
	this.Uid = aUid
}

func (this *TAuth) GetNickName() string {
	return this.NickName
}

func (this *TAuth) SetNickName(aNickName string) {
	this.NickName = aNickName
}

func (this *TAuth) GetAccount() string {
	return this.Account
}

func (this *TAuth) SetAccount(aAccount string) {
	this.Account = aAccount
}

func (this *TAuth) GetMail() string {
	return this.Mail
}

func (this *TAuth) SetMail(aMail string) {
	this.Mail = aMail
}

func (this *TAuth) GetPassWord() string {
	return this.PassWord
}

func (this *TAuth) SetPassWord(aPassWord string) {
	this.PassWord = aPassWord
}

func (this *TAuth) GetPhoneNum() string {
	return this.PhoneNum
}

func (this *TAuth) SetPhoneNum(aPhoneNum string) {
	this.PhoneNum = aPhoneNum
}

func (this *TAuth) GetCreateAt() int64 {
	return this.CreateAt
}

func (this *TAuth) SetCreateAt(aCreateAt int64) {
	this.CreateAt = aCreateAt
}

func (this *TAuth) GetUpdateAt() int64 {
	return this.UpdateAt
}

func (this *TAuth) SetUpdateAt(aUpdateAt int64) {
	this.UpdateAt = aUpdateAt
}
