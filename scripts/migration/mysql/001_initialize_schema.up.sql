SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

CREATE TABLE `p_attachment` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `user_id` bigint unsigned NOT NULL DEFAULT '0',
  `file_size` bigint unsigned NOT NULL,
  `img_width` bigint unsigned NOT NULL DEFAULT '0',
  `img_height` bigint unsigned NOT NULL DEFAULT '0',
  `type` tinyint unsigned NOT NULL DEFAULT '1' COMMENT '1图片，2视频，3其他附件',
  `content` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `created_on` bigint unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `modified_on` bigint unsigned NOT NULL DEFAULT '0' COMMENT '修改时间',
  `deleted_on` bigint unsigned NOT NULL DEFAULT '0' COMMENT '删除时间',
  `is_del` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '是否删除 0 为未删除、1 为已删除',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_user` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=100041 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='附件';

CREATE TABLE `p_captcha` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '验证码ID',
  `phone` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '手机号',
  `captcha` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '验证码',
  `use_times` int unsigned NOT NULL DEFAULT '0' COMMENT '使用次数',
  `expired_on` bigint unsigned NOT NULL DEFAULT '0' COMMENT '过期时间',
  `created_on` bigint unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `modified_on` bigint unsigned NOT NULL DEFAULT '0' COMMENT '修改时间',
  `deleted_on` bigint unsigned NOT NULL DEFAULT '0' COMMENT '删除时间',
  `is_del` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '是否删除 0 为未删除、1 为已删除',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_phone` (`phone`) USING BTREE,
  KEY `idx_expired_on` (`expired_on`) USING BTREE,
  KEY `idx_use_times` (`use_times`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1021 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='手机验证码';

CREATE TABLE `p_comment` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '评论ID',
  `post_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT 'POST ID',
  `user_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '用户ID',
  `ip` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'IP地址',
  `ip_loc` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'IP城市地址',
  `created_on` bigint unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `modified_on` bigint unsigned NOT NULL DEFAULT '0' COMMENT '修改时间',
  `deleted_on` bigint unsigned NOT NULL DEFAULT '0' COMMENT '删除时间',
  `is_del` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '是否删除 0 为未删除、1 为已删除',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_post` (`post_id`) USING BTREE,
  KEY `idx_user` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=6001736 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='评论';

CREATE TABLE `p_comment_content` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '内容ID',
  `comment_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '评论ID',
  `user_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '用户ID',
  `content` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '内容',
  `type` tinyint unsigned NOT NULL DEFAULT '2' COMMENT '类型，1标题，2文字段落，3图片地址，4视频地址，5语音地址，6链接地址',
  `sort` bigint unsigned NOT NULL DEFAULT '100' COMMENT '排序，越小越靠前',
  `created_on` bigint unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `modified_on` bigint unsigned NOT NULL DEFAULT '0' COMMENT '修改时间',
  `deleted_on` bigint unsigned NOT NULL DEFAULT '0' COMMENT '删除时间',
  `is_del` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '是否删除 0 为未删除、1 为已删除',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_reply` (`comment_id`) USING BTREE,
  KEY `idx_user` (`user_id`) USING BTREE,
  KEY `idx_type` (`type`) USING BTREE,
  KEY `idx_sort` (`sort`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=11001738 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='评论内容';

CREATE TABLE `p_comment_reply` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '回复ID',
  `comment_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '评论ID',
  `user_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '用户ID',
  `at_user_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '@用户ID',
  `content` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '内容',
  `ip` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'IP地址',
  `ip_loc` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'IP城市地址',
  `created_on` bigint unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `modified_on` bigint unsigned NOT NULL DEFAULT '0' COMMENT '修改时间',
  `deleted_on` bigint unsigned NOT NULL DEFAULT '0' COMMENT '删除时间',
  `is_del` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '是否删除 0 为未删除、1 为已删除',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_comment` (`comment_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=12000015 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='评论回复';

CREATE TABLE `p_message` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '消息通知ID',
  `sender_user_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '发送方用户ID',
  `receiver_user_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '接收方用户ID',
  `type` tinyint unsigned NOT NULL DEFAULT '1' COMMENT '通知类型，1动态，2评论，3回复，4私信，99系统通知',
  `brief` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '摘要说明',
  `content` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '详细内容',
  `post_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '动态ID',
  `comment_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '评论ID',
  `reply_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '回复ID',
  `is_read` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '是否已读',
  `created_on` bigint unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `modified_on` bigint unsigned NOT NULL DEFAULT '0' COMMENT '修改时间',
  `deleted_on` bigint unsigned NOT NULL DEFAULT '0' COMMENT '删除时间',
  `is_del` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '是否删除 0 为未删除、1 为已删除',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_receiver` (`receiver_user_id`) USING BTREE,
  KEY `idx_is_read` (`is_read`) USING BTREE,
  KEY `idx_type` (`type`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=16000033 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='消息通知';

CREATE TABLE `p_post` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主题ID',
  `user_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '用户ID',
  `comment_count` bigint unsigned NOT NULL DEFAULT '0' COMMENT '评论数',
  `collection_count` bigint unsigned NOT NULL DEFAULT '0' COMMENT '收藏数',
  `upvote_count` bigint unsigned NOT NULL DEFAULT '0' COMMENT '点赞数',
  `is_top` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '是否置顶',
  `is_essence` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '是否精华',
  `is_lock` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '是否锁定',
  `latest_replied_on` bigint unsigned NOT NULL DEFAULT '0' COMMENT '最新回复时间',
  `tags` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '标签',
  `attachment_price` bigint unsigned NOT NULL DEFAULT '0' COMMENT '附件价格(分)',
  `ip` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'IP地址',
  `ip_loc` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'IP城市地址',
  `created_on` bigint unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `modified_on` bigint unsigned NOT NULL DEFAULT '0' COMMENT '修改时间',
  `deleted_on` bigint unsigned NOT NULL DEFAULT '0' COMMENT '删除时间',
  `is_del` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '是否删除 0 为未删除、1 为已删除',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_user` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1080017989 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='冒泡/文章';

CREATE TABLE `p_post_attachment_bill` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '购买记录ID',
  `post_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT 'POST ID',
  `user_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '用户ID',
  `paid_amount` bigint unsigned NOT NULL DEFAULT '0' COMMENT '支付金额',
  `created_on` bigint unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `modified_on` bigint unsigned NOT NULL DEFAULT '0' COMMENT '修改时间',
  `deleted_on` bigint unsigned NOT NULL DEFAULT '0' COMMENT '删除时间',
  `is_del` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '是否删除 0 为未删除、1 为已删除',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_post` (`post_id`) USING BTREE,
  KEY `idx_user` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=5000002 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='冒泡/文章附件账单';

CREATE TABLE `p_post_collection` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '收藏ID',
  `post_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT 'POST ID',
  `user_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '用户ID',
  `created_on` bigint unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `modified_on` bigint unsigned NOT NULL DEFAULT '0' COMMENT '修改时间',
  `deleted_on` bigint unsigned NOT NULL DEFAULT '0' COMMENT '删除时间',
  `is_del` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '是否删除 0 为未删除、1 为已删除',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_post` (`post_id`) USING BTREE,
  KEY `idx_user` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=6000012 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='冒泡/文章收藏';

CREATE TABLE `p_post_content` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '内容ID',
  `post_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT 'POST ID',
  `user_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '用户ID',
  `content` varchar(2000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '内容',
  `type` tinyint unsigned NOT NULL DEFAULT '2' COMMENT '类型，1标题，2文字段落，3图片地址，4视频地址，5语音地址，6链接地址，7附件资源，8收费资源',
  `sort` int unsigned NOT NULL DEFAULT '100' COMMENT '排序，越小越靠前',
  `created_on` bigint unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `modified_on` bigint unsigned NOT NULL DEFAULT '0' COMMENT '修改时间',
  `deleted_on` bigint unsigned NOT NULL DEFAULT '0' COMMENT '删除时间',
  `is_del` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '是否删除 0 为未删除、1 为已删除',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_post` (`post_id`) USING BTREE,
  KEY `idx_user` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=180022546 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='冒泡/文章内容';

CREATE TABLE `p_post_star` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '收藏ID',
  `post_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT 'POST ID',
  `user_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '用户ID',
  `created_on` bigint unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `modified_on` bigint unsigned NOT NULL DEFAULT '0' COMMENT '修改时间',
  `deleted_on` bigint unsigned NOT NULL DEFAULT '0' COMMENT '删除时间',
  `is_del` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '是否删除 0 为未删除、1 为已删除',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_post` (`post_id`) USING BTREE,
  KEY `idx_user` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=6000028 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='冒泡/文章点赞';

CREATE TABLE `p_tag` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '标签ID',
  `user_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '创建者ID',
  `tag` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '标签名',
  `quote_num` bigint unsigned NOT NULL DEFAULT '0' COMMENT '引用数',
  `created_on` bigint unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `modified_on` bigint unsigned NOT NULL DEFAULT '0' COMMENT '修改时间',
  `deleted_on` bigint unsigned NOT NULL DEFAULT '0' COMMENT '删除时间',
  `is_del` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '是否删除 0 为未删除、1 为已删除',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `idx_tag` (`tag`) USING BTREE,
  KEY `idx_user` (`user_id`) USING BTREE,
  KEY `idx_num` (`quote_num`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=9000065 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='标签';

CREATE TABLE `p_user` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '用户ID',
  `nickname` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '昵称',
  `username` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户名',
  `phone` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '手机号',
  `password` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'MD5密码',
  `salt` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '盐值',
  `status` tinyint unsigned NOT NULL DEFAULT '1' COMMENT '状态，1正常，2停用',
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户头像',
  `balance` bigint unsigned NOT NULL COMMENT '用户余额（分）',
  `is_admin` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '是否管理员',
  `created_on` bigint unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `modified_on` bigint unsigned NOT NULL DEFAULT '0' COMMENT '修改时间',
  `deleted_on` bigint unsigned NOT NULL DEFAULT '0' COMMENT '删除时间',
  `is_del` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '是否删除 0 为未删除、1 为已删除',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `idx_username` (`username`) USING BTREE,
  KEY `idx_phone` (`phone`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=100058 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户';

CREATE TABLE `p_wallet_recharge` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '充值ID',
  `user_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '用户ID',
  `amount` bigint NOT NULL DEFAULT '0' COMMENT '充值金额',
  `trade_no` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '支付宝订单号',
  `trade_status` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '交易状态',
  `created_on` bigint unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `modified_on` bigint unsigned NOT NULL DEFAULT '0' COMMENT '修改时间',
  `deleted_on` bigint unsigned NOT NULL DEFAULT '0' COMMENT '删除时间',
  `is_del` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '是否删除 0 为未删除、1 为已删除',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_user` (`user_id`) USING BTREE,
  KEY `idx_trade_no` (`trade_no`) USING BTREE,
  KEY `idx_trade_status` (`trade_status`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=10023 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='钱包流水';

CREATE TABLE `p_wallet_statement` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '账单ID',
  `user_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '用户ID',
  `change_amount` bigint NOT NULL DEFAULT '0' COMMENT '变动金额',
  `balance_snapshot` bigint NOT NULL DEFAULT '0' COMMENT '资金快照',
  `reason` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '变动原因',
  `post_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '关联动态',
  `created_on` bigint unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `modified_on` bigint unsigned NOT NULL DEFAULT '0' COMMENT '修改时间',
  `deleted_on` bigint unsigned NOT NULL DEFAULT '0' COMMENT '删除时间',
  `is_del` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '是否删除 0 为未删除、1 为已删除',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_user` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=10010 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='钱包流水';

SET FOREIGN_KEY_CHECKS = 1;
