
package table

import (
	"sync"
)

var Memopool = sync.Pool{New: func() interface{} {
	return new(Memo)
}}

func NewMemo() *Memo {
	ret := Memopool.Get().(*Memo)
	*ret = Memo{}
	return ret
}

func (memo *Memo) Release() {
	*memo = Memo{}
	Authpool.Put(memo)
}

func (memo *Memo) TableName() string {
	return "memo"
}


func (memo *Memo) GetID() uint64 {
	return memo.ID
}

func (memo *Memo) SetID(aID uint64) {
	memo.ID = aID
}

func (memo *Memo) GetUid() uint64 {
	return memo.Uid
}

func (memo *Memo) SetUid(aUid uint64) {
	memo.Uid = aUid
}

func (memo *Memo) GetYear() int {
	return memo.Year
}

func (memo *Memo) SetYear(aYear int) {
	memo.Year = aYear
}

func (memo *Memo) GetMouth() int8 {
	return memo.Mouth
}

func (memo *Memo) SetMouth(aMouth int8) {
	memo.Mouth = aMouth
}

func (memo *Memo) GetCreatedAt() int64 {
	return memo.CreatedAt
}

func (memo *Memo) SetCreatedAt(aCreatedAt int64) {
	memo.CreatedAt = aCreatedAt
}

func (memo *Memo) GetDeletedAt() int64 {
	return memo.DeletedAt
}

func (memo *Memo) SetDeletedAt(aDeletedAt int64) {
	memo.DeletedAt = aDeletedAt
}

func (memo *Memo) GetText() string {
	return memo.Text
}

func (memo *Memo) SetText(aText string) {
	memo.Text = aText
}

