ALTER TABLE `p_attachment` 
    RENAME KEY `idx_user` TO `idx_attachment_user_id`;

ALTER TABLE `p_captcha` 
    RENAME KEY `idx_phone` TO `idx_captcha_phone`,
    RENAME KEY `idx_expired_on` TO `idx_captcha_expired_on`,
    RENAME KEY `idx_use_times` TO `idx_captcha_use_times`;

ALTER TABLE `p_comment`
    RENAME KEY `idx_post` TO `idx_comment_post_id`,
    RENAME KEY `idx_user` TO `idx_comment_user_id`;

ALTER TABLE `p_comment_content`
    RENAME KEY `idx_reply` TO `idx_comment_content_comment_id`,
    RENAME KEY `idx_user` TO `idx_comment_content_user_id`,
    RENAME KEY `idx_type` TO `idx_comment_content_type`,
    RENAME KEY `idx_sort` TO `idx_comment_content_sort`;

ALTER TABLE `p_comment_reply`
    RENAME KEY `idx_comment` TO `idx_comment_reply_comment_id`;

ALTER TABLE `p_message`
    RENAME KEY `idx_receiver` TO `idx_message_receiver_user_id`,
    RENAME KEY `idx_is_read` TO `idx_message_is_read`,
    RENAME KEY `idx_type` TO `idx_message_type`;

ALTER TABLE `p_post`
    RENAME KEY `idx_user` TO `idx_post_user_id`,
    RENAME KEY `idx_visibility` TO `idx_post_visibility`;

ALTER TABLE `p_post_attachment_bill`
    RENAME KEY `idx_post` TO `idx_post_attachment_bill_post_id`,
    RENAME KEY `idx_user` TO `idx_post_attachment_bill_user_id`;

ALTER TABLE `p_post_collection`
    RENAME KEY `idx_post` TO `idx_post_collection_post_id`,
    RENAME KEY `idx_user` TO `idx_post_collection_user_id`;

ALTER TABLE `p_post_content`
    RENAME KEY `idx_post` TO `idx_post_content_post_id`,
    RENAME KEY `idx_user` TO `idx_post_content_user_id`;

ALTER TABLE `p_post_star`
    RENAME KEY `idx_post` TO `idx_post_star_post_id`,
    RENAME KEY `idx_user` TO `idx_post_star_user_id`;

ALTER TABLE `p_tag`
    RENAME KEY `idx_user` TO `idx_tag_user_id`,
    RENAME KEY `idx_num` TO `idx_tag_quote_num`;

ALTER TABLE `p_user`
    RENAME KEY `idx_username` TO `idx_user_username`,
    RENAME KEY `idx_phone` TO `idx_user_phone`;

ALTER TABLE `p_wallet_recharge`
    RENAME KEY `idx_user` TO `idx_wallet_recharge_user_id`,
    RENAME KEY `idx_trade_no` TO `idx_wallet_recharge_trade_no`,
    RENAME KEY `idx_trade_status` TO `idx_wallet_recharge_trade_status`;

ALTER TABLE `p_wallet_statement`
    RENAME KEY `idx_user` To `idx_wallet_statement_user_id`;

ALTER TABLE `p_contact`
    RENAME KEY `idx_user_friend_id` TO `idx_contact_user_friend`,
    RENAME KEY `idx_user_friend_status` TO `idx_contact_user_friend_status`;
