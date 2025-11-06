-- 重新创建用户等级表
CREATE TABLE IF NOT EXISTS `p_user_level` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `level` int(11) NOT NULL COMMENT '等级',
  `min_experience` int(11) NOT NULL COMMENT '最小经验值',
  `max_experience` int(11) NOT NULL COMMENT '最大经验值',
  `created_on` int(11) NOT NULL DEFAULT 0,
  `updated_on` int(11) NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_level` (`level`),
  UNIQUE KEY `idx_min_experience` (`min_experience`),
  UNIQUE KEY `idx_max_experience` (`max_experience`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户等级配置表';

-- 插入初始等级数据
INSERT INTO `p_user_level` (`level`, `min_experience`, `max_experience`, `created_on`, `updated_on`) VALUES
(1, 0, 100, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(2, 101, 300, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(3, 301, 600, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(4, 601, 1000, UNIX_TIMESTAMP(), UNIX_TIMESTAMP()),
(5, 1001, 1500, UNIX_TIMESTAMP(), UNIX_TIMESTAMP());


