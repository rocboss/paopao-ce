--------------------------------------------------------------------------------
-- ship_index sql dml
--------------------------------------------------------------------------------

-- name: IndexByAdmin :many
SELECT * 
FROM p_post 
WHERE is_del=0 
ORDER BY is_top DESC, latest_replied_on DESC
LIMIT $1 OFFSET $2;
