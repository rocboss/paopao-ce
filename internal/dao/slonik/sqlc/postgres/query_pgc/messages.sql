--------------------------------------------------------------------------------
-- message sql dml
--------------------------------------------------------------------------------

-- name: GetUnreadCount :one
SELECT count(*) FROM p_message WHERE receiver_user_id=$1 AND is_read=0 AND is_del=0;

-- name: GetMessageById :one
SELECT * FROM p_message WHERE id=$1 AND is_del=0;

-- name: ReadMessage :exec
UPDATE p_message SET is_read=1, modified_on=$1 WHERE id=$2;

-- name: ReadAllMessage :exec
UPDATE p_message SET is_read=1, modified_on=$1 WHERE receiver_user_id=$2;

-- name: GetSystemMessages :many
SELECT * 
FROM p_message 
WHERE receiver_user_id=$1 AND type IN (1, 2, 3, 99) AND is_del=0 
ORDER BY id DESC
LIMIT $2 OFFSET $3;

-- name: CountSystemMessages :one
SELECT count(*) 
FROM p_message 
WHERE receiver_user_id=$1 AND type IN (1, 2, 3, 99) AND is_del=0; 

-- name: GetWhisperMessages :many
SELECT * 
FROM p_message 
WHERE ((receiver_user_id=$1 OR sender_user_id=$2) AND type=4) AND is_del=0 
ORDER BY id DESC
LIMIT $3 OFFSET $4;

-- name: CountWhisperMessages :one
SELECT count(*) 
FROM p_message 
WHERE ((receiver_user_id=$1 OR sender_user_id=$2) AND type=4) AND is_del=0;

-- name: GetRequestingMessages :many
SELECT * 
FROM p_message 
WHERE receiver_user_id=$1 AND type=5 AND is_del=0 
ORDER BY id DESC
LIMIT $2 OFFSET $3;

-- name: CountRequestingMessages :one
SELECT count(*) 
FROM p_message 
WHERE receiver_user_id=$1 AND type=5 AND is_del=0;

-- name: GetUnreadMessages :many
SELECT * 
FROM p_message 
WHERE receiver_user_id=$1 AND is_read=0 AND is_del=0 
ORDER BY id DESC
LIMIT $2 OFFSET $3;

-- name: CountUnreadMessages :one
SELECT count(*) 
FROM p_message 
WHERE receiver_user_id=$1 AND is_read=0 AND is_del=0;

-- name: GetAllMessages :many
SELECT * 
FROM p_message 
WHERE (receiver_user_id=$1 OR (sender_user_id=$2 AND type=4)) AND is_del=0 
ORDER BY id DESC
LIMIT $3 OFFSET $4;

-- name: CountAllMessages :one
SELECT count(*)
FROM p_message 
WHERE (receiver_user_id=$1 OR (sender_user_id=$2 AND type=4)) AND is_del=0;
