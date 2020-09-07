package memo

import (
	"backend/data"
	"backend/data/table"
	"context"
	"time"

	log "github.com/sirupsen/logrus"
)

func FindMemoByMouth(uid int64, year int64, mouth int64) ([]*table.TMemo, error) {
	memos := make(table.TMemos, 0)
	err := data.Manager.Query(&memos, memos.SelectByUidYearMouthStatusSql(),
		uid, year, mouth, 1)
	if err != nil {
		log.Errorf("find auth error : %s", err.Error())
		return nil, err
	}
	return memos, nil
}

func CreateMemo(uid int64, text string, remindTime int64) (int64, error) {
	memo := table.NewTMemo()
	defer memo.Put()
	memo.Uid = uid
	memo.RemindTime = remindTime
	t := time.Unix(remindTime, 0)
	y, m, _ := t.Date()
	memo.Text = text
	memo.Year, memo.Mouth = int64(y), int64(m)
	memo.Status = 0
	res, err := data.Manager.InsertTable(context.Background(), memo)
	if err != nil {
		log.Errorf("find auth error : %s", err.Error())
		return 0, err
	}
	return res.LastInsertId()
}
