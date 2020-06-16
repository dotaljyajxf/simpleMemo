package table

type Auth struct {
	Uid        uint64 `sql:"uid,primary_key"`
	NickName   string `sql:"nick_name"`
	Account    string `sql:"account"`
	Mail       string `sql:"mail"`
	PassWord   string `sql:"pass_word"`
	PhoneNum   string `sql:"phone_num"`
	CreateTime int64  `sql:"create_at"`
	UpdateTime int64  `sql:"update_at"`
}

/*
CREATE TABLE `auth` (
  `uid` int unsigned NOT NULL AUTO_INCREMENT COMMENT '用户id ',
  `nick_name` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '昵称',
  `account` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '账号',
  `mail` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '邮箱',
  `pass_word` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '密码',
  `phone` char(11) NOT NULL COMMENT 'phone',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`uid`),
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户表';
*/
