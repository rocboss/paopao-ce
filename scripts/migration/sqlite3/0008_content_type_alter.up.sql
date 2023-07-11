PRAGMA foreign_keys=off;

BEGIN TRANSACTION;

ALTER TABLE p_post_content RENAME TO _p_post_content;

ALTER TABLE p_comment_content RENAME TO _p_comment_content;

ALTER TABLE p_comment_reply RENAME TO _p_comment_reply;

CREATE TABLE p_post_content
( "id" integer NOT NULL,
  "post_id" integer NOT NULL,
  "user_id" integer NOT NULL,
  "content" text NOT NULL,
  "type" integer NOT NULL,
  "sort" integer NOT NULL,
  "created_on" integer NOT NULL,
  "modified_on" integer NOT NULL,
  "deleted_on" integer NOT NULL,
  "is_del" integer NOT NULL,
  PRIMARY KEY ("id")
);

CREATE TABLE p_comment_content
(  "id" integer NOT NULL,
  "comment_id" integer NOT NULL,
  "user_id" integer NOT NULL,
  "content" text NOT NULL,
  "type" integer NOT NULL,
  "sort" integer NOT NULL,
  "created_on" integer NOT NULL,
  "modified_on" integer NOT NULL,
  "deleted_on" integer NOT NULL,
  "is_del" integer NOT NULL,
  PRIMARY KEY ("id")
);

CREATE TABLE p_comment_reply
( "id" integer NOT NULL,
  "comment_id" integer NOT NULL,
  "user_id" integer NOT NULL,
  "at_user_id" integer NOT NULL,
  "content" text NOT NULL,
  "ip" text(64) NOT NULL,
  "ip_loc" text(64) NOT NULL,
  "thumbs_up_count" integer NOT NULL DEFAULT 0, -- 点赞数
  "thumbs_down_count" integer NOT NULL DEFAULT 0, -- 点踩数
  "created_on" integer NOT NULL,
  "modified_on" integer NOT NULL,
  "deleted_on" integer NOT NULL,
  "is_del" integer NOT NULL,
  PRIMARY KEY ("id")
);

-- ----------------------------
-- Indexes structure for table p_post_content
-- ----------------------------
CREATE INDEX "idx_post_content_post_id"
ON "p_post_content" (
  "post_id" ASC
);
CREATE INDEX "idx_post_content_user_id"
ON "p_post_content" (
  "user_id" ASC
);

-- ----------------------------
-- Indexes structure for table p_comment_reply
-- ----------------------------
CREATE INDEX "idx_comment_reply_comment_id"
ON "p_comment_reply" (
  "comment_id" ASC
);

-- ----------------------------
-- Indexes structure for table p_comment_content
-- ----------------------------
CREATE INDEX "idx_comment_content_comment_id"
ON "p_comment_content" (
  "comment_id" ASC
);
CREATE INDEX "idx_comment_content_sort"
ON "p_comment_content" (
  "sort" ASC
);
CREATE INDEX "idx_comment_content_type"
ON "p_comment_content" (
  "type" ASC
);
CREATE INDEX "idx_comment_content_user_id"
ON "p_comment_content" (
  "user_id" ASC
);

INSERT INTO p_post_content (id, post_id, user_id, content, `type`, sort, created_on, modified_on, deleted_on, is_del)
  SELECT id, post_id, user_id, content, `type`, sort, created_on, modified_on, deleted_on, is_del
  FROM _p_post_content;

INSERT INTO p_comment_content (id, comment_id, user_id, content, `type`, sort, created_on, modified_on, deleted_on, is_del)
  SELECT id, comment_id, user_id, content, `type`, sort, created_on, modified_on, deleted_on, is_del
  FROM _p_comment_content;

INSERT INTO p_comment_reply (id, comment_id, user_id, at_user_id, content, ip, ip_loc, thumbs_up_count, thumbs_down_count, created_on, modified_on, deleted_on, is_del)
  SELECT id, comment_id, user_id, at_user_id, content, ip, ip_loc, thumbs_up_count, thumbs_down_count, created_on, modified_on, deleted_on, is_del
  FROM _p_comment_reply;

DROP TABLE _p_post_content;
DROP TABLE _p_comment_content;
DROP TABLE _p_comment_reply;

-- ----------------------------
-- Indexes structure for table p_post_content
-- ----------------------------
CREATE INDEX "idx_post_content_post_id"
ON "p_post_content" (
  "post_id" ASC
);
CREATE INDEX "idx_post_content_user_id"
ON "p_post_content" (
  "user_id" ASC
);

-- ----------------------------
-- Indexes structure for table p_comment_reply
-- ----------------------------
CREATE INDEX "idx_comment_reply_comment_id"
ON "p_comment_reply" (
  "comment_id" ASC
);

-- ----------------------------
-- Indexes structure for table p_comment_content
-- ----------------------------
CREATE INDEX "idx_comment_content_comment_id"
ON "p_comment_content" (
  "comment_id" ASC
);
CREATE INDEX "idx_comment_content_sort"
ON "p_comment_content" (
  "sort" ASC
);
CREATE INDEX "idx_comment_content_type"
ON "p_comment_content" (
  "type" ASC
);
CREATE INDEX "idx_comment_content_user_id"
ON "p_comment_content" (
  "user_id" ASC
);

COMMIT;

PRAGMA foreign_keys=on;
