CREATE TABLE p_topic_user (
	ID BIGSERIAL PRIMARY KEY,
	topic_id BIGINT NOT NULL,-- 标签ID
	user_id BIGINT NOT NULL,-- 创建者ID
	alias_name VARCHAR ( 255 ),-- 别名
	remark VARCHAR ( 512 ),-- 备注
	quote_num BIGINT,-- 引用数
	is_top SMALLINT NOT NULL DEFAULT 0,-- 是否置顶 0 为未置顶、1 为已置顶
	created_on BIGINT NOT NULL DEFAULT 0,-- 创建时间
	modified_on BIGINT NOT NULL DEFAULT 0,-- 修改时间
	deleted_on BIGINT NOT NULL DEFAULT 0,-- 删除时间
	is_del SMALLINT NOT NULL DEFAULT 0,-- 是否删除 0 为未删除、1 为已删除
	reserve_a VARCHAR ( 255 ),-- 保留字段a
	reserve_b VARCHAR ( 255 ) -- 保留字段b
);
CREATE UNIQUE INDEX idx_topic_user_uid_tid ON p_topic_user USING btree ( topic_id, user_id );