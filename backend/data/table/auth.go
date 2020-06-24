package table

import "time"

type TAuth struct {
	Uid      int64     `sql:"uid"`
	NickName string    `sql:"nick_name"`
	Account  string    `sql:"account"`
	Mail     string    `sql:"mail"`
	PassWord string    `sql:"pass_word"`
	PhoneNum string    `sql:"phone_num"`
	CreateAt time.Time `sql:"create_at"`
	UpdateAt time.Time `sql:"update_at"`
}

/*
CREATE TABLE `auth` (
  `uid` int unsigned NOT NULL AUTO_INCREMENT COMMENT '用户id ',
  `nick_name` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 0 COMMENT '昵称',
  `account` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '0' COMMENT '账号',
  `mail` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '0' COMMENT '邮箱',
  `pass_word` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '0' COMMENT '密码',
  `phone_num` char(11) NOT NULL DEFAULT '0' COMMENT 'phone',
  `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`uid`),
  UNIQUE KEY `account` (`account`),
  UNIQUE KEY `nick_name` (`nick_name`),
  UNIQUE KEY `mail` (`mail`),
  KEY `idx_created` (`create_at`),
  KEY `idx_updated` (`update_at`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户表';
*/
