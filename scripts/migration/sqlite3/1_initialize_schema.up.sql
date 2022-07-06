PRAGMA foreign_keys = false;

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

CREATE TABLE "p_comment" (
  "id" integer NOT NULL,
  "post_id" integer NOT NULL,
  "user_id" integer NOT NULL,
  "ip" text(64) NOT NULL,
  "ip_loc" text(64) NOT NULL,
  "created_on" integer NOT NULL,
  "modified_on" integer NOT NULL,
  "deleted_on" integer NOT NULL,
  "is_del" integer NOT NULL,
  PRIMARY KEY ("id")
);

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

CREATE TABLE "p_comment_reply" (
  "id" integer NOT NULL,
  "comment_id" integer NOT NULL,
  "user_id" integer NOT NULL,
  "at_user_id" integer NOT NULL,
  "content" text(255) NOT NULL,
  "ip" text(64) NOT NULL,
  "ip_loc" text(64) NOT NULL,
  "created_on" integer NOT NULL,
  "modified_on" integer NOT NULL,
  "deleted_on" integer NOT NULL,
  "is_del" integer NOT NULL,
  PRIMARY KEY ("id")
);

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

CREATE TABLE "p_post" (
  "id" integer NOT NULL,
  "user_id" integer NOT NULL,
  "comment_count" integer NOT NULL,
  "collection_count" integer NOT NULL,
  "upvote_count" integer NOT NULL,
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
  PRIMARY KEY ("id")
);

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

CREATE INDEX "main"."idx_user"
ON "p_attachment" (
  "user_id" ASC
);

CREATE INDEX "main"."idx_expired_on"
ON "p_captcha" (
  "expired_on" ASC
);

CREATE INDEX "main"."idx_phone"
ON "p_captcha" (
  "phone" ASC
);

CREATE INDEX "main"."idx_use_times"
ON "p_captcha" (
  "use_times" ASC
);

CREATE INDEX "main"."idx_post"
ON "p_comment" (
  "post_id" ASC
);

PRAGMA foreign_keys = true;
