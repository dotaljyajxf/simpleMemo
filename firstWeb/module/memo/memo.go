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
package memo

import (
	"firstWeb/data"
	"firstWeb/data/table"
	log "github.com/sirupsen/logrus"
	"time"
)

func FindMemoByMouth(uid uint64, year int, mouth int8) ([]table.Memo, error) {
	memos := make([]table.Memo, 0)
	err := data.Db.Where("uid = ? and year = ? and mouth = ? and delete_at > 0", uid, year, mouth).Find(memos).Error
	if err != nil {
		log.Errorf("find auth error : %s", err.Error())
		return nil, err
	}
	return memos, nil
}

func CreateMemo(uid uint64, text string) error {
	curTime := time.Now()
	memo := table.NewMemo()
	memo.Uid = uid
	y, m, _ := curTime.Date()
	memo.CreatedAt = curTime.Unix()
	memo.DeletedAt = 0
	memo.Text = text
	memo.Year, memo.Mouth = int(y), int8(m)
	err := data.Db.Create(memo).Error
	if err != nil {
		log.Errorf("find auth error : %s", err.Error())
		return err
	}
	return nil
}
