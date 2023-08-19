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
