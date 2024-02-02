ALTER TABLE p_user ADD COLUMN salt text(32) NOT NULL DEFAULT '';
ALTER TABLE p_user ADD COLUMN password_copy text(64) NOT NULL DEFAULT '';

UPDATE
	p_user
SET
	salt = substr(password, instr(password, ':')+1),
	password_copy = substr(password, 1, instr(password, ':')-1);

ALTER TABLE p_user DROP COLUMN password;
ALTER TABLE p_user RERENAME COLUMN password_copy TO password;
