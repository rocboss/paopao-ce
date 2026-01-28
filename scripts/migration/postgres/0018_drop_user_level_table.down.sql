-- 重新创建用户等级表
CREATE TABLE IF NOT EXISTS "p_user_level" (
  "id" bigserial PRIMARY KEY,
  "level" integer NOT NULL,
  "min_experience" integer NOT NULL,
  "max_experience" integer NOT NULL,
  "created_on" integer NOT NULL DEFAULT 0,
  "updated_on" integer NOT NULL DEFAULT 0
);

-- 创建索引
CREATE UNIQUE INDEX IF NOT EXISTS "idx_user_level_level" ON "p_user_level" ("level");
CREATE UNIQUE INDEX IF NOT EXISTS "idx_user_level_min_experience" ON "p_user_level" ("min_experience");
CREATE UNIQUE INDEX IF NOT EXISTS "idx_user_level_max_experience" ON "p_user_level" ("max_experience");

-- 插入初始等级数据
INSERT INTO "p_user_level" ("level", "min_experience", "max_experience", "created_on", "updated_on") VALUES
(1, 0, 100, extract(epoch from now())::integer, extract(epoch from now())::integer),
(2, 101, 300, extract(epoch from now())::integer, extract(epoch from now())::integer),
(3, 301, 600, extract(epoch from now())::integer, extract(epoch from now())::integer),
(4, 601, 1000, extract(epoch from now())::integer, extract(epoch from now())::integer),
(5, 1001, 1500, extract(epoch from now())::integer, extract(epoch from now())::integer);


