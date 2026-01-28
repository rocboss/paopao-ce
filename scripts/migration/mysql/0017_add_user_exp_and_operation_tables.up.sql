-- 为用户表添加经验值字段
ALTER TABLE `p_user` ADD COLUMN `experience` int(11) NOT NULL DEFAULT 0 COMMENT '用户经验值' AFTER `balance`;

-- 创建运营设置表
CREATE TABLE IF NOT EXISTS `p_operation_settings` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `key` varchar(255) NOT NULL COMMENT '设置键',
  `value` text NOT NULL COMMENT '设置值',
  `description` varchar(500) DEFAULT NULL COMMENT '设置描述',
  `created_on` int(11) NOT NULL DEFAULT 0,
  `updated_on` int(11) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_key` (`key`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='运营设置表';

-- 创建用户签到记录表
CREATE TABLE IF NOT EXISTS `p_user_checkin` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) NOT NULL COMMENT '用户ID',
  `checkin_date` date NOT NULL COMMENT '签到日期',
  `experience_gained` int(11) NOT NULL DEFAULT 0 COMMENT '获得的经验值',
  `created_on` int(11) NOT NULL DEFAULT 0,
  `updated_on` int(11) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_user_date` (`user_id`, `checkin_date`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_checkin_date` (`checkin_date`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户签到记录表';

-- 插入默认运营设置
INSERT INTO `p_operation_settings` (`key`, `value`, `description`, `created_on`, `updated_on`) VALUES
('checkin_enabled', 'true', '签到功能是否启用', UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
('checkin_exp_range_min', '5', '签到最小经验值', UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
('checkin_exp_range_max', '15', '签到最大经验值', UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
('checkin_rank_length', '10', '签到排行榜显示人数', UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
('level_to_exp', '{"lv1":[0,100],"lv2":[101,300],"lv3":[301,600],"lv4":[601,1000],"lv5":[1001,1500]}', '经验-等级对应关系', UNIX_TIMESTAMP(), UNIX_TIMESTAMP());




