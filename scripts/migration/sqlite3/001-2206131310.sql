-- ----------------------------
-- Table p_post alter add visibility column
-- ----------------------------
ALTER TABLE `p_post` ADD COLUMN `visibility` integer NOT NULL DEFAULT '0';

-- ----------------------------
-- Indexes structure for table p_post
-- ----------------------------
CREATE INDEX "main"."idx_visibility"
ON "p_post" (
  "visibility" ASC
);
