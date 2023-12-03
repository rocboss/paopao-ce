CREATE TABLE p_post_metric (
	ID BIGSERIAL PRIMARY KEY,
	post_id BIGINT NOT NULL,
	rank_score BIGINT NOT NULL DEFAULT 0,
	incentive_score INT NOT NULL DEFAULT 0,
	decay_factor INT NOT NULL DEFAULT 0,
	motivation_factor INT NOT NULL DEFAULT 0,
	is_del SMALLINT NOT NULL DEFAULT 0,
	created_on BIGINT NOT NULL DEFAULT 0,
	modified_on BIGINT NOT NULL DEFAULT 0,
	deleted_on BIGINT NOT NULL DEFAULT 0 
);
CREATE INDEX idx_post_metric_post_id_rank_score ON p_post_metric USING btree ( post_id, rank_score );

INSERT INTO p_post_metric ( post_id, rank_score, created_on ) SELECT ID AS
post_id,
comment_count + upvote_count * 2 + collection_count * 4 AS rank_score,
created_on 
FROM
	p_post 
WHERE
	is_del = 0;

-- 原来的可见性: 0公开 1私密 2好友可见 3关注可见
-- 现在的可见性: 0私密 10充电可见 20订阅可见 30保留 40保留 50好友可见 60关注可见 70保留 80保留 90公开
UPDATE p_post a
SET visibility = (
	SELECT
	CASE visibility 
		WHEN 0 THEN 90 
		WHEN 1 THEN 0 
		WHEN 2 THEN 50 
		WHEN 3 THEN 60 
		ELSE 0
	END 
	FROM
		p_post b 
	WHERE
		a.ID = b.ID 
);
