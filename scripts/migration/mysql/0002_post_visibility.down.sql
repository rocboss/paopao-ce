ALTER TABLE `p_post` DROP COLUMN `visibility`;

DROP INDEX `idx_visibility` ON `p_post`;
