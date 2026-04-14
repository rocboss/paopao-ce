ALTER TABLE p_site_settings RENAME TO p_site_settings_legacy;

CREATE TABLE p_site_settings (
	key VARCHAR(191) PRIMARY KEY,
	value TEXT NOT NULL,
	is_encrypted BOOLEAN NOT NULL DEFAULT false,
	created_on BIGINT NOT NULL DEFAULT 0,
	modified_on BIGINT NOT NULL DEFAULT 0,
	deleted_on BIGINT NOT NULL DEFAULT 0,
	is_del SMALLINT NOT NULL DEFAULT 0
);

INSERT INTO p_site_settings (key, value, is_encrypted, created_on, modified_on, deleted_on, is_del)
SELECT key, value, false, created_on, modified_on, 0, 0 FROM (
	SELECT 'web_profile.use_friendship' AS key, CASE WHEN use_friendship THEN 'true' ELSE 'false' END AS value, created_on, modified_on FROM p_site_settings_legacy WHERE id = 1
	UNION ALL
	SELECT 'web_profile.enable_trends_bar', CASE WHEN enable_trends_bar THEN 'true' ELSE 'false' END, created_on, modified_on FROM p_site_settings_legacy WHERE id = 1
	UNION ALL
	SELECT 'web_profile.enable_wallet', CASE WHEN enable_wallet THEN 'true' ELSE 'false' END, created_on, modified_on FROM p_site_settings_legacy WHERE id = 1
	UNION ALL
	SELECT 'web_profile.allow_tweet_attachment', CASE WHEN allow_tweet_attachment THEN 'true' ELSE 'false' END, created_on, modified_on FROM p_site_settings_legacy WHERE id = 1
	UNION ALL
	SELECT 'web_profile.allow_tweet_attachment_price', CASE WHEN allow_tweet_attachment_price THEN 'true' ELSE 'false' END, created_on, modified_on FROM p_site_settings_legacy WHERE id = 1
	UNION ALL
	SELECT 'web_profile.allow_tweet_video', CASE WHEN allow_tweet_video THEN 'true' ELSE 'false' END, created_on, modified_on FROM p_site_settings_legacy WHERE id = 1
	UNION ALL
	SELECT 'web_profile.default_tweet_max_length', default_tweet_max_length::TEXT, created_on, modified_on FROM p_site_settings_legacy WHERE id = 1
	UNION ALL
	SELECT 'web_profile.tweet_web_ellipsis_size', tweet_web_ellipsis_size::TEXT, created_on, modified_on FROM p_site_settings_legacy WHERE id = 1
	UNION ALL
	SELECT 'web_profile.tweet_mobile_ellipsis_size', tweet_mobile_ellipsis_size::TEXT, created_on, modified_on FROM p_site_settings_legacy WHERE id = 1
	UNION ALL
	SELECT 'web_profile.default_tweet_visibility', default_tweet_visibility, created_on, modified_on FROM p_site_settings_legacy WHERE id = 1
	UNION ALL
	SELECT 'web_profile.default_msg_loop_interval', default_msg_loop_interval::TEXT, created_on, modified_on FROM p_site_settings_legacy WHERE id = 1
	UNION ALL
	SELECT 'web_profile.copyright_top', copyright_top, created_on, modified_on FROM p_site_settings_legacy WHERE id = 1
	UNION ALL
	SELECT 'web_profile.copyright_left', copyright_left, created_on, modified_on FROM p_site_settings_legacy WHERE id = 1
	UNION ALL
	SELECT 'web_profile.copyright_left_link', copyright_left_link, created_on, modified_on FROM p_site_settings_legacy WHERE id = 1
	UNION ALL
	SELECT 'web_profile.copyright_right', copyright_right, created_on, modified_on FROM p_site_settings_legacy WHERE id = 1
	UNION ALL
	SELECT 'web_profile.copyright_right_link', copyright_right_link, created_on, modified_on FROM p_site_settings_legacy WHERE id = 1
) legacy_rows;

DROP TABLE p_site_settings_legacy;
