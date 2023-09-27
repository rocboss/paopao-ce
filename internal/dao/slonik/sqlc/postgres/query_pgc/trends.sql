--------------------------------------------------------------------------------
-- trends_manager sql dml
--------------------------------------------------------------------------------

-- name: GetIndexTrends :many
SELECT u.username username,
    u.nickname nickname,
    u.avatar avatar
FROM p_contact c
JOIN p_user u 
ON c.friend_id=u.id
JOIN p_user_metric m 
ON c.friend_id=m.user_id 
WHERE c.user_id=$1 AND c.is_del=0 AND u.is_del=0 AND m.is_del=0 AND m.tweets_count>0
LIMIT $2 OFFSET $3;

-- name: CountIndexTrends :one
SELECT count(*)
FROM p_contact c 
JOIN p_user u 
ON c.friend_id=u.id 
JOIN p_user_metric m 
ON c.friend_id=m.user_id 
WHERE c.user_id=$1 AND c.is_del=0 AND u.is_del=0 AND m.is_del=0 AND m.tweets_count>0;
