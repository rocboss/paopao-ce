CREATE TABLE "p_post_metric" (
	"id" integer NOT NULL,
	"post_id" integer NOT NULL,
	"rank_score" integer NOT NULL,
	"incentive_score" integer NOT NULL,
	"decay_factor" integer NOT NULL,
	"motivation_factor" integer NOT NULL,
	"is_del" integer NOT NULL,
	"created_on" integer NOT NULL,
	"modified_on" integer NOT NULL,
	"deleted_on" integer NOT NULL,
	PRIMARY KEY ("id")
);

CREATE INDEX "idx_post_metric_post_id_rank_score"
ON "p_post_metric" (
  "post_id" ASC,
  "rank_score" ASC
);

INSERT INTO p_post_metric (post_id, rank_score, created_on) 
SELECT id AS post_id, 
	comment_count+upvote_count*2+collection_count*4 AS rank_score,
	created_on
FROM p_post
WHERE is_del=0;

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
