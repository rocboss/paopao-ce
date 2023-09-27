--------------------------------------------------------------------------------
-- tweet sql dml
--------------------------------------------------------------------------------

-- name: GetPostById :one
SELECT * FROM p_post WHERE id=$1 AND is_del=0;

-- name: GetUserPosts :many
SELECT * FROM p_post
WHERE user_id=$1 AND visibility IN ($1) AND is_del=0
ORDER BY latest_replied_on DESC
LIMIT $2 OFFSET $3;

-- name: GetAnyPosts :many
SELECT * FROM p_post WHERE visibility IN ($1) AND is_del=0 LIMIT $2 OFFSET $3;

-- name: CountUserPosts :one
SELECT count(*) FROM p_post WHERE user_id=$1 AND visibility IN ($2) AND is_del=0;

-- name: CountAnyPost :one
SELECT count(*) FROM p_post WHERE visibility IN ($1) AND is_del=0;

-- name: GetUserPostStar :one
SELECT
	s.*,
	P.ID "post.id",
	P.user_id "post.user_id",
	P.comment_count "post.comment_count",
	P.collection_count "post.collection_count",
	P.upvote_count "post.upvote_count",
	P.share_count "post.share_count",
	P.visibility "post.visibility",
	P.is_top "post.is_top",
	P.is_essence "post.is_essence",
	P.is_lock "post.is_lock",
	P.latest_replied_on "post.latest_replied_on",
	P.tags "post.tags",
	P.attachment_price "post.attachment_price",
	P.ip "post.ip",
	P.ip_loc "post.ip_loc",
	P.is_del "post.is_del",
	P.created_on "post.created_on",
	P.modified_on "post.modified_on",
	P.deleted_on "post.deleted_on" 
FROM
	p_post_star s
	JOIN p_post P ON s.post_id = P.ID 
WHERE
	s.post_id = $1
	AND s.user_id = $2 
	AND s.is_del = 0 
	AND (visibility >= 50 OR (visibility = 0 AND P.user_id = $3 ));

-- name: GetUserPostStars :many
SELECT
	s.*,
	P.ID "post.id",
	P.user_id "post.user_id",
	P.comment_count "post.comment_count",
	P.collection_count "post.collection_count",
	P.upvote_count "post.upvote_count",
	P.share_count "post.share_count",
	P.visibility "post.visibility",
	P.is_top "post.is_top",
	P.is_essence "post.is_essence",
	P.is_lock "post.is_lock",
	P.latest_replied_on "post.latest_replied_on",
	P.tags "post.tags",
	P.attachment_price "post.attachment_price",
	P.ip "post.ip",
	P.ip_loc "post.ip_loc",
	P.is_del "post.is_del",
	P.created_on "post.created_on",
	P.modified_on "post.modified_on",
	P.deleted_on "post.deleted_on" 
FROM
	p_post_star s
	JOIN p_post P ON s.post_id = P.ID 
WHERE
	s.user_id = $1 
	AND s.is_del = 0 
	AND (visibility >= 50 OR (visibility = 0 AND P.user_id = $2)) 
ORDER BY
    s.ID DESC,
	P.ID DESC
LIMIT $3 OFFSET $4;

-- name: CountUserPostStars :one
SELECT
    count(*) 
FROM
	p_post_star s
	JOIN p_post P ON s.post_id = P.ID 
WHERE
	s.user_id = $1 
	AND s.is_del = 0 
	AND (visibility >= 50 OR (visibility = 0 AND P.user_id = $2));

-- name: GetUserPostCollection :one
SELECT
	s.*,
	P.ID "post.id",
	P.user_id "post.user_id",
	P.comment_count "post.comment_count",
	P.collection_count "post.collection_count",
	P.upvote_count "post.upvote_count",
	P.share_count "post.share_count",
	P.visibility "post.visibility",
	P.is_top "post.is_top",
	P.is_essence "post.is_essence",
	P.is_lock "post.is_lock",
	P.latest_replied_on "post.latest_replied_on",
	P.tags "post.tags",
	P.attachment_price "post.attachment_price",
	P.ip "post.ip",
	P.ip_loc "post.ip_loc",
	P.is_del "post.is_del",
	P.created_on "post.created_on",
	P.modified_on "post.modified_on",
	P.deleted_on "post.deleted_on" 
FROM
	p_post_collection s
	JOIN p_post P ON s.post_id = P.ID 
WHERE
	s.post_id = $1
	AND s.user_id = $2 
	AND s.is_del = 0 
	AND (visibility >= 50 OR (visibility = 0 AND P.user_id = $3));

-- name: GetUserPostCollections :many
SELECT
	s.*,
	P.ID "post.id",
	P.user_id "post.user_id",
	P.comment_count "post.comment_count",
	P.collection_count "post.collection_count",
	P.upvote_count "post.upvote_count",
	P.share_count "post.share_count",
	P.visibility "post.visibility",
	P.is_top "post.is_top",
	P.is_essence "post.is_essence",
	P.is_lock "post.is_lock",
	P.latest_replied_on "post.latest_replied_on",
	P.tags "post.tags",
	P.attachment_price "post.attachment_price",
	P.ip "post.ip",
	P.ip_loc "post.ip_loc",
	P.is_del "post.is_del",
	P.created_on "post.created_on",
	P.modified_on "post.modified_on",
	P.deleted_on "post.deleted_on" 
FROM
	p_post_collection s
	JOIN p_post P ON s.post_id = P.ID 
WHERE
	s.user_id = $1 
	AND s.is_del = 0 
	AND (visibility >= 50 OR ( visibility = 0 AND P.user_id = $2 ) ) 
ORDER BY
    s.ID DESC,
	P.ID DESC
LIMIT $3 OFFSET $4;

-- CountGetUserPostCollection :one
SELECT
    count(*) 
FROM
	p_post_collection s
	JOIN p_post P ON s.post_id = P.ID 
WHERE
	s.user_id = $1 
	AND s.is_del = 0 
	AND (visibility >= 50 OR (visibility = 0 AND P.user_id = $2));

-- name: GetPostAttachementBill :one
SELECT * FROM p_post_attachment_bill WHERE post_id=$1 AND user_id=$2 AND is_del=0;

-- name: GetPostContentsByIds :many
SELECT * 
FROM p_post_content 
WHERE post_id = ANY(@ids::BIGINT[]) AND is_del=0;

-- name: GetPostContentById :one
SELECT * FROM p_post_content WHERE id=$1 AND is_del=0;

-- name: UserMediaTweetsByGuest :many
SELECT id,
	user_id,
	comment_count,
	collection_count,
	upvote_count,
	share_count,
	visibility,
	is_top,
	is_essence,
	is_lock,
	latest_replied_on,
	tags,
	attachment_price,
	ip,
	ip_loc,
	created_on,
	modified_on,
	deleted_on,
	is_del
FROM p_post_by_media
WHERE is_del=0 AND user_id=$1 AND visibility=90
ORDER BY latest_replied_on DESC
LIMIT $2 OFFSET $3;

-- name: CountUserMediaTweetsByGuest :one
SELECT count(*)
FROM p_post_by_media
WHERE is_del=0 AND user_id=$1 AND visibility>=90;

-- name: UserMediaTweetsByFriend :many
SELECT id,
	user_id,
	comment_count,
	collection_count,
	upvote_count,
	share_count,
	visibility,
	is_top,
	is_essence,
	is_lock,
	latest_replied_on,
	tags,
	attachment_price,
	ip,
	ip_loc,
	created_on,
	modified_on,
	deleted_on,
	is_del
FROM p_post_by_media
WHERE is_del=0 AND user_id=$1 AND visibility>=50
ORDER BY latest_replied_on DESC
LIMIT $2 OFFSET $3;

-- name: CountUserMediaTweetsByFriend :one
SELECT count(*)
FROM p_post_by_media
WHERE is_del=0 AND user_id=$1 AND visibility>=50;

-- name: UserMediaTweetsBySelf :many
SELECT id,
	user_id,
	comment_count,
	collection_count,
	upvote_count,
	share_count,
	visibility,
	is_top,
	is_essence,
	is_lock,
	latest_replied_on,
	tags,
	attachment_price,
	ip,
	ip_loc,
	created_on,
	modified_on,
	deleted_on,
	is_del
FROM p_post_by_media
WHERE is_del=0 AND user_id=$1
ORDER BY latest_replied_on DESC
LIMIT $2 OFFSET $3;

-- name: CountUserMediaTweetsBySelf :one
SELECT count(*)
FROM p_post_by_media
WHERE is_del=0 AND user_id=$1;

-- name: UserCommentTweetsByGuest :many
SELECT id,
	user_id,
	comment_count,
	collection_count,
	upvote_count,
	share_count,
	visibility,
	is_top,
	is_essence,
	is_lock,
	latest_replied_on,
	tags,
	attachment_price,
	ip,
	ip_loc,
	created_on,
	modified_on,
	deleted_on,
	is_del
FROM p_post_by_comment
WHERE is_del=0 AND comment_user_id=$1 AND visibility>=90
ORDER BY latest_replied_on DESC
LIMIT $2 OFFSET $3;

-- name: CountUserCommentTweetsByGuest :one
SELECT count(*)
FROM p_post_by_comment
WHERE is_del=0 AND comment_user_id=$1 AND visibility>=90;

-- name: UserCommentTweetsByFriend :many
SELECT id,
	user_id,
	comment_count,
	collection_count,
	upvote_count,
	share_count,
	visibility,
	is_top,
	is_essence,
	is_lock,
	latest_replied_on,
	tags,
	attachment_price,
	ip,
	ip_loc,
	created_on,
	modified_on,
	deleted_on,
	is_del
FROM p_post_by_comment
WHERE is_del=0 AND comment_user_id=$1 AND visibility>=50
ORDER BY latest_replied_on DESC
LIMIT $2 OFFSET $3;

-- name: CountUserCommentTweetsByFriend :one
SELECT count(*)
FROM p_post_by_comment
WHERE is_del=0 AND comment_user_id=$1 AND visibility>=50;

-- name: UserCommentTweetsBySelf :many
SELECT id,
	user_id,
	comment_count,
	collection_count,
	upvote_count,
	share_count,
	visibility,
	is_top,
	is_essence,
	is_lock,
	latest_replied_on,
	tags,
	attachment_price,
	ip,
	ip_loc,
	created_on,
	modified_on,
	deleted_on,
	is_del
FROM p_post_by_comment
WHERE is_del=0 AND comment_user_id=$1
ORDER BY latest_replied_on DESC
LIMIT $2 OFFSET $3;

-- name: CountUserCommentTweetsBySelf :one
SELECT count(*)
FROM p_post_by_comment
WHERE is_del=0 AND comment_user_id=$1;

-- name: UserStarTweetsByGuest :many
SELECT
	star.*,
	post.ID "post.id",
	post.created_on "post.created_on",
	post.modified_on "post.modified_on",
	post.deleted_on "post.deleted_on",
	post.is_del "post.is_del",
	post.user_id "post.user_id",
	post.comment_count "post.comment_count",
	post.collection_count "post.collection_count",
	post.share_count "post.share_count",
	post.upvote_count "post.upvote_count",
	post.visibility "post.visibility",
	post.is_top "post.is_top",
	post.is_essence "post.is_essence",
	post.is_lock "post.is_lock",
	post.latest_replied_on "post.latest_replied_on",
	post.tags "post.tags",
	post.attachment_price "post.attachment_price",
	post.ip "post.ip",
	post.ip_loc "post.ip_loc" 
FROM
	p_post_star star
	JOIN p_post post ON star.post_id = post.ID 
WHERE
	star.is_del=0 
	AND star.user_id=$1 
	AND post.visibility>=90
ORDER BY post.latest_replied_on DESC
LIMIT $2 OFFSET $3;

-- name: CountUserStarTweetsByGuest :one
SELECT count(*)
FROM
	p_post_star star
	JOIN p_post post ON star.post_id = post.ID 
WHERE star.is_del=0 AND star.user_id=$1 AND post.visibility>=90;

-- name: UserStarTweetsByFriend :many
SELECT
	star.*,
	post.ID "post.id",
	post.created_on "post.created_on",
	post.modified_on "post.modified_on",
	post.deleted_on "post.deleted_on",
	post.is_del "post.is_del",
	post.user_id "post.user_id",
	post.comment_count "post.comment_count",
	post.collection_count "post.collection_count",
	post.share_count "post.share_count",
	post.upvote_count "post.upvote_count",
	post.visibility "post.visibility",
	post.is_top "post.is_top",
	post.is_essence "post.is_essence",
	post.is_lock "post.is_lock",
	post.latest_replied_on "post.latest_replied_on",
	post.tags "post.tags",
	post.attachment_price "post.attachment_price",
	post.ip "post.ip",
	post.ip_loc "post.ip_loc" 
FROM
	p_post_star star
	JOIN p_post post ON star.post_id = post.ID 
WHERE star.is_del=0 AND star.user_id=$1 AND post.visibility>=50
ORDER BY post.latest_replied_on DESC
LIMIT $2 OFFSET $3;

-- name: CountUserStarTweetsByFriend :one
SELECT count(*)
FROM
	p_post_star star
	JOIN p_post post ON star.post_id = post.ID 
WHERE star.is_del=0 AND star.user_id=$1 AND post.visibility>=50;

-- name: UserStarTweetsBySelf :many
SELECT
	star.*,
	post.ID "post.id",
	post.created_on "post.created_on",
	post.modified_on "post.modified_on",
	post.deleted_on "post.deleted_on",
	post.is_del "post.is_del",
	post.user_id "post.user_id",
	post.comment_count "post.comment_count",
	post.collection_count "post.collection_count",
	post.share_count "post.share_count",
	post.upvote_count "post.upvote_count",
	post.visibility "post.visibility",
	post.is_top "post.is_top",
	post.is_essence "post.is_essence",
	post.is_lock "post.is_lock",
	post.latest_replied_on "post.latest_replied_on",
	post.tags "post.tags",
	post.attachment_price "post.attachment_price",
	post.ip "post.ip",
	post.ip_loc "post.ip_loc" 
FROM
	p_post_star star
	JOIN p_post post ON star.post_id = post.ID 
WHERE star.is_del=0 AND star.user_id=$1 AND (post.visibility<>90 OR (post.visibility>=90 AND post.user_id=$2))
ORDER BY post.latest_replied_on DESC
LIMIT $3 OFFSET $4;

-- name: CountUserStarTweetsBySelf :one
SELECT count(*)
FROM
	p_post_star star
	JOIN p_post post ON star.post_id = post.ID 
WHERE star.is_del=0 AND star.user_id=$1 AND (post.visibility<>90 OR (post.visibility>=90 AND post.user_id=$2));

-- name: UserStarTweetsByAdmin :many
SELECT
	star.*,
	post.ID "post.id",
	post.created_on "post.created_on",
	post.modified_on "post.modified_on",
	post.deleted_on "post.deleted_on",
	post.is_del "post.is_del",
	post.user_id "post.user_id",
	post.comment_count "post.comment_count",
	post.collection_count "post.collection_count",
	post.share_count "post.share_count",
	post.upvote_count "post.upvote_count",
	post.visibility "post.visibility",
	post.is_top "post.is_top",
	post.is_essence "post.is_essence",
	post.is_lock "post.is_lock",
	post.latest_replied_on "post.latest_replied_on",
	post.tags "post.tags",
	post.attachment_price "post.attachment_price",
	post.ip "post.ip",
	post.ip_loc "post.ip_loc" 
FROM
	p_post_star star
	JOIN p_post post ON star.post_id = post.ID 
WHERE star.is_del=0 AND star.user_id=$1
ORDER BY post.latest_replied_on DESC
LIMIT $2 OFFSET $3;

-- name: CountUserStarTweetsByAdmin :one
SELECT count(*)
FROM
	p_post_star star
	JOIN p_post post ON star.post_id = post.ID 
WHERE star.is_del=0 AND star.user_id=$1;

-- name: ListUserTweets :many
SELECT * 
FROM p_post
WHERE user_id=$1 AND visibility>=$2 AND is_essence=$3 AND is_del=0
ORDER BY is_top DESC, latest_replied_on DESC
LIMIT $4 OFFSET $5;

-- name: CountUserTweets :one
SELECT count(*) 
FROM p_post
WHERE user_id=$1 AND visibility>=$2 AND is_essence=$3 AND is_del=0;

-- name: ListIndexNewestTweets :many
SELECT *
FROM p_post
WHERE visibility>=90 AND is_del=0
ORDER BY is_top DESC, latest_replied_on DESC
LIMIT $1 OFFSET $2;

-- name: CountIndexNewestTweets :one
SELECT count(*)
FROM p_post
WHERE visibility>=90 AND is_del=0;

-- name: ListIndexHotsTweets :many
SELECT post.*
FROM p_post post
LEFT JOIN p_post_metric metric
ON post.id=metric.post_id
WHERE post.visibility>=90 AND post.is_del=0
ORDER BY post.is_top DESC, metric.rank_score DESC, post.latest_replied_on DESC
LIMIT $1 OFFSET $2;

-- name: CountIndexHotsTweets :one
SELECT count(*)
FROM p_post post
LEFT JOIN p_post_metric metric
ON post.id=metric.post_id AND metric.is_del=0
WHERE post.visibility>=90 AND post.is_del=0;

-- name: ListSyncSearchTweets :many
SELECT *
FROM p_post
WHERE visibility>=50 AND is_del=0
LIMIT $1 OFFSET $2;

-- name: CountSyncSearchTweets :one
SELECT count(*)
FROM p_post
WHERE visibility>=50 AND is_del=0;

-- name: ListFollowingTweetsFriendFollow :many
SELECT *
FROM p_post
WHERE (user_id=$1 OR (visibility>=50 AND user_id = ANY(@friendIds::BIGINT[])) OR (visibility>=60 AND user_id = ANY(@followIds::BIGINT[]))) AND is_del=0
ORDER BY is_top DESC, latest_replied_on DESC
LIMIT $2 OFFSET $3;

-- name: CountFollowingTweetsFriendFollow :one
SELECT count(*)
FROM p_post
WHERE (user_id=$1 OR (visibility>=50 AND user_id = ANY(@fiendIds::BIGINT[])) OR (visibility>=60 AND user_id = ANY(@followIds::BIGINT[]))) AND is_del=0;

-- name: ListFollowingTweetsFriend :many
SELECT *
FROM p_post
WHERE (user_id=$1 OR (visibility>=50 AND user_id = ANY(@friendIds::BIGINT[]))) AND is_del=0
ORDER BY is_top DESC, latest_replied_on DESC
LIMIT $2 OFFSET $3;

-- name: CountListFollowingTweetsFriend :one
SELECT count(*)
FROM p_post
WHERE (user_id=$1 OR (visibility>=50 AND user_id = ANY(@friendIds::BIGINT[]))) AND is_del=0;

-- name: ListFollowingTweetsFollow :many
SELECT *
FROM p_post
WHERE (user_id=$1 OR (visibility>=60 AND user_id = ANY(@followIds::BIGINT[]))) AND is_del=0
ORDER BY is_top DESC, latest_replied_on DESC
LIMIT $2 OFFSET $3;

-- name: CountFollowingTweetsFollow :one
SELECT count(*)
FROM p_post
WHERE (user_id=$1 OR (visibility>=60 AND user_id = ANY(@followIds::BIGINT[]))) AND is_del=0;

-- name: ListFollowingTweets :many
SELECT *
FROM p_post
WHERE user_id=$1 AND is_del=0
ORDER BY is_top DESC, latest_replied_on DESC
LIMIT $2 OFFSET $3;

-- name: CountFollowingTweets :one
SELECT count(*)
FROM p_post
WHERE user_id=$1 AND is_del=0;

-- name: GetBeFriendIds :many
SELECT user_id FROM p_contact WHERE friend_id=$1 AND is_del=0;

-- name: GetBeFollowIds :many
SELECT follow_id FROM p_following WHERE user_id=$1 AND is_del=0;

--------------------------------------------------------------------------------
-- tweet_manage sql dml
--------------------------------------------------------------------------------

-- name: AddPost :exec
INSERT INTO p_post (user_id, tags, ip, ip_loc, attachment_price, visibility, latest_replied_on, created_on)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING id;

-- name: MediaContentByPostId :many
SELECT content FROM p_post_content WHERE post_id=$1 AND is_del=0 AND type IN (3, 4, 5, 7, 8);

-- name: DeletePostById :exec
UPDATE p_post SET is_del=1, deleted_on=$1 WHERE id=$2 AND is_del=0;

-- name: LockPost :one
UPDATE p_post 
SET is_lock=1-is_lock, modified_on=$1 
WHERE id=$2 AND is_del=0 
RETURNING is_lock;

-- name: StickPost :one
UPDATE p_post 
SET is_top=1-is_top, modified_on=$1 
WHERE id=$2 AND is_del=0
RETURNING is_top;

-- name: VisiblePost :one
UPDATE p_post 
SET visibility=$1, is_top=$2, modified_on=$3 
WHERE id=$4 AND is_del=0
RETURNING visibility;

-- name: HighlightPost :one
UPDATE p_post
SET is_essence=1-is_essence 
WHERE id=$1 AND is_del=0
RETURNING is_essence;

-- name: PostHighlightStatus :one
SELECT user_id, is_essence 
FROM p_post 
WHERE id=$1 AND is_del=0;

-- name: UpdatePost :exec
UPDATE p_post SET comment_count=$1, 
    upvote_count=$2,
    collection_count=$3,
    latest_replied_on=$4,
    modified_on=$5
WHERE id=$6 AND is_del=0;

-- name: AddPostStar :one
INSERT INTO p_post_star (post_id, user_id, created_on)
VALUES ($1, $2, $3)
RETURNING id;

-- name: DeletePostStar :exec
UPDATE p_post_star 
SET is_del=1, deleted_on=$1 
WHERE id=$2 AND is_del=0;

-- name: AddPostCollection :one
INSERT INTO p_post_collection (post_id, user_id, created_on)
VALUES ($1, $2, $3)
RETURNING id;

-- name: DeletePostCollecton :exec
UPDATE p_post_collection SET is_del=1, deleted_on=$1 WHERE id=$2 AND is_del=0;

-- name: AddPostContent :one
INSERT INTO p_post_content (post_id, user_id, content, type, sort, created_on)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id;

-- name: AddAttachment :one
INSERT INTO p_attachment (user_id, file_size, img_width, img_height, type, content, created_on)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING id;

-- name: CommentIdsByPostId :many
SELECT id FROM p_comment WHERE post_id=$1 AND is_del=0;

-- name: CommentMediaFromCommentIds :many
SELECT content FROM p_comment_content WHERE comment_id = ANY(@ids::BIGINT[]) AND type=3 AND is_del=0;

-- name: DeleteCommentByPostId :exec
UPDATE p_comment SET deleted_on=$1, is_del=1 WHERE post_id=$2 AND is_del=0;

-- name: DeleteCommentContentByCommentIds :exec
UPDATE p_comment_content SET deleted_on=$1, is_del=1 WHERE comment_id = ANY(@ids::BIGINT[]) AND is_del=0;

-- name: DeleteReplyByCommentIds :exec
UPDATE p_comment_reply SET deleted_on=$1, is_del=1 WHERE comment_id = ANY(@ids::BIGINT[]) AND is_del=0;

--------------------------------------------------------------------------------
-- tweet_help sql dml
--------------------------------------------------------------------------------

-- name: GetPostConetentByIds :many
SELECT id, post_id, content, type, sort 
FROM p_post_content
WHERE post_id = ANY(@ids::BIGINT[]) AND is_del=0;
