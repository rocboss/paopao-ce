ALTER TABLE `p_attachment` 
    RENAME KEY `idx_attachment_user_id` TO `idx_user`;

ALTER TABLE `p_captcha` 
    RENAME KEY `idx_captcha_phone` TO `idx_phone`,
    RENAME KEY `idx_captcha_expired_on` TO `idx_expired_on`,
    RENAME KEY `idx_captcha_use_times` TO `idx_use_times`;

ALTER TABLE `p_comment`
    RENAME KEY `idx_comment_post_id` TO `idx_post`,
    RENAME KEY `idx_comment_user_id` TO `idx_user`;

ALTER TABLE `p_comment_content`
    RENAME KEY `idx_comment_content_comment_id` TO `idx_reply`,
    RENAME KEY `idx_comment_content_user_id` TO `idx_user`,
    RENAME KEY `idx_comment_content_type` TO `idx_type`,
    RENAME KEY `idx_comment_content_sort` TO `idx_sort`;

ALTER TABLE `p_comment_reply`
    RENAME KEY `idx_comment_reply_comment_id` TO `idx_comment`;

ALTER TABLE `p_message`
    RENAME KEY `idx_message_receiver_user_id` TO `idx_receiver`,
    RENAME KEY `idx_message_is_read` TO `idx_is_read`,
    RENAME KEY `idx_message_type` TO `idx_type`;

ALTER TABLE `p_post`
    RENAME KEY `idx_post_user_id` TO `idx_user`,
    RENAME KEY `idx_post_visibility` TO `idx_visibility`;

ALTER TABLE `p_post_attachment_bill`
    RENAME KEY `idx_post_attachment_bill_post_id` TO `idx_post`,
    RENAME KEY `idx_post_attachment_bill_user_id` TO `idx_user`;

ALTER TABLE `p_post_collection`
    RENAME KEY `idx_post_collection_post_id` TO `idx_post`,
    RENAME KEY `idx_post_collection_user_id` TO `idx_user`;

ALTER TABLE `p_post_content`
    RENAME KEY `idx_post_content_post_id` TO `idx_post`,
    RENAME KEY `idx_post_content_user_id` TO `idx_user`;

ALTER TABLE `p_post_star`
    RENAME KEY `idx_post_star_post_id` TO `idx_post`,
    RENAME KEY `idx_post_star_user_id` TO `idx_user`;

ALTER TABLE `p_tag`
    RENAME KEY `idx_tag_user_id` TO `idx_user`,
    RENAME KEY `idx_tag_quote_num` TO `idx_num`;

ALTER TABLE `p_user`
    RENAME KEY `idx_user_username` TO `idx_username`,
    RENAME KEY `idx_user_phone` TO `idx_phone`;

ALTER TABLE `p_wallet_recharge`
    RENAME KEY `idx_wallet_recharge_user_id` TO `idx_user`,
    RENAME KEY `idx_wallet_recharge_trade_no` TO `idx_trade_no`,
    RENAME KEY `idx_wallet_recharge_trade_status` TO `idx_trade_status`;

ALTER TABLE `p_wallet_statement`
    RENAME KEY `idx_wallet_statement_user_id` TO `idx_user`;

ALTER TABLE `p_contact`
    RENAME KEY `idx_contact_user_friend` TO `idx_user_friend_id`,
    RENAME KEY `idx_contact_user_friend_status` TO `idx_user_friend_status`;
 