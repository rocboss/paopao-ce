ALTER TABLE "p_site_settings" RENAME TO "p_site_settings_legacy";

CREATE TABLE "p_site_settings" (
	"key" text NOT NULL,
	"value" text NOT NULL DEFAULT '',
	"is_encrypted" integer NOT NULL DEFAULT 0,
	"created_on" integer NOT NULL DEFAULT 0,
	"modified_on" integer NOT NULL DEFAULT 0,
	"deleted_on" integer NOT NULL DEFAULT 0,
	"is_del" integer NOT NULL DEFAULT 0,
	PRIMARY KEY ("key")
);

INSERT INTO "p_site_settings" ("key", "value", "is_encrypted", "created_on", "modified_on", "deleted_on", "is_del")
SELECT "key", "value", 0, created_on, modified_on, 0, 0 FROM (
	SELECT 'web_profile.use_friendship' AS "key", CASE WHEN use_friendship <> 0 THEN 'true' ELSE 'false' END AS "value", created_on, modified_on FROM "p_site_settings_legacy" WHERE id = 1
	UNION ALL
	SELECT 'web_profile.enable_trends_bar', CASE WHEN enable_trends_bar <> 0 THEN 'true' ELSE 'false' END, created_on, modified_on FROM "p_site_settings_legacy" WHERE id = 1
	UNION ALL
	SELECT 'web_profile.enable_wallet', CASE WHEN enable_wallet <> 0 THEN 'true' ELSE 'false' END, created_on, modified_on FROM "p_site_settings_legacy" WHERE id = 1
	UNION ALL
	SELECT 'web_profile.allow_tweet_attachment', CASE WHEN allow_tweet_attachment <> 0 THEN 'true' ELSE 'false' END, created_on, modified_on FROM "p_site_settings_legacy" WHERE id = 1
	UNION ALL
	SELECT 'web_profile.allow_tweet_attachment_price', CASE WHEN allow_tweet_attachment_price <> 0 THEN 'true' ELSE 'false' END, created_on, modified_on FROM "p_site_settings_legacy" WHERE id = 1
	UNION ALL
	SELECT 'web_profile.allow_tweet_video', CASE WHEN allow_tweet_video <> 0 THEN 'true' ELSE 'false' END, created_on, modified_on FROM "p_site_settings_legacy" WHERE id = 1
	UNION ALL
	SELECT 'web_profile.default_tweet_max_length', CAST(default_tweet_max_length AS TEXT), created_on, modified_on FROM "p_site_settings_legacy" WHERE id = 1
	UNION ALL
	SELECT 'web_profile.tweet_web_ellipsis_size', CAST(tweet_web_ellipsis_size AS TEXT), created_on, modified_on FROM "p_site_settings_legacy" WHERE id = 1
	UNION ALL
	SELECT 'web_profile.tweet_mobile_ellipsis_size', CAST(tweet_mobile_ellipsis_size AS TEXT), created_on, modified_on FROM "p_site_settings_legacy" WHERE id = 1
	UNION ALL
	SELECT 'web_profile.default_tweet_visibility', default_tweet_visibility, created_on, modified_on FROM "p_site_settings_legacy" WHERE id = 1
	UNION ALL
	SELECT 'web_profile.default_msg_loop_interval', CAST(default_msg_loop_interval AS TEXT), created_on, modified_on FROM "p_site_settings_legacy" WHERE id = 1
	UNION ALL
	SELECT 'web_profile.copyright_top', copyright_top, created_on, modified_on FROM "p_site_settings_legacy" WHERE id = 1
	UNION ALL
	SELECT 'web_profile.copyright_left', copyright_left, created_on, modified_on FROM "p_site_settings_legacy" WHERE id = 1
	UNION ALL
	SELECT 'web_profile.copyright_left_link', copyright_left_link, created_on, modified_on FROM "p_site_settings_legacy" WHERE id = 1
	UNION ALL
	SELECT 'web_profile.copyright_right', copyright_right, created_on, modified_on FROM "p_site_settings_legacy" WHERE id = 1
	UNION ALL
	SELECT 'web_profile.copyright_right_link', copyright_right_link, created_on, modified_on FROM "p_site_settings_legacy" WHERE id = 1
);

DROP TABLE "p_site_settings_legacy";
