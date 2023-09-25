--------------------------------------------------------------------------------
-- file yesql.sql
-- this sql file contain some PostgreSQL special sql dml
-- version is c*
--------------------------------------------------------------------------------

--------------------------------------------------------------------------------
-- tweet sql dml
--------------------------------------------------------------------------------

-- name: list_index_hots_tweets@tweet
-- prepare: stmt
SELECT post.*
FROM @post post
LEFT JOIN @post_metric metric
ON post.id=metric.post_id
WHERE post.visibility>=90 AND post.is_del=0
ORDER BY post.is_top DESC, metric.rank_score DESC NULLS LAST, post.latest_replied_on DESC
LIMIT ? OFFSET ?;

--------------------------------------------------------------------------------
-- tweet_manage sql dml
--------------------------------------------------------------------------------

-- name: add_post@tweet_manage
-- prepare: named_stmt
INSERT INTO @post (user_id, tags, ip, ip_loc, attachment_price, visibility, latest_replied_on, created_on)
VALUES (:user_id, :tags, :ip, :ip_loc, :attachment_price, :visibility, :latest_replied_on, :created_on)
RETURNING id;

-- name: add_post_content@tweet_manage
-- prepare: named_stmt
INSERT INTO @post_content (post_id, user_id, content, type, sort, created_on)
VALUES (:post_id, :user_id, :content, :type, :sort, :created_on)
RETURNING id;

-- name: add_attachment@tweet_manage
-- prepare: stmt
INSERT INTO @attachment (user_id, file_size, img_width, img_height, type, content, created_on)
VALUES (?, ?, ?, ?, ?, ?, ?)
RETURNING id;

-- name: add_post_star@tweet_manage
-- prepare: stmt
INSERT INTO @post_star (post_id, user_id, created_on)
VALUES (?, ?, ?)
RETURNING *;

-- name: add_post_collection@tweet_manage
-- prepare: stmt
INSERT INTO @post_collection (post_id, user_id, created_on)
VALUES (?, ?, ?)
RETURNING *;

--------------------------------------------------------------------------------
-- comment sql dml
--------------------------------------------------------------------------------

-- name: get_hots_comments@comment
-- prepare: stmt
SELECT c.* 
FROM @comment c
LEFT JOIN @comment_metric m
ON c.id=m.comment_id
WHERE c.post_id=? AND c.is_del=0 AND m.is_del=0
ORDER BY is_essence DESC, m.rank_score DESC NULLS LAST, id DESC
LIMIT ? OFFSET ?; 

--------------------------------------------------------------------------------
-- user_manage sql dml
--------------------------------------------------------------------------------

-- name: create_user@user_manage
-- prepare: named_stmt
INSERT INTO @user (username, nickname, password, salt, avatar, status, created_on, balance) 
VALUES (:username, :nickname, :password, :salt, :avatar, :status, :created_on, 0)
RETURNING id;

--------------------------------------------------------------------------------
-- comment_manage sql dml
--------------------------------------------------------------------------------

-- name: create_comment@comment_manage
-- prepare: stmt
INSERT INTO @comment (post_id, user_id, ip, ip_loc, created_on) 
VALUES (?, ?, ?, ?, ?)
RETURNING *;

-- name: create_comment_content@comment_manage
-- prepare: stmt
INSERT INTO @comment_content (comment_id, user_id, content, type, sort, created_on) 
VALUES (?, ?, ?, ?, ?, ?)
RETURNING *;

-- name: create_comment_reply@comment_manage
-- prepare: stmt
INSERT INTO @comment_reply (comment_id, user_id, content, at_user_id, ip, ip_loc, created_on) 
VALUES (?, ?, ?, ?, ?, ?, ?)
RETURNING *;

--------------------------------------------------------------------------------
-- contact_manager sql dml
--------------------------------------------------------------------------------

-- name: create_contact@contact_manager
-- prepare: stmt
INSERT INTO @contact (user_id, friend_id, status, created_on) VALUES (?, ?, ?, ?)
RETURNING *;

--------------------------------------------------------------------------------
-- message sql dml
--------------------------------------------------------------------------------

-- name: create_message@message
-- prepare: named_stmt
INSERT INTO @message (sender_user_id, receiver_user_id, type, brief, content, post_id, comment_id, reply_id, created_on) 
VALUES (:sender_user_id, :receiver_user_id, :type, :brief, :content, :post_id, :comment_id, :reply_id, :created_on)
RETURNING id;

--------------------------------------------------------------------------------
-- wallet sql dml
--------------------------------------------------------------------------------

-- name: create_recharge@wallet
-- prepare: stmt
INSERT INTO @wallet_recharge (user_id, amount, created_on) VALUES (?, ?, ?)
RETURNING *;
