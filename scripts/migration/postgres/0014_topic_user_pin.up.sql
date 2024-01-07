ALTER TABLE p_topic_user ADD COLUMN is_pin SMALLINT NOT NULL DEFAULT 0; -- 是否钉住 0 为未钉住、1 为已钉住
CREATE INDEX idx_topic_user_uid_ispin ON p_topic_user USING btree ( user_id, is_pin );
