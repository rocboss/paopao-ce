-- 删除用户签到记录表
DROP TABLE IF EXISTS "p_user_checkin";

-- 删除运营设置表
DROP TABLE IF EXISTS "p_operation_settings";

-- 删除用户表的经验值字段
ALTER TABLE "p_user" DROP COLUMN "experience";




