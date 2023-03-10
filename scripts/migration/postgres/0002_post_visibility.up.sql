ALTER TABLE p_post ADD COLUMN visibility SMALLINT NOT NULL DEFAULT 0; -- 可见性 0公开 1私密 2好友可见
CREATE INDEX idx_visibility ON p_post USING btree (visibility)
