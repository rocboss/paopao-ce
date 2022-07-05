ALTER TABLE `p_post` ADD COLUMN `visibility` integer NOT NULL DEFAULT '0';

CREATE INDEX "main"."idx_visibility"
ON "p_post" (
  "visibility" ASC
);
