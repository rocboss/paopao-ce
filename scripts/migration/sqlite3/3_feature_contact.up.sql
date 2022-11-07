PRAGMA foreign_keys = false;

CREATE TABLE "p_contact" (
  "id" integer NOT NULL,
  "user_id" integer NOT NULL,
  "friend_id" integer NOT NULL,
  "group_id" integer NOT NULL,
  "remark" text(32) NOT NULL,
  "status" integer NOT NULL,
   "notice_enable" integer NOT NULL,
  "is_top" integer NOT NULL,
  "is_black" integer NOT NULL,
  "is_del" integer NOT NULL,
  "created_on" integer NOT NULL,
  "modified_on" integer NOT NULL,
  "deleted_on" integer NOT NULL,
  PRIMARY KEY ("id")
);

CREATE TABLE "p_contact_group" (
  "id" integer NOT NULL,
  "user_id" integer NOT NULL,
  "name" text(32) NOT NULL,
  "is_del" integer NOT NULL,
  "created_on" integer NOT NULL,
  "modified_on" integer NOT NULL,
  "deleted_on" integer NOT NULL,
  PRIMARY KEY ("id")
);

CREATE UNIQUE INDEX "main"."idx_user_friend_id"
ON "p_contact" (
  "user_id" ASC,
  "friend_id" ASC
);

CREATE INDEX "main"."idx_user_friend_status"
ON "p_contact" (
  "user_id" ASC,
  "friend_id" ASC,
  "status" ASC
);

PRAGMA foreign_keys = true;
