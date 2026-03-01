CREATE TABLE IF NOT EXISTS p_emoji (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL DEFAULT 0,
    name VARCHAR(64) NOT NULL DEFAULT '',
    url VARCHAR(255) NOT NULL DEFAULT '',
    width INTEGER NOT NULL DEFAULT 0,
    height INTEGER NOT NULL DEFAULT 0,
    size BIGINT NOT NULL DEFAULT 0,
    type VARCHAR(32) NOT NULL DEFAULT '',
    collection_count BIGINT NOT NULL DEFAULT 0,
    created_on BIGINT NOT NULL DEFAULT 0,
    modified_on BIGINT NOT NULL DEFAULT 0,
    deleted_on BIGINT NOT NULL DEFAULT 0,
    is_del SMALLINT NOT NULL DEFAULT 0
);

CREATE INDEX IF NOT EXISTS idx_emoji_user ON p_emoji(user_id);
CREATE INDEX IF NOT EXISTS idx_emoji_collection ON p_emoji(collection_count);

CREATE TABLE IF NOT EXISTS p_user_emoji (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL DEFAULT 0,
    emoji_id BIGINT NOT NULL DEFAULT 0,
    created_on BIGINT NOT NULL DEFAULT 0,
    modified_on BIGINT NOT NULL DEFAULT 0,
    deleted_on BIGINT NOT NULL DEFAULT 0,
    is_del SMALLINT NOT NULL DEFAULT 0
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_user_emoji ON p_user_emoji(user_id, emoji_id);
CREATE INDEX IF NOT EXISTS idx_user_emoji_emoji ON p_user_emoji(emoji_id);