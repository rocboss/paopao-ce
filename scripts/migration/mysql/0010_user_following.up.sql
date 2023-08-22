CREATE TABLE `p_following` (
	`id` bigint unsigned NOT NULL AUTO_INCREMENT,
	`user_id` bigint unsigned NOT NULL,
	`follow_id` bigint unsigned NOT NULL,
	`is_del` tinyint NOT NULL DEFAULT 0, -- 是否删除, 0否, 1是
	`created_on` bigint unsigned NOT NULL DEFAULT '0',
	`modified_on` bigint unsigned NOT NULL DEFAULT '0',
	`deleted_on` bigint unsigned NOT NULL DEFAULT '0',
	PRIMARY KEY (`id`) USING BTREE,
	KEY `idx_following_user_follow` (`user_id`,`follow_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
