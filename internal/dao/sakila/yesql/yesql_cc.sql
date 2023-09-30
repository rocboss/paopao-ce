--------------------------------------------------------------------------------
-- file yesql_cc.sql
-- this sql file define all sql dml for old version data logic
--------------------------------------------------------------------------------

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

-- name: get_newest_comments@comment
-- prepare: stmt
SELECT * 
FROM @comment 
WHERE post_id=? AND is_del=0 
ORDER BY is_essence DESC, id DESC 
LIMIT ? OFFSET ?; 

-- name: get_hots_comments@comment
-- prepare: stmt
SELECT c.* 
FROM @comment c
LEFT JOIN @comment_metric m
ON c.id=m.comment_id
WHERE c.post_id=? AND c.is_del=0 AND m.is_del=0
ORDER BY is_essence DESC, m.rank_score DESC, c.id DESC
LIMIT ? OFFSET ?; 

-- name: get_default_comments@comment
-- prepare: stmt
SELECT * 
FROM @comment 
WHERE post_id=? AND is_del=0 
ORDER BY is_essence DESC, id ASC 
LIMIT ? OFFSET ?; 

-- name: get_comment_by_id@comment
-- prepare: stmt
SELECT * FROM @comment WHERE id=? AND is_del=0;

-- name: get_comment_count@comment
-- prepare: stmt
SELECT count(*) FROM @comment WHERE post_id=? AND is_del=0;

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
SELECT user_id, 
	tweet_id, 
	comment_id, 
	reply_id, 
	comment_type, 
	is_thumbs_up, 
	is_thumbs_down 
FROM @tweet_comment_thumbs WHERE user_id=? AND tweet_id=?;

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

-- name: highlight_comment@comment_manage
UPDATE @comment
SET is_essence=1-is_essence,
	modified_on=?
WHERE id=? AND user_id=? AND is_del=0;

-- name: getIssenceStatusById@comment_manage
SELECT is_essence FROM @comment WHERE id=?;

-- name: incr_comment_reply_count@comment_manage
-- prepare: stmt
UPDATE @comment 
SET reply_count=reply_count+1,
	modified_on=?
WHERE id=? AND is_del=0;

-- name: decr_comment_reply_count@comment_manage
-- prepare: stmt
UPDATE @comment 
SET reply_count=reply_count-1,
	modified_on=?
WHERE id=? AND is_del=0;

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
SET is_del=1, deleted_on=? 
WHERE user_id=? AND follow_id=? AND is_del=0;

-- name: list_follows@following_manager
-- prepare: stmt
SELECT u.id user_id, 
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
SELECT u.id user_id, 
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
-- trends_manager sql dml
--------------------------------------------------------------------------------

-- name: get_index_trends@trends_manager
-- prepare: stmt
SELECT u.username username, 
	u.nickname nickname,
	u.avatar avatar
FROM @user_relation r
JOIN @user u 
ON r.he_uid=u.id
JOIN @user_metric m 
ON r.he_uid=m.user_id
WHERE r.user_id=? 
	AND m.is_del=0
	AND m.tweets_count>0
ORDER BY r.style ASC
LIMIT ? OFFSET ?;

-- name: count_index_trends@trends_manager
-- prepare: stmt
SELECT count(*)
FROM @user_relation r
JOIN @user u 
ON r.he_uid=u.id
JOIN @user_metric m 
ON r.he_uid=m.user_id
WHERE r.user_id=? 
	AND m.is_del=0
	AND m.tweets_count>0;

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
WHERE c.user_id=? AND c.status=2 AND c.is_del=0 
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

-- name: read_all_message@message
-- prepare: stmt
UPDATE @message SET is_read=1, modified_on=? WHERE receiver_user_id=?;

-- name: get_system_messages@message
-- prepare: stmt
SELECT * 
FROM @message 
WHERE receiver_user_id=? AND type IN (1, 2, 3, 99) AND is_del=0 
ORDER BY id DESC
LIMIT ? OFFSET ?;

-- name: count_system_messages@message
-- prepare: stmt
SELECT count(*) 
FROM @message 
WHERE receiver_user_id=? AND type IN (1, 2, 3, 99) AND is_del=0; 

-- name: get_whisper_messages@message
-- prepare: stmt
SELECT * 
FROM @message 
WHERE ((receiver_user_id=? OR sender_user_id=?) AND type=4) AND is_del=0 
ORDER BY id DESC
LIMIT ? OFFSET ?;

-- name: count_whisper_messages@message
-- prepare: stmt
SELECT count(*) 
FROM @message 
WHERE ((receiver_user_id=? OR sender_user_id=?) AND type=4) AND is_del=0;

-- name: get_requesting_messages@message
-- prepare: stmt
SELECT * 
FROM @message 
WHERE receiver_user_id=? AND type=5 AND is_del=0 
ORDER BY id DESC
LIMIT ? OFFSET ?;

-- name: count_requesting_messages@message
-- prepare: stmt
SELECT count(*) 
FROM @message 
WHERE receiver_user_id=? AND type=5 AND is_del=0;

-- name: get_unread_messages@message
-- prepare: stmt
SELECT * 
FROM @message 
WHERE receiver_user_id=? AND is_read=0 AND is_del=0 
ORDER BY id DESC
LIMIT ? OFFSET ?;

-- name: count_unread_messages@message
-- prepare: stmt
SELECT count(*) 
FROM @message 
WHERE receiver_user_id=? AND is_read=0 AND is_del=0;

-- name: get_all_messages@message
-- prepare: stmt
SELECT * 
FROM @message 
WHERE (receiver_user_id=? OR (sender_user_id=? AND type=4)) AND is_del=0 
ORDER BY id DESC
LIMIT ? OFFSET ?;

-- name: count_all_messages@message
-- prepare: stmt
SELECT count(*)
FROM @message 
WHERE (receiver_user_id=? OR (sender_user_id=? AND type=4)) AND is_del=0;

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
FROM @post 
WHERE is_del=0 
ORDER BY is_top DESC, latest_replied_on DESC
LIMIT ? OFFSET ?;

-- name: index_count_by_admin@ship_index
-- prepare: stmt
SELECT count(*)
FROM @post
WHERE is_del=0;

-- name: index_by_guest@ship_index
-- prepare: stmt
SELECT * 
FROM @post 
WHERE visibility=90 AND is_del=0 
ORDER BY is_top DESC, latest_replied_on DESC
LIMIT ? OFFSET ?;

-- name: index_count_by_guest@ship_index
-- prepare: stmt
SELECT count(*)
FROM @post
WHERE visibility=90 AND is_del=0;

-- name: index_by_self@ship_index
-- prepare: raw
-- clause: in
SELECT *
FROM @post
WHERE is_del=0 AND 
	(visibility=90 OR
	(visibility=0 AND user_id=?) OR
	(visibility=50 AND user_id IN (?)))
ORDER BY is_top DESC, latest_replied_on DESC
LIMIT ? OFFSET ?;

-- name: index_count_by_self@ship_index
-- prepare: raw
-- clause: in
SELECT count(*)
FROM @post
WHERE is_del=0 AND 
	(visibility=90 OR
	(visibility=0 AND user_id=?) OR
	(visibility=50 AND user_id IN (?))); 

--------------------------------------------------------------------------------
-- simple_index sql dml
--------------------------------------------------------------------------------

-- name: index@simple_index
-- prepare: stmt
SELECT * 
FROM @post 
WHERE visibility=90
ORDER BY is_top DESC, latest_replied_on DESC
LIMIT ? OFFSET ?;

-- name: index_count@simple_index
-- prepare: stmt
SELECT count(*) 
FROM @post 
WHERE visibility=90;

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
ORDER BY latest_replied_on DESC
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
	AND ( visibility >= 50 OR ( visibility = 0 AND P.user_id = ? ) ) 
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
	AND ( visibility >= 50 OR ( visibility = 0 AND P.user_id = ? ) ) 
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
	AND ( visibility >= 50 OR ( visibility = 0 AND P.user_id = ? ) );

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
	AND ( visibility >= 50 OR ( visibility = 0 AND P.user_id = ? ) ) 
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
	AND ( visibility >= 50 OR ( visibility = 0 AND P.user_id = ? ) ) 
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
	AND ( visibility >= 50 OR ( visibility = 0 AND P.user_id = ? ) );

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
FROM @post_by_media
WHERE is_del=0 AND user_id=? AND visibility=90
ORDER BY latest_replied_on DESC
LIMIT ? OFFSET ?;

-- name: user_media_tweets_count_by_guest@tweet
-- prepare: stmt
SELECT count(*)
FROM @post_by_media
WHERE is_del=0 AND user_id=? AND visibility>=90;

-- name: user_media_tweets_by_friend@tweet
-- prepare: stmt
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
FROM @post_by_media
WHERE is_del=0 AND user_id=? AND visibility>=50
ORDER BY latest_replied_on DESC
LIMIT ? OFFSET ?;

-- name: user_media_tweets_count_by_friend@tweet
-- prepare: stmt
SELECT count(*)
FROM @post_by_media
WHERE is_del=0 AND user_id=? AND visibility>=50;

-- name: user_media_tweets_by_self@tweet
-- prepare: stmt
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
FROM @post_by_comment
WHERE is_del=0 AND comment_user_id=? AND visibility>=90
ORDER BY latest_replied_on DESC
LIMIT ? OFFSET ?;

-- name: user_comment_tweets_count_by_guest@tweet
-- prepare: stmt
SELECT count(*)
FROM @post_by_comment
WHERE is_del=0 AND comment_user_id=? AND visibility>=90;

-- name: user_comment_tweets_by_friend@tweet
-- prepare: stmt
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
FROM @post_by_comment
WHERE is_del=0 AND comment_user_id=? AND visibility>=50
ORDER BY latest_replied_on DESC
LIMIT ? OFFSET ?;

-- name: user_comment_tweets_count_by_friend@tweet
-- prepare: stmt
SELECT count(*)
FROM @post_by_comment
WHERE is_del=0 AND comment_user_id=? AND visibility>=50

-- name: user_comment_tweets_by_self@tweet
-- prepare: stmt
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
	star.is_del=0 
	AND star.user_id=? 
	AND post.visibility>=90
ORDER BY post.latest_replied_on DESC
LIMIT ? OFFSET ?;

-- name: user_star_tweets_count_by_guest@tweet
-- prepare: stmt
SELECT count(*)
FROM
	@post_star star
	JOIN @post post ON star.post_id = post.ID 
WHERE star.is_del=0 AND star.user_id=? AND post.visibility>=90;

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
WHERE star.is_del=0 AND star.user_id=? AND post.visibility>=50
ORDER BY post.latest_replied_on DESC
LIMIT ? OFFSET ?;

-- name: user_star_tweets_count_by_friend@tweet
-- prepare: stmt
SELECT count(*)
FROM
	@post_star star
	JOIN @post post ON star.post_id = post.ID 
WHERE star.is_del=0 AND star.user_id=? AND post.visibility>=50;

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
WHERE star.is_del=0 AND star.user_id=? AND (post.visibility<>90 OR (post.visibility>=90 AND post.user_id=?))
ORDER BY post.latest_replied_on DESC
LIMIT ? OFFSET ?;

-- name: user_star_tweets_count_by_self@tweet
-- prepare: stmt
SELECT count(*)
FROM
	@post_star star
	JOIN @post post ON star.post_id = post.ID 
WHERE star.is_del=0 AND star.user_id=? AND (post.visibility<>90 OR (post.visibility>=90 AND post.user_id=?));

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
ORDER BY post.latest_replied_on DESC
LIMIT ? OFFSET ?;

-- name: user_star_tweets_count_by_admin@tweet
-- prepare: stmt
SELECT count(*)
FROM
	@post_star star
	JOIN @post post ON star.post_id = post.ID 
WHERE star.is_del=0 AND star.user_id=?;

-- name: list_user_tweets@tweet
-- prepare: stmt
SELECT * 
FROM @post
WHERE user_id=? AND visibility>=? AND is_essence=? AND is_del=0
ORDER BY is_top DESC, latest_replied_on DESC
LIMIT ? OFFSET ?;

-- name: count_user_tweets@tweet
-- prepare: stmt
SELECT count(*) 
FROM @post
WHERE user_id=? AND visibility>=? AND is_essence=? AND is_del=0;

-- name: list_index_newest_tweets@tweet
-- prepare: stmt
SELECT *
FROM @post
WHERE visibility>=90 AND is_del=0
ORDER BY is_top DESC, latest_replied_on DESC
LIMIT ? OFFSET ?;

-- name: count_index_newest_tweets@tweet
-- prepare: stmt
SELECT count(*)
FROM @post
WHERE visibility>=90 AND is_del=0;

-- name: list_index_hots_tweets@tweet
-- prepare: stmt
SELECT post.*
FROM @post post
LEFT JOIN @post_metric metric
ON post.id=metric.post_id
WHERE post.visibility>=90 AND post.is_del=0
ORDER BY post.is_top DESC, metric.rank_score DESC, post.latest_replied_on DESC
LIMIT ? OFFSET ?;

-- name: count_index_hots_tweets@tweet
-- prepare: stmt
SELECT count(*)
FROM @post post
LEFT JOIN @post_metric metric
ON post.id=metric.post_id AND metric.is_del=0
WHERE post.visibility>=90 AND post.is_del=0;

-- name: list_sync_search_tweets@tweet
-- prepare: stmt
SELECT *
FROM @post
WHERE visibility>=50 AND is_del=0
LIMIT ? OFFSET ?;

-- name: count_sync_search_tweets@tweet
-- prepare: stmt
SELECT count(*)
FROM @post
WHERE visibility>=50 AND is_del=0;

-- name: list_following_tweets_friend_follow@tweet
-- prepare: raw
-- clause: in
SELECT *
FROM @post
WHERE (user_id=? OR (visibility>=50 AND user_id IN(?)) OR (visibility>=60 AND user_id IN(?))) AND is_del=0
ORDER BY is_top DESC, latest_replied_on DESC
LIMIT ? OFFSET ?;

-- name: count_following_tweets_friend_follow@tweet
-- prepare: raw
-- clause: in
SELECT count(*)
FROM @post
WHERE (user_id=? OR (visibility>=50 AND user_id IN(?)) OR (visibility>=60 AND user_id IN(?))) AND is_del=0;

-- name: list_following_tweets_friend@tweet
-- prepare: raw
-- clause: in
SELECT *
FROM @post
WHERE (user_id=? OR (visibility>=50 AND user_id IN(?))) AND is_del=0
ORDER BY is_top DESC, latest_replied_on DESC
LIMIT ? OFFSET ?;

-- name: count_following_tweets_friend@tweet
-- prepare: raw
-- clause: in
SELECT count(*)
FROM @post
WHERE (user_id=? OR (visibility>=50 AND user_id IN(?))) AND is_del=0;

-- name: list_following_tweets_follow@tweet
-- prepare: raw
-- clause: in
SELECT *
FROM @post
WHERE (user_id=? OR (visibility>=60 AND user_id IN(?))) AND is_del=0
ORDER BY is_top DESC, latest_replied_on DESC
LIMIT ? OFFSET ?;

-- name: count_following_tweets_follow@tweet
-- prepare: raw
-- clause: in
SELECT count(*)
FROM @post
WHERE (user_id=? OR (visibility>=60 AND user_id IN(?))) AND is_del=0;

-- name: list_following_tweets@tweet
-- prepare: stmt
SELECT *
FROM @post
WHERE user_id=? AND is_del=0
ORDER BY is_top DESC, latest_replied_on DESC
LIMIT ? OFFSET ?;

-- name: count_following_tweets@tweet
-- prepare: stmt
SELECT count(*)
FROM @post
WHERE user_id=? AND is_del=0;

-- name: get_be_friend_ids@tweet
-- prepare: stmt
SELECT user_id FROM @contact WHERE friend_id=? AND is_del=0;

-- name: get_be_follow_ids@tweet
-- prepare: stmt
SELECT follow_id FROM @following WHERE user_id=? AND is_del=0;

--------------------------------------------------------------------------------
-- tweet_manage sql dml
--------------------------------------------------------------------------------

-- name: add_post@tweet_manage
-- prepare: named_stmt
INSERT INTO @post (user_id, tags, ip, ip_loc, attachment_price, visibility, latest_replied_on, created_on)
VALUES (:user_id, :tags, :ip, :ip_loc, :attachment_price, :visibility, :latest_replied_on, :created_on);

-- name: media_content_by_post_id@tweet_manage
-- prepare stmt
SELECT content FROM @post_content WHERE post_id=? AND is_del=0 AND type IN (3, 4, 5, 7, 8);

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
    latest_replied_on=:latest_replied_on,
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
-- tweet_metrics sql dml
--------------------------------------------------------------------------------

-- name: update_rank_score@tweet_metrics
-- prepare: stmt
UPDATE @post_metric SET rank_score=?, modified_on=? WHERE post_id=? AND is_del=0;

-- name: get_motivation_factor@tweet_metrics
-- prepare: stmt
SELECT motivation_factor FROM @post_metric WHERE post_id=? AND is_del=0;

-- name: add_tweet_metric@tweet_metrics
-- prepare: stmt
INSERT INTO @post_metric (post_id, created_on) VALUES (?, ?);

-- name: upsert_tweet_metric@tweet_metrics
-- prepare: stmt
INSERT INTO @post_metric (post_id, rank_score, created_on) VALUES (?, ?, ?);

-- name: delete_tweet_metric@tweet_metrics
-- prepare: stmt
UPDATE @post_metric SET is_del=1, deleted_on=? WHERE post_id=? AND is_del=0;

--------------------------------------------------------------------------------
-- comment_metrics sql dml
--------------------------------------------------------------------------------

-- name: update_rank_score@comment_metrics
-- prepare: stmt
UPDATE @comment_metric SET rank_score=?, modified_on=? WHERE comment_id=? AND is_del=0;

-- name: get_motivation_factor@comment_metrics
-- prepare: stmt
SELECT motivation_factor FROM @comment_metric WHERE comment_id=? AND is_del=0;

-- name: add_comment_metric@comment_metrics
-- prepare: stmt
INSERT INTO @comment_metric (comment_id, created_on) VALUES (?, ?);

-- name: upsert_comment_metric@comment_metrics
-- prepare: stmt
INSERT INTO @comment_metric (comment_id, rank_score, created_on) VALUES (?, ?, ?);

-- name: delete_comment_metric@comment_metrics
-- prepare: stmt
UPDATE @comment_metric SET is_del=1, deleted_on=? WHERE comment_id=? AND is_del=0;

--------------------------------------------------------------------------------
-- user_metrics sql dml
--------------------------------------------------------------------------------

-- name: update_user_metric@user_metrics
-- prepare: stmt
UPDATE @user_metric SET tweets_count=?, modified_on=? WHERE user_id=? AND is_del=0;

-- name: get_tweets_count@user_metrics
-- prepare: stmt
SELECT tweets_count FROM @user_metric WHERE user_id=? AND is_del=0;

-- name: add_user_metric@user_metrics
-- prepare: stmt
INSERT INTO @user_metric (user_id, created_on) VALUES (?, ?);

-- name: delete_user_metric@user_metrics
-- prepare: stmt
UPDATE @user_metric SET is_del=1, deleted_on=? WHERE user_id=? AND is_del=0;

--------------------------------------------------------------------------------
-- user_manage sql dml
--------------------------------------------------------------------------------

-- name: get_register_user_count@user_manage
-- prepare: stmt
SELECT count(*) FROM @user WHERE is_del=0;

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
INSERT INTO @user (username, nickname, password, salt, avatar, status, created_on, balance) 
VALUES (:username, :nickname, :password, :salt, :avatar, :status, :created_on, 0);

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
-- user_relation sql dml
--------------------------------------------------------------------------------

-- name: my_friend_ids@user_relation
-- prepare: stmt
SELECT friend_id FROM @contact WHERE user_id=? AND status=2 AND is_del=0;

-- name: my_follow_ids@user_relation
-- prepare: stmt
SELECT follow_id FROM @following WHERE user_id=? AND is_del=0;

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
