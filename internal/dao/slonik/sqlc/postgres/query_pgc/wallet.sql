--------------------------------------------------------------------------------
-- wallet sql dml
--------------------------------------------------------------------------------

-- name: GetUserWalletBills :many
SELECT * 
FROM p_wallet_statement 
WHERE user_id=$1 AND is_del=0 
ORDER BY id DESC 
LIMIT $2 OFFSET $3;

-- name: CountUserWalletBill :one
SELECT count(*) FROM p_wallet_statement WHERE user_id=$1 AND is_del=0;

-- name: GetRechargeById :one
SELECT * FROM p_wallet_recharge WHERE id=$1 AND is_del=0;

-- name: CreateRecharge :one
INSERT INTO p_wallet_recharge (user_id, amount, created_on) 
VALUES ($1, $2, $3)
RETURNING id;

-- name: AddUserBalance :one
UPDATE p_user 
SET balance=balance+$1, modified_on=$2 
WHERE id=$3 AND is_del=0
RETURNING balance;

-- name: MinusUserBalance :one
UPDATE p_user 
SET balance=balance-$1, modified_on=$2 
WHERE id=$3 AND is_del=0
RETURNING balance;

-- name: CreateWalletStatement :one
INSERT INTO p_wallet_statement (user_id, change_amount, balance_snapshot, reason, created_on)
VALUES ($1, $2, $3, $4, $5)
RETURNING id;

-- name: GetUserBalance :one
SELECT balance FROM p_user WHERE id=$1 AND is_del=0;

-- name: MarkSuccessRecharge :exec
UPDATE p_wallet_recharge 
SET trade_no=$1, trade_status='TRADE_SUCCESS', modified_on=$2
WHERE id=$3 AND is_del=0;

-- name: CreatePostAttachmentBill :one
INSERT INTO p_post_attachment_bill (post_id, user_id, paid_amount, created_on)
VALUES ($1, $2, $3, $4)
RETURNING id;

-- name: CreatePostBill :one
INSERT INTO p_wallet_statement (post_id, user_id, change_amount, balance_snapshot, reason, created_on) 
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id;
