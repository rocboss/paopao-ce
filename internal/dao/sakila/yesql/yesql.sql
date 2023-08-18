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

--------------------------------------------------------------------------------
-- authorization_manage sql dml
--------------------------------------------------------------------------------

-- name: be_friend_ids@authorization_manage
-- prepare: stmt
SELECT user_id FROM @contact WHERE friend_id=? AND status=2 AND is_del=0;

-- name: my_friend_set@authorization_manage
-- prepare: stmt
SELECT friend_id FROM @contact WHERE user_id=? AND status=2 AND is_del=0;

-- name: is_friend@authorization_manage
-- prepare: stmt
SELECT status FROM @contact WHERE user_id=? AND friend_id=? AND is_del=0;

--------------------------------------------------------------------------------
-- comment sql dml
--------------------------------------------------------------------------------

-- name: get_comments@comment
-- prepare: raw
SELECT * FROM @comment WHERE post_id=? AND is_del=0 ORDER BY %order% LIMIT ? OFFSET ?; 

-- name: get_comment_by_id@comment
-- prepare: stmt
SELECT * FROM @comment WHERE id=? AND is_del=0;

-- name: get_comment_count@comment
-- prepare: named_stmt
SELECT count(*) FROM @comment WHERE post_id=:post_id AND is_del=0;

-- name: get_comment_reply_by_id@comment
-- prepare: stmt
SELECT * FROM @comment_reply WHERE id=? AND is_del=0;

-- name: get_comment_contents_by_ids@comment
-- prepare: raw
-- clause: in
SELECT * FROM @comment_content WHERE comment_id IN (?);

-- name: get_commment_replies_by_ids@comment
-- prepare: raw
-- clause: in
SELECT * FROM @comment_reply WHERE comment_id IN (?) ORDER BY id ASC;

-- name: get_users_by_ids@comment
-- prepare: raw
-- clause: in
SELECT id, nickname, username, status, avatar, is_admin FROM @user WHERE id IN (?);

-- name: get_comment_thumbs@comment
-- prepare: stmt
SELECT * FROM @tweet_comment_thumbs WHERE user_id=? AND tweet_id=?;

--------------------------------------------------------------------------------
-- comment_manage sql dml
--------------------------------------------------------------------------------

-- name: delete_comment@comment_manage
-- prepare: stmt
UPDATE @comment SET deleted_on=?, is_del=1 WHERE id=? AND is_del=0;

-- name: delete_comment_thumbs@comment_manage
-- prepare: stmt
UPDATE @tweet_comment_thumbs 
SET deleted_on=?, is_del=1 
WHERE user_id=? AND tweet_id=? AND comment_id=? AND is_del=0;

-- name: create_comment@comment_manage
-- prepare: stmt
INSERT INTO @comment (post_id, user_id, ip, ip_loc, created_on) 
VALUES (?, ?, ?, ?, ?);

-- name: create_comment_reply@comment_manage
-- prepare: stmt
INSERT INTO @comment_reply (comment_id, user_id, content, at_user_id, ip, ip_loc, created_on) 
VALUES (?, ?, ?, ?, ?, ?, ?);

-- name: delete_comment_reply@comment_manage
-- prepare: stmt
UPDATE @comment_reply SET deleted_on=?, is_del=1 WHERE id=? AND is_del=0;

-- name: delete_reply_thumbs@comment_manage
-- prepare: stmt
UPDATE @tweet_comment_thumbs 
SET deleted_on=?, is_del=1 
WHERE user_id=? AND comment_id=? AND reply_id=? AND is_del=0;

-- name: create_comment_content@comment_manage
-- prepare: stmt
INSERT INTO @comment_content (comment_id, user_id, content, type, sort, created_on) 
VALUES (?, ?, ?, ?, ?, ?);

-- name: update_thumbs_updown_comment@comment_manage
-- prepare: named_stmt
UPDATE @tweet_comment_thumbs 
SET is_thumbs_up=:is_thumbs_up, is_thumbs_down=:is_thumbs_down, modified_on=:modified_on 
WHERE id=:id AND is_del=0;

-- name: create_thumbs_updown_comment@comment_manage
-- prepare: named_stmt
INSERT INTO @tweet_comment_thumbs (user_id, tweet_id, comment_id, reply_id, is_thumbs_up, is_thumbs_down, comment_type, created_on) 
VALUES (:user_id, :tweet_id, :comment_id, :reply_id, :is_thumbs_up, :is_thumbs_down, :comment_type, :created_on);

-- name: update_comment_thumbs_count@comment_manage
-- prepare: stmt
UPDATE @comment 
SET thumbs_up_count=?, thumbs_down_count=?, modified_on=? 
WHERE id=? AND is_del=0;

-- name: get_tweet_comment_thumb@comment_manage
-- prepare: stmt
SELECT * 
FROM @tweet_comment_thumbs 
WHERE user_id=? AND tweet_id=? AND comment_id=? AND comment_type=0 AND is_del=0;

-- name: get_comment_reply_thumb@comment_manage
-- prepare: stmt
SELECT * 
FROM @tweet_comment_thumbs 
WHERE user_id=? AND tweet_id=? AND comment_id=? AND reply_id=? AND comment_type=1 AND is_del=0;

-- name: update_reply_thumbs_count@comment_manage
-- prepare: stmt
UPDATE @comment_reply 
SET thumbs_up_count=?, thumbs_down_count=?, modified_on=? 
WHERE id=? AND is_del=0;

--------------------------------------------------------------------------------
-- following_manager sql dml
--------------------------------------------------------------------------------

-- name: create_following@following_manager
-- prepare: stmt
INSERT INTO @following (user_id, follow_id, created_on) VALUES (?, ?, ?);

-- name: exist_following@following_manager
-- prepare: stmt
SELECT 1 FROM @following WHERE user_id=? AND follow_id=? AND is_del=0;

-- name: delete_following@following_manager
-- prepare: stmt
UPDATE @following 
SET is_del=0, deleted_on=? 
WHERE user_id=? AND follow_id=? AND is_del=0;

-- name: list_follows@following_manager
-- prepare: stmt
SELECT u.user_id user_id, 
	u.username username,
	u.nickname nickname,
	u.avatar avatar,
	u.created_on created_on
FROM @following f JOIN @user u ON f.follow_id=u.id
WHERE f.user_id=? AND f.is_del=0
ORDER BY u.nickname ASC
LIMIT ? OFFSET ?;

-- name: count_follows@following_manager
-- prepare: stmt
SELECT count(*) FROM @following WHERE user_id=? AND is_del=0;

-- name: list_followings@following_manager
-- prepare: stmt
SELECT u.user_id user_id, 
	u.username username,
	u.nickname nickname,
	u.avatar avatar,
	u.created_on created_on
FROM @following f JOIN @user u ON f.user_id=u.id
WHERE f.follow_id=? AND f.is_del=0
ORDER BY u.nickname ASC
LIMIT ? OFFSET ?;

-- name: count_followings@following_manager
-- prepare: stmt
SELECT count(*) FROM @following WHERE follow_id=? AND is_del=0;

--------------------------------------------------------------------------------
-- contact_manager sql dml
--------------------------------------------------------------------------------

-- name: create_contact@contact_manager
-- prepare: stmt
INSERT INTO @contact (user_id, friend_id, status, created_on) VALUES (?, ?, ?, ?);

-- name: fresh_contact_status@contact_manager
-- prepare: stmt
UPDATE @contact SET status=?, modified_on=?, is_del=0 WHERE id=?;

-- name: create_message@contact_manager
-- prepare: named_stmt
INSERT INTO @message (sender_user_id, receiver_user_id, type, brief, content, reply_id, created_on) 
VALUES (:sender_user_id, :receiver_user_id, :type, :brief, :content, :reply_id, :created_on);

-- name: add_friend_msgs_update@contact_manager
-- prepare: stmt
UPDATE @message 
SET reply_id=?, modified_on=? 
WHERE ((sender_user_id = ? AND receiver_user_id = ?) OR (sender_user_id = ? AND receiver_user_id = ?)) AND type = ? AND reply_id = ?;

-- name: reject_friend_msgs_update@contact_manager
-- prepare: stmt
UPDATE @message 
SET reply_id=?, modified_on=? 
WHERE sender_user_id = ? AND receiver_user_id = ? AND type = ? AND reply_id = ?;

-- name: del_friend@contact_manager
-- prepare: stmt
UPDATE @contact SET status=4, is_del=1, deleted_on=? WHERE id=?;

-- name: list_friend@contact_manager
-- prepare: stmt
SELECT c.friend_id user_id, u.username username, u.nickname nickname, u.avatar avatar, u.phone phone 
FROM @contact c JOIN @user u 
ON c.friend_id=u.id 
WHERE user_id=? AND status=2 AND is_del=0 
ORDER BY u.nickname ASC 
LIMIT ? OFFSET ?;

-- name: total_friends_by_id@contact_manager
-- prepare: stmt
SELECT count(*) 
FROM @contact 
WHERE user_id=? AND status=2 AND is_del=0;

-- name: get_contacts@contact_manager
-- prepare: stmt
SELECT id, user_id, friend_id, group_id, remark, status, is_top, is_black, notice_enable, is_del 
FROM @contact 
WHERE (user_id=? AND friend_id=?) OR (user_id=? AND friend_id=?);

-- name: get_user_friend@contact_manager
-- prepare: stmt
SELECT id, user_id, friend_id, group_id, remark, status, is_top, is_black, notice_enable, is_del 
FROM @contact 
WHERE user_id=? AND friend_id=? AND is_del=0;

-- name: get_contact@contact_manager
-- prepare: stmt
SELECT id, user_id, friend_id, group_id, remark, status, is_top, is_black, notice_enable, is_del 
FROM @contact 
WHERE user_id=? AND friend_id=?;

-- name: is_friend@contact_manager
-- prepare: stmt
SELECT true FROM @contact WHERE user_id=? AND friend_id=? AND is_del=0 AND status=2;

--------------------------------------------------------------------------------
-- message sql dml
--------------------------------------------------------------------------------

-- name: create_message@message
-- prepare: named_stmt
INSERT INTO @message (sender_user_id, receiver_user_id, type, brief, content, post_id, comment_id, reply_id, created_on) 
VALUES (:sender_user_id, :receiver_user_id, :type, :brief, :content, :post_id, :comment_id, :reply_id, :created_on);

-- name: get_unread_count@message
-- prepare: stmt
SELECT count(*) FROM @message WHERE receiver_user_id=? AND is_read=0 AND is_del=0;

-- name: get_message_by_id@message
-- prepare: stmt
SELECT * FROM @message WHERE id=? AND is_del=0;

-- name: read_message@message
-- prepare: stmt
UPDATE @message SET is_read=1, modified_on=? WHERE id=?;

-- name: get_messages@message
-- prepare: named_stmt
SELECT * 
FROM @message 
WHERE receiver_user_id=:recerver_user_id AND is_del=0 
ORDER BY id DESC
LIMIT :limit OFFSET :offset

-- name: get_message_count@message
-- prepare: named_stmt
SELECT count(*) FROM @message WHERE receiver_user_id=:recerver_user_id AND is_del=0;

--------------------------------------------------------------------------------
-- security sql dml
--------------------------------------------------------------------------------

-- name: get_latest_phone_captcha@security
-- prepare: stmt
SELECT * FROM @captcha WHERE phone=? AND is_del=0;

-- name: use_phone_captcha@security
-- prepare: stmt
UPDATE @captcha SET use_times=use_times+1, modified_on=? WHERE id=? AND is_del=0;

-- name: create_phone_captcha@security
-- prepare: named_stmt
INSERT INTO @captcha (phone, captcha, expired_on, created_on) 
VALUES (:phone, :captcha, :expired_on, :created_on);

--------------------------------------------------------------------------------
-- ship_index sql dml
--------------------------------------------------------------------------------

-- name: index_by_admin@ship_index
-- prepare: stmt
SELECT * 
FROM @p_post 
WHERE is_del=0 
ORDER BY is_top DESC, latest_replied_on DESC
LIMIT ? OFFSET ?;

-- name: index_count_by_admin@ship_index
-- prepare: stmt
SELECT count(*)
FROM @p_post
WHERE is_del=0;

-- name: index_by_guest@ship_index
-- prepare: stmt
SELECT * 
FROM @p_post 
WHERE visibility=0 AND is_del=0 
ORDER BY is_top DESC, latest_replied_on DESC
LIMIT ? OFFSET ?;

-- name: index_count_by_guest@ship_index
-- prepare: stmt
SELECT count(*)
FROM @p_post
WHERE visibility=0 AND is_del=0;

-- name: index_by_self@ship_index
-- prepare: raw
-- clause: in
SELECT *
FROM @p_post
WHERE is_del=0 AND 
	(visibility=0 OR
	(visibility=1 AND user_id=?) OR
	(visibility=2 AND user_id IN ?))
ORDER BY is_top DESC, latest_replied_on DESC
LIMIT ? OFFSET ?;

-- name: index_count_by_self@ship_index
-- prepare: raw
-- clause: in
SELECT count(*)
FROM @p_post
WHERE is_del=0 AND 
	(visibility=0 OR
	(visibility=1 AND user_id=?) OR
	(visibility=2 AND user_id IN ?)); 

--------------------------------------------------------------------------------
-- simple_index sql dml
--------------------------------------------------------------------------------

-- name: index@simple_index
-- prepare: stmt
SELECT * 
FROM @p_post 
WHERE visibility=0
ORDER BY is_top DESC, latest_replied_on DESC
LIMIT ? OFFSET ?;

-- name: index_count@simple_index
-- prepare: stmt
SELECT count(*) 
FROM @p_post 
WHERE visibility=0;

--------------------------------------------------------------------------------
-- ship_index_a sql dml
--------------------------------------------------------------------------------

-- name: user_info@ship_index_a
-- prepare: raw
SELECT * FROM @user WHERE username=?

--------------------------------------------------------------------------------
-- simple_index_a sql dml
--------------------------------------------------------------------------------

-- name: user_info@simple_index_a
-- prepare: raw
SELECT * FROM @user WHERE username=?

--------------------------------------------------------------------------------
-- tweet sql dml
--------------------------------------------------------------------------------

-- name: get_post_by_id@tweet
-- prepare: stmt
SELECT * FROM @post WHERE id=? AND is_del=0;

-- name: get_user_posts@tweet
-- prepare: raw
-- clause: in
SELECT * FROM @post
WHERE user_id=? AND visibility IN (?)
ORDER BY latest_replies_on DESC
LIMIT ? OFFSET ?;

-- name: get_any_posts@tweet
-- prepare: raw
-- clause: in
SELECT * FROM @post WHERE visibility IN (?) AND is_del=0 LIMIT ? OFFSET ?;

-- name: get_user_post_count@tweet
-- prepare: raw
-- clause: in
SELECT count(*) FROM @post WHERE user_id=? AND visibility IN (?);

-- name: get_any_post_count@tweet
-- prepare: raw
-- clause: in
SELECT count(*) FROM @post WHERE visibility IN (?);

-- name: get_user_post_star@tweet
-- prepare: stmt
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
	@post_star s
	JOIN @post P ON s.post_id = P.ID 
WHERE
	s.post_id = ? 
	AND s.user_id = ? 
	AND s.is_del = 0 
	AND ( visibility <> 1 OR ( visibility = 1 AND P.user_id = ? ) ) 
ORDER BY
	P.ID DESC;

-- name: get_user_post_stars@tweet
-- prepare: stmt
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
	@post_star s
	JOIN @post P ON s.post_id = P.ID 
WHERE
	s.user_id = ? 
	AND s.is_del = 0 
	AND ( visibility <> 1 OR ( visibility = 1 AND P.user_id = ? ) ) 
ORDER BY
    s.ID DESC,
	P.ID DESC
LIMIT ? OFFSET ?;

-- name: get_user_post_star_count@tweet
-- prepare: stmt
SELECT
    count(*) 
FROM
	@post_star s
	JOIN @post P ON s.post_id = P.ID 
WHERE
	s.user_id = ? 
	AND s.is_del = 0 
	AND ( visibility <> 1 OR ( visibility = 1 AND P.user_id = ? ) );

-- name: get_user_post_collection@tweet
-- prepare: stmt
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
	@post_collection s
	JOIN @post P ON s.post_id = P.ID 
WHERE
	s.post_id = ? 
	AND s.user_id = ? 
	AND s.is_del = 0 
	AND ( visibility <> 1 OR ( visibility = 1 AND P.user_id = ? ) ) 
ORDER BY
	P.ID DESC;

-- name: get_user_post_collections@tweet
-- prepare: stmt
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
	@post_collection s
	JOIN @post P ON s.post_id = P.ID 
WHERE
	s.user_id = ? 
	AND s.is_del = 0 
	AND ( visibility <> 1 OR ( visibility = 1 AND P.user_id = ? ) ) 
ORDER BY
    s.ID DESC,
	P.ID DESC
LIMIT ? OFFSET ?;

-- name: get_user_post_collection_count@tweet
-- prepare: stmt
SELECT
    count(*) 
FROM
	@post_collection s
	JOIN @post P ON s.post_id = P.ID 
WHERE
	s.user_id = ? 
	AND s.is_del = 0 
	AND ( visibility <> 1 OR ( visibility = 1 AND P.user_id = ? ) );

-- name: get_post_attachment_bill@tweet
-- prepare: stmt
SELECT * FROM @post_attachment_bill WHERE post_id=? AND user_id=? AND is_del=0;

-- name: get_post_contents_by_ids@tweet
-- prepare: raw
-- clause: in
SELECT * 
FROM @post_content 
WHERE post_id IN (?) AND is_del=0;

-- name: get_post_content_by_id@tweet
-- prepare: stmt
SELECT * FROM @post_content WHERE id=? AND is_del=0;

-- name: user_media_tweets_by_guest@tweet
-- prepare: stmt
SELECT *
FROM @post_by_media
WHERE is_del=0 AND user_id=? AND visibility=0
ORDER BY latest_replied_on DESC
LIMIT ? OFFSET ?;

-- name: user_media_tweets_count_by_guest@tweet
-- prepare: stmt
SELECT count(*)
FROM @post_by_media
WHERE is_del=0 AND user_id=? AND visibility=0;

-- name: user_media_tweets_by_friend@tweet
-- prepare: stmt
SELECT *
FROM @post_by_media
WHERE is_del=0 AND user_id=? AND (visibility=0 OR visibility=2)
ORDER BY latest_replied_on DESC
LIMIT ? OFFSET ?;

-- name: user_media_tweets_count_by_friend@tweet
-- prepare: stmt
SELECT count(*)
FROM @post_by_media
WHERE is_del=0 AND user_id=? AND (visibility=0 OR visibility=2);

-- name: user_media_tweets_by_self@tweet
-- prepare: stmt
SELECT *
FROM @post_by_media
WHERE is_del=0 AND user_id=?
ORDER BY latest_replied_on DESC
LIMIT ? OFFSET ?;

-- name: user_media_tweets_count_by_self@tweet
-- prepare: stmt
SELECT count(*)
FROM @post_by_media
WHERE is_del=0 AND user_id=?;

-- name: user_comment_tweets_by_guest@tweet
-- prepare: stmt
SELECT *
FROM @post_by_comment
WHERE is_del=0 AND comment_user_id=? AND visibility=0
ORDER BY latest_replied_on DESC
LIMIT ? OFFSET ?;

-- name: user_comment_tweets_count_by_guest@tweet
-- prepare: stmt
SELECT count(*)
FROM @post_by_comment
WHERE is_del=0 AND comment_user_id=? AND visibility=0;

-- name: user_comment_tweets_by_friend@tweet
-- prepare: stmt
SELECT *
FROM @post_by_comment
WHERE is_del=0 AND comment_user_id=? AND (visibility=0 OR visibility=2)
ORDER BY latest_replied_on DESC
LIMIT ? OFFSET ?;

-- name: user_comment_tweets_count_by_friend@tweet
-- prepare: stmt
SELECT count(*)
FROM @post_by_comment
WHERE is_del=0 AND comment_user_id=? AND (visibility=0 OR visibility=2);

-- name: user_comment_tweets_by_self@tweet
-- prepare: stmt
SELECT *
FROM @post_by_comment
WHERE is_del=0 AND comment_user_id=?
ORDER BY latest_replied_on DESC
LIMIT ? OFFSET ?;

-- name: user_comment_tweets_count_by_self@tweet
-- prepare: stmt
SELECT count(*)
FROM @post_by_comment
WHERE is_del=0 AND comment_user_id=?;

-- name: user_star_tweets_by_guest@tweet
-- prepare: stmt
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
	@post_star star
	JOIN @post post ON star.post_id = post.ID 
WHERE
	star.is_del = 0 
	AND star.user_id =? 
	AND post.visibility = 0 
ORDER BY post.latest_replied_on DESC
LIMIT ? OFFSET ?;

-- name: user_star_tweets_count_by_guest@tweet
-- prepare: stmt
SELECT count(*)
FROM
	@post_star star
	JOIN @post post ON star.post_id = post.ID 
WHERE star.is_del=0 AND star.user_id=? AND post.visibility=0;

-- name: user_star_tweets_by_friend@tweet
-- prepare: stmt
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
	@post_star star
	JOIN @post post ON star.post_id = post.ID 
WHERE star.is_del=0 AND star.user_id=? AND (post.visibility=0 OR post.visibility=2)
ORDER BY post.latest_replied_on DESC
LIMIT ? OFFSET ?;

-- name: user_star_tweets_count_by_friend@tweet
-- prepare: stmt
SELECT count(*)
FROM
	@post_star star
	JOIN @post post ON star.post_id = post.ID 
WHERE star.is_del=0 AND star.user_id=? AND (post.visibility=0 OR post.visibility=2);

-- name: user_star_tweets_by_self@tweet
-- prepare: stmt
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
	@post_star star
	JOIN @post post ON star.post_id = post.ID 
WHERE star.is_del=0 AND star.user_id=? AND (post.visibility<>0 (post.visibility=0 AND post.user_id=?))
ORDER BY post.latest_replied_on DESC;
LIMIT ? OFFSET ?;

-- name: user_star_tweets_count_by_self@tweet
-- prepare: stmt
SELECT count(*)
FROM
	@post_star star
	JOIN @post post ON star.post_id = post.ID 
WHERE star.is_del=0 AND star.user_id=? AND (post.visibility<>0 (post.visibility=0 AND post.user_id=?));

-- name: user_star_tweets_by_admin@tweet
-- prepare: stmt
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
	@post_star star
	JOIN @post post ON star.post_id = post.ID 
WHERE star.is_del=0 AND star.user_id=?
ORDER BY post.latest_replied_on DESC;
LIMIT ? OFFSET ?;

-- name: user_star_tweets_count_by_admin@tweet
-- prepare: stmt
SELECT count(*)
FROM
	@post_star star
	JOIN @post post ON star.post_id = post.ID 
WHERE star.is_del=0 AND star.user_id=?;

--------------------------------------------------------------------------------
-- tweet_manage sql dml
--------------------------------------------------------------------------------

-- name: add_post@tweet_manage
-- prepare: named_stmt
INSERT INTO @post (user_id, tags, ip, ip_loc, attachment_price, visibility, latest_replies_on, created_on)
VALUES (:user_id, :tags, :ip, :ip_loc, :attachment_price, :visibility, :latest_replies_on, :created_on);

-- name: media_content_by_post_id@tweet_manage
-- prepare stmt
SELECT content FROM post_content WHERE post_id=? AND is_del=0 AND type IN (3, 4, 5, 7, 8);

-- name: del_post_by_id@tweet_manage
-- prepare: stmt
UPDATE @post SET is_del=1, deleted_on=? WHERE id=? AND is_del=0;

-- name: lock_post@tweet_manage
-- prepare: stmt
UPDATE @post SET is_lock=1-is_lock, modified_on=? WHERE id=? AND is_del=0;

-- name: stick_post@tweet_manage
-- prepare: stmt
UPDATE @post SET is_top=1-is_top, modified_on=? WHERE id=? AND is_del=0;

-- name: visible_post@tweet_manage
-- prepare: stmt
UPDATE @post SET visibility=?, is_top=?, modified_on=? WHERE id=? AND is_del=0;

-- name: highlight_post@tweet_manage
-- prepare: stmt
UPDATE @post SET is_essence=1-is_essence WHERE id=? AND is_del=0;

-- name: post_highlight_status@tweet_manage
-- prepare: stmt
SELECT user_id, is_essence FROM @post WHERE id=? AND is_del=0;

-- name: update_post@tweet_manage
-- prepare: named_stmt
UPDATE @post SET comment_count=:comment_count, 
    upvote_count=:upvote_count,
    collection_count=:collection_count,
    latest_replies_on=:latest_replies_on,
    modified_on=:modified_on
WHERE id=:id AND is_del=0;

-- name: add_post_star@tweet_manage
-- prepare: stmt
INSERT INTO @post_star (post_id, user_id, created_on)
VALUES (?, ?, ?);

-- name: del_post_star@tweet_manage
-- prepare: stmt
UPDATE @post_star SET is_del=1, deleted_on=? WHERE id=? AND is_del=0;

-- name: add_post_collection@tweet_manage
-- prepare: stmt
INSERT INTO @post_collection (post_id, user_id, created_on)
VALUES (?, ?, ?);

-- name: del_post_collection@tweet_manage
-- prepare: stmt
UPDATE @post_collection SET is_del=1, deleted_on=? WHERE id=? AND is_del=0;

-- name: add_post_content@tweet_manage
-- prepare: named_stmt
INSERT INTO @post_content (post_id, user_id, content, type, sort, created_on)
VALUES (:post_id, :user_id, :content, :type, :sort, :created_on);

-- name: add_attachment@tweet_manage
-- prepare: stmt
INSERT INTO @attachment (user_id, file_size, img_width, img_height, type, content, created_on)
VALUES (?, ?, ?, ?, ?, ?, ?);

-- name: comment_ids_by_post_id@tweet_manage
-- prepare: stmt
SELECT id FROM @comment WHERE post_id=? AND is_del=0;

-- name: comment_media_from_comment_ids@tweet_manage
-- prepare: raw
-- clause: in
SELECT content FROM @comment_content WHERE comment_id IN (?) AND type=3 AND is_del=0;

-- name: del_comment_by_post_id@tweet_manage
-- prepare: stmt
UPDATE @comment SET deleted_on=?, is_del=1 WHERE post_id=? AND is_del=0;

-- name: del_comment_content_by_comment_ids@tweet_manage
-- prepare: raw
-- clause: in
UPDATE @comment_content SET deleted_on=?, is_del=1 WHERE comment_id IN (?) AND is_del=0;

-- name: del_reply_by_comment_ids@tweet_manage
-- prepare: raw
-- clause: in
UPDATE @comment_reply SET deleted_on=?, is_del=1 WHERE comment_id IN (?) AND is_del=0;

--------------------------------------------------------------------------------
-- tweet_help sql dml
--------------------------------------------------------------------------------

-- name: get_post_content_by_ids@tweet_help
-- prepare: raw
-- clause: in
SELECT id, post_id, content, type, sort 
FROM @post_content
WHERE post_id IN (?) AND is_del=0;

-- name: get_users_by_ids@tweet_help
-- prepare: raw
-- clause: in
SELECT id, username, nickname, status, avatar, is_admin
FROM @user
WHERE id IN (?) AND is_del=0;

--------------------------------------------------------------------------------
-- tweet_a sql dml
--------------------------------------------------------------------------------

-- name: user_info@tweet_a
-- prepare: raw
SELECT * FROM @user WHERE username=?

-- name: tweet_info_by_id@tweet_a
-- prepare: raw
SELECT * FROM @user WHERE username=?

-- name: tweet_item_by_id@tweet_a
-- prepare: raw
SELECT * FROM @user WHERE username=?

-- name: user_tweets_by_admin@tweet_a
-- prepare: raw
SELECT * FROM @user WHERE username=?

-- name: user_tweets_by_self@tweet_a
-- prepare: raw
SELECT * FROM @user WHERE username=?

-- name: user_tweets_by_friend@tweet_a
-- prepare: raw
SELECT * FROM @user WHERE username=?

-- name: user_tweets_by_guest@tweet_a
-- prepare: raw
SELECT * FROM @user WHERE username=?

-- name: reaction_by_tweet_id@tweet_a
-- prepare: raw
SELECT * FROM @user WHERE username=?

-- name: user_reactions@tweet_a
-- prepare: raw
SELECT * FROM @user WHERE username=?

-- name: favorite_by_tweet_id@tweet_a
-- prepare: raw
SELECT * FROM @user WHERE username=?

-- name: user_favorites@tweet_a
-- prepare: raw
SELECT * FROM @user WHERE username=?

-- name: attachment_by_tweet_id@tweet_a
-- prepare: raw
SELECT * FROM @user WHERE username=?

--------------------------------------------------------------------------------
-- tweet_manage_a sql dml
--------------------------------------------------------------------------------

-- name: user_info@tweet_manage_a
-- prepare: raw
SELECT * FROM @user WHERE username=?

--------------------------------------------------------------------------------
-- tweet_help_a sql dml
--------------------------------------------------------------------------------

-- name: user_info@tweet_help_a
-- prepare: raw
SELECT * FROM @user WHERE username=?

--------------------------------------------------------------------------------
-- topic_a sql dml
--------------------------------------------------------------------------------

-- name: newest_tags@topic_a
-- get newest tag information
SELECT t.id id, t.user_id user_id, t.tag tag, t.quote_num quote_num, u.id, u.nickname, u.username, u.status, u.avatar, u.is_admin 
FROM @tag t 
JOIN @user u 
ON t.user_id = u.id 
WHERE t.is_del = 0 AND t.quote_num > 0 
ORDER BY t.id DESC 
LIMIT ? OFFSET ?;

-- name: hot_tags@topic_a
-- get get host tag information
SELECT t.id id, t.user_id user_id, t.tag tag, t.quote_num quote_num, u.id, u.nickname, u.username, u.status, u.avatar, u.is_admin 
FROM @tag t 
JOIN @user u 
ON t.user_id = u.id 
WHERE t.is_del = 0 AND t.quote_num > 0 
ORDER BY t.quote_num DESC 
LIMIT ? OFFSET ?;

-- name: tags_by_keyword_a@topic_a
-- get tags by keyword
SELECT id, user_id, tag, quote_num 
FROM @tag 
WHERE is_del = 0 
ORDER BY quote_num DESC 
LIMIT 6;

-- name: tags_by_keyword_b@topic_a
SELECT id, user_id, tag, quote_num 
FROM @tag WHERE is_del = 0 AND tag LIKE ? 
ORDER BY quote_num DESC 
LIMIT 6;

-- name: insert_tag@topic_a
INSERT INTO @tag (user_id, tag, created_on, modified_on, quote_num)
VALUES (?, ?, ?, ?, 1);

-- name: tags_by_id_a@topic_a
-- prepare: raw
-- clause: in
SELECT id FROM @tag WHERE id IN (?) AND is_del = 0 AND quote_num > 0;

-- name: tags_by_id_b@topic_a
-- prepare: raw
-- clause: in
SELECT id, user_id, tag, quote_num FROM @tag WHERE id IN (?);

-- name: decr_tags_by_id@topic_a
-- prepare: raw
-- clause: in
UPDATE @tag SET quote_num=quote_num-1, modified_on=? WHERE id IN (?);

-- name: tags_for_incr@topic_a
-- prepare: raw
-- clause: in
SELECT id, user_id, tag, quote_num FROM @tag WHERE tag IN (?);

-- name: incr_tags_by_id@topic_a
-- prepare: raw
-- clause: in
UPDATE @tag SET quote_num=quote_num+1, is_del=0, modified_on=? WHERE id IN (?);

--------------------------------------------------------------------------------
-- user_manage sql dml
--------------------------------------------------------------------------------

-- name: get_user_by_id@user_manage
-- prepare: stmt
SELECT * FROM @user WHERE id=? AND is_del=0;

-- name: get_user_by_username@user_manage
-- prepare: stmt
SELECT * FROM @user WHERE username=? AND is_del=0;

-- name: get_user_by_phone@user_manage
-- prepare: stmt
SELECT * FROM @user WHERE phone=? AND is_del=0;

-- name: get_users_by_ids@user_manage
-- prepare: raw
-- clause: in
SELECT * FROM @user WHERE id IN (?) AND is_del=0;

-- name: get_users_by_keyword@user_manage
-- prepare: stmt
SELECT * FROM @user WHERE username LIKE ? AND is_del=0 limit 6;

-- name: get_any_users@user_manage
-- prepare: stmt
SELECT * FROM @user WHERE is_del=0 ORDER BY id ASC limit 6;

-- name: create_user@user_manage
-- prepare: named_stmt
INSERT INTO @user (username, nickname, password, salt, avatar, status, created_on) 
VALUES (:username, :nickname, :password, :salt, :avatar, :status, :created_on);

-- name: update_user@user_manage
-- prepare: named_stmt
UPDATE @user 
SET username=:username, 
    nickname=:nickname, 
    phone=:phone,
    password=:password,
    salt=:salt,
    status=:status,
    avatar=:avatar,
    balance=:balance,
    is_admin=:is_admin,
    modified_on=:modified_on 
WHERE id=? AND is_del=0;

--------------------------------------------------------------------------------
-- wallet sql dml
--------------------------------------------------------------------------------

-- name: get_user_wallet_bills@wallet
-- prepare: stmt
SELECT * 
FROM @wallet_statement 
WHERE user_id=? AND is_del=0 
ORDER BY id DESC 
LIMIT ? OFFSET ?;

-- name: get_user_wallet_bill_count@wallet
-- prepare: stmt
SELECT count(*) FROM @wallet_statement WHERE user_id=? AND is_del=0;

-- name: get_recharge_by_id@wallet
-- prepare: stmt
SELECT * FROM @wallet_recharge WHERE id=? AND is_del=?;

-- name: create_recharge@wallet
-- prepare: stmt
INSERT INTO @wallet_recharge (user_id, amount, created_on) VALUES (?, ?, ?);

-- name: get_user_by_uid@wallet
-- prepare: stmt
SELECT * FROM @user WHERE id=? AND is_del=0;

-- name: add_user_balance@wallet
-- prepare: stmt
UPDATE @user SET balance=balance+?, modified_on=? WHERE id=? AND is_del=0;

-- name: minus_user_balance@wallet
-- prepare: stmt
UPDATE @user SET balance=balance-?, modified_on=? WHERE id=? AND is_del=0;

-- name: create_wallet_statement@wallet
-- prepare: stmt
INSERT INTO @wallet_statement (user_id, change_amount, balance_snapshot, reason, created_on)
VALUES (?, ?, ?, ?, ?);

-- name: get_user_balance@wallet
-- prepare: stmt
SELECT balance FROM @user WHERE id=? AND is_del=0;

-- name: mark_success_recharge@wallet
-- prepare: stmt
UPDATE @wallet_recharge 
SET trade_no=?, trade_status='TRADE_SUCCESS', modified_on=? 
WHERE id=? AND is_del=0;

-- name: new_post_attachment_bill@wallet
-- prepare: stmt
INSERT INTO @post_attachment_bill (post_id, user_id, paid_amount, created_on)
VALUES (?, ?, ?, ?);

-- name: new_post_bill@wallet
-- prepare: stmt
INSERT INTO @wallet_statement (post_id, user_id, change_amount, balance_snapshot, reason, created_on) 
VALUES (?, ?, ?, ?, ?, ?);
