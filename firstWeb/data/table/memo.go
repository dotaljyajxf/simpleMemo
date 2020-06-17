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

type TMemo struct {
	ID        uint64 `sql:"id,primary_key"`
	Uid       uint64 `sql:"uid"`
	Year      int    `sql:"year"`
	Mouth     int8   `sql:"mouth"`
	CreatedAt int64  `sql:"create_at"`
	DeletedAt int64  `sql:"delete_at"`
	Text      string `sql:"text"`
}

/*
CREATE TABLE `memo` (
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '用户id ',
  `uid` int unsigned NOT NULL  COMMENT '用户id ',
  `year` smallint unsigned NOT NULL  COMMENT '年份 ',
  `mouth` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '账号',
  `text` varchar(256) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '文本',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='备忘录表';
*/
