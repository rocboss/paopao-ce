ALTER TABLE `p_comment` ADD COLUMN `reply_count` int unsigned NOT NULL DEFAULT 0 COMMENT '回复数';

UPDATE p_comment comment 
SET reply_count = (
    SELECT count(*) FROM p_comment_reply reply WHERE reply.comment_id=comment.id AND reply.is_del=0
)
WHERE is_del=0;

CREATE TABLE `p_comment_metric` (
	`id` bigint unsigned NOT NULL AUTO_INCREMENT,
	`comment_id` bigint unsigned NOT NULL,
	`rank_score` bigint unsigned NOT NULL DEFAULT 0,
	`incentive_score` int unsigned NOT NULL DEFAULT 0,
	`decay_factor` int unsigned NOT NULL DEFAULT 0,
	`motivation_factor` int unsigned NOT NULL DEFAULT 0,
	`is_del` tinyint NOT NULL DEFAULT 0,
	`created_on` bigint unsigned NOT NULL DEFAULT 0,
	`modified_on` bigint unsigned NOT NULL DEFAULT 0,
	`deleted_on` bigint unsigned NOT NULL DEFAULT 0,
	PRIMARY KEY (`id`) USING BTREE,
	KEY `idx_comment_metric_comment_id_rank_score` (`comment_id`, `rank_score`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

INSERT INTO p_comment_metric (comment_id, rank_score, created_on) 
SELECT id AS comment_id, 
	reply_count*2 + thumbs_up_count*4 - thumbs_down_count AS rank_score,
	created_on
FROM p_comment
WHERE is_del=0;

CREATE TABLE `p_user_metric` (
	`id` bigint unsigned NOT NULL AUTO_INCREMENT,
	`user_id` bigint unsigned NOT NULL,
	`tweets_count` int unsigned NOT NULL DEFAULT 0,
    `latest_trends_on` bigint unsigned NOT NULL DEFAULT 0 COMMENT '最新动态时间',
	`is_del` tinyint NOT NULL DEFAULT 0,
	`created_on` bigint unsigned NOT NULL DEFAULT 0,
	`modified_on` bigint unsigned NOT NULL DEFAULT 0,
	`deleted_on` bigint unsigned NOT NULL DEFAULT 0,
	PRIMARY KEY (`id`) USING BTREE,
	KEY `idx_user_metric_user_id_tweets_count_trends` (`user_id`, `tweets_count`, `latest_trends_on`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

INSERT INTO p_user_metric (user_id, tweets_count) 
SELECT user_id, count(*) AS tweets_count
FROM p_post
WHERE is_del=0
GROUP BY user_id;