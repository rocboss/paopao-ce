--------------------------------------------------------------------------------
-- file yesql.sql
-- this sql file contain some PostgreSQL special sql dml
--------------------------------------------------------------------------------

-- name: add_post@tweet_manage
-- prepare: named_stmt
INSERT INTO @post (user_id, tags, ip, ip_loc, attachment_price, visibility, latest_replied_on, created_on)
VALUES (:user_id, :tags, :ip, :ip_loc, :attachment_price, :visibility, :latest_replied_on, :created_on)
RETURNING id;

