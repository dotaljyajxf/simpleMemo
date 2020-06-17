package table

import (
	"encoding/json"
	"fmt"
	"sync"
)

var memoPool = &sync.Pool{New: func() interface{} {
	return new(TMemo)
}}

func NewTMemo() *TMemo {
	ret := memoPool.Get().(*TMemo)
	*ret = TMemo{}
	return ret
}

func (this *TMemo) Release() {
	*this = TMemo{}
	memoPool.Put(this)
}

func (this *TMemo) GetStringKey() string {
	return fmt.Sprintf("%d", this.ID)
}

func (this *TMemo) Decode(v []byte) error {
	return json.Unmarshal(v, this)
}

func (this *TMemo) Encode() []byte {
	b, _ := json.Marshal(this)
	return b
}

func (this *TMemo) SelectSql() (string, []interface{}) {
	sql := "select `id`,`uid`,`year`,`mouth`,`create_at`,`delete_at`,`text` from memo where id = ?"
	return sql, []interface{}{this.ID}
}

func (this *TMemo) InsertSql() (string, []interface{}) {
	sql := "insert into memo(`id`,`uid`,`year`,`mouth`,`delete_at`,`text`)  values(?,?,?,?,?,?)"
	return sql, []interface{}{this.ID, this.Uid, this.Year, this.Mouth, this.DeletedAt, this.Text}
}

func (this *TMemo) UpdateSql() (string, []interface{}) {
	sql := "update memo set`id` = ?,`uid` = ?,`year` = ?,`mouth` = ?,`delete_at` = ?,`text` = ? where id = ?"
	return sql, []interface{}{this.ID}
}

func (this *TMemo) GetID() uint64 {
	return this.ID
}

func (this *TMemo) SetID(aID uint64) {
	this.ID = aID
}

func (this *TMemo) GetUid() uint64 {
	return this.Uid
}

func (this *TMemo) SetUid(aUid uint64) {
	this.Uid = aUid
}

func (this *TMemo) GetYear() int {
	return this.Year
}

func (this *TMemo) SetYear(aYear int) {
	this.Year = aYear
}

func (this *TMemo) GetMouth() int8 {
	return this.Mouth
}

func (this *TMemo) SetMouth(aMouth int8) {
	this.Mouth = aMouth
}

func (this *TMemo) GetCreatedAt() int64 {
	return this.CreatedAt
}

func (this *TMemo) SetCreatedAt(aCreatedAt int64) {
	this.CreatedAt = aCreatedAt
}

func (this *TMemo) GetDeletedAt() int64 {
	return this.DeletedAt
}

func (this *TMemo) SetDeletedAt(aDeletedAt int64) {
	this.DeletedAt = aDeletedAt
}

func (this *TMemo) GetText() string {
	return this.Text
}

func (this *TMemo) SetText(aText string) {
	this.Text = aText
}
