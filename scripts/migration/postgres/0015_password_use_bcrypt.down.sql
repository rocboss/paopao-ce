ALTER TABLE p_user ADD COLUMN salt VARCHAR(32) NOT NULL DEFAULT '';

UPDATE
	p_user
SET
	salt = split_part(password, ':', -1),
	password = split_part(password, ':', 1);

ALTER TABLE p_user 
ALTER COLUMN password TYPE VARCHAR(64),
ALTER COLUMN password SET NOT NULL,
ALTER COLUMN password SET DEFAULT '';
