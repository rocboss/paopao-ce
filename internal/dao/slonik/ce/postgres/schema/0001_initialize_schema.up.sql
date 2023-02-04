CREATE TABLE p_tag (
  id bigserial PRIMARY KEY,
  user_id bigserial NOT NULL DEFAULT 0,
  tag varchar(255) NOT NULL,
  quote_num bigint NOT NULL DEFAULT 0,
  created_on int NOT NULL DEFAULT 0,
  modified_on int NOT NULL DEFAULT 0,
  deleted_on int NOT NULL DEFAULT 0,
  is_del boolean NOT NULL DEFAULT false,
  UNIQUE (tag)
);
COMMENT ON TABLE p_tag IS '主题标签';
COMMENT ON COLUMN p_tag.is_del IS '是否删除';
CREATE UNIQUE INDEX p_tag_tag_idx ON p_tag (tag);
CREATE INDEX p_tag_user_idx ON p_tag (user_id);
CREATE INDEX p_tag_quote_num_idx On p_tag (quote_num);

