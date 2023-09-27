--------------------------------------------------------------------------------
-- tweet_metrics sql dml
--------------------------------------------------------------------------------

-- name: UpdateTweetMetric :exec
UPDATE p_post_metric SET rank_score=$1, modified_on=$2 WHERE post_id=$3 AND is_del=0;

-- name: TweetMotivationFactor :one
SELECT motivation_factor FROM p_post_metric WHERE post_id=$1 AND is_del=0;

-- name: AddTweetMetric :one
INSERT INTO p_post_metric (post_id, created_on) 
VALUES ($1, $2)
RETURNING id;

-- name: UpsertTweetMetric :exec
INSERT INTO p_post_metric (post_id, rank_score, created_on) VALUES ($1, $2, $3);

-- name: DeleteTweetMetric :exec
UPDATE p_post_metric SET is_del=1, deleted_on=$1 WHERE post_id=$2 AND is_del=0;

--------------------------------------------------------------------------------
-- comment_metrics sql dml
--------------------------------------------------------------------------------

-- name: UpdateCommentMetric :exec
UPDATE p_comment_metric SET rank_score=$1, modified_on=$2 WHERE comment_id=$3 AND is_del=0;

-- name: CommentMotivationFactor :one
SELECT motivation_factor FROM p_comment_metric WHERE comment_id=$1 AND is_del=0;

-- name: AddCommentMetric :one
INSERT INTO p_comment_metric (comment_id, created_on) 
VALUES ($1, $2)
RETURNING id;

-- name: UpsertCommentMetric :exec
INSERT INTO p_comment_metric (comment_id, rank_score, created_on) 
VALUES ($1, $2, $3)
RETURNING id;

-- name: DeleteCommentMetric :exec
UPDATE p_comment_metric SET is_del=1, deleted_on=$1 WHERE comment_id=$2 AND is_del=0;

--------------------------------------------------------------------------------
-- user_metrics sql dml
--------------------------------------------------------------------------------

-- name: UpdateUserMetric :exec
UPDATE p_user_metric SET tweets_count=$1, modified_on=$2 WHERE user_id=$3 AND is_del=0;

-- name: GetUserTweetsCount :one
SELECT tweets_count FROM p_user_metric WHERE user_id=$1 AND is_del=0;

-- name: AddUserMetric :one
INSERT INTO p_user_metric (user_id, created_on) VALUES ($1, $2) RETURNING id;

-- name: DeleteUserMetric :exec
UPDATE p_user_metric SET is_del=1, deleted_on=$1 WHERE user_id=$2 AND is_del=0;
