ALTER TABLE `p_user` MODIFY COLUMN `password` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '密码';

UPDATE
	p_user
SET
	password = CONCAT_WS(':', password, salt);

ALTER TABLE `p_user` DROP COLUMN `salt`;
