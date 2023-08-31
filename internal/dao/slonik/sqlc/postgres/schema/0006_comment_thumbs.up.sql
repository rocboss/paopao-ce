ALTER TABLE p_comment ADD COLUMN thumbs_up_count INT NOT NULL DEFAULT 0;
ALTER TABLE p_comment ADD COLUMN thumbs_down_count INT NOT NULL DEFAULT 0;
ALTER TABLE p_comment_reply ADD COLUMN thumbs_up_count INT NOT NULL DEFAULT 0;
ALTER TABLE p_comment_reply ADD COLUMN thumbs_down_count INT NOT NULL DEFAULT 0;

CREATE TABLE p_tweet_comment_thumbs (
	ID BIGSERIAL PRIMARY KEY,
	user_id BIGINT NOT NULL,
	tweet_id BIGINT NOT NULL,
	comment_id BIGINT NOT NULL,
	reply_id BIGINT,
	comment_type SMALLINT NOT NULL DEFAULT 0,-- 评论类型 0为推文评论、1为评论回复
	is_thumbs_up SMALLINT NOT NULL DEFAULT 0,-- 是否点赞 0 为否 1为是
	is_thumbs_down SMALLINT NOT NULL DEFAULT 0,-- 是否点踩 0 为否 1为是
	created_on BIGINT NOT NULL DEFAULT 0,
	modified_on BIGINT NOT NULL DEFAULT 0,
	deleted_on BIGINT NOT NULL DEFAULT 0,
	is_del SMALLINT NOT NULL DEFAULT 0 -- 是否删除 0 为未删除、1 为已删除
);
CREATE INDEX idx_tweet_comment_thumbs_uid_tid ON p_tweet_comment_thumbs USING btree ( user_id, tweet_id );