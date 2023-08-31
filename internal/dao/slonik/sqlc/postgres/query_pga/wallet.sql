--------------------------------------------------------------------------------
-- wallet sql dml
--------------------------------------------------------------------------------

-- name: GetUserWalletBills :one
SELECT * 
FROM p_wallet_statement 
WHERE user_id=$1 AND is_del=0 
ORDER BY id DESC 
LIMIT $2 OFFSET $3;
