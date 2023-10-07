ALTER TABLE "p_comment" ADD COLUMN "reply_count" integer NOT NULL DEFAULT 0;

UPDATE p_comment AS comment
SET reply_count = (
	SELECT count(*)
	FROM
		p_comment_reply AS reply
	WHERE
		comment.id=reply.comment_id AND comment.is_del=0 AND reply.is_del=0
);

CREATE TABLE p_comment_metric (
	"id" integer,
	"comment_id" integer NOT NULL,
	"rank_score" integer NOT NULL DEFAULT 0,
	"incentive_score" integer NOT NULL DEFAULT 0,
	"decay_factor" integer NOT NULL DEFAULT 0,
	"motivation_factor" integer NOT NULL DEFAULT 0,
	"is_del" integer NOT NULL DEFAULT 0,
	"created_on" integer NOT NULL DEFAULT 0,
	"modified_on" integer NOT NULL DEFAULT 0,
	"deleted_on" integer NOT NULL DEFAULT 0,
	PRIMARY KEY ("id")
);

CREATE INDEX "idx_comment_metric_comment_id_rank_score"
ON "p_comment_metric" (
	"comment_id" ASC,
	"rank_score" ASC
);

INSERT INTO p_comment_metric (comment_id, rank_score, created_on) 
SELECT id AS comment_id, 
	reply_count*2 + thumbs_up_count*4 - thumbs_down_count AS rank_score,
	created_on
FROM p_comment
WHERE is_del=0;

CREATE TABLE "p_user_metric" (
	"id" integer,
	"user_id" integer NOT NULL,
	"tweets_count" integer NOT NULL DEFAULT 0,
	"latest_trends_on" integer NOT NULL DEFAULT 0,
	"is_del" integer NOT NULL DEFAULT 0,
	"created_on" integer NOT NULL DEFAULT 0,
	"modified_on" integer NOT NULL DEFAULT 0,
	"deleted_on" integer NOT NULL DEFAULT 0,
	PRIMARY KEY ("id")
);

CREATE INDEX "idx_user_metric_user_id_tweets_count_trends"
ON "p_user_metric" (
  	"user_id" ASC,
  	"tweets_count" ASC,
	"latest_trends_on" ASC
);

INSERT INTO p_user_metric (user_id, tweets_count) 
SELECT user_id, count(*) AS tweets_count
FROM p_post
WHERE is_del=0
GROUP BY user_id;
