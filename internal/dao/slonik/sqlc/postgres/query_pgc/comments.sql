--------------------------------------------------------------------------------
-- comment sql dml
--------------------------------------------------------------------------------

-- name: GetNewestComments :many
SELECT * 
FROM p_comment 
WHERE post_id=$1 AND is_del=0 
ORDER BY is_essence DESC, id DESC 
LIMIT $2 OFFSET $3; 

-- name: GetHotsComments :many
SELECT c.* 
FROM p_comment c
LEFT JOIN p_comment_metric m
ON c.id=m.comment_id
WHERE c.post_id=$1 AND c.is_del=0 AND m.is_del=0
ORDER BY is_essence DESC, m.rank_score DESC, c.id DESC
LIMIT $2 OFFSET $3;

-- name: GetDefaultComments :many
SELECT * 
FROM p_comment 
WHERE post_id=$1 AND is_del=0 
ORDER BY is_essence DESC, id ASC 
LIMIT $2 OFFSET $3; 

-- name: GetCommentById :one
SELECT * FROM p_comment WHERE id=$1 AND is_del=0;

-- name: GetCommentCount :one
SELECT count(*) FROM p_comment WHERE post_id=$1 AND is_del=0;

-- name: GetCommentReplyById :one
SELECT * FROM p_comment_reply WHERE id=$1 AND is_del=0;

-- name: GetCommentContentsByIds :many
SELECT * FROM p_comment_content WHERE comment_id = ANY(@ids::BIGINT[]);

-- name: GetCommentRepliesByIds :many
SELECT * 
FROM p_comment_reply 
WHERE comment_id = ANY(@ids::BIGINT[])
ORDER BY id ASC;

-- name: GetCommentThumbs :many
SELECT user_id, 
    tweet_id, 
    comment_id, 
    reply_id, 
    comment_type, 
    is_thumbs_up, 
    is_thumbs_down
FROM p_tweet_comment_thumbs
WHERE user_id=$1 AND tweet_id=$2;

--------------------------------------------------------------------------------
-- comment_manage sql dml
--------------------------------------------------------------------------------

-- name: DeleteComment :exec
UPDATE p_comment SET deleted_on=$1, is_del=1 WHERE id=$2 AND is_del=0;

-- name: DeleteCommentThumbs :exec
UPDATE p_tweet_comment_thumbs
SET deleted_on=$1, is_del=1
WHERE user_id=$2 AND tweet_id=$3 AND comment_id=$4 AND is_del=0;

-- name: CreateComment :one
INSERT INTO p_comment (post_id, user_id, ip, ip_loc, created_on)
VALUES ($1, $2, $3, $4, $5)
RETURNING id;

-- name: CreateCommentReply :one
INSERT INTO p_comment_reply (comment_id, user_id, content, at_user_id, ip, ip_loc, created_on)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING id;

-- name: HighlightComment :one
UPDATE p_comment
SET is_essence=1-is_essence,
    modified_on=$1
WHERE id=$2 AND user_id=$3 AND is_del=0
RETURNING is_essence;

-- name: IncrCommentReplyCount :exec
UPDATE p_comment
SET reply_count=reply_count+1,
    modified_on=$1
WHERE id=$2 AND is_del=0;

-- name: DecrCommentReplyCount :exec
UPDATE p_comment
SET reply_count=reply_count-1,
    modified_on=$1
WHERE id=$2 AND is_del=0;

-- name: DeleteCommentReply :exec
UPDATE p_comment_reply SET deleted_on=$1, is_del=1 WHERE id=$2 AND is_del=0;

-- name: DeleteReplyThumbs :exec
UPDATE p_tweet_comment_thumbs
SET deleted_on=$1, is_del=1
WHERE user_id=$2 AND comment_id=$3 AND reply_id=$4 AND is_del=0;

-- name: CreateCommentContent :one
INSERT INTO p_comment_content (comment_id, user_id, content, type, sort, created_on)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id;

-- name: UpdateThumbsUpDownComment :exec
UPDATE p_tweet_comment_thumbs
SET is_thumbs_up=$1,
    is_thumbs_down=$2,
    modified_on=$3
WHERE id=$4 AND is_del=0;

-- name: CreateThumbsUpDownComment :one
INSERT INTO p_tweet_comment_thumbs (user_id, tweet_id, comment_id, 
    reply_id, is_thumbs_up, is_thumbs_down, comment_type, created_on)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING id;

-- name: UpdateCommentThumbsCount :exec
UPDATE p_comment
SET thumbs_up_count=$1, Thumbs_down_count=$2, modified_on=$3
WHERE id=$4 AND is_del=0;

-- name: GetTweetCommentThumb :one
SELECT *
FROM p_tweet_comment_thumbs
WHERE user_id=$1 AND tweet_id=$2 AND comment_id=$3 AND comment_type=0 AND is_del=0;

-- name: GetCommentReplyThumb :one
SELECT *
FROM p_tweet_comment_thumbs
WHERE user_id=$1 AND tweet_id=$2 AND comment_id=$3 AND reply_id=$4 AND comment_type=1 AND is_del=0;

-- name: UpdateReplyThumbsCount :exec
UPDATE p_comment_reply
SET thumbs_up_count=$1, thumbs_down_count=$2, modified_on=$3
WHERE id=$4 AND is_del=0;
