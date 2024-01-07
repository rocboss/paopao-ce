--------------------------------------------------------------------------------
-- file yesql.sql
-- this sql file contain some PostgreSQL special sql dml
-- version is a*
--------------------------------------------------------------------------------

--------------------------------------------------------------------------------
-- topic_a sql dml
--------------------------------------------------------------------------------

-- name: insert_tag@topic_a
INSERT INTO @tag (user_id, tag, created_on, modified_on, quote_num)
VALUES (?, ?, ?, ?, 1)
RETURNING id;

-- name: stick_topic@topic_a
-- prepare: stmt
UPDATE @topic_user
SET is_top=1-is_top, modified_on=?
WHERE user_id=? AND topic_id=? AND is_del=0;
RETURNING is_top;

-- name: pin_topic@topic_a
-- prepare: stmt
UPDATE @topic_user
SET is_pin=1-is_pin, modified_on=?
WHERE user_id=? AND topic_id=? AND is_del=0;
RETURNING is_pin;
