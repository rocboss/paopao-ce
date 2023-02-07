-- user ddl --
CREATE TABLE p_user (
	id bigserial PRIMARY KEY,
	nickname VARCHAR(32) NOT NULL DEFAULT '',
	username VARCHAR(32) NOT NULL DEFAULT '',
	phone VARCHAR(16) NOT NULL DEFAULT '',
	PASSWORD VARCHAR(32) NOT NULL DEFAULT '',
	salt VARCHAR(16) NOT NULL DEFAULT '',
	status SMALLINT NOT NULL DEFAULT 1,
	avatar VARCHAR(255) NOT NULL DEFAULT '',
	balance BIGINT NOT NULL,
	is_admin BOOLEAN NOT NULL DEFAULT FALSE,
	created_on BIGINT NOT NULL,
	modified_on BIGINT NOT NULL,
	deleted_on BIGINT NOT NULL,
	is_del BOOLEAN NOT NULL DEFAULT FALSE,
	UNIQUE (username)
);
COMMENT ON TABLE p_user IS '用户';
COMMENT ON COLUMN p_user.password IS 'MD5密码';
COMMENT ON COLUMN p_user.status IS '状态, 1正常, 2停用';
COMMENT ON COLUMN p_user.balance IS '用户余额（分）';
CREATE UNIQUE INDEX p_user_username_idx ON p_user (username);
CREATE INDEX p_user_phone_idx ON p_user (phone);

-- tag ddl --
CREATE TABLE p_tag (
	id bigserial PRIMARY KEY,
	user_id bigserial NOT NULL,
	tag VARCHAR(255) NOT NULL,
	quote_num BIGINT NOT NULL DEFAULT 0,
	created_on BIGINT NOT NULL DEFAULT 0,
	modified_on BIGINT NOT NULL DEFAULT 0,
	deleted_on BIGINT NOT NULL DEFAULT 0,
	is_del BOOLEAN NOT NULL DEFAULT FALSE,
	UNIQUE (tag)
);
COMMENT ON TABLE p_tag IS '主题标签';
COMMENT ON COLUMN p_tag.is_del IS '是否删除';
CREATE UNIQUE INDEX p_tag_tag_idx ON p_tag (tag);
CREATE INDEX p_tag_user_idx ON p_tag (user_id);
CREATE INDEX p_tag_quote_num_idx ON p_tag (quote_num);
