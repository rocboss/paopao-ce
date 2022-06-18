-- ----------------------------
-- Table p_post alter add visibility column
-- ----------------------------
ALTER TABLE `p_post` ADD COLUMN `visibility` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '可见性 0公开 1私密 2好友可见';

-- ----------------------------
-- Indexes structure for table p_post
-- ----------------------------
CREATE INDEX `idx_visibility` ON `p_post` ( `visibility` ) USING BTREE;
