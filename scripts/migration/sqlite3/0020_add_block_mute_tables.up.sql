-- 添加用户拉黑和屏蔽功能相关表

-- 用户拉黑表
CREATE TABLE IF NOT EXISTS `p_user_blocks` (
  `id` INTEGER PRIMARY KEY AUTOINCREMENT,
  `user_id` INTEGER NOT NULL,
  `target_user_id` INTEGER NOT NULL,
  `reason` TEXT DEFAULT '',
  `created_on` INTEGER NOT NULL DEFAULT 0,
  `modified_on` INTEGER NOT NULL DEFAULT 0,
  `is_del` INTEGER NOT NULL DEFAULT 0
);

-- 创建索引
CREATE UNIQUE INDEX IF NOT EXISTS `uk_user_blocks_user_target` ON `p_user_blocks` (`user_id`, `target_user_id`);
CREATE INDEX IF NOT EXISTS `idx_user_blocks_user_id` ON `p_user_blocks` (`user_id`);
CREATE INDEX IF NOT EXISTS `idx_user_blocks_target_user_id` ON `p_user_blocks` (`target_user_id`);
CREATE INDEX IF NOT EXISTS `idx_user_blocks_created_on` ON `p_user_blocks` (`created_on`);

-- 用户屏蔽表
CREATE TABLE IF NOT EXISTS `p_user_mutes` (
  `id` INTEGER PRIMARY KEY AUTOINCREMENT,
  `user_id` INTEGER NOT NULL,
  `target_user_id` INTEGER NOT NULL,
  `reason` TEXT DEFAULT '',
  `created_on` INTEGER NOT NULL DEFAULT 0,
  `modified_on` INTEGER NOT NULL DEFAULT 0,
  `is_del` INTEGER NOT NULL DEFAULT 0
);

-- 创建索引
CREATE UNIQUE INDEX IF NOT EXISTS `uk_user_mutes_user_target` ON `p_user_mutes` (`user_id`, `target_user_id`);
CREATE INDEX IF NOT EXISTS `idx_user_mutes_user_id` ON `p_user_mutes` (`user_id`);
CREATE INDEX IF NOT EXISTS `idx_user_mutes_target_user_id` ON `p_user_mutes` (`target_user_id`);
CREATE INDEX IF NOT EXISTS `idx_user_mutes_created_on` ON `p_user_mutes` (`created_on`);
