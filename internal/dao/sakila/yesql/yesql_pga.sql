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
