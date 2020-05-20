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

import "time"

type Memo struct {
	ID        uint64    `gorm:"primary_key;AUTO_INCREMENT"`
	CreatedAt string    `gorm:"column:create_time"`
	DeletedAt string    `gorm:"column:delete_time;index:uid_data_del"`
	Uid       uint64    `gorm:"column:uid;index:uid_data_del"`
	Date      time.Time `gorm:"column:date;index:uid_data_del"`
	Text      string    `gorm:"size:255"`
}
