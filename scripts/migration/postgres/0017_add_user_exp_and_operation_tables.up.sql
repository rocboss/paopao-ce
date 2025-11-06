-- 为用户表添加经验值字段
ALTER TABLE "p_user" ADD COLUMN "experience" integer NOT NULL DEFAULT 0;

-- 创建运营设置表
CREATE TABLE IF NOT EXISTS "p_operation_settings" (
  "id" bigserial PRIMARY KEY,
  "key" varchar(255) NOT NULL UNIQUE,
  "value" text NOT NULL,
  "description" varchar(500),
  "created_on" integer NOT NULL DEFAULT 0,
  "updated_on" integer NOT NULL DEFAULT 0
);

-- 创建用户签到记录表
CREATE TABLE IF NOT EXISTS "p_user_checkin" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "checkin_date" date NOT NULL,
  "experience_gained" integer NOT NULL DEFAULT 0,
  "created_on" integer NOT NULL DEFAULT 0,
  "updated_on" integer NOT NULL DEFAULT 0,
  UNIQUE ("user_id", "checkin_date")
);

-- 创建索引
CREATE INDEX IF NOT EXISTS "idx_user_checkin_user_id" ON "p_user_checkin" ("user_id");
CREATE INDEX IF NOT EXISTS "idx_user_checkin_checkin_date" ON "p_user_checkin" ("checkin_date");

-- 插入默认运营设置
INSERT INTO "p_operation_settings" ("key", "value", "description", "created_on", "updated_on") VALUES
('checkin_enabled', 'true', '签到功能是否启用', extract(epoch from now())::integer, extract(epoch from now())::integer),
('checkin_exp_range_min', '5', '签到最小经验值', extract(epoch from now())::integer, extract(epoch from now())::integer),
('checkin_exp_range_max', '15', '签到最大经验值', extract(epoch from now())::integer, extract(epoch from now())::integer),
('checkin_rank_length', '10', '签到排行榜显示人数', extract(epoch from now())::integer, extract(epoch from now())::integer),
('level_to_exp', '{"lv1":[0,100],"lv2":[101,300],"lv3":[301,600],"lv4":[601,1000],"lv5":[1001,1500]}', '经验-等级对应关系', extract(epoch from now())::integer, extract(epoch from now())::integer);




