ALTER TABLE `p_user` ADD COLUMN `salt` VARCHAR(32) NOT NULL DEFAULT '' COMMENT '盐值';

UPDATE
	`p_user`
SET
	`salt` = SUBSTRING_INDEX(`password`, ':', -1),
	`password` = SUBSTRING_INDEX(`password`, ':', 1);

ALTER TABLE `p_user` MODIFY COLUMN `password` VARCHAR(64) NOT NULL DEFAULT '' COMMENT '密码';
