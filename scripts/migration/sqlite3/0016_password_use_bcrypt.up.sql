ALTER TABLE p_user ADD COLUMN password_copy text(255) NOT NULL DEFAULT '';
UPDATE p_user SET password_copy = concat_ws(':', password, salt);
ALTER TABLE p_user DROP COLUMN password;
ALTER TABLE p_user RERENAME COLUMN password_copy TO password;
ALTER TABLE p_user DROP COLUMN salt;
