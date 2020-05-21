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

type Memo struct {
	ID        uint64 `gorm:"primary_key;AUTO_INCREMENT"`
	Uid       uint64 `gorm:"column:uid;index:u_y_m_del"`
	Year      int    `gorm:"index:u_y_m_del"`
	Mouth     int8   `gorm:"index:u_y_m_del"`
	CreatedAt int64  `gorm:"column:create_at"`
	DeletedAt int64  `gorm:"column:delete_at;index:u_y_m_del"`
	Text      string `gorm:"size:255"`
}
