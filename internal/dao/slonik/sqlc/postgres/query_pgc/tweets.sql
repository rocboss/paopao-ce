--------------------------------------------------------------------------------
-- tweet sql dml
--------------------------------------------------------------------------------

-- name: GetPostById :one
SELECT * FROM p_post WHERE id=$1 AND is_del=0;
