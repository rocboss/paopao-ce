ALTER TABLE p_post_content ALTER COLUMN content SET DATA TYPE TEXT NOT NULL DEFAULT '';
ALTER TABLE p_comment_content ALTER COLUMN content SET DATA TYPE TEXT NOT NULL DEFAULT '';
ALTER TABLE p_comment_reply ALTER COLUMN content SET DATA TYPE TEXT NOT NULL DEFAULT '';
