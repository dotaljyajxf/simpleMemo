package table

type TMemo struct {
	ID        uint64 `sql:"id,primary_key,auto_increment"`
	Uid       uint64 `sql:"uid"`
	Year      int    `sql:"year"`
	Mouth     int8   `sql:"mouth"`
	Status    int8   `sql:"status"`
	CreatedAt int64  `sql:"create_at"`
	DeletedAt int64  `sql:"delete_at"`
	Text      string `sql:"text"`
}

/*
CREATE TABLE `memo` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '用户id ',
  `uid` int unsigned NOT NULL  DEFAULT 0 COMMENT '用户id ',
  `year` tinyint unsigned NOT NULL  DEFAULT 0 COMMENT '年份 ',
  `status` tinyint unsigned NOT NULL DEFAULT 0 COMMENT '状态 ',
  `mouth` tinyint unsigned NOT NULL DEFAULT 0 COMMENT '账号',
  `text` varchar(256) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '文本',
  `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `delete_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `u_year_m_d` (`uid`,`year`,`mouth`,`status`),
  KEY `idx_created` (`create_at`),
  KEY `idx_deleted` (`create_at`),
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='备忘录表';
*/
