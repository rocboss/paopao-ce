CREATE TABLE p_following (
	id BIGSERIAL PRIMARY KEY,
	user_id BIGINT NOT NULL,
	follow_id BIGINT NOT NULL,
	is_del SMALLINT NOT NULL DEFAULT 0, -- 是否删除, 0否, 1是
	created_on BIGINT NOT NULL DEFAULT 0,
	modified_on BIGINT NOT NULL DEFAULT 0,
	deleted_on BIGINT NOT NULL DEFAULT 0
);
CREATE INDEX idx_following_user_follow ON p_following USING btree (user_id, follow_id);
