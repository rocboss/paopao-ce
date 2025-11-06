-- 添加用户拉黑和屏蔽功能相关表

-- 用户拉黑表
CREATE TABLE IF NOT EXISTS `p_user_blocks` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) NOT NULL COMMENT '发起拉黑的用户ID',
  `target_user_id` bigint(20) NOT NULL COMMENT '被拉黑的用户ID',
  `reason` varchar(500) DEFAULT '' COMMENT '拉黑原因',
  `created_on` int(11) NOT NULL DEFAULT 0,
  `modified_on` int(11) NOT NULL DEFAULT 0,
  `is_del` int(11) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_user_target` (`user_id`, `target_user_id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_target_user_id` (`target_user_id`),
  KEY `idx_created_on` (`created_on`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户拉黑表';

-- 用户屏蔽表
CREATE TABLE IF NOT EXISTS `p_user_mutes` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) NOT NULL COMMENT '发起屏蔽的用户ID',
  `target_user_id` bigint(20) NOT NULL COMMENT '被屏蔽的用户ID',
  `reason` varchar(500) DEFAULT '' COMMENT '屏蔽原因',
  `created_on` int(11) NOT NULL DEFAULT 0,
  `modified_on` int(11) NOT NULL DEFAULT 0,
  `is_del` int(11) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_user_target` (`user_id`, `target_user_id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_target_user_id` (`target_user_id`),
  KEY `idx_created_on` (`created_on`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户屏蔽表';
