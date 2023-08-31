--------------------------------------------------------------------------------
-- following_manager sql dml
--------------------------------------------------------------------------------

-- name: CreateFollowing :exec
INSERT INTO p_following (user_id, follow_id, created_on) VALUES ($1, $2, $3);
