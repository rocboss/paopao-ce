DROP TABLE IF EXISTS p_user;
DROP INDEX IF EXISTS p_user_username_idx, p_user_phone_idx;

DROP TABLE IF EXISTS p_tag;
DROP INDEX IF EXISTS p_tag_tag_idx, p_tag_user_idx, p_tag_quote_num_idx;
