CREATE TABLE `p_post_metric` (
	`id` bigint unsigned NOT NULL AUTO_INCREMENT,
	`post_id` bigint unsigned NOT NULL,
	`rank_score` bigint unsigned NOT NULL DEFAULT 0,
	`incentive_score` int unsigned NOT NULL DEFAULT 0,
	`decay_factor` int unsigned NOT NULL DEFAULT 0,
	`motivation_factor` int unsigned NOT NULL DEFAULT 0,
	`is_del` tinyint NOT NULL DEFAULT 0, -- 是否删除, 0否, 1是
	`created_on` bigint unsigned NOT NULL DEFAULT '0',
	`modified_on` bigint unsigned NOT NULL DEFAULT '0',
	`deleted_on` bigint unsigned NOT NULL DEFAULT '0',
	PRIMARY KEY (`id`) USING BTREE,
	KEY `idx_post_metric_post_id_rank_score` (`post_id`,`rank_score`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

INSERT INTO p_post_metric (post_id, rank_score, created_on) 
SELECT id AS post_id, 
	comment_count + upvote_count*2 + collection_count*4 AS rank_score,
	created_on
FROM p_post
WHERE is_del=0;

-- 原来的可见性: 0公开 1私密 2好友可见 3关注可见
-- 现在的可见性: 0私密 10充电可见 20订阅可见 30保留 40保留 50好友可见 60关注可见 70保留 80保留 90公开
UPDATE p_post a, p_post b 
SET a.visibility = (
	CASE b.visibility 
		WHEN 0 THEN 90 
		WHEN 1 THEN 0 
		WHEN 2 THEN 50 
		WHEN 3 THEN 60 
		ELSE 0
	END 
)
WHERE a.ID = b.ID;
