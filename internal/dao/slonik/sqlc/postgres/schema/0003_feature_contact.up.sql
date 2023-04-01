CREATE TABLE p_contact (
	id BIGSERIAL PRIMARY KEY,
	user_id BIGINT NOT NULL,
	friend_id BIGINT NOT NULL,
	group_id BIGINT NOT NULL DEFAULT 0, -- 好友分组ID:默认为0无分组
	remark VARCHAR(32) NOT NULL DEFAULT '', -- 好友备注
	status SMALLINT NOT NULL DEFAULT 0, -- 好友状态: 1请求好友, 2已好友, 3拒绝好友, 4已删好友
	is_top SMALLINT NOT NULL DEFAULT 0, -- 是否置顶, 0否, 1是
	is_black SMALLINT NOT NULL DEFAULT 0, -- 是否为黑名单, 0否, 1是
	is_del SMALLINT NOT NULL DEFAULT 0, -- 否删除好友, 0否, 1是
	notice_enable SMALLINT NOT NULL DEFAULT 0, -- 是否有消息提醒, 0否, 1是
	created_on BIGINT NOT NULL DEFAULT 0,
	modified_on BIGINT NOT NULL DEFAULT 0,
	deleted_on BIGINT NOT NULL DEFAULT 0
);
CREATE UNIQUE INDEX idx_contact_user_friend ON p_contact USING btree (user_id,friend_id);
CREATE INDEX idx_contact_user_friend_status ON p_contact USING btree (user_id, friend_id, status);

CREATE TABLE p_contact_group (
	id BIGSERIAL PRIMARY KEY,
	user_id int NOT NULL DEFAULT 0,
	name VARCHAR(32) NOT NULL DEFAULT '', -- 分组名称
	is_del SMALLINT NOT NULL DEFAULT 1, -- 是否删除, 0否, 1是
	created_on BIGINT NOT NULL DEFAULT 0,
	modified_on BIGINT NOT NULL DEFAULT 0,
	deleted_on BIGINT NOT NULL DEFAULT 0
);
