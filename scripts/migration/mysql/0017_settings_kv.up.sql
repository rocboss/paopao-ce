RENAME TABLE `p_site_settings` TO `p_site_settings_legacy`;

CREATE TABLE `p_site_settings` (
	`key` VARCHAR(191) NOT NULL,
	`value` TEXT NOT NULL,
	`is_encrypted` TINYINT NOT NULL DEFAULT 0,
	`created_on` BIGINT NOT NULL DEFAULT 0,
	`modified_on` BIGINT NOT NULL DEFAULT 0,
	`deleted_on` BIGINT NOT NULL DEFAULT 0,
	`is_del` TINYINT NOT NULL DEFAULT 0,
	PRIMARY KEY (`key`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='Admin settings overrides';

INSERT INTO `p_site_settings` (`key`, `value`, `is_encrypted`, `created_on`, `modified_on`, `deleted_on`, `is_del`)
SELECT `key`, `value`, 0, created_on, modified_on, 0, 0 FROM (
	SELECT 'web_profile.use_friendship' AS `key`, IF(use_friendship <> 0, 'true', 'false') AS `value`, created_on, modified_on FROM `p_site_settings_legacy` WHERE id = 1
	UNION ALL
	SELECT 'web_profile.enable_trends_bar', IF(enable_trends_bar <> 0, 'true', 'false'), created_on, modified_on FROM `p_site_settings_legacy` WHERE id = 1
	UNION ALL
	SELECT 'web_profile.enable_wallet', IF(enable_wallet <> 0, 'true', 'false'), created_on, modified_on FROM `p_site_settings_legacy` WHERE id = 1
	UNION ALL
	SELECT 'web_profile.allow_tweet_attachment', IF(allow_tweet_attachment <> 0, 'true', 'false'), created_on, modified_on FROM `p_site_settings_legacy` WHERE id = 1
	UNION ALL
	SELECT 'web_profile.allow_tweet_attachment_price', IF(allow_tweet_attachment_price <> 0, 'true', 'false'), created_on, modified_on FROM `p_site_settings_legacy` WHERE id = 1
	UNION ALL
	SELECT 'web_profile.allow_tweet_video', IF(allow_tweet_video <> 0, 'true', 'false'), created_on, modified_on FROM `p_site_settings_legacy` WHERE id = 1
	UNION ALL
	SELECT 'web_profile.default_tweet_max_length', CAST(default_tweet_max_length AS CHAR), created_on, modified_on FROM `p_site_settings_legacy` WHERE id = 1
	UNION ALL
	SELECT 'web_profile.tweet_web_ellipsis_size', CAST(tweet_web_ellipsis_size AS CHAR), created_on, modified_on FROM `p_site_settings_legacy` WHERE id = 1
	UNION ALL
	SELECT 'web_profile.tweet_mobile_ellipsis_size', CAST(tweet_mobile_ellipsis_size AS CHAR), created_on, modified_on FROM `p_site_settings_legacy` WHERE id = 1
	UNION ALL
	SELECT 'web_profile.default_tweet_visibility', default_tweet_visibility, created_on, modified_on FROM `p_site_settings_legacy` WHERE id = 1
	UNION ALL
	SELECT 'web_profile.default_msg_loop_interval', CAST(default_msg_loop_interval AS CHAR), created_on, modified_on FROM `p_site_settings_legacy` WHERE id = 1
	UNION ALL
	SELECT 'web_profile.copyright_top', copyright_top, created_on, modified_on FROM `p_site_settings_legacy` WHERE id = 1
	UNION ALL
	SELECT 'web_profile.copyright_left', copyright_left, created_on, modified_on FROM `p_site_settings_legacy` WHERE id = 1
	UNION ALL
	SELECT 'web_profile.copyright_left_link', copyright_left_link, created_on, modified_on FROM `p_site_settings_legacy` WHERE id = 1
	UNION ALL
	SELECT 'web_profile.copyright_right', copyright_right, created_on, modified_on FROM `p_site_settings_legacy` WHERE id = 1
	UNION ALL
	SELECT 'web_profile.copyright_right_link', copyright_right_link, created_on, modified_on FROM `p_site_settings_legacy` WHERE id = 1
) AS legacy_rows;

DROP TABLE `p_site_settings_legacy`;
