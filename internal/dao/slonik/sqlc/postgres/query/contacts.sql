--------------------------------------------------------------------------------
-- contact_manager sql dml
--------------------------------------------------------------------------------

-- name: CreateContact :exec
INSERT INTO p_contact (user_id, friend_id, status, created_on) VALUES ($1, $2, $3, $4);
