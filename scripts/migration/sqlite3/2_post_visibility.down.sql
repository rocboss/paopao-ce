ALTER TABLE `p_post` DROP COLUMN `visibility` integer NOT NULL DEFAULT '0';

DROP INDEX "main"."idx_visibility" ON "p_post"
