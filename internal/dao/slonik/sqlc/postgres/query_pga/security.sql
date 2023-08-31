--------------------------------------------------------------------------------
-- security sql dml
--------------------------------------------------------------------------------

-- name: GetLatestPhoneCaptcha :one
SELECT * FROM p_captcha WHERE phone=$1 AND is_del=0;
