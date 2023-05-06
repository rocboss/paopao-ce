CREATE TABLE "p_topic_user" (
	"id" integer,
	"topic_id" integer NOT NULL,-- 标签ID
	"user_id" integer NOT NULL,-- 创建者ID
	"alias_name" text ( 255 ),-- 别名
	"remark" text ( 512 ),-- 备注
	"quote_num" integer,-- 引用数
	"is_top" integer NOT NULL DEFAULT 0,-- 是否置顶 0 为未置顶、1 为已置顶
	"created_on" integer NOT NULL DEFAULT 0,-- 创建时间
	"modified_on" integer NOT NULL DEFAULT 0,-- 修改时间
	"deleted_on" integer NOT NULL DEFAULT 0,-- 删除时间
	"is_del" integer NOT NULL DEFAULT 0,-- 是否删除 0 为未删除、1 为已删除
	"reserve_a" text,-- 保留字段a
	"reserve_b" text,-- 保留字段b
	PRIMARY KEY ( "id" ) 
);
CREATE UNIQUE INDEX "idx_topic_user_uid_tid" ON "p_topic_user" ( "topic_id", "user_id" );