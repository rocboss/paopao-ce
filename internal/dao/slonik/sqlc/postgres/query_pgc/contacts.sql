--------------------------------------------------------------------------------
-- contact_manager sql dml
--------------------------------------------------------------------------------

-- name: CreateContact :exec
INSERT INTO p_contact (user_id, friend_id, status, created_on) VALUES ($1, $2, $3, $4);

-- name: FreshContactStatus :exec
UPDATE p_contact SET status=$1, modified_on=$2, is_del=0 WHERE id=$3;

-- name: CreateMessage :one
INSERT INTO p_message (sender_user_id, receiver_user_id, type, brief, content, reply_id, created_on)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING id;

-- name: AddFriendMsgsUpdate :exec
UPDATE p_message
SET reply_id=$1, modified_on=$2
WHERE ((sender_user_id = $3 AND receiver_user_id = $4) OR
        (sender_user_id = $4 AND receiver_user_id = $3)) AND
    type = $5 AND reply_id = $6;

-- name: RejectFriendMsgsUpdate :exec
UPDATE p_message
SET reply_id=$1, modified_on=$2
WHERE sender_user_id=$3 AND receiver_user_id=$4 AND type=$5 AND reply_id=$6;

-- name: DeleteFriend :exec
UPDATE p_contact SET status=4, is_del=1, deleted_on=$1 WHERE id=$2;

-- name: ListFriend :many
SELECT c.friend_id user_id,
    u.username username,
    u.nickname nickname,
    u.avatar avatar,
    u.phone phone
FROM p_contact c 
JOIN p_user u 
ON c.friend_id=u.id
WHERE c.user_id=$1 AND c.status=2 AND c.is_del=0
ORDER BY u.nickname ASC
LIMIT $2 OFFSET $3;

-- name: CountFriendsById :one
SELECT count(*) FROM p_contact WHERE user_id=$1 AND status=2 AND is_del=0;

-- name: GetContacts :many
SELECT id, user_id, friend_id, group_id, remark, status, is_top, is_black, notice_enable, is_del
FROM p_contact
WHERE (user_id=$1 AND friend_id=$2) OR (user_id=$3 AND friend_id=$4);

-- name: GetUserFriend :one
SELECT id, user_id, friend_id, group_id, remark, status, is_top, is_black, notice_enable, is_del
FROM p_contact
WHERE user_id=$1 AND friend_id=$2 AND is_del=0;

-- name: GetContact :one
SELECT id, user_id, friend_id, group_id, remark, status, is_top, is_black, notice_enable, is_del
FROM p_contact
WHERE user_id=$1 AND friend_id=$2;

-- name: IsFriend :one
SELECT true FROM p_contact WHERE user_id=$1 AND friend_id=$2 AND is_del=0 AND status=2;
