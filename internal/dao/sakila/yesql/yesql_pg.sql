--------------------------------------------------------------------------------
-- file yesql_pg.sql
-- this sql file just define global sql dml
-- this sql file contain some PostgreSQL special sql dml
--------------------------------------------------------------------------------

--------------------------------------------------------------------------------
-- global sql dml 
--------------------------------------------------------------------------------

-- name: insert_tag
-- prepare: stmt
INSERT INTO @tag (user_id, tag, created_on, modified_on, quote_num)
VALUES (?, ?, ?, ?, 1)
RETURNING id;

