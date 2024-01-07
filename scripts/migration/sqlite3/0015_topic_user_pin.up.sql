ALTER TABLE "p_topic_user" ADD COLUMN "is_pin" integer NOT NULL DEFAULT 0;
CREATE INDEX "main"."idx_topic_user_uid_ispin"
ON "p_topic_user" (
  "user_id" ASC,
  "is_pin" ASC
);
