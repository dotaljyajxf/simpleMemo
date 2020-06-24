package table

import "time"

type TMemo struct {
	ID         int64     `sql:"id"`
	Uid        int64     `sql:"uid"`
	Year       int64     `sql:"year"`
	Mouth      int64     `sql:"mouth"`
	Status     int64     `sql:"status"` //0未删除
	RemindTime int64     `sql:"remind_time"`
	Text       string    `sql:"text"`
	CreatedAt  time.Time `sql:"create_at"`
	UpdateAt   time.Time `sql:"update_at"`
}

/*
CREATE TABLE `memo` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '用户id ',
  `uid` int unsigned NOT NULL  DEFAULT 0 COMMENT '用户id ',
  `year` tinyint unsigned NOT NULL  DEFAULT 0 COMMENT '年份 ',
  `status` tinyint unsigned NOT NULL DEFAULT 0 COMMENT '状态 ',
  `mouth` tinyint unsigned NOT NULL DEFAULT 0 COMMENT '账号',
  `remind_time` int unsigned NOT NULL DEFAULT 0 COMMENT '提醒时间',
  `text` varchar(256) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '文本',
  `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `u_year_m_d` (`uid`,`year`,`mouth`,`status`),
  KEY `idx_created` (`create_at`),
  KEY `idx_updated` (`update_at`),
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='备忘录表';
*/
