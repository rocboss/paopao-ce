CREATE VIEW p_user_relation AS 
SELECT user_id, friend_id he_uid, 5 AS style 
FROM p_contact WHERE status=2 AND is_del=0
UNION
SELECT user_id, follow_id he_uid, 10 AS style 
FROM p_following WHERE is_del=0;
