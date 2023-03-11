PRAGMA foreign_keys = false;

DROP INDEX IF EXISTS "main"."idx_user";
DROP INDEX IF EXISTS "main"."idx_expired_on";
DROP INDEX IF EXISTS "main"."idx_phone";
DROP INDEX IF EXISTS "main"."idx_use_times";
DROP INDEX IF EXISTS "main"."idx_post";
DROP INDEX IF EXISTS "main"."idx_user_friend_id";
DROP INDEX IF EXISTS "main"."idx_user_friend_status";
DROP INDEX IF EXISTS "main"."idx_visibility";

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
