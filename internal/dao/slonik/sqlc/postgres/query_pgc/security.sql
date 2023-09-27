--------------------------------------------------------------------------------
-- security sql dml
--------------------------------------------------------------------------------

-- name: GetLatestPhoneCaptcha :one
SELECT * FROM p_captcha WHERE phone=$1 AND is_del=0;

-- name: UsePhoneCaptcha :exec
UPDATE p_captcha SET use_times=use_times+1, modified_on=$1 WHERE id=$2 AND is_del=0;

-- name: CreatePhoneCaptcha :one
INSERT INTO p_captcha (phone, captcha, expired_on, created_on) 
VALUES ($1, $2, $3, $4)
RETURNING id;
