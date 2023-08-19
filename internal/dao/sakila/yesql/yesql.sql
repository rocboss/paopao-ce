--------------------------------------------------------------------------------
-- file yesql.sql
-- this sql file just define global sql dml
--------------------------------------------------------------------------------

--------------------------------------------------------------------------------
-- global sql dml 
--------------------------------------------------------------------------------

-- name: tags_from_names
-- prepare: raw
-- clause: in
SELECT * FROM @tag WHERE tag IN (?) AND is_del=0;

-- name: update_tag_quote
-- prepare: stmt
UPDATE @tag SET quote_num=?, modified_on=?
WHERE id=? AND is_del=0;

-- name: tags_for_incr
-- prepare: raw
-- clause: in
SELECT id, user_id, tag, quote_num FROM @tag WHERE tag IN (?);

-- name: incr_tags_by_id
-- prepare: raw
-- clause: in
UPDATE @tag SET quote_num=quote_num+1, is_del=0, modified_on=? WHERE id IN (?);

-- name: insert_tag
-- prepare: stmt
INSERT INTO @tag (user_id, tag, created_on, modified_on, quote_num)
VALUES (?, ?, ?, ?, 1);

-- name: tags_by_id_b
-- prepare: raw
-- clause: in
SELECT id, user_id, tag, quote_num FROM @tag WHERE id IN (?);
