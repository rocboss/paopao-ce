-- 创建长文章相关表

-- 长文章主表
CREATE TABLE IF NOT EXISTS `p_articles` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) NOT NULL COMMENT '用户ID',
  `series_id` bigint(20) DEFAULT NULL COMMENT '系列ID，如果属于某个系列',
  `title` varchar(200) NOT NULL COMMENT '文章标题',
  `summary` text COMMENT '文章摘要',
  `word_count` int(11) NOT NULL DEFAULT 0 COMMENT '字数统计',
  `visibility` tinyint(4) NOT NULL DEFAULT 0 COMMENT '可见性: 0私密 10充电可见 20订阅可见 30保留 40保留 50好友可见 60关注可见 70保留 80保留 90公开',
  `price` bigint(20) DEFAULT 0 COMMENT '付费价格(分)',
  `is_top` int(11) NOT NULL DEFAULT 0 COMMENT '是否置顶',
  `is_essence` int(11) NOT NULL DEFAULT 0 COMMENT '是否精华',
  `is_lock` int(11) NOT NULL DEFAULT 0 COMMENT '是否锁定',
  `comment_count` bigint(20) NOT NULL DEFAULT 0 COMMENT '评论数',
  `collection_count` bigint(20) NOT NULL DEFAULT 0 COMMENT '收藏数',
  `share_count` bigint(20) NOT NULL DEFAULT 0 COMMENT '分享数',
  `upvote_count` bigint(20) NOT NULL DEFAULT 0 COMMENT '点赞数',
  `view_count` bigint(20) NOT NULL DEFAULT 0 COMMENT '阅读数',
  `latest_replied_on` int(11) NOT NULL DEFAULT 0 COMMENT '最后回复时间',
  `tags` varchar(500) DEFAULT '' COMMENT '标签，逗号分隔',
  `ip` varchar(64) DEFAULT '' COMMENT 'IP地址',
  `ip_loc` varchar(64) DEFAULT '' COMMENT 'IP地理位置',
  `created_on` int(11) NOT NULL DEFAULT 0,
  `modified_on` int(11) NOT NULL DEFAULT 0,
  `is_del` int(11) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_series_id` (`series_id`),
  KEY `idx_visibility` (`visibility`),
  KEY `idx_is_top` (`is_top`),
  KEY `idx_is_essence` (`is_essence`),
  KEY `idx_created_on` (`created_on`),
  KEY `idx_modified_on` (`modified_on`),
  KEY `idx_tags` (`tags`(191))
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='长文章表';

-- 长文章内容表
CREATE TABLE IF NOT EXISTS `p_article_contents` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `article_id` bigint(20) NOT NULL COMMENT '文章ID',
  `user_id` bigint(20) NOT NULL COMMENT '用户ID',
  `content` longtext NOT NULL COMMENT '内容',
  `content_type` int(11) NOT NULL DEFAULT 1 COMMENT '内容类型: 1标题 2文字段落 3图片地址 4视频地址 5语音地址 6链接地址 7附件资源 8收费附件',
  `sort` bigint(20) NOT NULL DEFAULT 0 COMMENT '排序',
  `created_on` int(11) NOT NULL DEFAULT 0,
  `modified_on` int(11) NOT NULL DEFAULT 0,
  `is_del` int(11) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`),
  KEY `idx_article_id` (`article_id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_content_type` (`content_type`),
  KEY `idx_sort` (`sort`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='长文章内容表';

-- 长文章系列表
CREATE TABLE IF NOT EXISTS `p_article_series` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) NOT NULL COMMENT '用户ID',
  `title` varchar(200) NOT NULL COMMENT '系列标题',
  `description` text COMMENT '系列描述',
  `cover` varchar(500) DEFAULT '' COMMENT '封面图片',
  `article_count` int(11) NOT NULL DEFAULT 0 COMMENT '文章数量',
  `total_word_count` int(11) NOT NULL DEFAULT 0 COMMENT '总字数',
  `is_completed` tinyint(4) NOT NULL DEFAULT 0 COMMENT '是否完结',
  `visibility` tinyint(4) NOT NULL DEFAULT 0 COMMENT '可见性',
  `tags` varchar(500) DEFAULT '' COMMENT '标签，逗号分隔',
  `created_on` int(11) NOT NULL DEFAULT 0,
  `modified_on` int(11) NOT NULL DEFAULT 0,
  `is_del` int(11) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_visibility` (`visibility`),
  KEY `idx_is_completed` (`is_completed`),
  KEY `idx_created_on` (`created_on`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='长文章系列表';

-- 系列文章关联表
CREATE TABLE IF NOT EXISTS `p_article_series_items` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `series_id` bigint(20) NOT NULL COMMENT '系列ID',
  `article_id` bigint(20) NOT NULL COMMENT '文章ID',
  `sort` int(11) NOT NULL DEFAULT 0 COMMENT '排序',
  `created_on` int(11) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_series_article` (`series_id`, `article_id`),
  KEY `idx_series_id` (`series_id`),
  KEY `idx_article_id` (`article_id`),
  KEY `idx_sort` (`sort`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='系列文章关联表';

-- 长文章收藏表
CREATE TABLE IF NOT EXISTS `p_article_collections` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) NOT NULL COMMENT '用户ID',
  `article_id` bigint(20) NOT NULL COMMENT '文章ID',
  `created_on` int(11) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_user_article` (`user_id`, `article_id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_article_id` (`article_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='长文章收藏表';

-- 长文章点赞表
CREATE TABLE IF NOT EXISTS `p_article_stars` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) NOT NULL COMMENT '用户ID',
  `article_id` bigint(20) NOT NULL COMMENT '文章ID',
  `created_on` int(11) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_user_article` (`user_id`, `article_id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_article_id` (`article_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='长文章点赞表';

-- 长文章评论表
CREATE TABLE IF NOT EXISTS `p_article_comments` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) NOT NULL COMMENT '用户ID',
  `article_id` bigint(20) NOT NULL COMMENT '文章ID',
  `content` text NOT NULL COMMENT '评论内容',
  `ip` varchar(64) DEFAULT '' COMMENT 'IP地址',
  `ip_loc` varchar(64) DEFAULT '' COMMENT 'IP地理位置',
  `created_on` int(11) NOT NULL DEFAULT 0,
  `modified_on` int(11) NOT NULL DEFAULT 0,
  `is_del` int(11) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_article_id` (`article_id`),
  KEY `idx_created_on` (`created_on`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='长文章评论表';

-- 长文章评论回复表
CREATE TABLE IF NOT EXISTS `p_article_comment_replies` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) NOT NULL COMMENT '用户ID',
  `article_id` bigint(20) NOT NULL COMMENT '文章ID',
  `comment_id` bigint(20) NOT NULL COMMENT '评论ID',
  `reply_to_id` bigint(20) DEFAULT NULL COMMENT '回复目标ID',
  `content` text NOT NULL COMMENT '回复内容',
  `ip` varchar(64) DEFAULT '' COMMENT 'IP地址',
  `ip_loc` varchar(64) DEFAULT '' COMMENT 'IP地理位置',
  `created_on` int(11) NOT NULL DEFAULT 0,
  `modified_on` int(11) NOT NULL DEFAULT 0,
  `is_del` int(11) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_article_id` (`article_id`),
  KEY `idx_comment_id` (`comment_id`),
  KEY `idx_reply_to_id` (`reply_to_id`),
  KEY `idx_created_on` (`created_on`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='长文章评论回复表';
