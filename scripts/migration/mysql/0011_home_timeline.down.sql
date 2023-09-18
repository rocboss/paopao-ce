DROP TABLE IF EXISTS `p_post_metric`;

-- 原来的可见性: 0公开 1私密 2好友可见 3关注可见
-- 现在的可见性: 0私密 10充电可见 20订阅可见 30保留 40保留 50好友可见 60关注可见 70保留 80保留 90公开
UPDATE p_post a, p_post b 
SET a.visibility = (
	CASE b.visibility 
		WHEN 90 THEN 0 
		WHEN 0 THEN 1 
		WHEN 50 THEN 2 
		WHEN 60 THEN 3 
		ELSE 0
	END 
)
WHERE a.ID = b.ID;
