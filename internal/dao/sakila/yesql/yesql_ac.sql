--------------------------------------------------------------------------------
-- file yesql_ac.sql
-- this sql file define all sql dml for next "A" version data logic
--------------------------------------------------------------------------------

--------------------------------------------------------------------------------
-- ship_index_a sql dml
--------------------------------------------------------------------------------

-- name: user_info@ship_index_a
-- prepare: raw
SELECT * FROM @user WHERE username=?

--------------------------------------------------------------------------------
-- simple_index_a sql dml
--------------------------------------------------------------------------------

-- name: user_info@simple_index_a
-- prepare: raw
SELECT * FROM @user WHERE username=?

--------------------------------------------------------------------------------
-- tweet_a sql dml
--------------------------------------------------------------------------------

-- name: user_info@tweet_a
-- prepare: raw
SELECT * FROM @user WHERE username=?

-- name: tweet_info_by_id@tweet_a
-- prepare: raw
SELECT * FROM @user WHERE username=?

-- name: tweet_item_by_id@tweet_a
-- prepare: raw
SELECT * FROM @user WHERE username=?

-- name: user_tweets_by_admin@tweet_a
-- prepare: raw
SELECT * FROM @user WHERE username=?

-- name: user_tweets_by_self@tweet_a
-- prepare: raw
SELECT * FROM @user WHERE username=?

-- name: user_tweets_by_friend@tweet_a
-- prepare: raw
SELECT * FROM @user WHERE username=?

-- name: user_tweets_by_guest@tweet_a
-- prepare: raw
SELECT * FROM @user WHERE username=?

-- name: reaction_by_tweet_id@tweet_a
-- prepare: raw
SELECT * FROM @user WHERE username=?

-- name: user_reactions@tweet_a
-- prepare: raw
SELECT * FROM @user WHERE username=?

-- name: favorite_by_tweet_id@tweet_a
-- prepare: raw
SELECT * FROM @user WHERE username=?

-- name: user_favorites@tweet_a
-- prepare: raw
SELECT * FROM @user WHERE username=?

-- name: attachment_by_tweet_id@tweet_a
-- prepare: raw
SELECT * FROM @user WHERE username=?

--------------------------------------------------------------------------------
-- tweet_manage_a sql dml
--------------------------------------------------------------------------------

-- name: user_info@tweet_manage_a
-- prepare: raw
SELECT * FROM @user WHERE username=?

--------------------------------------------------------------------------------
-- tweet_help_a sql dml
--------------------------------------------------------------------------------

-- name: user_info@tweet_help_a
-- prepare: raw
SELECT * FROM @user WHERE username=?

--------------------------------------------------------------------------------
-- topic_a sql dml
--------------------------------------------------------------------------------

-- name: newest_tags@topic_a
-- get newest tag information
-- prepare: stmt
SELECT t.id id, 
	t.user_id user_id, 
	t.tag tag, 
	t.quote_num quote_num, 
	u.id "u.id", 
	u.nickname "u.nickname",
	u.username "u.username", 
	u.status "u.status",
	u.avatar "u.avatar",
	u.is_admin "u.is_admin" 
FROM @tag t 
JOIN @user u ON t.user_id = u.id 
WHERE t.is_del = 0 AND t.quote_num > 0 
ORDER BY t.id DESC 
LIMIT ? OFFSET ?;

-- name: hot_tags@topic_a
-- get get host tag information
-- prepare: stmt
SELECT t.id id, 
	t.user_id user_id, 
	t.tag tag, 
	t.quote_num quote_num, 
	u.id "u.id", 
	u.nickname "u.nickname",
	u.username "u.username", 
	u.status "u.status",
	u.avatar "u.avatar",
	u.is_admin "u.is_admin" 
FROM @tag t 
JOIN @user u ON t.user_id = u.id 
WHERE t.is_del = 0 AND t.quote_num > 0 
ORDER BY t.quote_num DESC 
LIMIT ? OFFSET ?;

-- name: follow_tags@topic_a
-- get get follow tag information
-- prepare: stmt
SELECT t.id id, 
	t.user_id user_id, 
	t.tag tag, 
	t.quote_num quote_num, 
	u.id "u.id", 
	1 as is_following,
	c.is_top,
	u.nickname "u.nickname",
	u.username "u.username", 
	u.status "u.status",
	u.avatar "u.avatar",
	u.is_admin "u.is_admin" 
FROM @topic_user c
JOIN @user u ON c.user_id = u.id 
JOIN @tag t ON c.topic_id = t.id
WHERE c.is_del = 0 AND t.quote_num > 0 AND c.user_id=?
ORDER BY c.is_top DESC, t.quote_num DESC
LIMIT ? OFFSET ?;

-- name: topic_infos@topic_a
-- prepare: raw
-- clause: in
SELECT topic_id, is_top 
FROM @topic_user 
WHERE is_del=0 AND user_id=? AND topic_id IN (?);

-- name: exist_topic_user@topic_a
-- prepare: stmt
SELECT 1 FROM @topic_user WHERE user_id=? AND topic_id=? AND is_del=0;

-- name: follow_topic@topic_a
-- prepare: stmt
INSERT INTO @topic_user(user_id, topic_id, created_on) VALUES (?, ?, ?)

-- name: unfollow_topic@topic_a
-- prepare: stmt
DELETE FROM @topic_user 
WHERE user_id=? AND topic_id=? AND is_del=0;

-- name: stick_topic@topic_a
-- prepare: stmt
UPDATE @topic_user
SET is_top=1-is_top, modified_on=?
WHERE user_id=? AND topic_id=? AND is_del=0;

-- name: topic_is_top@topic_a
-- prepare: stmt
SELECT is_top FROM @topic_user WHERE user_id=? AND topic_id=? AND is_del=0;

-- name: tags_by_keyword_a@topic_a
-- get tags by keyword
SELECT id, user_id, tag, quote_num 
FROM @tag 
WHERE is_del = 0 
ORDER BY quote_num DESC 
LIMIT 6;

-- name: tags_by_keyword_b@topic_a
SELECT id, user_id, tag, quote_num 
FROM @tag WHERE is_del = 0 AND tag LIKE ? 
ORDER BY quote_num DESC 
LIMIT 6;

-- name: insert_tag@topic_a
INSERT INTO @tag (user_id, tag, created_on, modified_on, quote_num)
VALUES (?, ?, ?, ?, 1);

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
