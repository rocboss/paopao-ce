--------------------------------------------------------------------------------
-- following_manager sql dml
--------------------------------------------------------------------------------

-- name: CreateFollowing :exec
INSERT INTO p_following (user_id, follow_id, created_on) 
VALUES ($1, $2, $3);

-- name: ExistFollowing :one
SELECT true FROM p_following WHERE user_id=$1 AND follow_id=$2 AND is_del=0;

-- name: DeleteFollowing :exec
UPDATE p_following
SET is_del=1, deleted_on=$1
WHERE user_id=$2 AND follow_id=$3 AND is_del=0;

-- name: ListFollows :many
SELECT u.id user_id,
    u.username username,
    u.nickname nickname,
    u.avatar avatar,
    u.created_on created_on
FROM p_following f JOIN p_user u ON f.follow_id=u.id
WHERE f.user_id=$1 AND f.is_del=0
ORDER BY u.nickname ASC
LIMIT $2 OFFSET $3;

-- name: CountFollows :one
SELECT count(*) FROM p_following WHERE user_id=$1 AND is_del=0;

-- name: ListFollowings :many
SELECT u.id user_id,
    u.username username,
    u.nickname nickname,
    u.avatar avatar,
    u.created_on created_on
FROM p_following f JOIN p_user u ON f.user_id=u.id
WHERE f.follow_id=$1 AND f.is_del=0
ORDER BY u.nickname ASC
LIMIT $2 OFFSET $3;

-- name: CountFollowings :one
SELECT count(*) FROM p_following WHERE follow_id=$1 AND is_del=0;
