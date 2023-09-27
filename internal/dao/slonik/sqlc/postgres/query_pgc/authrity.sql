--------------------------------------------------------------------------------
-- authorization_manage sql dml
--------------------------------------------------------------------------------

-- name: BeFriendIds :many
SELECT user_id FROM p_contact WHERE friend_id=$1 AND status=2 AND is_del=0;

-- name: MyFriendSet :many
SELECT friend_id FROM p_contact WHERE user_id=$1 AND status=2 AND is_del=0;
