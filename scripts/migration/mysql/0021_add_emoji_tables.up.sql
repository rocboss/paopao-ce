SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

CREATE TABLE `p_emoji` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '表情包ID',
  `user_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '上传用户ID',
  `name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '表情包名称',
  `url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '表情包URL',
  `width` int unsigned NOT NULL DEFAULT '0' COMMENT '宽度',
  `height` int unsigned NOT NULL DEFAULT '0' COMMENT '高度',
  `size` bigint unsigned NOT NULL DEFAULT '0' COMMENT '文件大小',
  `type` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '文件类型',
  `collection_count` bigint unsigned NOT NULL DEFAULT '0' COMMENT '收藏数',
  `created_on` bigint unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `modified_on` bigint unsigned NOT NULL DEFAULT '0' COMMENT '修改时间',
  `deleted_on` bigint unsigned NOT NULL DEFAULT '0' COMMENT '删除时间',
  `is_del` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '是否删除 0 为未删除、1 为已删除',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_user` (`user_id`) USING BTREE,
  KEY `idx_collection` (`collection_count`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='表情包';

CREATE TABLE `p_user_emoji` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '记录ID',
  `user_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '用户ID',
  `emoji_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '表情包ID',
  `created_on` bigint unsigned NOT NULL DEFAULT '0' COMMENT '收藏时间',
  `modified_on` bigint unsigned NOT NULL DEFAULT '0' COMMENT '修改时间',
  `deleted_on` bigint unsigned NOT NULL DEFAULT '0' COMMENT '删除时间',
  `is_del` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '是否删除 0 为未删除、1 为已删除',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `idx_user_emoji` (`user_id`,`emoji_id`) USING BTREE,
  KEY `idx_emoji` (`emoji_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户表情包收藏';

SET FOREIGN_KEY_CHECKS = 1;