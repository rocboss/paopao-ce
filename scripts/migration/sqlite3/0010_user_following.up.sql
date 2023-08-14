CREATE TABLE "p_following" (
	"id" integer NOT NULL,
	"user_id" integer NOT NULL,
	"follow_id" integer NOT NULL,
	"is_del" integer NOT NULL,
	"created_on" integer NOT NULL,
	"modified_on" integer NOT NULL,
	"deleted_on" integer NOT NULL,
	PRIMARY KEY ("id")
);
CREATE INDEX "idx_following_user_follow"
ON "p_following" (
  "user_id" ASC,
  "follow_id" ASC
);
