--------------------------------------------------------------------------------
-- comment sql dml
--------------------------------------------------------------------------------

-- name: GetNewestComments :many
SELECT * FROM p_comment WHERE post_id=$1 AND is_del=0 ORDER BY id DESC LIMIT $2 OFFSET $3; 

-- name: GetDefaultComments :many
SELECT * FROM p_comment WHERE post_id=$1 AND is_del=0 ORDER BY id ASC LIMIT $2 OFFSET $3; 

--------------------------------------------------------------------------------
-- comment_manage sql dml
--------------------------------------------------------------------------------

-- name: DeleteComment :exec
UPDATE p_comment SET deleted_on=$1, is_del=1 WHERE id=$2 AND is_del=0;
