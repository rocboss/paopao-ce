ALTER TABLE "p_comment" ADD COLUMN "thumbs_up_count" integer NOT NULL DEFAULT 0;
ALTER TABLE "p_comment" ADD COLUMN "thumbs_down_count" integer NOT NULL DEFAULT 0;
ALTER TABLE "p_comment_reply" ADD COLUMN "thumbs_up_count" integer NOT NULL DEFAULT 0;
ALTER TABLE "p_comment_reply" ADD COLUMN "thumbs_down_count" integer NOT NULL DEFAULT 0;

CREATE TABLE "p_tweet_comment_thumbs" (
  "id"  integer PRIMARY KEY,
  "user_id" integer NOT NULL,
  "tweet_id" integer NOT NULL,
  "comment_id" integer NOT NULL,
  "reply_id" integer,
  "comment_type" integer NOT NULL DEFAULT 0, -- 评论类型 0为推文评论、1为评论回复
  "is_thumbs_up" integer NOT NULL DEFAULT 0, -- 是否点赞 0 为否 1为是
  "is_thumbs_down" integer NOT NULL DEFAULT 0, -- 是否点踩 0 为否 1为是
  "created_on" integer NOT NULL DEFAULT 0,
  "modified_on" integer NOT NULL DEFAULT 0,
  "deleted_on" integer NOT NULL DEFAULT 0,
  "is_del" integer NOT NULL DEFAULT 0 -- 是否删除 0 为未删除、1 为已删除
);
