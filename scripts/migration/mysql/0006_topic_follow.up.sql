CREATE TABLE `p_topic_user` (
	`id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
	`topic_id` BIGINT UNSIGNED NOT NULL COMMENT '标签ID',
	`user_id` BIGINT UNSIGNED NOT NULL COMMENT '创建者ID',
	`alias_name` VARCHAR ( 255 ) COMMENT '别名',
	`remark` VARCHAR ( 512 ) COMMENT '备注',
	`quote_num` BIGINT UNSIGNED COMMENT '引用数',
	`is_top` TINYINT UNSIGNED NOT NULL DEFAULT '0' COMMENT '是否置顶 0 为未置顶、1 为已置顶',
	`created_on` BIGINT UNSIGNED NOT NULL DEFAULT '0' COMMENT '创建时间',
	`modified_on` BIGINT UNSIGNED NOT NULL DEFAULT '0' COMMENT '修改时间',
	`deleted_on` BIGINT UNSIGNED NOT NULL DEFAULT '0' COMMENT '删除时间',
	`is_del` TINYINT UNSIGNED NOT NULL DEFAULT '0' COMMENT '是否删除 0 为未删除、1 为已删除',
	`reserve_a` VARCHAR ( 255 ) COMMENT '保留字段a',
	`reserve_b` VARCHAR ( 255 ) COMMENT '保留字段b',
	PRIMARY KEY ( `id` ) USING BTREE,
    UNIQUE KEY `idx_topic_user_uid_tid` ( `topic_id`, `user_id` ) USING BTREE 
) ENGINE = INNODB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户话题';