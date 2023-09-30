SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for p_attachment
-- ----------------------------
DROP TABLE IF EXISTS `p_attachment`;
CREATE TABLE `p_attachment` (
	`id` BIGINT NOT NULL AUTO_INCREMENT,
	`user_id` BIGINT NOT NULL DEFAULT '0',
	`file_size` BIGINT NOT NULL,
	`img_width` BIGINT NOT NULL DEFAULT '0',
	`img_height` BIGINT NOT NULL DEFAULT '0',
	`type` tinyint NOT NULL DEFAULT '1' COMMENT '1图片，2视频，3其他附件',
	`content` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
	`created_on` BIGINT NOT NULL DEFAULT '0' COMMENT '创建时间',
	`modified_on` BIGINT NOT NULL DEFAULT '0' COMMENT '修改时间',
	`deleted_on` BIGINT NOT NULL DEFAULT '0' COMMENT '删除时间',
	`is_del` tinyint NOT NULL DEFAULT '0' COMMENT '是否删除 0 为未删除、1 为已删除',
	PRIMARY KEY (`id`) USING BTREE,
	KEY `idx_attachment_user` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=100041 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='附件';

-- ----------------------------
-- Table structure for p_captcha
-- ----------------------------
DROP TABLE IF EXISTS `p_captcha`;
CREATE TABLE `p_captcha` (
	`id` BIGINT NOT NULL AUTO_INCREMENT COMMENT '验证码ID',
	`phone` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '手机号',
	`captcha` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '验证码',
	`use_times` int NOT NULL DEFAULT '0' COMMENT '使用次数',
	`expired_on` BIGINT NOT NULL DEFAULT '0' COMMENT '过期时间',
	`created_on` BIGINT NOT NULL DEFAULT '0' COMMENT '创建时间',
	`modified_on` BIGINT NOT NULL DEFAULT '0' COMMENT '修改时间',
	`deleted_on` BIGINT NOT NULL DEFAULT '0' COMMENT '删除时间',
	`is_del` tinyint NOT NULL DEFAULT '0' COMMENT '是否删除 0 为未删除、1 为已删除',
	PRIMARY KEY (`id`) USING BTREE,
	KEY `idx_captcha_phone` (`phone`) USING BTREE,
	KEY `idx_captcha_expired_on` (`expired_on`) USING BTREE,
	KEY `idx_captcha_use_times` (`use_times`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1021 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='手机验证码';

-- ----------------------------
-- Table structure for p_comment
-- ----------------------------
DROP TABLE IF EXISTS `p_comment`;
CREATE TABLE `p_comment` (
	`id` BIGINT NOT NULL AUTO_INCREMENT COMMENT '评论ID',
	`post_id` BIGINT NOT NULL DEFAULT '0' COMMENT 'POST ID',
	`user_id` BIGINT NOT NULL DEFAULT '0' COMMENT '用户ID',
	`ip` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'IP地址',
	`ip_loc` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'IP城市地址',
	`is_essence` tinyint NOT NULL DEFAULT 0 COMMENT '是否精选',
	`reply_count` int NOT NULL DEFAULT 0 COMMENT '回复数',
	`thumbs_up_count` int NOT NULL DEFAULT 0 COMMENT '点赞数',
	`thumbs_down_count` int NOT NULL DEFAULT 0 COMMENT '点踩数',
	`created_on` BIGINT NOT NULL DEFAULT '0' COMMENT '创建时间',
	`modified_on` BIGINT NOT NULL DEFAULT '0' COMMENT '修改时间',
	`deleted_on` BIGINT NOT NULL DEFAULT '0' COMMENT '删除时间',
	`is_del` tinyint NOT NULL DEFAULT '0' COMMENT '是否删除 0 为未删除、1 为已删除',
	PRIMARY KEY (`id`) USING BTREE,
	KEY `idx_comment_post_id` (`post_id`) USING BTREE,
	KEY `idx_comment_user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=6001736 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='评论';

-- ----------------------------
-- Table structure for p_comment_content
-- ----------------------------
DROP TABLE IF EXISTS `p_comment_content`;
CREATE TABLE `p_comment_content` (
	`id` BIGINT NOT NULL AUTO_INCREMENT COMMENT '内容ID',
	`comment_id` BIGINT NOT NULL DEFAULT '0' COMMENT '评论ID',
	`user_id` BIGINT NOT NULL DEFAULT '0' COMMENT '用户ID',
	`content` varchar(4000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '内容',
	`type` tinyint NOT NULL DEFAULT '2' COMMENT '类型，1标题，2文字段落，3图片地址，4视频地址，5语音地址，6链接地址',
	`sort` BIGINT NOT NULL DEFAULT '100' COMMENT '排序，越小越靠前',
	`created_on` BIGINT NOT NULL DEFAULT '0' COMMENT '创建时间',
	`modified_on` BIGINT NOT NULL DEFAULT '0' COMMENT '修改时间',
	`deleted_on` BIGINT NOT NULL DEFAULT '0' COMMENT '删除时间',
	`is_del` tinyint NOT NULL DEFAULT '0' COMMENT '是否删除 0 为未删除、1 为已删除',
	PRIMARY KEY (`id`) USING BTREE,
	KEY `idx_comment_content_comment_id` (`comment_id`) USING BTREE,
	KEY `idx_comment_content_user_id` (`user_id`) USING BTREE,
	KEY `idx_comment_content_type` (`type`) USING BTREE,
	KEY `idx_comment_content_sort` (`sort`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=11001738 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='评论内容';

-- ----------------------------
-- Table structure for p_comment_reply
-- ----------------------------
DROP TABLE IF EXISTS `p_comment_reply`;
CREATE TABLE `p_comment_reply` (
	`id` BIGINT NOT NULL AUTO_INCREMENT COMMENT '回复ID',
	`comment_id` BIGINT NOT NULL DEFAULT '0' COMMENT '评论ID',
	`user_id` BIGINT NOT NULL DEFAULT '0' COMMENT '用户ID',
	`at_user_id` BIGINT NOT NULL DEFAULT '0' COMMENT '@用户ID',
	`content` varchar(4000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '内容',
	`ip` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'IP地址',
	`ip_loc` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'IP城市地址',
	`thumbs_up_count` int NOT NULL DEFAULT '0' COMMENT '点赞数',
	`thumbs_down_count` int NOT NULL DEFAULT '0' COMMENT '点踩数',
	`created_on` BIGINT NOT NULL DEFAULT '0' COMMENT '创建时间',
	`modified_on` BIGINT NOT NULL DEFAULT '0' COMMENT '修改时间',
	`deleted_on` BIGINT NOT NULL DEFAULT '0' COMMENT '删除时间',
	`is_del` tinyint NOT NULL DEFAULT '0' COMMENT '是否删除 0 为未删除、1 为已删除',
	PRIMARY KEY (`id`) USING BTREE,
	KEY `idx_comment_reply_comment_id` (`comment_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=12000015 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='评论回复';

-- ----------------------------
-- Table structure for p_comment_metric
-- ----------------------------
DROP TABLE IF EXISTS `p_comment_metric`;
CREATE TABLE `p_comment_metric` (
	`id` BIGINT NOT NULL AUTO_INCREMENT,
	`comment_id` BIGINT NOT NULL,
	`rank_score` BIGINT NOT NULL DEFAULT 0,
	`incentive_score` int NOT NULL DEFAULT 0,
	`decay_factor` int NOT NULL DEFAULT 0,
	`motivation_factor` int NOT NULL DEFAULT 0,
	`is_del` tinyint NOT NULL DEFAULT 0,
	`created_on` BIGINT NOT NULL DEFAULT 0,
	`modified_on` BIGINT NOT NULL DEFAULT 0,
	`deleted_on` BIGINT NOT NULL DEFAULT 0,
	PRIMARY KEY (`id`) USING BTREE,
	KEY `idx_comment_metric_comment_id_rank_score` (`comment_id`, `rank_score`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Table structure for p_tweet_comment_thumbs
-- ----------------------------
DROP TABLE IF EXISTS `p_tweet_comment_thumbs`;
CREATE TABLE `p_tweet_comment_thumbs` (
  `id` BIGINT NOT NULL AUTO_INCREMENT COMMENT 'thumbs ID',
  `user_id` BIGINT NOT NULL,
  `tweet_id` BIGINT NOT NULL COMMENT '推文ID',
  `comment_id` BIGINT NOT NULL COMMENT '评论ID',
  `reply_id` BIGINT COMMENT '评论回复ID',
  `comment_type` TINYINT NOT NULL DEFAULT '0' COMMENT '评论类型 0为推文评论、1为评论回复',
  `is_thumbs_up` TINYINT NOT NULL DEFAULT '0' COMMENT '是否点赞',
  `is_thumbs_down` TINYINT NOT NULL DEFAULT '0' COMMENT '是否点踩',
  `created_on` BIGINT NOT NULL DEFAULT '0' COMMENT '创建时间',
  `modified_on` BIGINT NOT NULL DEFAULT '0' COMMENT '修改时间',
  `deleted_on` BIGINT NOT NULL DEFAULT '0' COMMENT '删除时间',
  `is_del` TINYINT NOT NULL DEFAULT '0' COMMENT '是否删除 0 为未删除、1 为已删除',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_tweet_comment_thumbs_uid_tid` (`user_id`, `tweet_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='推文评论点赞';

-- ----------------------------
-- Table structure for p_message
-- ----------------------------
DROP TABLE IF EXISTS `p_message`;
CREATE TABLE `p_message` (
	`id` BIGINT NOT NULL AUTO_INCREMENT COMMENT '消息通知ID',
	`sender_user_id` BIGINT NOT NULL DEFAULT '0' COMMENT '发送方用户ID',
	`receiver_user_id` BIGINT NOT NULL DEFAULT '0' COMMENT '接收方用户ID',
	`type` tinyint NOT NULL DEFAULT '1' COMMENT '通知类型，1动态，2评论，3回复，4私信，99系统通知',
	`brief` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '摘要说明',
	`content` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '详细内容',
	`post_id` BIGINT NOT NULL DEFAULT '0' COMMENT '动态ID',
	`comment_id` BIGINT NOT NULL DEFAULT '0' COMMENT '评论ID',
	`reply_id` BIGINT NOT NULL DEFAULT '0' COMMENT '回复ID',
	`is_read` tinyint NOT NULL DEFAULT '0' COMMENT '是否已读',
	`created_on` BIGINT NOT NULL DEFAULT '0' COMMENT '创建时间',
	`modified_on` BIGINT NOT NULL DEFAULT '0' COMMENT '修改时间',
	`deleted_on` BIGINT NOT NULL DEFAULT '0' COMMENT '删除时间',
	`is_del` tinyint NOT NULL DEFAULT '0' COMMENT '是否删除 0 为未删除、1 为已删除',
	PRIMARY KEY (`id`) USING BTREE,
	KEY `idx_message_receiver_user_id` (`receiver_user_id`) USING BTREE,
	KEY `idx_message_is_read` (`is_read`) USING BTREE,
	KEY `idx_message_type` (`type`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=16000033 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='消息通知';

-- ----------------------------
-- Table structure for p_post
-- ----------------------------
DROP TABLE IF EXISTS `p_post`;
CREATE TABLE `p_post` (
	`id` BIGINT NOT NULL AUTO_INCREMENT COMMENT '主题ID',
	`user_id` BIGINT NOT NULL DEFAULT '0' COMMENT '用户ID',
	`comment_count` BIGINT NOT NULL DEFAULT '0' COMMENT '评论数',
	`collection_count` BIGINT NOT NULL DEFAULT '0' COMMENT '收藏数',
	`upvote_count` BIGINT NOT NULL DEFAULT '0' COMMENT '点赞数',
	`share_count` BIGINT NOT NULL DEFAULT '0' COMMENT '分享数',
	`visibility` tinyint NOT NULL DEFAULT '0' COMMENT '可见性: 0私密 10充电可见 20订阅可见 30保留 40保留 50好友可见 60关注可见 70保留 80保留 90公开',
	`is_top` tinyint NOT NULL DEFAULT '0' COMMENT '是否置顶',
	`is_essence` tinyint NOT NULL DEFAULT '0' COMMENT '是否精华',
	`is_lock` tinyint NOT NULL DEFAULT '0' COMMENT '是否锁定',
	`latest_replied_on` BIGINT NOT NULL DEFAULT '0' COMMENT '最新回复时间',
	`tags` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '标签',
	`attachment_price` BIGINT NOT NULL DEFAULT '0' COMMENT '附件价格(分)',
	`ip` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'IP地址',
	`ip_loc` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'IP城市地址',
	`created_on` BIGINT NOT NULL DEFAULT '0' COMMENT '创建时间',
	`modified_on` BIGINT NOT NULL DEFAULT '0' COMMENT '修改时间',
	`deleted_on` BIGINT NOT NULL DEFAULT '0' COMMENT '删除时间',
	`is_del` tinyint NOT NULL DEFAULT '0' COMMENT '是否删除 0 为未删除、1 为已删除',
	PRIMARY KEY (`id`) USING BTREE,
	KEY `idx_post_user_id` (`user_id`) USING BTREE,
	KEY `idx_post_visibility` (`visibility`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1080017989 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='冒泡/文章';

-- ----------------------------
-- Table structure for p_post_metric
-- ----------------------------
DROP TABLE IF EXISTS `p_post_metric`;
CREATE TABLE `p_post_metric` (
	`id` BIGINT NOT NULL AUTO_INCREMENT,
	`post_id` BIGINT NOT NULL,
	`rank_score` BIGINT NOT NULL DEFAULT 0,
	`incentive_score` int NOT NULL DEFAULT 0,
	`decay_factor` int NOT NULL DEFAULT 0,
	`motivation_factor` int NOT NULL DEFAULT 0,
	`is_del` tinyint NOT NULL DEFAULT 0, -- 是否删除, 0否, 1是
	`created_on` BIGINT NOT NULL DEFAULT '0',
	`modified_on` BIGINT NOT NULL DEFAULT '0',
	`deleted_on` BIGINT NOT NULL DEFAULT '0',
	PRIMARY KEY (`id`) USING BTREE,
	KEY `idx_post_metric_post_id_rank_score` (`post_id`,`rank_score`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Table structure for p_post_attachment_bill
-- ----------------------------
DROP TABLE IF EXISTS `p_post_attachment_bill`;
CREATE TABLE `p_post_attachment_bill` (
	`id` BIGINT NOT NULL AUTO_INCREMENT COMMENT '购买记录ID',
	`post_id` BIGINT NOT NULL DEFAULT '0' COMMENT 'POST ID',
	`user_id` BIGINT NOT NULL DEFAULT '0' COMMENT '用户ID',
	`paid_amount` BIGINT NOT NULL DEFAULT '0' COMMENT '支付金额',
	`created_on` BIGINT NOT NULL DEFAULT '0' COMMENT '创建时间',
	`modified_on` BIGINT NOT NULL DEFAULT '0' COMMENT '修改时间',
	`deleted_on` BIGINT NOT NULL DEFAULT '0' COMMENT '删除时间',
	`is_del` tinyint NOT NULL DEFAULT '0' COMMENT '是否删除 0 为未删除、1 为已删除',
	PRIMARY KEY (`id`) USING BTREE,
	KEY `idx_post_attachment_bill_post_id` (`post_id`) USING BTREE,
	KEY `idx_post_attachment_bill_user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=5000002 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='冒泡/文章附件账单';

-- ----------------------------
-- Table structure for p_post_collection
-- ----------------------------
DROP TABLE IF EXISTS `p_post_collection`;
CREATE TABLE `p_post_collection` (
	`id` BIGINT NOT NULL AUTO_INCREMENT COMMENT '收藏ID',
	`post_id` BIGINT NOT NULL DEFAULT '0' COMMENT 'POST ID',
	`user_id` BIGINT NOT NULL DEFAULT '0' COMMENT '用户ID',
	`created_on` BIGINT NOT NULL DEFAULT '0' COMMENT '创建时间',
	`modified_on` BIGINT NOT NULL DEFAULT '0' COMMENT '修改时间',
	`deleted_on` BIGINT NOT NULL DEFAULT '0' COMMENT '删除时间',
	`is_del` tinyint NOT NULL DEFAULT '0' COMMENT '是否删除 0 为未删除、1 为已删除',
	PRIMARY KEY (`id`) USING BTREE,
	KEY `idx_post_collection_post_id` (`post_id`) USING BTREE,
	KEY `idx_post_collection_user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=6000012 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='冒泡/文章收藏';

-- ----------------------------
-- Table structure for p_post_content
-- ----------------------------
DROP TABLE IF EXISTS `p_post_content`;
CREATE TABLE `p_post_content` (
	`id` BIGINT NOT NULL AUTO_INCREMENT COMMENT '内容ID',
	`post_id` BIGINT NOT NULL DEFAULT '0' COMMENT 'POST ID',
	`user_id` BIGINT NOT NULL DEFAULT '0' COMMENT '用户ID',
	`content` varchar(4000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '内容',
	`type` tinyint NOT NULL DEFAULT '2' COMMENT '类型，1标题，2文字段落，3图片地址，4视频地址，5语音地址，6链接地址，7附件资源，8收费资源',
	`sort` int NOT NULL DEFAULT '100' COMMENT '排序，越小越靠前',
	`created_on` BIGINT NOT NULL DEFAULT '0' COMMENT '创建时间',
	`modified_on` BIGINT NOT NULL DEFAULT '0' COMMENT '修改时间',
	`deleted_on` BIGINT NOT NULL DEFAULT '0' COMMENT '删除时间',
	`is_del` tinyint NOT NULL DEFAULT '0' COMMENT '是否删除 0 为未删除、1 为已删除',
	PRIMARY KEY (`id`) USING BTREE,
	KEY `idx_post_content_post_id` (`post_id`) USING BTREE,
	KEY `idx_post_content_user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=180022546 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='冒泡/文章内容';

-- ----------------------------
-- Table structure for p_post_star
-- ----------------------------
DROP TABLE IF EXISTS `p_post_star`;
CREATE TABLE `p_post_star` (
	`id` BIGINT NOT NULL AUTO_INCREMENT COMMENT '收藏ID',
	`post_id` BIGINT NOT NULL DEFAULT '0' COMMENT 'POST ID',
	`user_id` BIGINT NOT NULL DEFAULT '0' COMMENT '用户ID',
	`created_on` BIGINT NOT NULL DEFAULT '0' COMMENT '创建时间',
	`modified_on` BIGINT NOT NULL DEFAULT '0' COMMENT '修改时间',
	`deleted_on` BIGINT NOT NULL DEFAULT '0' COMMENT '删除时间',
	`is_del` tinyint NOT NULL DEFAULT '0' COMMENT '是否删除 0 为未删除、1 为已删除',
	PRIMARY KEY (`id`) USING BTREE,
	KEY `idx_post_star_post_id` (`post_id`) USING BTREE,
	KEY `idx_post_star_user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=6000028 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='冒泡/文章点赞';

-- ----------------------------
-- Table structure for p_tag
-- ----------------------------
DROP TABLE IF EXISTS `p_tag`;
CREATE TABLE `p_tag` (
	`id` BIGINT NOT NULL AUTO_INCREMENT COMMENT '标签ID',
	`user_id` BIGINT NOT NULL DEFAULT '0' COMMENT '创建者ID',
	`tag` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '标签名',
	`quote_num` BIGINT NOT NULL DEFAULT '0' COMMENT '引用数',
	`created_on` BIGINT NOT NULL DEFAULT '0' COMMENT '创建时间',
	`modified_on` BIGINT NOT NULL DEFAULT '0' COMMENT '修改时间',
	`deleted_on` BIGINT NOT NULL DEFAULT '0' COMMENT '删除时间',
	`is_del` tinyint NOT NULL DEFAULT '0' COMMENT '是否删除 0 为未删除、1 为已删除',
	PRIMARY KEY (`id`) USING BTREE,
	UNIQUE KEY `idx_tag_tag` (`tag`) USING BTREE,
	KEY `idx_tag_user_id` (`user_id`) USING BTREE,
	KEY `idx_tag_quote_num` (`quote_num`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=9000065 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='标签';

-- ----------------------------
-- Table structure for p_topic_user
-- ----------------------------
DROP TABLE IF EXISTS `p_topic_user`;
CREATE TABLE `p_topic_user` (
	`id` BIGINT NOT NULL AUTO_INCREMENT,
	`topic_id` BIGINT NOT NULL COMMENT '标签ID',
	`user_id` BIGINT NOT NULL COMMENT '创建者ID',
	`alias_name` VARCHAR ( 255 ) COMMENT '别名',
	`remark` VARCHAR ( 512 ) COMMENT '备注',
	`quote_num` BIGINT COMMENT '引用数',
	`is_top` TINYINT NOT NULL DEFAULT '0' COMMENT '是否置顶 0 为未置顶、1 为已置顶',
	`created_on` BIGINT NOT NULL DEFAULT '0' COMMENT '创建时间',
	`modified_on` BIGINT NOT NULL DEFAULT '0' COMMENT '修改时间',
	`deleted_on` BIGINT NOT NULL DEFAULT '0' COMMENT '删除时间',
	`is_del` TINYINT NOT NULL DEFAULT '0' COMMENT '是否删除 0 为未删除、1 为已删除',
	`reserve_a` VARCHAR ( 255 ) COMMENT '保留字段a',
	`reserve_b` VARCHAR ( 255 ) COMMENT '保留字段b',
	PRIMARY KEY ( `id` ) USING BTREE,
	UNIQUE KEY `idx_topic_user_uid_tid` ( `topic_id`, `user_id` ) USING BTREE 
) ENGINE = INNODB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户话题';

-- ----------------------------
-- Table structure for p_user
-- ----------------------------
DROP TABLE IF EXISTS `p_user`;
CREATE TABLE `p_user` (
	`id` BIGINT NOT NULL AUTO_INCREMENT COMMENT '用户ID',
	`nickname` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '昵称',
	`username` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户名',
	`phone` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '手机号',
	`password` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT 'MD5密码',
	`salt` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '盐值',
	`status` tinyint NOT NULL DEFAULT '1' COMMENT '状态，1正常，2停用',
	`avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '用户头像',
	`balance` BIGINT NOT NULL COMMENT '用户余额（分）',
	`is_admin` tinyint NOT NULL DEFAULT '0' COMMENT '是否管理员',
	`created_on` BIGINT NOT NULL DEFAULT '0' COMMENT '创建时间',
	`modified_on` BIGINT NOT NULL DEFAULT '0' COMMENT '修改时间',
	`deleted_on` BIGINT NOT NULL DEFAULT '0' COMMENT '删除时间',
	`is_del` tinyint NOT NULL DEFAULT '0' COMMENT '是否删除 0 为未删除、1 为已删除',
	PRIMARY KEY (`id`) USING BTREE,
	UNIQUE KEY `idx_user_username` (`username`) USING BTREE,
	KEY `idx_user_phone` (`phone`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=100058 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户';

-- ----------------------------
-- Table structure for p_user_metric
-- ----------------------------
DROP TABLE IF EXISTS `p_user_metric`;
CREATE TABLE `p_user_metric` (
	`id` BIGINT NOT NULL AUTO_INCREMENT,
	`user_id` BIGINT NOT NULL,
	`tweets_count` int NOT NULL DEFAULT 0,
    `latest_trends_on` BIGINT NOT NULL DEFAULT 0 COMMENT '最新动态时间',
	`is_del` tinyint NOT NULL DEFAULT 0,
	`created_on` BIGINT NOT NULL DEFAULT 0,
	`modified_on` BIGINT NOT NULL DEFAULT 0,
	`deleted_on` BIGINT NOT NULL DEFAULT 0,
	PRIMARY KEY (`id`) USING BTREE,
	KEY `idx_user_metric_user_id_tweets_count_trends` (`user_id`, `tweets_count`, `latest_trends_on`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Table structure for p_following
-- ----------------------------
DROP TABLE IF EXISTS `p_following`;
CREATE TABLE `p_following` (
	`id` BIGINT NOT NULL AUTO_INCREMENT,
	`user_id` BIGINT NOT NULL,
	`follow_id` BIGINT NOT NULL,
	`is_del` tinyint NOT NULL DEFAULT 0, -- 是否删除, 0否, 1是
	`created_on` BIGINT NOT NULL DEFAULT '0',
	`modified_on` BIGINT NOT NULL DEFAULT '0',
	`deleted_on` BIGINT NOT NULL DEFAULT '0',
	PRIMARY KEY (`id`) USING BTREE,
	KEY `idx_following_user_follow` (`user_id`,`follow_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Table structure for p_contact
-- ----------------------------
DROP TABLE IF EXISTS `p_contact`;
CREATE TABLE `p_contact` (
	`id` BIGINT NOT NULL AUTO_INCREMENT COMMENT '联系人ID',
	`user_id` BIGINT NOT NULL COMMENT '用户ID',
	`friend_id` BIGINT NOT NULL COMMENT '好友ID',
	`group_id` BIGINT NOT NULL DEFAULT '0' COMMENT '好友分组ID:默认为0无分组',
	`remark` varchar(32) NOT NULL DEFAULT '' COMMENT '好友备注',
	`status` tinyint NOT NULL DEFAULT '0' COMMENT '好友状态: 1请求好友, 2已好友, 3拒绝好友, 4已删好友',
	`is_top` tinyint NOT NULL DEFAULT '0' COMMENT '是否置顶, 0否, 1是',
	`is_black` tinyint NOT NULL DEFAULT '0' COMMENT '是否为黑名单, 0否, 1是',
	`is_del` tinyint NOT NULL DEFAULT '0' COMMENT '是否删除好友, 0否, 1是',
	`notice_enable` tinyint NOT NULL DEFAULT '0' COMMENT '是否有消息提醒, 0否, 1是',
	`created_on` BIGINT NOT NULL DEFAULT '0' COMMENT '创建时间',
	`modified_on` BIGINT NOT NULL DEFAULT '0' COMMENT '修改时间',
	`deleted_on` BIGINT NOT NULL DEFAULT '0' COMMENT '删除时间',
	PRIMARY KEY (`id`) USING BTREE,
	UNIQUE KEY `idx_contact_user_friend` (`user_id`,`friend_id`) USING BTREE,
	KEY `idx_contact_user_friend_status` (`user_id`, `friend_id`, `status`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='联系人';

-- ----------------------------
-- Table structure for p_contact_group
-- ----------------------------
DROP TABLE IF EXISTS `p_contact_group`;
CREATE TABLE `p_contact_group` (
	`id` BIGINT NOT NULL AUTO_INCREMENT COMMENT '联系人ID',
	`user_id` int NOT NULL DEFAULT '0' COMMENT '用户id',
	`name` varchar(32) NOT NULL DEFAULT '' COMMENT '分组名称',
	`is_del` tinyint NOT NULL DEFAULT '1' COMMENT '是否删除, 0否, 1是',
	`created_on` BIGINT NOT NULL DEFAULT '0' COMMENT '创建时间',
	`modified_on` BIGINT NOT NULL DEFAULT '0' COMMENT '修改时间',
	`deleted_on` BIGINT NOT NULL DEFAULT '0' COMMENT '删除时间',
	PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='联系人分组';

-- ----------------------------
-- Table structure for p_wallet_recharge
-- ----------------------------
DROP TABLE IF EXISTS `p_wallet_recharge`;
CREATE TABLE `p_wallet_recharge` (
	`id` BIGINT NOT NULL AUTO_INCREMENT COMMENT '充值ID',
	`user_id` BIGINT NOT NULL DEFAULT '0' COMMENT '用户ID',
	`amount` BIGINT NOT NULL DEFAULT '0' COMMENT '充值金额',
	`trade_no` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '支付宝订单号',
	`trade_status` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '交易状态',
	`created_on` BIGINT NOT NULL DEFAULT '0' COMMENT '创建时间',
	`modified_on` BIGINT NOT NULL DEFAULT '0' COMMENT '修改时间',
	`deleted_on` BIGINT NOT NULL DEFAULT '0' COMMENT '删除时间',
	`is_del` tinyint NOT NULL DEFAULT '0' COMMENT '是否删除 0 为未删除、1 为已删除',
	PRIMARY KEY (`id`) USING BTREE,
	KEY `idx_wallet_recharge_user_id` (`user_id`) USING BTREE,
	KEY `idx_wallet_recharge_trade_no` (`trade_no`) USING BTREE,
	KEY `idx_wallet_recharge_trade_status` (`trade_status`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=10023 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='钱包流水';

-- ----------------------------
-- Table structure for p_wallet_statement
-- ----------------------------
DROP TABLE IF EXISTS `p_wallet_statement`;
CREATE TABLE `p_wallet_statement` (
	`id` BIGINT NOT NULL AUTO_INCREMENT COMMENT '账单ID',
	`user_id` BIGINT NOT NULL DEFAULT '0' COMMENT '用户ID',
	`change_amount` BIGINT NOT NULL DEFAULT '0' COMMENT '变动金额',
	`balance_snapshot` BIGINT NOT NULL DEFAULT '0' COMMENT '资金快照',
	`reason` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '变动原因',
	`post_id` BIGINT NOT NULL DEFAULT '0' COMMENT '关联动态',
	`created_on` BIGINT NOT NULL DEFAULT '0' COMMENT '创建时间',
	`modified_on` BIGINT NOT NULL DEFAULT '0' COMMENT '修改时间',
	`deleted_on` BIGINT NOT NULL DEFAULT '0' COMMENT '删除时间',
	`is_del` tinyint NOT NULL DEFAULT '0' COMMENT '是否删除 0 为未删除、1 为已删除',
	PRIMARY KEY (`id`) USING BTREE,
	KEY `idx_wallet_statement_user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=10010 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='钱包流水';

DROP VIEW IF EXISTS p_post_by_media;
CREATE VIEW p_post_by_media AS 
SELECT post.* 
FROM
	( SELECT DISTINCT post_id FROM p_post_content WHERE ( TYPE = 3 OR TYPE = 4 OR TYPE = 7 OR TYPE = 8 ) AND is_del = 0 ) media
	JOIN p_post post ON media.post_id = post.ID 
WHERE
	post.is_del = 0;

DROP VIEW IF EXISTS p_post_by_comment;
CREATE VIEW p_post_by_comment AS 
SELECT P.*, C.user_id comment_user_id
FROM
	(
	SELECT
		post_id,
		user_id
	FROM
		p_comment 
	WHERE
		is_del = 0 UNION
	SELECT
		post_id,
		reply.user_id user_id
	FROM
		p_comment_reply reply
		JOIN p_comment COMMENT ON reply.comment_id = COMMENT.ID 
	WHERE
		reply.is_del = 0 
		AND COMMENT.is_del = 0 
	)
	C JOIN p_post P ON C.post_id = P.ID 
WHERE
	P.is_del = 0;

DROP VIEW IF EXISTS p_user_relation;
CREATE VIEW p_user_relation AS 
SELECT user_id, friend_id he_uid, 5 AS style 
FROM p_contact WHERE status=2 AND is_del=0
UNION
SELECT user_id, follow_id he_uid, 10 AS style 
FROM p_following WHERE is_del=0;

SET FOREIGN_KEY_CHECKS = 1;
