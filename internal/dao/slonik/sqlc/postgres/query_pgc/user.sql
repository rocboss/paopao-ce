
--------------------------------------------------------------------------------
-- user_manage sql dml
--------------------------------------------------------------------------------

-- name: GetRegisterUserCount :one
SELECT count(*) FROM p_user WHERE is_del=0;

-- name: GetUserById :one
SELECT * FROM p_user WHERE id=$1 AND is_del=0;

-- name: GetUserByUsername :one
SELECT * FROM p_user WHERE username=$1 AND is_del=0;

-- name: GetUserByPhone :one
SELECT * FROM p_user WHERE phone=$1 AND is_del=0;

-- name: GetUsersByIds :many
SELECT * FROM p_user WHERE id = ANY(@ids::BIGINT[]) AND is_del=0;

-- name: GetUsersByKeyword :many
SELECT * FROM p_user WHERE username LIKE $1 AND is_del=0 limit 6;

-- name: GetAnyusers :many
SELECT * FROM p_user WHERE is_del=0 ORDER BY id ASC limit 6;

-- name: CreateUser :one
INSERT INTO p_user (username, nickname, password, salt, avatar, status, created_on, balance) 
VALUES ($1, $2, $3, $4, $5, $6, $7, 0)
RETURNING id;

-- name: UpdateUser :exec
UPDATE p_user 
SET username=$1, 
    nickname=$2, 
    phone=$3,
    password=$4,
    salt=$5,
    status=$6,
    avatar=$7,
    balance=$8,
    is_admin=$9,
    modified_on=$10
WHERE id=$11 AND is_del=0;

--------------------------------------------------------------------------------
-- user_relation sql dml
--------------------------------------------------------------------------------

-- name: MyFriendIds :many
SELECT friend_id FROM p_contact WHERE user_id=$1 AND is_del=0;

-- name: MyFollowIds :many
SELECT follow_id FROM p_following WHERE user_id=$1 AND is_del=0;
