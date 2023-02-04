-- name: NewestTags :many
SELECT id, user_id, tag, quote_num
FROM p_tag
WHERE is_del = false AND quote_num > 0
ORDER BY id DESC
OFFSET $1 LIMIT $2;

-- name: HotTags :many
SELECT id, user_id, tag, quote_num
FROM p_tag
WHERE  is_del = false AND quote_num > 0
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
RETURNING id, user_id, tag, quote_num;

-- name: TagsByIdA :many
SELECT id
FROM p_tag
WHERE id = ANY(@ids::bigserial[]) AND is_del = false AND quote_num >= 0;

-- name: TagsByIdB :many
SELECT id, user_id, tag, quote_num
FROM p_tag
WHERE id = ANY(@ids::bigserial[]);

-- name: DecrTagsById :exec
UPDATE p_tag
SET quote_num = quote_num-1,
	modified_on=$1
WHERE id = ANY(@ids::bigserial[]);

-- name: TagsByName :many
SELECT id, user_id, tag, quote_num
FROM p_tag
WHERE tag = ANY(@tags::varchar[]) AND is_del = false AND quote_num >= 0;

-- name: IncrTagsById :exec
UPDATE p_tag
SET quote_num = quote_num+1, modified_on = $1
WHERE id = ANY(@ids::bigserial[]);
