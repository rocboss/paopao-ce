ALTER TABLE `p_topic_user` DROP COLUMN `is_pin`;
DROP INDEX IF EXISTS `idx_topic_user_uid_ispin`;
