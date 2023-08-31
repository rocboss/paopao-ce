--------------------------------------------------------------------------------
-- message sql dml
--------------------------------------------------------------------------------

-- name: GetUnreadCount :one
SELECT count(*) FROM p_message WHERE receiver_user_id=$1 AND is_read=0 AND is_del=0;
