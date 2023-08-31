CREATE TABLE p_attachment (
	id BIGSERIAL PRIMARY KEY,
	user_id BIGINT NOT NULL DEFAULT 0,
	file_size BIGINT NOT NULL,
	img_width BIGINT NOT NULL DEFAULT 0,
	img_height BIGINT NOT NULL DEFAULT 0,
	"type" SMALLINT  NOT NULL DEFAULT 1, -- 1图片、2视频、3其他附件
	content VARCHAR(255) NOT NULL DEFAULT '',
	created_on BIGINT NOT NULL DEFAULT 0,
	modified_on BIGINT NOT NULL DEFAULT 0,
	deleted_on BIGINT NOT NULL DEFAULT 0,
	is_del SMALLINT NOT NULL DEFAULT 0  -- 是否删除 0为未删除、1为已删除
);
CREATE INDEX idx_attachment_user_id ON p_attachment USING btree (id);

CREATE TABLE p_captcha (
	id BIGSERIAL PRIMARY KEY,
	phone VARCHAR(16),
	captcha VARCHAR(16),
	use_times INTEGER NOT NULL DEFAULT 0,
	expired_on BIGINT NOT NULL DEFAULT 0,
	created_on BIGINT NOT NULL DEFAULT 0,
	modified_on BIGINT NOT NULL DEFAULT 0,
	deleted_on BIGINT NOT NULL DEFAULT 0,
	is_del SMALLINT NOT NULL DEFAULT 0 -- 是否删除 0为未删除、1为已删除'
);
CREATE INDEX idx_captcha_phone ON p_captcha USING btree (phone);
CREATE INDEX idx_captcha_expired_on ON p_captcha USING btree (expired_on);
CREATE INDEX idx_captcha_use_times ON p_captcha USING btree (use_times);

CREATE TABLE p_comment (
	id BIGSERIAL PRIMARY KEY,
	post_id BIGINT NOT NULL DEFAULT 0,
	user_id BIGINT NOT NULL DEFAULT 0,
	ip VARCHAR(64) NOT NULL DEFAULT '',
	ip_loc VARCHAR(64) NOT NULL DEFAULT '',
	created_on BIGINT NOT NULL DEFAULT 0,
	modified_on BIGINT NOT NULL DEFAULT 0,
	deleted_on BIGINT NOT NULL DEFAULT 0,
	is_del SMALLINT NOT NULL DEFAULT 0
);
CREATE INDEX idx_comment_post_id ON p_comment USING btree (post_id);
CREATE INDEX idx_comment_user_id ON p_comment USING btree (user_id);

CREATE TABLE p_comment_content (
	id BIGSERIAL PRIMARY KEY,
	comment_id BIGINT NOT NULL DEFAULT 0,
	user_id BIGINT NOT NULL DEFAULT 0,
	content VARCHAR(255) NOT NULL DEFAULT '',
	"type" SMALLINT NOT NULL DEFAULT 2, -- 类型，1标题，2文字段落，3图片地址，4视频地址，5语音地址，6链接地址,
	sort BIGINT NOT NULL DEFAULT 100,
	created_on BIGINT NOT NULL DEFAULT 0,
	modified_on BIGINT NOT NULL DEFAULT 0,
	deleted_on BIGINT NOT NULL DEFAULT 0,
	is_del SMALLINT NOT NULL DEFAULT 0
);
CREATE INDEX idx_comment_content_comment_id ON p_comment_content USING btree (comment_id);
CREATE INDEX idx_comment_content_user_id ON p_comment_content USING btree (user_id);
CREATE INDEX idx_comment_content_type ON p_comment_content USING btree ("type");
CREATE INDEX idx_comment_content_sort ON p_comment_content USING btree (sort);

CREATE TABLE p_comment_reply (
	id BIGSERIAL PRIMARY KEY,
	comment_id BIGINT NOT NULL DEFAULT 0,
	user_id BIGINT NOT NULL DEFAULT 0,
	at_user_id BIGINT NOT NULL DEFAULT 0,
	content VARCHAR(255) NOT NULL DEFAULT '',
	ip VARCHAR(64) NOT NULL DEFAULT '',
	ip_loc VARCHAR(64) NOT NULL DEFAULT '',
	created_on BIGINT NOT NULL DEFAULT 0,
	modified_on BIGINT NOT NULL DEFAULT 0,
	deleted_on BIGINT NOT NULL DEFAULT 0,
	is_del SMALLINT NOT NULL DEFAULT 0
);
CREATE INDEX idx_comment_reply_comment_id ON p_comment_reply USING btree (comment_id);

CREATE TABLE p_message (
	id BIGSERIAL PRIMARY KEY,
	sender_user_id BIGINT NOT NULL DEFAULT 0,
	receiver_user_id BIGINT NOT NULL DEFAULT 0,
	type SMALLINT NOT NULL DEFAULT 1,
	brief VARCHAR(255) NOT NULL DEFAULT '',
	content VARCHAR(255) NOT NULL DEFAULT '',
	post_id BIGINT NOT NULL DEFAULT 0,
	comment_id BIGINT NOT NULL DEFAULT 0,
	reply_id BIGINT NOT NULL DEFAULT 0,
	is_read SMALLINT NOT NULL DEFAULT 0,
	created_on BIGINT NOT NULL DEFAULT 0,
	modified_on BIGINT NOT NULL DEFAULT 0,
	deleted_on BIGINT NOT NULL DEFAULT 0,
	is_del SMALLINT NOT NULL DEFAULT 0
);
CREATE INDEX idx_message_receiver_user_id ON p_message USING btree (receiver_user_id);
CREATE INDEX idx_message_is_read ON p_message USING btree (is_read);
CREATE INDEX idx_message_type ON p_message USING btree ("type");

CREATE SEQUENCE IF NOT EXISTS post_id_seq AS BIGINT MINVALUE 1080017989 NO MAXVALUE;
CREATE TABLE p_post (
	id BIGINT NOT NULL DEFAULT nextval('post_id_seq'::regclass),
	user_id BIGINT NOT NULL DEFAULT 0,
	comment_count BIGINT NOT NULL DEFAULT 0,
	collection_count BIGINT NOT NULL DEFAULT 0,
	upvote_count BIGINT NOT NULL DEFAULT 0,
	is_top SMALLINT NOT NULL DEFAULT 0, -- 是否置顶
	is_essence SMALLINT NOT NULL DEFAULT 0, -- 是否精华
	is_lock SMALLINT NOT NULL DEFAULT 0, -- 是否锁定
	latest_replied_on BIGINT NOT NULL DEFAULT 0, -- 最新回复时间
	tags VARCHAR(255) NOT NULL DEFAULT '',
	attachment_price BIGINT NOT NULL DEFAULT 0, -- 附件价格(分)
	ip VARCHAR(64) NOT NULL DEFAULT '', -- IP地址
	ip_loc VARCHAR(64) NOT NULL DEFAULT '', -- IP城市地址
	created_on BIGINT NOT NULL DEFAULT 0,
	modified_on BIGINT NOT NULL DEFAULT 0,
	deleted_on BIGINT NOT NULL DEFAULT 0,
	is_del SMALLINT NOT NULL DEFAULT 0,
	PRIMARY KEY (id)
);
CREATE INDEX idx_post_user_id ON p_post USING btree (user_id);

CREATE TABLE p_post_attachment_bill (
	id BIGSERIAL PRIMARY KEY,
	post_id BIGINT NOT NULL DEFAULT 0,
	user_id BIGINT NOT NULL DEFAULT 0,
	paid_amount BIGINT NOT NULL DEFAULT 0, -- 支付金额
	created_on BIGINT NOT NULL DEFAULT 0,
	modified_on BIGINT NOT NULL DEFAULT 0,
	deleted_on BIGINT NOT NULL DEFAULT 0,
	is_del SMALLINT NOT NULL DEFAULT 0
);
CREATE INDEX idx_post_attachment_bill_post_id ON p_post_attachment_bill USING btree (post_id);
CREATE INDEX idx_post_attachment_bill_user_id ON p_post_attachment_bill USING btree (user_id);

CREATE TABLE p_post_collection (
	id BIGSERIAL PRIMARY KEY,
	post_id BIGINT NOT NULL DEFAULT 0,
	user_id BIGINT NOT NULL DEFAULT 0,
	created_on BIGINT NOT NULL DEFAULT 0,
	modified_on BIGINT NOT NULL DEFAULT 0,
	deleted_on BIGINT NOT NULL DEFAULT 0,
	is_del SMALLINT NOT NULL DEFAULT 0
);
CREATE INDEX idx_post_collection_post_id ON p_post_collection USING btree (post_id);
CREATE INDEX idx_post_collection_user_id ON p_post_collection USING btree (user_id);

CREATE TABLE p_post_content (
	id BIGSERIAL PRIMARY KEY,
	post_id BIGINT NOT NULL DEFAULT 0,
	user_id BIGINT NOT NULL DEFAULT 0,
	content VARCHAR(2000) NOT NULL DEFAULT '',
	"type" SMALLINT NOT NULL DEFAULT 2, -- 类型，1标题，2文字段落，3图片地址，4视频地址，5语音地址，6链接地址，7附件资源，8收费资源
	sort SMALLINT NOT NULL DEFAULT 100,
	created_on BIGINT NOT NULL DEFAULT 0,
	modified_on BIGINT NOT NULL DEFAULT 0,
	deleted_on BIGINT NOT NULL DEFAULT 0,
	is_del SMALLINT NOT NULL DEFAULT 0
);
CREATE INDEX idx_post_content_post_id ON p_post_content USING btree (post_id);
CREATE INDEX idx_post_content_user_id ON p_post_content USING btree (user_id);

CREATE TABLE p_post_star (
	id BIGSERIAL PRIMARY KEY,
	post_id BIGINT NOT NULL DEFAULT 0,
	user_id BIGINT NOT NULL DEFAULT 0,
	created_on BIGINT NOT NULL DEFAULT 0,
	modified_on BIGINT NOT NULL DEFAULT 0,
	deleted_on BIGINT NOT NULL DEFAULT 0,
	is_del SMALLINT NOT NULL DEFAULT 0
);
CREATE INDEX idx_post_star_post_id ON p_post_star USING btree (post_id);
CREATE INDEX idx_post_star_user_id ON p_post_star USING btree (user_id);

CREATE TABLE p_tag (
	id BIGSERIAL PRIMARY KEY,
	user_id BIGINT NOT NULL DEFAULT 0,
	tag VARCHAR(255) NOT NULL,
	quote_num BIGINT NOT NULL DEFAULT 0, -- 引用数
	created_on BIGINT NOT NULL DEFAULT 0,
	modified_on BIGINT NOT NULL DEFAULT 0,
	deleted_on BIGINT NOT NULL DEFAULT 0,
	is_del SMALLINT NOT NULL DEFAULT 0
);
CREATE UNIQUE INDEX idx_tag_tag ON p_tag USING btree (tag);
CREATE INDEX idx_tag_user_id ON p_tag USING btree (user_id);
CREATE INDEX idx_tag_quote_num ON p_tag USING btree (quote_num);

CREATE TABLE p_user (
	id BIGSERIAL PRIMARY KEY,
	nickname VARCHAR(32) NOT NULL DEFAULT '',
	username VARCHAR(32) NOT NULL DEFAULT '',
	phone VARCHAR(16) NOT NULL DEFAULT '', -- 手机号
	password VARCHAR(32) NOT NULL DEFAULT '', -- MD5密码
	salt VARCHAR(16) NOT NULL DEFAULT '', -- 盐值
	status SMALLINT NOT NULL DEFAULT 1, -- 状态，1正常，2停用
	avatar VARCHAR(255) NOT NULL DEFAULT '',
	balance BIGINT NOT NULL, -- 用户余额（分）
	is_admin BOOLEAN NOT NULL DEFAULT false, -- 是否管理员
	created_on BIGINT NOT NULL DEFAULT 0,
	modified_on BIGINT NOT NULL DEFAULT 0,
	deleted_on BIGINT NOT NULL DEFAULT 0,
	is_del SMALLINT NOT NULL DEFAULT 0
);
CREATE UNIQUE INDEX idx_user_username ON p_user USING btree (username);
CREATE INDEX idx_user_phone ON p_user USING btree (phone);

CREATE TABLE p_wallet_recharge (
	id BIGSERIAL PRIMARY KEY,
	user_id BIGINT NOT NULL DEFAULT 0,
	amount BIGINT NOT NULL DEFAULT 0, -- 充值金额
	trade_no VARCHAR(64) NOT NULL DEFAULT '', -- 支付宝订单号
	trade_status VARCHAR(32) NOT NULL DEFAULT '', -- 交易状态
	created_on BIGINT NOT NULL DEFAULT 0,
	modified_on BIGINT NOT NULL DEFAULT 0,
	deleted_on BIGINT NOT NULL DEFAULT 0,
	is_del SMALLINT NOT NULL DEFAULT 0 -- 是否删除 0 为未删除、1 为已删除
);
CREATE INDEX idx_wallet_recharge_user_id ON p_wallet_recharge USING btree (user_id);
CREATE INDEX idx_wallet_recharge_trade_no ON p_wallet_recharge USING btree (trade_no);
CREATE INDEX idx_wallet_recharge_trade_status ON p_wallet_recharge USING btree (trade_status);

CREATE TABLE p_wallet_statement (
	id BIGSERIAL PRIMARY KEY,
	user_id BIGINT NOT NULL DEFAULT 0,
	change_amount BIGINT NOT NULL DEFAULT 0, -- 变动金额
	balance_snapshot BIGINT NOT NULL DEFAULT 0, -- 资金快照
	reason VARCHAR(255) NOT NULL, -- 变动原因
	post_id BIGINT NOT NULL DEFAULT 0, -- 关联动态
	created_on BIGINT NOT NULL DEFAULT 0,
	modified_on BIGINT NOT NULL DEFAULT 0,
	deleted_on BIGINT NOT NULL DEFAULT 0,
	is_del SMALLINT NOT NULL DEFAULT 0
);
CREATE INDEX idx_wallet_statement_user_id ON p_wallet_statement USING btree (user_id);
