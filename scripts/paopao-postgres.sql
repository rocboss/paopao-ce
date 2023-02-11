/*
 Navicat Premium Data Transfer
*/

-- ----------------------------
-- Table structure for p_anouncement
-- ----------------------------
DROP TABLE IF EXISTS "public"."p_anouncement";
CREATE TABLE "public"."p_anouncement" (
  "id" numeric(20,0) NOT NULL,
  "user_id" numeric(20,0) NOT NULL,
  "comment_count" numeric(20,0) NOT NULL,
  "collection_count" numeric(20,0) NOT NULL,
  "upvote_count" numeric(20,0) NOT NULL,
  "is_top" int2 NOT NULL,
  "is_essence" int2 NOT NULL,
  "is_lock" int2 NOT NULL,
  "latest_replied_on" numeric(20,0) NOT NULL,
  "tags" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "attachment_price" numeric(20,0) NOT NULL,
  "ip" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "ip_loc" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "created_on" numeric(20,0) NOT NULL,
  "modified_on" numeric(20,0) NOT NULL,
  "deleted_on" numeric(20,0) NOT NULL,
  "is_del" int2 NOT NULL
);
COMMENT ON COLUMN "public"."p_anouncement"."user_id" IS '用户ID';
COMMENT ON COLUMN "public"."p_anouncement"."comment_count" IS '评论数';
COMMENT ON COLUMN "public"."p_anouncement"."collection_count" IS '收藏数';
COMMENT ON COLUMN "public"."p_anouncement"."upvote_count" IS '点赞数';
COMMENT ON COLUMN "public"."p_anouncement"."is_top" IS '是否置顶';
COMMENT ON COLUMN "public"."p_anouncement"."is_essence" IS '是否精华';
COMMENT ON COLUMN "public"."p_anouncement"."is_lock" IS '是否锁定';
COMMENT ON COLUMN "public"."p_anouncement"."latest_replied_on" IS '最新回复时间';
COMMENT ON COLUMN "public"."p_anouncement"."tags" IS '标签';
COMMENT ON COLUMN "public"."p_anouncement"."attachment_price" IS '附件价格(分)';
COMMENT ON COLUMN "public"."p_anouncement"."ip" IS 'IP地址';
COMMENT ON COLUMN "public"."p_anouncement"."ip_loc" IS 'IP城市地址';
COMMENT ON COLUMN "public"."p_anouncement"."created_on" IS '创建时间';
COMMENT ON COLUMN "public"."p_anouncement"."modified_on" IS '修改时间';
COMMENT ON COLUMN "public"."p_anouncement"."deleted_on" IS '删除时间';
COMMENT ON COLUMN "public"."p_anouncement"."is_del" IS '是否删除 0 为未删除、1 为已删除';
COMMENT ON TABLE "public"."p_anouncement" IS '公告';

-- ----------------------------
-- Table structure for p_anouncement_content
-- ----------------------------
DROP TABLE IF EXISTS "public"."p_anouncement_content";
CREATE TABLE "public"."p_anouncement_content" (
  "id" numeric(20,0) NOT NULL,
  "anouncement_id" numeric(20,0) NOT NULL,
  "user_id" numeric(20,0) NOT NULL,
  "content" varchar(2000) COLLATE "pg_catalog"."default" NOT NULL,
  "type" int2 NOT NULL,
  "sort" int8 NOT NULL,
  "created_on" numeric(20,0) NOT NULL,
  "modified_on" numeric(20,0) NOT NULL,
  "deleted_on" numeric(20,0) NOT NULL,
  "is_del" int2 NOT NULL
);
COMMENT ON COLUMN "public"."p_anouncement_content"."id" IS '内容ID';
COMMENT ON COLUMN "public"."p_anouncement_content"."anouncement_id" IS '公告ID';
COMMENT ON COLUMN "public"."p_anouncement_content"."user_id" IS '用户ID';
COMMENT ON COLUMN "public"."p_anouncement_content"."content" IS '内容';
COMMENT ON COLUMN "public"."p_anouncement_content"."type" IS '类型，1标题，2文字段落，3图片地址，4视频地址，5语音地址，6链接地址，7附件资源，8收费资源';
COMMENT ON COLUMN "public"."p_anouncement_content"."sort" IS '排序，越小越靠前';
COMMENT ON COLUMN "public"."p_anouncement_content"."created_on" IS '创建时间';
COMMENT ON COLUMN "public"."p_anouncement_content"."modified_on" IS '修改时间';
COMMENT ON COLUMN "public"."p_anouncement_content"."deleted_on" IS '删除时间';
COMMENT ON COLUMN "public"."p_anouncement_content"."is_del" IS '是否删除 0 为未删除、1 为已删除';
COMMENT ON TABLE "public"."p_anouncement_content" IS '公告内容';

-- ----------------------------
-- Table structure for p_attachment
-- ----------------------------
DROP TABLE IF EXISTS "public"."p_attachment";
CREATE TABLE "public"."p_attachment" (
  "id" numeric(20,0) NOT NULL,
  "user_id" numeric(20,0) NOT NULL,
  "file_size" numeric(20,0) NOT NULL,
  "img_width" numeric(20,0) NOT NULL,
  "img_height" numeric(20,0) NOT NULL,
  "type" int2 NOT NULL,
  "content" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "created_on" numeric(20,0) NOT NULL,
  "modified_on" numeric(20,0) NOT NULL,
  "deleted_on" numeric(20,0) NOT NULL,
  "is_del" int2 NOT NULL,
  "scope" int2 NOT NULL
);
COMMENT ON COLUMN "public"."p_attachment"."type" IS '1图片，2视频，3其他附件';
COMMENT ON COLUMN "public"."p_attachment"."created_on" IS '创建时间';
COMMENT ON COLUMN "public"."p_attachment"."modified_on" IS '修改时间';
COMMENT ON COLUMN "public"."p_attachment"."deleted_on" IS '删除时间';
COMMENT ON COLUMN "public"."p_attachment"."is_del" IS '是否删除 0 为未删除、1 为已删除';
COMMENT ON COLUMN "public"."p_attachment"."scope" IS '领域 0未知 1泡泡 2公告';
COMMENT ON TABLE "public"."p_attachment" IS '附件';

-- ----------------------------
-- Table structure for p_captcha
-- ----------------------------
DROP TABLE IF EXISTS "public"."p_captcha";
CREATE TABLE "public"."p_captcha" (
  "id" numeric(20,0) NOT NULL,
  "phone" varchar(16) COLLATE "pg_catalog"."default" NOT NULL,
  "captcha" varchar(16) COLLATE "pg_catalog"."default" NOT NULL,
  "use_times" int8 NOT NULL,
  "expired_on" numeric(20,0) NOT NULL,
  "created_on" numeric(20,0) NOT NULL,
  "modified_on" numeric(20,0) NOT NULL,
  "deleted_on" numeric(20,0) NOT NULL,
  "is_del" int2 NOT NULL
);
COMMENT ON COLUMN "public"."p_captcha"."id" IS '验证码ID';
COMMENT ON COLUMN "public"."p_captcha"."phone" IS '手机号';
COMMENT ON COLUMN "public"."p_captcha"."captcha" IS '验证码';
COMMENT ON COLUMN "public"."p_captcha"."use_times" IS '使用次数';
COMMENT ON COLUMN "public"."p_captcha"."expired_on" IS '过期时间';
COMMENT ON COLUMN "public"."p_captcha"."created_on" IS '创建时间';
COMMENT ON COLUMN "public"."p_captcha"."modified_on" IS '修改时间';
COMMENT ON COLUMN "public"."p_captcha"."deleted_on" IS '删除时间';
COMMENT ON COLUMN "public"."p_captcha"."is_del" IS '是否删除 0 为未删除、1 为已删除';
COMMENT ON TABLE "public"."p_captcha" IS '手机验证码';

-- ----------------------------
-- Table structure for p_comment
-- ----------------------------
DROP TABLE IF EXISTS "public"."p_comment";
CREATE TABLE "public"."p_comment" (
  "id" numeric(20,0) NOT NULL,
  "post_id" numeric(20,0) NOT NULL,
  "user_id" numeric(20,0) NOT NULL,
  "ip" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "ip_loc" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "created_on" numeric(20,0) NOT NULL,
  "modified_on" numeric(20,0) NOT NULL,
  "deleted_on" numeric(20,0) NOT NULL,
  "is_del" int2 NOT NULL,
  "scope" int2 NOT NULL
);
COMMENT ON COLUMN "public"."p_comment"."id" IS '评论ID';
COMMENT ON COLUMN "public"."p_comment"."post_id" IS 'POST ID';
COMMENT ON COLUMN "public"."p_comment"."user_id" IS '用户ID';
COMMENT ON COLUMN "public"."p_comment"."ip" IS 'IP地址';
COMMENT ON COLUMN "public"."p_comment"."ip_loc" IS 'IP城市地址';
COMMENT ON COLUMN "public"."p_comment"."created_on" IS '创建时间';
COMMENT ON COLUMN "public"."p_comment"."modified_on" IS '修改时间';
COMMENT ON COLUMN "public"."p_comment"."deleted_on" IS '删除时间';
COMMENT ON COLUMN "public"."p_comment"."is_del" IS '是否删除 0 为未删除、1 为已删除';
COMMENT ON COLUMN "public"."p_comment"."scope" IS '领域 0未知 1泡泡 2公告';
COMMENT ON TABLE "public"."p_comment" IS '评论';

-- ----------------------------
-- Table structure for p_comment_content
-- ----------------------------
DROP TABLE IF EXISTS "public"."p_comment_content";
CREATE TABLE "public"."p_comment_content" (
  "id" numeric(20,0) NOT NULL,
  "comment_id" numeric(20,0) NOT NULL,
  "user_id" numeric(20,0) NOT NULL,
  "content" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "type" int2 NOT NULL,
  "sort" numeric(20,0) NOT NULL,
  "created_on" numeric(20,0) NOT NULL,
  "modified_on" numeric(20,0) NOT NULL,
  "deleted_on" numeric(20,0) NOT NULL,
  "is_del" int2 NOT NULL
);
COMMENT ON COLUMN "public"."p_comment_content"."id" IS '内容ID';
COMMENT ON COLUMN "public"."p_comment_content"."comment_id" IS '评论ID';
COMMENT ON COLUMN "public"."p_comment_content"."user_id" IS '用户ID';
COMMENT ON COLUMN "public"."p_comment_content"."content" IS '内容';
COMMENT ON COLUMN "public"."p_comment_content"."type" IS '类型，1标题，2文字段落，3图片地址，4视频地址，5语音地址，6链接地址';
COMMENT ON COLUMN "public"."p_comment_content"."sort" IS '排序，越小越靠前';
COMMENT ON COLUMN "public"."p_comment_content"."created_on" IS '创建时间';
COMMENT ON COLUMN "public"."p_comment_content"."modified_on" IS '修改时间';
COMMENT ON COLUMN "public"."p_comment_content"."deleted_on" IS '删除时间';
COMMENT ON COLUMN "public"."p_comment_content"."is_del" IS '是否删除 0 为未删除、1 为已删除';
COMMENT ON TABLE "public"."p_comment_content" IS '评论内容';

-- ----------------------------
-- Table structure for p_comment_reply
-- ----------------------------
DROP TABLE IF EXISTS "public"."p_comment_reply";
CREATE TABLE "public"."p_comment_reply" (
  "id" numeric(20,0) NOT NULL,
  "comment_id" numeric(20,0) NOT NULL,
  "user_id" numeric(20,0) NOT NULL,
  "at_user_id" numeric(20,0) NOT NULL,
  "content" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "ip" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "ip_loc" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "created_on" numeric(20,0) NOT NULL,
  "modified_on" numeric(20,0) NOT NULL,
  "deleted_on" numeric(20,0) NOT NULL,
  "is_del" int2 NOT NULL
);
COMMENT ON COLUMN "public"."p_comment_reply"."id" IS '回复ID';
COMMENT ON COLUMN "public"."p_comment_reply"."comment_id" IS '评论ID';
COMMENT ON COLUMN "public"."p_comment_reply"."user_id" IS '用户ID';
COMMENT ON COLUMN "public"."p_comment_reply"."at_user_id" IS '@用户ID';
COMMENT ON COLUMN "public"."p_comment_reply"."content" IS '内容';
COMMENT ON COLUMN "public"."p_comment_reply"."ip" IS 'IP地址';
COMMENT ON COLUMN "public"."p_comment_reply"."ip_loc" IS 'IP城市地址';
COMMENT ON COLUMN "public"."p_comment_reply"."created_on" IS '创建时间';
COMMENT ON COLUMN "public"."p_comment_reply"."modified_on" IS '修改时间';
COMMENT ON COLUMN "public"."p_comment_reply"."deleted_on" IS '删除时间';
COMMENT ON COLUMN "public"."p_comment_reply"."is_del" IS '是否删除 0 为未删除、1 为已删除';
COMMENT ON TABLE "public"."p_comment_reply" IS '评论回复';

-- ----------------------------
-- Table structure for p_contact
-- ----------------------------
DROP TABLE IF EXISTS "public"."p_contact";
CREATE TABLE "public"."p_contact" (
  "id" numeric(20,0) NOT NULL,
  "user_id" numeric(20,0) NOT NULL,
  "friend_id" numeric(20,0) NOT NULL,
  "group_id" numeric(20,0) NOT NULL,
  "remark" varchar(32) COLLATE "pg_catalog"."default" NOT NULL,
  "status" int2 NOT NULL,
  "notice_enable" int2 NOT NULL,
  "is_top" int2 NOT NULL,
  "is_black" int2 NOT NULL,
  "is_del" int2 NOT NULL,
  "created_on" numeric(20,0) NOT NULL,
  "modified_on" numeric(20,0) NOT NULL,
  "deleted_on" numeric(20,0) NOT NULL
);
COMMENT ON COLUMN "public"."p_contact"."id" IS '联系人ID';
COMMENT ON COLUMN "public"."p_contact"."user_id" IS '用户ID';
COMMENT ON COLUMN "public"."p_contact"."friend_id" IS '好友ID';
COMMENT ON COLUMN "public"."p_contact"."group_id" IS '好友分组ID:默认为0无分组';
COMMENT ON COLUMN "public"."p_contact"."remark" IS '好友备注';
COMMENT ON COLUMN "public"."p_contact"."status" IS '好友状态: 1请求好友, 2已好友, 3拒绝好友, 4已删好友';
COMMENT ON COLUMN "public"."p_contact"."notice_enable" IS '是否有消息提醒, 0否, 1是';
COMMENT ON COLUMN "public"."p_contact"."is_top" IS '是否置顶, 0否, 1是';
COMMENT ON COLUMN "public"."p_contact"."is_black" IS '是否为黑名单, 0否, 1是';
COMMENT ON COLUMN "public"."p_contact"."is_del" IS '是否删除好友, 0否, 1是';
COMMENT ON COLUMN "public"."p_contact"."created_on" IS '创建时间';
COMMENT ON COLUMN "public"."p_contact"."modified_on" IS '修改时间';
COMMENT ON COLUMN "public"."p_contact"."deleted_on" IS '删除时间';
COMMENT ON TABLE "public"."p_contact" IS '联系人';

-- ----------------------------
-- Table structure for p_contact_group
-- ----------------------------
DROP TABLE IF EXISTS "public"."p_contact_group";
CREATE TABLE "public"."p_contact_group" (
  "id" numeric(20,0) NOT NULL,
  "user_id" int4 NOT NULL,
  "name" varchar(32) COLLATE "pg_catalog"."default" NOT NULL,
  "is_del" int2 NOT NULL,
  "created_on" numeric(20,0) NOT NULL,
  "modified_on" numeric(20,0) NOT NULL,
  "deleted_on" numeric(20,0) NOT NULL
);
COMMENT ON COLUMN "public"."p_contact_group"."id" IS '联系人ID';
COMMENT ON COLUMN "public"."p_contact_group"."user_id" IS '用户id';
COMMENT ON COLUMN "public"."p_contact_group"."name" IS '分组名称';
COMMENT ON COLUMN "public"."p_contact_group"."is_del" IS '是否删除, 0否, 1是';
COMMENT ON COLUMN "public"."p_contact_group"."created_on" IS '创建时间';
COMMENT ON COLUMN "public"."p_contact_group"."modified_on" IS '修改时间';
COMMENT ON COLUMN "public"."p_contact_group"."deleted_on" IS '删除时间';
COMMENT ON TABLE "public"."p_contact_group" IS '联系人分组';

-- ----------------------------
-- Table structure for p_message
-- ----------------------------
DROP TABLE IF EXISTS "public"."p_message";
CREATE TABLE "public"."p_message" (
  "id" numeric(20,0) NOT NULL,
  "sender_user_id" numeric(20,0) NOT NULL,
  "receiver_user_id" numeric(20,0) NOT NULL,
  "type" int2 NOT NULL,
  "brief" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "content" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "post_id" numeric(20,0) NOT NULL,
  "comment_id" numeric(20,0) NOT NULL,
  "reply_id" numeric(20,0) NOT NULL,
  "is_read" int2 NOT NULL,
  "created_on" numeric(20,0) NOT NULL,
  "modified_on" numeric(20,0) NOT NULL,
  "deleted_on" numeric(20,0) NOT NULL,
  "is_del" int2 NOT NULL
);
COMMENT ON COLUMN "public"."p_message"."id" IS '消息通知ID';
COMMENT ON COLUMN "public"."p_message"."sender_user_id" IS '发送方用户ID';
COMMENT ON COLUMN "public"."p_message"."receiver_user_id" IS '接收方用户ID';
COMMENT ON COLUMN "public"."p_message"."type" IS '通知类型，1动态，2评论，3回复，4私信，99系统通知';
COMMENT ON COLUMN "public"."p_message"."brief" IS '摘要说明';
COMMENT ON COLUMN "public"."p_message"."content" IS '详细内容';
COMMENT ON COLUMN "public"."p_message"."post_id" IS '动态ID';
COMMENT ON COLUMN "public"."p_message"."comment_id" IS '评论ID';
COMMENT ON COLUMN "public"."p_message"."reply_id" IS '回复ID';
COMMENT ON COLUMN "public"."p_message"."is_read" IS '是否已读';
COMMENT ON COLUMN "public"."p_message"."created_on" IS '创建时间';
COMMENT ON COLUMN "public"."p_message"."modified_on" IS '修改时间';
COMMENT ON COLUMN "public"."p_message"."deleted_on" IS '删除时间';
COMMENT ON COLUMN "public"."p_message"."is_del" IS '是否删除 0 为未删除、1 为已删除';
COMMENT ON TABLE "public"."p_message" IS '消息通知';

-- ----------------------------
-- Table structure for p_post
-- ----------------------------
DROP TABLE IF EXISTS "public"."p_post";
CREATE TABLE "public"."p_post" (
  "id" numeric(20,0) NOT NULL,
  "user_id" numeric(20,0) NOT NULL,
  "comment_count" numeric(20,0) NOT NULL,
  "collection_count" numeric(20,0) NOT NULL,
  "upvote_count" numeric(20,0) NOT NULL,
  "is_top" int2 NOT NULL,
  "is_essence" int2 NOT NULL,
  "is_lock" int2 NOT NULL,
  "latest_replied_on" numeric(20,0) NOT NULL,
  "tags" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "attachment_price" numeric(20,0) NOT NULL,
  "ip" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "ip_loc" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "created_on" numeric(20,0) NOT NULL,
  "modified_on" numeric(20,0) NOT NULL,
  "deleted_on" numeric(20,0) NOT NULL,
  "is_del" int2 NOT NULL,
  "visibility" int2 NOT NULL
);
COMMENT ON COLUMN "public"."p_post"."id" IS '主题ID';
COMMENT ON COLUMN "public"."p_post"."user_id" IS '用户ID';
COMMENT ON COLUMN "public"."p_post"."comment_count" IS '评论数';
COMMENT ON COLUMN "public"."p_post"."collection_count" IS '收藏数';
COMMENT ON COLUMN "public"."p_post"."upvote_count" IS '点赞数';
COMMENT ON COLUMN "public"."p_post"."is_top" IS '是否置顶';
COMMENT ON COLUMN "public"."p_post"."is_essence" IS '是否精华';
COMMENT ON COLUMN "public"."p_post"."is_lock" IS '是否锁定';
COMMENT ON COLUMN "public"."p_post"."latest_replied_on" IS '最新回复时间';
COMMENT ON COLUMN "public"."p_post"."tags" IS '标签';
COMMENT ON COLUMN "public"."p_post"."attachment_price" IS '附件价格(分)';
COMMENT ON COLUMN "public"."p_post"."ip" IS 'IP地址';
COMMENT ON COLUMN "public"."p_post"."ip_loc" IS 'IP城市地址';
COMMENT ON COLUMN "public"."p_post"."created_on" IS '创建时间';
COMMENT ON COLUMN "public"."p_post"."modified_on" IS '修改时间';
COMMENT ON COLUMN "public"."p_post"."deleted_on" IS '删除时间';
COMMENT ON COLUMN "public"."p_post"."is_del" IS '是否删除 0 为未删除、1 为已删除';
COMMENT ON COLUMN "public"."p_post"."visibility" IS '可见性 0公开 1私密 2好友可见';
COMMENT ON TABLE "public"."p_post" IS '冒泡/文章';

-- ----------------------------
-- Table structure for p_post_attachment_bill
-- ----------------------------
DROP TABLE IF EXISTS "public"."p_post_attachment_bill";
CREATE TABLE "public"."p_post_attachment_bill" (
  "id" numeric(20,0) NOT NULL,
  "post_id" numeric(20,0) NOT NULL,
  "user_id" numeric(20,0) NOT NULL,
  "paid_amount" numeric(20,0) NOT NULL,
  "created_on" numeric(20,0) NOT NULL,
  "modified_on" numeric(20,0) NOT NULL,
  "deleted_on" numeric(20,0) NOT NULL,
  "is_del" int2 NOT NULL,
  "scope" int2 NOT NULL
);
COMMENT ON COLUMN "public"."p_post_attachment_bill"."id" IS '购买记录ID';
COMMENT ON COLUMN "public"."p_post_attachment_bill"."post_id" IS 'POST ID';
COMMENT ON COLUMN "public"."p_post_attachment_bill"."user_id" IS '用户ID';
COMMENT ON COLUMN "public"."p_post_attachment_bill"."paid_amount" IS '支付金额';
COMMENT ON COLUMN "public"."p_post_attachment_bill"."created_on" IS '创建时间';
COMMENT ON COLUMN "public"."p_post_attachment_bill"."modified_on" IS '修改时间';
COMMENT ON COLUMN "public"."p_post_attachment_bill"."deleted_on" IS '删除时间';
COMMENT ON COLUMN "public"."p_post_attachment_bill"."is_del" IS '是否删除 0 为未删除、1 为已删除';
COMMENT ON COLUMN "public"."p_post_attachment_bill"."scope" IS '领域 0未知 1泡泡 2公告';
COMMENT ON TABLE "public"."p_post_attachment_bill" IS '冒泡/文章附件账单';

-- ----------------------------
-- Table structure for p_post_collection
-- ----------------------------
DROP TABLE IF EXISTS "public"."p_post_collection";
CREATE TABLE "public"."p_post_collection" (
  "id" numeric(20,0) NOT NULL,
  "post_id" numeric(20,0) NOT NULL,
  "user_id" numeric(20,0) NOT NULL,
  "created_on" numeric(20,0) NOT NULL,
  "modified_on" numeric(20,0) NOT NULL,
  "deleted_on" numeric(20,0) NOT NULL,
  "is_del" int2 NOT NULL,
  "scope" int2 NOT NULL
);
COMMENT ON COLUMN "public"."p_post_collection"."id" IS '收藏ID';
COMMENT ON COLUMN "public"."p_post_collection"."post_id" IS 'POST ID';
COMMENT ON COLUMN "public"."p_post_collection"."user_id" IS '用户ID';
COMMENT ON COLUMN "public"."p_post_collection"."created_on" IS '创建时间';
COMMENT ON COLUMN "public"."p_post_collection"."modified_on" IS '修改时间';
COMMENT ON COLUMN "public"."p_post_collection"."deleted_on" IS '删除时间';
COMMENT ON COLUMN "public"."p_post_collection"."is_del" IS '是否删除 0 为未删除、1 为已删除';
COMMENT ON COLUMN "public"."p_post_collection"."scope" IS '领域 0未知 1泡泡 2公告';
COMMENT ON TABLE "public"."p_post_collection" IS '冒泡/文章收藏';

-- ----------------------------
-- Table structure for p_post_content
-- ----------------------------
DROP TABLE IF EXISTS "public"."p_post_content";
CREATE TABLE "public"."p_post_content" (
  "id" numeric(20,0) NOT NULL,
  "post_id" numeric(20,0) NOT NULL,
  "user_id" numeric(20,0) NOT NULL,
  "content" varchar(2000) COLLATE "pg_catalog"."default" NOT NULL,
  "type" int2 NOT NULL,
  "sort" int8 NOT NULL,
  "created_on" numeric(20,0) NOT NULL,
  "modified_on" numeric(20,0) NOT NULL,
  "deleted_on" numeric(20,0) NOT NULL,
  "is_del" int2 NOT NULL
);
COMMENT ON COLUMN "public"."p_post_content"."id" IS '内容ID';
COMMENT ON COLUMN "public"."p_post_content"."post_id" IS 'POST ID';
COMMENT ON COLUMN "public"."p_post_content"."user_id" IS '用户ID';
COMMENT ON COLUMN "public"."p_post_content"."content" IS '内容';
COMMENT ON COLUMN "public"."p_post_content"."type" IS '类型，1标题，2文字段落，3图片地址，4视频地址，5语音地址，6链接地址，7附件资源，8收费资源';
COMMENT ON COLUMN "public"."p_post_content"."sort" IS '排序，越小越靠前';
COMMENT ON COLUMN "public"."p_post_content"."created_on" IS '创建时间';
COMMENT ON COLUMN "public"."p_post_content"."modified_on" IS '修改时间';
COMMENT ON COLUMN "public"."p_post_content"."deleted_on" IS '删除时间';
COMMENT ON COLUMN "public"."p_post_content"."is_del" IS '是否删除 0 为未删除、1 为已删除';
COMMENT ON TABLE "public"."p_post_content" IS '冒泡/文章内容';

-- ----------------------------
-- Table structure for p_post_star
-- ----------------------------
DROP TABLE IF EXISTS "public"."p_post_star";
CREATE TABLE "public"."p_post_star" (
  "id" numeric(20,0) NOT NULL,
  "post_id" numeric(20,0) NOT NULL,
  "user_id" numeric(20,0) NOT NULL,
  "created_on" numeric(20,0) NOT NULL,
  "modified_on" numeric(20,0) NOT NULL,
  "deleted_on" numeric(20,0) NOT NULL,
  "is_del" int2 NOT NULL,
  "scope" int2 NOT NULL
);
COMMENT ON COLUMN "public"."p_post_star"."id" IS '收藏ID';
COMMENT ON COLUMN "public"."p_post_star"."post_id" IS 'POST ID';
COMMENT ON COLUMN "public"."p_post_star"."user_id" IS '用户ID';
COMMENT ON COLUMN "public"."p_post_star"."created_on" IS '创建时间';
COMMENT ON COLUMN "public"."p_post_star"."modified_on" IS '修改时间';
COMMENT ON COLUMN "public"."p_post_star"."deleted_on" IS '删除时间';
COMMENT ON COLUMN "public"."p_post_star"."is_del" IS '是否删除 0 为未删除、1 为已删除';
COMMENT ON COLUMN "public"."p_post_star"."scope" IS '领域 0未知 1泡泡 2公告';
COMMENT ON TABLE "public"."p_post_star" IS '冒泡/文章点赞';

-- ----------------------------
-- Table structure for p_tag
-- ----------------------------
DROP TABLE IF EXISTS "public"."p_tag";
CREATE TABLE "public"."p_tag" (
  "id" numeric(20,0) NOT NULL,
  "user_id" numeric(20,0) NOT NULL,
  "tag" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "quote_num" numeric(20,0) NOT NULL,
  "created_on" numeric(20,0) NOT NULL,
  "modified_on" numeric(20,0) NOT NULL,
  "deleted_on" numeric(20,0) NOT NULL,
  "is_del" int2 NOT NULL
);
COMMENT ON COLUMN "public"."p_tag"."id" IS '标签ID';
COMMENT ON COLUMN "public"."p_tag"."user_id" IS '创建者ID';
COMMENT ON COLUMN "public"."p_tag"."tag" IS '标签名';
COMMENT ON COLUMN "public"."p_tag"."quote_num" IS '引用数';
COMMENT ON COLUMN "public"."p_tag"."created_on" IS '创建时间';
COMMENT ON COLUMN "public"."p_tag"."modified_on" IS '修改时间';
COMMENT ON COLUMN "public"."p_tag"."deleted_on" IS '删除时间';
COMMENT ON COLUMN "public"."p_tag"."is_del" IS '是否删除 0 为未删除、1 为已删除';
COMMENT ON TABLE "public"."p_tag" IS '标签';

-- ----------------------------
-- Table structure for p_user
-- ----------------------------
DROP TABLE IF EXISTS "public"."p_user";
CREATE TABLE "public"."p_user" (
  "id" numeric(20,0) NOT NULL,
  "nickname" varchar(32) COLLATE "pg_catalog"."default" NOT NULL,
  "username" varchar(32) COLLATE "pg_catalog"."default" NOT NULL,
  "phone" varchar(16) COLLATE "pg_catalog"."default" NOT NULL,
  "password" varchar(32) COLLATE "pg_catalog"."default" NOT NULL,
  "salt" varchar(16) COLLATE "pg_catalog"."default" NOT NULL,
  "status" int2 NOT NULL,
  "avatar" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "balance" numeric(20,0) NOT NULL,
  "is_admin" int2 NOT NULL,
  "created_on" numeric(20,0) NOT NULL,
  "modified_on" numeric(20,0) NOT NULL,
  "deleted_on" numeric(20,0) NOT NULL,
  "is_del" int2 NOT NULL
);
COMMENT ON COLUMN "public"."p_user"."id" IS '用户ID';
COMMENT ON COLUMN "public"."p_user"."nickname" IS '昵称';
COMMENT ON COLUMN "public"."p_user"."username" IS '用户名';
COMMENT ON COLUMN "public"."p_user"."phone" IS '手机号';
COMMENT ON COLUMN "public"."p_user"."password" IS 'MD5密码';
COMMENT ON COLUMN "public"."p_user"."salt" IS '盐值';
COMMENT ON COLUMN "public"."p_user"."status" IS '状态，1正常，2停用';
COMMENT ON COLUMN "public"."p_user"."avatar" IS '用户头像';
COMMENT ON COLUMN "public"."p_user"."balance" IS '用户余额（分）';
COMMENT ON COLUMN "public"."p_user"."is_admin" IS '是否管理员';
COMMENT ON COLUMN "public"."p_user"."created_on" IS '创建时间';
COMMENT ON COLUMN "public"."p_user"."modified_on" IS '修改时间';
COMMENT ON COLUMN "public"."p_user"."deleted_on" IS '删除时间';
COMMENT ON COLUMN "public"."p_user"."is_del" IS '是否删除 0 为未删除、1 为已删除';
COMMENT ON TABLE "public"."p_user" IS '用户';

-- ----------------------------
-- Table structure for p_wallet_recharge
-- ----------------------------
DROP TABLE IF EXISTS "public"."p_wallet_recharge";
CREATE TABLE "public"."p_wallet_recharge" (
  "id" numeric(20,0) NOT NULL,
  "user_id" numeric(20,0) NOT NULL,
  "amount" int8 NOT NULL,
  "trade_no" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
  "trade_status" varchar(32) COLLATE "pg_catalog"."default" NOT NULL,
  "created_on" numeric(20,0) NOT NULL,
  "modified_on" numeric(20,0) NOT NULL,
  "deleted_on" numeric(20,0) NOT NULL,
  "is_del" int2 NOT NULL
);
COMMENT ON COLUMN "public"."p_wallet_recharge"."id" IS '充值ID';
COMMENT ON COLUMN "public"."p_wallet_recharge"."user_id" IS '用户ID';
COMMENT ON COLUMN "public"."p_wallet_recharge"."amount" IS '充值金额';
COMMENT ON COLUMN "public"."p_wallet_recharge"."trade_no" IS '支付宝订单号';
COMMENT ON COLUMN "public"."p_wallet_recharge"."trade_status" IS '交易状态';
COMMENT ON COLUMN "public"."p_wallet_recharge"."created_on" IS '创建时间';
COMMENT ON COLUMN "public"."p_wallet_recharge"."modified_on" IS '修改时间';
COMMENT ON COLUMN "public"."p_wallet_recharge"."deleted_on" IS '删除时间';
COMMENT ON COLUMN "public"."p_wallet_recharge"."is_del" IS '是否删除 0 为未删除、1 为已删除';
COMMENT ON TABLE "public"."p_wallet_recharge" IS '钱包流水';

-- ----------------------------
-- Table structure for p_wallet_statement
-- ----------------------------
DROP TABLE IF EXISTS "public"."p_wallet_statement";
CREATE TABLE "public"."p_wallet_statement" (
  "id" numeric(20,0) NOT NULL,
  "user_id" numeric(20,0) NOT NULL,
  "change_amount" int8 NOT NULL,
  "balance_snapshot" int8 NOT NULL,
  "reason" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "post_id" numeric(20,0) NOT NULL,
  "created_on" numeric(20,0) NOT NULL,
  "modified_on" numeric(20,0) NOT NULL,
  "deleted_on" numeric(20,0) NOT NULL,
  "is_del" int2 NOT NULL
);
COMMENT ON COLUMN "public"."p_wallet_statement"."id" IS '账单ID';
COMMENT ON COLUMN "public"."p_wallet_statement"."user_id" IS '用户ID';
COMMENT ON COLUMN "public"."p_wallet_statement"."change_amount" IS '变动金额';
COMMENT ON COLUMN "public"."p_wallet_statement"."balance_snapshot" IS '资金快照';
COMMENT ON COLUMN "public"."p_wallet_statement"."reason" IS '变动原因';
COMMENT ON COLUMN "public"."p_wallet_statement"."post_id" IS '关联动态';
COMMENT ON COLUMN "public"."p_wallet_statement"."created_on" IS '创建时间';
COMMENT ON COLUMN "public"."p_wallet_statement"."modified_on" IS '修改时间';
COMMENT ON COLUMN "public"."p_wallet_statement"."deleted_on" IS '删除时间';
COMMENT ON COLUMN "public"."p_wallet_statement"."is_del" IS '是否删除 0 为未删除、1 为已删除';
COMMENT ON TABLE "public"."p_wallet_statement" IS '钱包流水';

-- ----------------------------
-- Primary Key structure for table p_anouncement
-- ----------------------------
ALTER TABLE "public"."p_anouncement" ADD CONSTRAINT "p_anouncement_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Indexes structure for table p_anouncement_content
-- ----------------------------
CREATE INDEX "idx_b_anouncement" ON "public"."p_anouncement_content" USING btree (
  "anouncement_id" "pg_catalog"."numeric_ops" ASC NULLS LAST
);
CREATE INDEX "idx_b_user" ON "public"."p_anouncement_content" USING btree (
  "user_id" "pg_catalog"."numeric_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table p_anouncement_content
-- ----------------------------
ALTER TABLE "public"."p_anouncement_content" ADD CONSTRAINT "p_anouncement_content_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Indexes structure for table p_attachment
-- ----------------------------
CREATE INDEX "idx_user" ON "public"."p_attachment" USING btree (
  "user_id" "pg_catalog"."numeric_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table p_attachment
-- ----------------------------
ALTER TABLE "public"."p_attachment" ADD CONSTRAINT "p_attachment_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Indexes structure for table p_captcha
-- ----------------------------
CREATE INDEX "idx_expired_on" ON "public"."p_captcha" USING btree (
  "expired_on" "pg_catalog"."numeric_ops" ASC NULLS LAST
);
CREATE INDEX "idx_phone" ON "public"."p_captcha" USING btree (
  "phone" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);
CREATE INDEX "idx_use_times" ON "public"."p_captcha" USING btree (
  "use_times" "pg_catalog"."int8_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table p_captcha
-- ----------------------------
ALTER TABLE "public"."p_captcha" ADD CONSTRAINT "p_captcha_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Indexes structure for table p_comment
-- ----------------------------
CREATE INDEX "idx_post" ON "public"."p_comment" USING btree (
  "post_id" "pg_catalog"."numeric_ops" ASC NULLS LAST
);
