ALTER TABLE `p_post` DROP COLUMN `visibility`;

DROP INDEX "main"."idx_visibility" ON "p_post"
