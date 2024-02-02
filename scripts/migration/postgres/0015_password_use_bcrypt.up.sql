ALTER TABLE p_user 
ALTER COLUMN password TYPE VARCHAR(255) 
ALTER COLUMN password SET NOT NULL
ALTER COLUMN password SET DEFAULT '';

UPDATE
	p_user
SET
	password = concat_ws(':', password, salt);

ALTER TABLE p_user DROP COLUMN salt;
