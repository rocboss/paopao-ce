-- 添加用户拉黑和屏蔽功能相关表

-- 用户拉黑表
CREATE TABLE IF NOT EXISTS "p_user_blocks" (
  "id" bigserial NOT NULL PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "target_user_id" bigint NOT NULL,
  "reason" varchar(500) DEFAULT '',
  "created_on" integer NOT NULL DEFAULT 0,
  "modified_on" integer NOT NULL DEFAULT 0,
  "is_del" integer NOT NULL DEFAULT 0
);

-- 创建索引
CREATE UNIQUE INDEX IF NOT EXISTS "uk_user_blocks_user_target" ON "p_user_blocks" ("user_id", "target_user_id");
CREATE INDEX IF NOT EXISTS "idx_user_blocks_user_id" ON "p_user_blocks" ("user_id");
CREATE INDEX IF NOT EXISTS "idx_user_blocks_target_user_id" ON "p_user_blocks" ("target_user_id");
CREATE INDEX IF NOT EXISTS "idx_user_blocks_created_on" ON "p_user_blocks" ("created_on");

-- 添加注释
COMMENT ON TABLE "p_user_blocks" IS '用户拉黑表';
COMMENT ON COLUMN "p_user_blocks"."user_id" IS '发起拉黑的用户ID';
COMMENT ON COLUMN "p_user_blocks"."target_user_id" IS '被拉黑的用户ID';
COMMENT ON COLUMN "p_user_blocks"."reason" IS '拉黑原因';

-- 用户屏蔽表
CREATE TABLE IF NOT EXISTS "p_user_mutes" (
  "id" bigserial NOT NULL PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "target_user_id" bigint NOT NULL,
  "reason" varchar(500) DEFAULT '',
  "created_on" integer NOT NULL DEFAULT 0,
  "modified_on" integer NOT NULL DEFAULT 0,
  "is_del" integer NOT NULL DEFAULT 0
);

-- 创建索引
CREATE UNIQUE INDEX IF NOT EXISTS "uk_user_mutes_user_target" ON "p_user_mutes" ("user_id", "target_user_id");
CREATE INDEX IF NOT EXISTS "idx_user_mutes_user_id" ON "p_user_mutes" ("user_id");
CREATE INDEX IF NOT EXISTS "idx_user_mutes_target_user_id" ON "p_user_mutes" ("target_user_id");
CREATE INDEX IF NOT EXISTS "idx_user_mutes_created_on" ON "p_user_mutes" ("created_on");

-- 添加注释
COMMENT ON TABLE "p_user_mutes" IS '用户屏蔽表';
COMMENT ON COLUMN "p_user_mutes"."user_id" IS '发起屏蔽的用户ID';
COMMENT ON COLUMN "p_user_mutes"."target_user_id" IS '被屏蔽的用户ID';
COMMENT ON COLUMN "p_user_mutes"."reason" IS '屏蔽原因';
