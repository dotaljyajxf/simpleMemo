package table

import (
	"encoding/json"
	"fmt"
	"sync"
)

var aTMemoPool = &sync.Pool{New: func() interface{} {
	return new(TMemo)
}}

func NewTMemo() *TMemo {
	ret := aTMemoPool.Get().(*TMemo)
	*ret = TMemo{}
	return ret
}

func (this *TMemo) Put() {
	*this = TMemo{}
	aTMemoPool.Put(this)
}
func (this *TMemo) GetStringKey() string {
	return fmt.Sprintf("memo#%v", this.ID)
}

func (this *TMemo) Decode(v []byte) error {
	return json.Unmarshal(v, this)
}

func (this *TMemo) Encode() []byte {
	b, _ := json.Marshal(this)
	return b
}

func (this *TMemo) UpdateSql() (string, []interface{}) {
	sql := "update memo set  `uid` = ? and `year` = ? and `mouth` = ? and `status` = ? and `remind_time` = ? and `text` = ? where `id` = ?"
	return sql, []interface{}{this.Uid, this.Year, this.Mouth, this.Status, this.RemindTime, this.Text, this.ID}
}

func (this *TMemo) InsertSql() (string, []interface{}) {
	sql := "insert into memo(`uid`,`year`,`mouth`,`status`,`remind_time`,`text`) values(?,?,?,?,?,?)"
	return sql, []interface{}{this.Uid, this.Year, this.Mouth, this.Status, this.RemindTime, this.Text}
}

func (this *TMemo) TableName() string {
	return "memo"
}

func (this *TMemo) SelectStr() string {
	return "`id`,`uid`,`year`,`mouth`,`status`,`remind_time`,`text`,`create_at`,`update_at`"
}
func (this *TMemo) SelectByCreatedAtSql() string {
	return "select `id`,`uid`,`year`,`mouth`,`status`,`remind_time`,`text`,`create_at`,`update_at` from memo where `create_at` = ?"
}
func (this *TMemo) SelectByIDSql() string {
	return "select `id`,`uid`,`year`,`mouth`,`status`,`remind_time`,`text`,`create_at`,`update_at` from memo where `id` = ?"
}
func (this *TMemo) SelectByUidSql() string {
	return "select `id`,`uid`,`year`,`mouth`,`status`,`remind_time`,`text`,`create_at`,`update_at` from memo where `uid` = ?"
}
func (this *TMemo) SelectByUidYearSql() string {
	return "select `id`,`uid`,`year`,`mouth`,`status`,`remind_time`,`text`,`create_at`,`update_at` from memo where `uid` = ? and `year` = ?"
}
func (this *TMemo) SelectByUidYearMouthSql() string {
	return "select `id`,`uid`,`year`,`mouth`,`status`,`remind_time`,`text`,`create_at`,`update_at` from memo where `uid` = ? and `year` = ? and `mouth` = ?"
}
func (this *TMemo) SelectByUidYearMouthStatusSql() string {
	return "select `id`,`uid`,`year`,`mouth`,`status`,`remind_time`,`text`,`create_at`,`update_at` from memo where `uid` = ? and `year` = ? and `mouth` = ? and `status` = ?"
}
func (this *TMemo) SelectByUpdateAtSql() string {
	return "select `id`,`uid`,`year`,`mouth`,`status`,`remind_time`,`text`,`create_at`,`update_at` from memo where `update_at` = ?"
}
