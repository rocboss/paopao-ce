CREATE VIEW p_post_by_media AS 
SELECT post.* 
FROM
	( SELECT DISTINCT post_id FROM p_post_content WHERE ( TYPE = 3 OR TYPE = 4 OR TYPE = 7 OR TYPE = 8 ) AND is_del = 0 ) media
	JOIN p_post post ON media.post_id = post.ID 
WHERE
	post.is_del = 0;

CREATE VIEW p_post_by_comment AS 
SELECT P.*, C.user_id comment_user_id
FROM
	(
	SELECT
		post_id,
		user_id
	FROM
		p_comment 
	WHERE
		is_del = 0 UNION
	SELECT
		post_id,
		reply.user_id user_id
	FROM
		p_comment_reply reply
		JOIN p_comment COMMENT ON reply.comment_id = COMMENT.ID 
	WHERE
		reply.is_del = 0 
		AND COMMENT.is_del = 0 
	)
	C JOIN p_post P ON C.post_id = P.ID 
WHERE
	P.is_del = 0;
