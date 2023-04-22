ALTER TABLE `p_comment` ADD COLUMN `thumbs_up_count` INT unsigned NOT NULL DEFAULT '0' COMMENT '点赞数';
ALTER TABLE `p_comment` ADD COLUMN `thumbs_down_count` INT unsigned NOT NULL DEFAULT '0' COMMENT '点踩数';
ALTER TABLE `p_comment_reply` ADD COLUMN `thumbs_up_count` INT unsigned NOT NULL DEFAULT '0' COMMENT '点赞数';
ALTER TABLE `p_comment_reply` ADD COLUMN `thumbs_down_count` INT unsigned NOT NULL DEFAULT '0' COMMENT '点踩数';

CREATE TABLE `p_tweet_comment_thumbs` (
  `id` BIGINT unsigned NOT NULL AUTO_INCREMENT COMMENT 'thumbs ID',
  `user_id` BIGINT unsigned NOT NULL,
  `tweet_id` BIGINT unsigned NOT NULL COMMENT '推文ID',
  `comment_id` BIGINT unsigned NOT NULL COMMENT '评论ID',
  `reply_id` BIGINT unsigned COMMENT '评论回复ID',
  `comment_type` TINYINT NOT NULL DEFAULT '0' COMMENT '评论类型 0为推文评论、1为评论回复',
  `is_thumbs_up` TINYINT unsigned NOT NULL DEFAULT '0' COMMENT '是否点赞',
  `is_thumbs_down` TINYINT unsigned NOT NULL DEFAULT '0' COMMENT '是否点踩',
  `created_on` BIGINT unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `modified_on` BIGINT unsigned NOT NULL DEFAULT '0' COMMENT '修改时间',
  `deleted_on` BIGINT unsigned NOT NULL DEFAULT '0' COMMENT '删除时间',
  `is_del` TINYINT unsigned NOT NULL DEFAULT '0' COMMENT '是否删除 0 为未删除、1 为已删除',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_tweet_comment_thumbs_uid_tid` (`user_id`, `tweet_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='推文评论点赞';
