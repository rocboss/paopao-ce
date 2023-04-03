-- name: NewestTags :many
SELECT t.id, t.tag, t.quote_num, u.id user_id, u.nickname, u.username, u.status, u.avatar, u.is_admin 
FROM p_tag t JOIN p_user u ON t.user_id = u.id
WHERE t.is_del = false AND t.quote_num > 0 
ORDER BY t.id DESC
OFFSET $1 LIMIT $2;

-- name: HotTags :many
SELECT  t.id, t.tag, t.quote_num, u.id user_id, u.nickname, u.username, u.status, u.avatar, u.is_admin 
FROM p_tag t JOIN p_user u ON t.user_id = u.id
WHERE  t.is_del = false AND t.quote_num > 0
ORDER BY quote_num DESC
OFFSET $1 LIMIT $2;

-- name: TagsByKeywordA :many
SELECT id, user_id, tag, quote_num
FROM p_tag
WHERE is_del AND quote_num > 0
ORDER BY quote_num DESC
OFFSET 0 LIMIT 6;

-- name: TagsByKeywordB :many
SELECT id, user_id, tag, quote_num
FROM p_tag
WHERE is_del = false AND tag LIKE $1
ORDER BY quote_num DESC
OFFSET 0 LIMIT 6;

-- name: InsertTags :one
INSERT INTO p_tag (user_id, tag, created_on, modified_on, quote_num)
VALUES ($1, $2, $3, $3, 1)
RETURNING id;

-- name: DecrTagsById :exec
UPDATE p_tag
SET quote_num = quote_num-1,
	modified_on=$1
WHERE id IN (
	SELECT id
	FROM p_tag
	WHERE id = ANY(@ids::BIGINT[]) AND is_del = false AND quote_num >= 1
);

-- name: IncrTags :many
UPDATE p_tag
SET quote_num = quote_num+1, 
	modified_on = $1,
	id_del = false
WHERE id IN (
	SELECT id
	FROM p_tag
	WHERE tag = ANY(@tags::VARCHAR[])
)
RETURNING id, user_id, tag, quote_num;
