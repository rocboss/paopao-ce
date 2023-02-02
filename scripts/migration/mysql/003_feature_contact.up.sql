SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

CREATE TABLE `p_contact` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '联系人ID',
  `user_id` bigint unsigned NOT NULL COMMENT '用户ID',
  `friend_id` bigint unsigned NOT NULL COMMENT '好友ID',
  `group_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '好友分组ID:默认为0无分组',
  `remark` varchar(32) NOT NULL DEFAULT '' COMMENT '好友备注',
  `status` tinyint NOT NULL DEFAULT '0' COMMENT '好友状态: 1请求好友, 2已好友, 3拒绝好友, 4已删好友',
  `notice_enable` tinyint NOT NULL DEFAULT '0' COMMENT '是否有消息提醒, 0否, 1是',
  `is_top` tinyint NOT NULL DEFAULT '0' COMMENT '是否置顶, 0否, 1是',
  `is_black` tinyint NOT NULL DEFAULT '0' COMMENT '是否为黑名单, 0否, 1是',
  `is_del` tinyint NOT NULL DEFAULT '0' COMMENT '是否删除好友, 0否, 1是',
  `created_on` bigint unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `modified_on` bigint unsigned NOT NULL DEFAULT '0' COMMENT '修改时间',
  `deleted_on` bigint unsigned NOT NULL DEFAULT '0' COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `idx_user_friend_id` (`user_id`,`friend_id`) USING BTREE,
  KEY `idx_user_friend_status` (`user_id`, `friend_id`, `status`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='联系人';

CREATE TABLE `p_contact_group` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '联系人ID',
  `user_id` int NOT NULL DEFAULT '0' COMMENT '用户id',
  `name` varchar(32) NOT NULL DEFAULT '' COMMENT '分组名称',
  `is_del` tinyint NOT NULL DEFAULT '1' COMMENT '是否删除, 0否, 1是',
  `created_on` bigint unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `modified_on` bigint unsigned NOT NULL DEFAULT '0' COMMENT '修改时间',
  `deleted_on` bigint unsigned NOT NULL DEFAULT '0' COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='联系人分组';

SET FOREIGN_KEY_CHECKS = 1;