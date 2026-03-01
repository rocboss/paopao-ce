CREATE TABLE IF NOT EXISTS p_emoji (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL DEFAULT 0,
    name TEXT NOT NULL DEFAULT '',
    url TEXT NOT NULL DEFAULT '',
    width INTEGER NOT NULL DEFAULT 0,
    height INTEGER NOT NULL DEFAULT 0,
    size INTEGER NOT NULL DEFAULT 0,
    type TEXT NOT NULL DEFAULT '',
    collection_count INTEGER NOT NULL DEFAULT 0,
    created_on INTEGER NOT NULL DEFAULT 0,
    modified_on INTEGER NOT NULL DEFAULT 0,
    deleted_on INTEGER NOT NULL DEFAULT 0,
    is_del INTEGER NOT NULL DEFAULT 0
);

CREATE INDEX IF NOT EXISTS idx_emoji_user ON p_emoji(user_id);
CREATE INDEX IF NOT EXISTS idx_emoji_collection ON p_emoji(collection_count);

CREATE TABLE IF NOT EXISTS p_user_emoji (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL DEFAULT 0,
    emoji_id INTEGER NOT NULL DEFAULT 0,
    created_on INTEGER NOT NULL DEFAULT 0,
    modified_on INTEGER NOT NULL DEFAULT 0,
    deleted_on INTEGER NOT NULL DEFAULT 0,
    is_del INTEGER NOT NULL DEFAULT 0
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_user_emoji ON p_user_emoji(user_id, emoji_id);
CREATE INDEX IF NOT EXISTS idx_user_emoji_emoji ON p_user_emoji(emoji_id);