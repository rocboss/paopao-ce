ALTER TABLE p_comment ADD COLUMN reply_count INT NOT NULL DEFAULT 0;

WITH comment_reply AS (
	SELECT comment_id, count(*) AS count 
	FROM p_comment_reply 
	WHERE is_del=0 
	GROUP By comment_id
)
UPDATE p_comment comment
SET reply_count = reply.count
FROM comment_reply reply
WHERE comment.id = reply.comment_id;

CREATE TABLE p_comment_metric (
	id BIGSERIAL PRIMARY KEY,
	comment_id BIGINT NOT NULL,
	rank_score BIGINT NOT NULL DEFAULT 0,
	incentive_score INT NOT NULL DEFAULT 0,
	decay_factor INT NOT NULL DEFAULT 0,
	motivation_factor INT NOT NULL DEFAULT 0,
	is_del SMALLINT NOT NULL DEFAULT 0,
	created_on BIGINT NOT NULL DEFAULT 0,
	modified_on BIGINT NOT NULL DEFAULT 0,
	deleted_on BIGINT NOT NULL DEFAULT 0
);
CREATE INDEX idx_comment_metric_comment_id_rank_score ON p_comment_metric USING btree (comment_id, rank_score);

INSERT INTO p_comment_metric (comment_id, rank_score, created_on) 
SELECT id AS comment_id, 
	reply_count*2 + thumbs_up_count*4 - thumbs_down_count AS rank_score,
	created_on
FROM p_comment
WHERE is_del=0;

CREATE TABLE p_user_metric (
	id BIGSERIAL PRIMARY KEY,
	user_id BIGINT NOT NULL,
	tweets_count INT NOT NULL DEFAULT 0,
	latest_trends_on BIGINT NOT NULL DEFAULT 0,
	is_del SMALLINT NOT NULL DEFAULT 0,
	created_on BIGINT NOT NULL DEFAULT 0,
	modified_on BIGINT NOT NULL DEFAULT 0,
	deleted_on BIGINT NOT NULL DEFAULT 0
);
CREATE INDEX idx_user_metric_user_id_tweets_count_trends ON p_user_metric USING btree (user_id, tweets_count, latest_trends_on);

INSERT INTO p_user_metric (user_id, tweets_count) 
SELECT user_id, count(*) AS tweets_count
FROM p_post
WHERE is_del=0
GROUP BY user_id;
