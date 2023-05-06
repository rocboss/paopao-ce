ALTER TABLE `p_comment` DROP COLUMN `thumbs_up_count`;
ALTER TABLE `p_comment` DROP COLUMN `thumbs_down_count`;
ALTER TABLE `p_comment_reply` DROP COLUMN `thumbs_up_count`;
ALTER TABLE `p_comment_reply` DROP COLUMN `thumbs_down_count`;

DROP TABLE IF EXISTS `p_tweet_comment_thumbs`;
