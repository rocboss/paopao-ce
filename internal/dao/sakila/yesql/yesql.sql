--------------------------------------------------------------------------------
-- authorization_manage sql dml
--------------------------------------------------------------------------------

-- name: user_info@authorization_manage
-- prepare: stmt
SELECT * FROM @user WHERE username=?

--------------------------------------------------------------------------------
-- comment sql dml
--------------------------------------------------------------------------------

-- name: user_info@comment
-- prepare: stmt
SELECT * FROM @user WHERE username=?

--------------------------------------------------------------------------------
-- comment_manage sql dml
--------------------------------------------------------------------------------

-- name: user_info@comment_manage
-- prepare: stmt
SELECT * FROM @user WHERE username=?

--------------------------------------------------------------------------------
-- contact_manager sql dml
--------------------------------------------------------------------------------

-- name: user_info@contact_manager
-- prepare: stmt
SELECT * FROM @user WHERE username=?

--------------------------------------------------------------------------------
-- message sql dml
--------------------------------------------------------------------------------

-- name: user_info@message
-- prepare: stmt
SELECT * FROM @user WHERE username=?

--------------------------------------------------------------------------------
-- security sql dml
--------------------------------------------------------------------------------

-- name: user_info@security
-- prepare: stmt
SELECT * FROM @user WHERE username=?

--------------------------------------------------------------------------------
-- friend_index sql dml
--------------------------------------------------------------------------------

-- name: user_info@friend_index
-- prepare: stmt
SELECT * FROM @user WHERE username=?

--------------------------------------------------------------------------------
-- follow_index sql dml
--------------------------------------------------------------------------------

-- name: user_info@follow_index
-- prepare: stmt
SELECT * FROM @user WHERE username=?

--------------------------------------------------------------------------------
-- light_index sql dml
--------------------------------------------------------------------------------

-- name: user_info@light_index
-- prepare: stmt
SELECT * FROM @user WHERE username=?

--------------------------------------------------------------------------------
-- simple_index sql dml
--------------------------------------------------------------------------------

-- name: user_info@simple_index
-- prepare: stmt
SELECT * FROM @user WHERE username=?

--------------------------------------------------------------------------------
-- friend_index_a sql dml
--------------------------------------------------------------------------------

-- name: user_info@friend_index_a
-- prepare: stmt
SELECT * FROM @user WHERE username=?

--------------------------------------------------------------------------------
-- follow_index_a sql dml
--------------------------------------------------------------------------------

-- name: user_info@follow_index_a
-- prepare: stmt
SELECT * FROM @user WHERE username=?

--------------------------------------------------------------------------------
-- light_index_a sql dml
--------------------------------------------------------------------------------

-- name: user_info@light_index_a
-- prepare: stmt
SELECT * FROM @user WHERE username=?

--------------------------------------------------------------------------------
-- simple_index_a sql dml
--------------------------------------------------------------------------------

-- name: user_info@simple_index_a
-- prepare: stmt
SELECT * FROM @user WHERE username=?

--------------------------------------------------------------------------------
-- tweet sql dml
--------------------------------------------------------------------------------

-- name: user_info@tweet
-- prepare: stmt
SELECT * FROM @user WHERE username=?

--------------------------------------------------------------------------------
-- tweet_manage sql dml
--------------------------------------------------------------------------------

-- name: user_info@tweet_manage
-- prepare: stmt
SELECT * FROM @user WHERE username=?

--------------------------------------------------------------------------------
-- tweet_help sql dml
--------------------------------------------------------------------------------

-- name: user_info@tweet_help
-- prepare: stmt
SELECT * FROM @user WHERE username=?

--------------------------------------------------------------------------------
-- tweet_a sql dml
--------------------------------------------------------------------------------

-- name: user_info@tweet_a
-- prepare: stmt
SELECT * FROM @user WHERE username=?

--------------------------------------------------------------------------------
-- tweet_manage_a sql dml
--------------------------------------------------------------------------------

-- name: user_info@tweet_manage_a
-- prepare: stmt
SELECT * FROM @user WHERE username=?

--------------------------------------------------------------------------------
-- tweet_help_a sql dml
--------------------------------------------------------------------------------

-- name: user_info@tweet_help_a
-- prepare: stmt
SELECT * FROM @user WHERE username=?

--------------------------------------------------------------------------------
-- topic_a sql dml
--------------------------------------------------------------------------------

-- name: newest_tags@topic_a
-- get newest tag information
SELECT t.id id, t.user_id user_id, t.tag tag, t.quote_num quote_num, u.id, u.nickname, u.username, u.status, u.avatar, u.is_admin 
FROM @tag t 
JOIN @user u 
ON t.user_id = u.id 
WHERE t.is_del = 0 AND t.quote_num > 0 
ORDER BY t.id DESC 
LIMIT ? OFFSET ?;

-- name: hot_tags@topic_a
-- get get host tag information
SELECT t.id id, t.user_id user_id, t.tag tag, t.quote_num quote_num, u.id, u.nickname, u.username, u.status, u.avatar, u.is_admin 
FROM @tag t 
JOIN @user u 
ON t.user_id = u.id 
WHERE t.is_del = 0 AND t.quote_num > 0 
ORDER BY t.quote_num DESC 
LIMIT ? OFFSET ?;

-- name: tags_by_keyword_a@topic_a
-- get tags by keyword
SELECT id, user_id, tag, quote_num FROM @tag WHERE is_del = 0 ORDER BY quote_num DESC LIMIT 6;

-- name: tags_by_keyword_b@topic_a
SELECT id, user_id, tag, quote_num FROM @tag WHERE is_del = 0 AND tag LIKE ? ORDER BY quote_num DESC LIMIT 6;

-- name: insert_tag@topic_a
INSERT INTO @tag (user_id, tag, created_on, modified_on, quote_num) VALUES (?, ?, ?, ?, 1);

-- name: tags_by_id_a@topic_a
-- prepare: raw
-- clause: in
SELECT id FROM @tag WHERE id IN (?) AND is_del = 0 AND quote_num > 0;

-- name: tags_by_id_b@topic_a
-- prepare: raw
-- clause: in
SELECT id, user_id, tag, quote_num FROM @tag WHERE id IN (?);

-- name: decr_tags_by_id@topic_a
-- prepare: raw
-- clause: in
UPDATE @tag SET quote_num=quote_num-1, modified_on=? WHERE id IN (?);

-- name: tags_for_incr@topic_a
-- prepare: raw
-- clause: in
SELECT id, user_id, tag, quote_num FROM @tag WHERE tag IN (?);

-- name: incr_tags_by_id@topic_a
-- prepare: raw
-- clause: in
UPDATE @tag SET quote_num=quote_num+1, is_del=0, modified_on=? WHERE id IN (?);

--------------------------------------------------------------------------------
-- user_manage sql dml
--------------------------------------------------------------------------------

-- name: user_info@user_manage
-- prepare: stmt
SELECT * FROM @user WHERE username=?

--------------------------------------------------------------------------------
-- wallet sql dml
--------------------------------------------------------------------------------

-- name: user_info@wallet
-- prepare: stmt
SELECT * FROM @user WHERE username=?