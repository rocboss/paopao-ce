PRAGMA foreign_keys = false;

-- ----------------------------
-- Table structure for p_attachment
-- ----------------------------
DROP TABLE IF EXISTS "p_attachment";
CREATE TABLE "p_attachment" (
  "id" integer NOT NULL,
  "user_id" integer NOT NULL,
  "file_size" integer NOT NULL,
  "img_width" integer NOT NULL,
  "img_height" integer NOT NULL,
  "type" integer NOT NULL,
  "content" text(255) NOT NULL,
  "created_on" integer NOT NULL,
  "modified_on" integer NOT NULL,
  "deleted_on" integer NOT NULL,
  "is_del" integer NOT NULL,
  PRIMARY KEY ("id")
);

-- ----------------------------
-- Table structure for p_captcha
-- ----------------------------
DROP TABLE IF EXISTS "p_captcha";
CREATE TABLE "p_captcha" (
  "id" integer NOT NULL,
  "phone" text(16) NOT NULL,
  "captcha" text(16) NOT NULL,
  "use_times" integer NOT NULL,
  "expired_on" integer NOT NULL,
  "created_on" integer NOT NULL,
  "modified_on" integer NOT NULL,
  "deleted_on" integer NOT NULL,
  "is_del" integer NOT NULL,
  PRIMARY KEY ("id")
);

-- ----------------------------
-- Table structure for p_comment
-- ----------------------------
DROP TABLE IF EXISTS "p_comment";
CREATE TABLE "p_comment" (
  "id" integer NOT NULL,
  "post_id" integer NOT NULL,
  "user_id" integer NOT NULL,
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
-- Table structure for p_comment_content
-- ----------------------------
DROP TABLE IF EXISTS "p_comment_content";
CREATE TABLE "p_comment_content" (
  "id" integer NOT NULL,
  "comment_id" integer NOT NULL,
  "user_id" integer NOT NULL,
  "content" text(255) NOT NULL,
  "type" integer NOT NULL,
  "sort" integer NOT NULL,
  "created_on" integer NOT NULL,
  "modified_on" integer NOT NULL,
  "deleted_on" integer NOT NULL,
  "is_del" integer NOT NULL,
  PRIMARY KEY ("id")
);

-- ----------------------------
-- Table structure for p_comment_reply
-- ----------------------------
DROP TABLE IF EXISTS "p_comment_reply";
CREATE TABLE "p_comment_reply" (
  "id" integer NOT NULL,
  "comment_id" integer NOT NULL,
  "user_id" integer NOT NULL,
  "at_user_id" integer NOT NULL,
  "content" text(255) NOT NULL,
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
-- Table structure for p_tweet_comment_thumbs
-- ----------------------------
DROP TABLE IF EXISTS p_tweet_comment_thumbs;
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

-- ----------------------------
-- Table structure for p_contact
-- ----------------------------
DROP TABLE IF EXISTS "p_contact";
CREATE TABLE "p_contact" (
  "id" integer NOT NULL,
  "user_id" integer NOT NULL,
  "friend_id" integer NOT NULL,
  "group_id" integer NOT NULL,
  "remark" text(32) NOT NULL,
  "status" integer NOT NULL,
  "notice_enable" integer NOT NULL,
  "is_top" integer NOT NULL,
  "is_black" integer NOT NULL,
  "is_del" integer NOT NULL,
  "created_on" integer NOT NULL,
  "modified_on" integer NOT NULL,
  "deleted_on" integer NOT NULL,
  PRIMARY KEY ("id")
);

-- ----------------------------
-- Table structure for p_contact_group
-- ----------------------------
DROP TABLE IF EXISTS "p_contact_group";
CREATE TABLE "p_contact_group" (
  "id" integer NOT NULL,
  "user_id" integer NOT NULL,
  "name" text(32) NOT NULL,
  "is_del" integer NOT NULL,
  "created_on" integer NOT NULL,
  "modified_on" integer NOT NULL,
  "deleted_on" integer NOT NULL,
  PRIMARY KEY ("id")
);

-- ----------------------------
-- Table structure for p_message
-- ----------------------------
DROP TABLE IF EXISTS "p_message";
CREATE TABLE "p_message" (
  "id" integer NOT NULL,
  "sender_user_id" integer NOT NULL,
  "receiver_user_id" integer NOT NULL,
  "type" integer NOT NULL,
  "brief" text(255) NOT NULL,
  "content" text(255) NOT NULL,
  "post_id" integer NOT NULL,
  "comment_id" integer NOT NULL,
  "reply_id" integer NOT NULL,
  "is_read" integer NOT NULL,
  "created_on" integer NOT NULL,
  "modified_on" integer NOT NULL,
  "deleted_on" integer NOT NULL,
  "is_del" integer NOT NULL,
  PRIMARY KEY ("id")
);

-- ----------------------------
-- Table structure for p_post
-- ----------------------------
DROP TABLE IF EXISTS "p_post";
CREATE TABLE "p_post" (
  "id" integer NOT NULL,
  "user_id" integer NOT NULL,
  "comment_count" integer NOT NULL,
  "collection_count" integer NOT NULL,
  "upvote_count" integer NOT NULL,
  "share_count" integer NOT NULL,
  "is_top" integer NOT NULL,
  "is_essence" integer NOT NULL,
  "is_lock" integer NOT NULL,
  "latest_replied_on" integer NOT NULL,
  "tags" text(255) NOT NULL,
  "attachment_price" integer NOT NULL,
  "ip" text(64) NOT NULL,
  "ip_loc" text(64) NOT NULL,
  "created_on" integer NOT NULL,
  "modified_on" integer NOT NULL,
  "deleted_on" integer NOT NULL,
  "is_del" integer NOT NULL,
  "visibility" integer NOT NULL,
  PRIMARY KEY ("id")
);

-- ----------------------------
-- Table structure for p_post_attachment_bill
-- ----------------------------
DROP TABLE IF EXISTS "p_post_attachment_bill";
CREATE TABLE "p_post_attachment_bill" (
  "id" integer NOT NULL,
  "post_id" integer NOT NULL,
  "user_id" integer NOT NULL,
  "paid_amount" integer NOT NULL,
  "created_on" integer NOT NULL,
  "modified_on" integer NOT NULL,
  "deleted_on" integer NOT NULL,
  "is_del" integer NOT NULL,
  PRIMARY KEY ("id")
);

-- ----------------------------
-- Table structure for p_post_collection
-- ----------------------------
DROP TABLE IF EXISTS "p_post_collection";
CREATE TABLE "p_post_collection" (
  "id" integer NOT NULL,
  "post_id" integer NOT NULL,
  "user_id" integer NOT NULL,
  "created_on" integer NOT NULL,
  "modified_on" integer NOT NULL,
  "deleted_on" integer NOT NULL,
  "is_del" integer NOT NULL,
  PRIMARY KEY ("id")
);

-- ----------------------------
-- Table structure for p_post_content
-- ----------------------------
DROP TABLE IF EXISTS "p_post_content";
CREATE TABLE "p_post_content" (
  "id" integer NOT NULL,
  "post_id" integer NOT NULL,
  "user_id" integer NOT NULL,
  "content" text(2000) NOT NULL,
  "type" integer NOT NULL,
  "sort" integer NOT NULL,
  "created_on" integer NOT NULL,
  "modified_on" integer NOT NULL,
  "deleted_on" integer NOT NULL,
  "is_del" integer NOT NULL,
  PRIMARY KEY ("id")
);

-- ----------------------------
-- Table structure for p_post_star
-- ----------------------------
DROP TABLE IF EXISTS "p_post_star";
CREATE TABLE "p_post_star" (
  "id" integer NOT NULL,
  "post_id" integer NOT NULL,
  "user_id" integer NOT NULL,
  "created_on" integer NOT NULL,
  "modified_on" integer NOT NULL,
  "deleted_on" integer NOT NULL,
  "is_del" integer NOT NULL,
  PRIMARY KEY ("id")
);

-- ----------------------------
-- Table structure for p_tag
-- ----------------------------
DROP TABLE IF EXISTS "p_tag";
CREATE TABLE "p_tag" (
  "id" integer NOT NULL,
  "user_id" integer NOT NULL,
  "tag" text(255) NOT NULL,
  "quote_num" integer NOT NULL,
  "created_on" integer NOT NULL,
  "modified_on" integer NOT NULL,
  "deleted_on" integer NOT NULL,
  "is_del" integer NOT NULL,
  PRIMARY KEY ("id")
);

-- ----------------------------
-- Table structure for p_topic_user
-- ----------------------------
DROP TABLE IF EXISTS "p_topic_user";
CREATE TABLE "p_topic_user" (
	"id" integer,
	"topic_id" integer NOT NULL,-- 标签ID
	"user_id" integer NOT NULL,-- 创建者ID
	"alias_name" text ( 255 ),-- 别名
	"remark" text ( 512 ),-- 备注
	"quote_num" integer,-- 引用数
	"is_top" integer NOT NULL DEFAULT 0,-- 是否置顶 0 为未置顶、1 为已置顶
	"created_on" integer NOT NULL DEFAULT 0,-- 创建时间
	"modified_on" integer NOT NULL DEFAULT 0,-- 修改时间
	"deleted_on" integer NOT NULL DEFAULT 0,-- 删除时间
	"is_del" integer NOT NULL DEFAULT 0,-- 是否删除 0 为未删除、1 为已删除
	"reserve_a" text,-- 保留字段a
	"reserve_b" text,-- 保留字段b
	PRIMARY KEY ( "id" ) 
);

-- ----------------------------
-- Table structure for p_user
-- ----------------------------
DROP TABLE IF EXISTS "p_user";
CREATE TABLE "p_user" (
  "id" integer NOT NULL,
  "nickname" text(32) NOT NULL,
  "username" text(32) NOT NULL,
  "phone" text(16) NOT NULL,
  "password" text(32) NOT NULL,
  "salt" text(16) NOT NULL,
  "status" integer NOT NULL,
  "avatar" text(255) NOT NULL,
  "balance" integer NOT NULL,
  "is_admin" integer NOT NULL,
  "created_on" integer NOT NULL,
  "modified_on" integer NOT NULL,
  "deleted_on" integer NOT NULL,
  "is_del" integer NOT NULL,
  PRIMARY KEY ("id")
);

-- ----------------------------
-- Table structure for p_wallet_recharge
-- ----------------------------
DROP TABLE IF EXISTS "p_wallet_recharge";
CREATE TABLE "p_wallet_recharge" (
  "id" integer NOT NULL,
  "user_id" integer NOT NULL,
  "amount" integer NOT NULL,
  "trade_no" text(64) NOT NULL,
  "trade_status" text(32) NOT NULL,
  "created_on" integer NOT NULL,
  "modified_on" integer NOT NULL,
  "deleted_on" integer NOT NULL,
  "is_del" integer NOT NULL,
  PRIMARY KEY ("id")
);

-- ----------------------------
-- Table structure for p_wallet_statement
-- ----------------------------
DROP TABLE IF EXISTS "p_wallet_statement";
CREATE TABLE "p_wallet_statement" (
  "id" integer NOT NULL,
  "user_id" integer NOT NULL,
  "change_amount" integer NOT NULL,
  "balance_snapshot" integer NOT NULL,
  "reason" text(255) NOT NULL,
  "post_id" integer NOT NULL,
  "created_on" integer NOT NULL,
  "modified_on" integer NOT NULL,
  "deleted_on" integer NOT NULL,
  "is_del" integer NOT NULL,
  PRIMARY KEY ("id")
);

-- ----------------------------
-- Indexes structure for table p_attachment
-- ----------------------------
CREATE INDEX "idx_attachment_user_id"
ON "p_attachment" (
  "user_id" ASC
);

-- ----------------------------
-- Indexes structure for table p_captcha
-- ----------------------------
CREATE INDEX "idx_captcha_expired_on"
ON "p_captcha" (
  "expired_on" ASC
);
CREATE INDEX "idx_captcha_phone"
ON "p_captcha" (
  "phone" ASC
);
CREATE INDEX "idx_captcha_use_times"
ON "p_captcha" (
  "use_times" ASC
);

-- ----------------------------
-- Indexes structure for table p_comment
-- ----------------------------
CREATE INDEX "idx_comment_post_id"
ON "p_comment" (
  "post_id" ASC
);
CREATE INDEX "idx_comment_user_id"
ON "p_comment" (
  "user_id" ASC
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

-- ----------------------------
-- Indexes structure for table p_comment_reply
-- ----------------------------
CREATE INDEX "idx_comment_reply_comment_id"
ON "p_comment_reply" (
  "comment_id" ASC
);

-- ----------------------------
-- Indexes structure for table idx_tweet_comment_thumbs_uid_tid
-- ----------------------------
CREATE INDEX "idx_tweet_comment_thumbs_uid_tid"
ON "p_tweet_comment_thumbs"(
  "user_id" ASC,
  "tweet_id" ASC
);

-- ----------------------------
-- Indexes structure for table p_contact
-- ----------------------------
CREATE UNIQUE INDEX "idx_contact_user_friend"
ON "p_contact" (
  "user_id" ASC,
  "friend_id" ASC
);
CREATE INDEX "idx_contact_user_friend_status"
ON "p_contact" (
  "user_id" ASC,
  "friend_id" ASC,
  "status" ASC
);

-- ----------------------------
-- Indexes structure for table p_message
-- ----------------------------
CREATE INDEX "idx_message_is_read"
ON "p_message" (
  "is_read" ASC
);
CREATE INDEX "idx_message_receiver_user_id"
ON "p_message" (
  "receiver_user_id" ASC
);
CREATE INDEX "idx_message_type"
ON "p_message" (
  "type" ASC
);

-- ----------------------------
-- Indexes structure for table p_post
-- ----------------------------
CREATE INDEX "idx_post_user_id"
ON "p_post" (
  "user_id" ASC
);
CREATE INDEX "idx_post_visibility"
ON "p_post" (
  "visibility" ASC
);

-- ----------------------------
-- Indexes structure for table p_post_attachment_bill
-- ----------------------------
CREATE INDEX "idx_post_attachment_bill_post_id"
ON "p_post_attachment_bill" (
  "post_id" ASC
);
CREATE INDEX "idx_post_attachment_bill_user_id"
ON "p_post_attachment_bill" (
  "user_id" ASC
);

-- ----------------------------
-- Indexes structure for table p_post_collection
-- ----------------------------
CREATE INDEX "idx_post_collection_post_id"
ON "p_post_collection" (
  "post_id" ASC
);
CREATE INDEX "idx_post_collection_user_id"
ON "p_post_collection" (
  "user_id" ASC
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
-- Indexes structure for table p_post_star
-- ----------------------------
CREATE INDEX "idx_post_star_post_id"
ON "p_post_star" (
  "post_id" ASC
);
CREATE INDEX "idx_post_star_user_id"
ON "p_post_star" (
  "user_id" ASC
);

-- ----------------------------
-- Indexes structure for table p_tag
-- ----------------------------
CREATE UNIQUE INDEX "idx_tag"
ON "p_tag" (
  "tag" ASC
);
CREATE INDEX "idx_tag_quote_num"
ON "p_tag" (
  "quote_num" ASC
);
CREATE INDEX "idx_tag_user_id"
ON "p_tag" (
  "user_id" ASC
);

-- ----------------------------
-- Indexes structure for table p_topic_user
-- ----------------------------
CREATE UNIQUE INDEX "idx_topic_user_uid_tid" 
ON "p_topic_user" (
  "topic_id",
  "user_id"
);

-- ----------------------------
-- Indexes structure for table p_user
-- ----------------------------
CREATE INDEX "idx_user_phone"
ON "p_user" (
  "phone" ASC
);
CREATE UNIQUE INDEX "idx_user_username"
ON "p_user" (
  "username" ASC
);

-- ----------------------------
-- Indexes structure for table p_wallet_recharge
-- ----------------------------
CREATE INDEX "idx_wallet_recharge_trade_no"
ON "p_wallet_recharge" (
  "trade_no" ASC
);
CREATE INDEX "idx_wallet_recharge_trade_status"
ON "p_wallet_recharge" (
  "trade_status" ASC
);
CREATE INDEX "idx_wallet_recharge_user_id"
ON "p_wallet_recharge" (
  "user_id" ASC
);

-- ----------------------------
-- Indexes structure for table p_wallet_statement
-- ----------------------------
CREATE INDEX "idx_wallet_statement_user_id"
ON "p_wallet_statement" (
  "user_id" ASC
);

PRAGMA foreign_keys = true;
