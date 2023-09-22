// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package jinzhu

import (
	"github.com/rocboss/paopao-ce/internal/conf"
)

// 数据库表名，统一使用 _<table name>_ 的形式命名， 比如tag表 => _tag_
var (
	_anouncement_        string
	_anouncementContent_ string
	_attachment_         string
	_captcha_            string
	_comment_            string
	_commentMetric_      string
	_commentContent_     string
	_commentReply_       string
	_following_          string
	_contact_            string
	_contactGroup_       string
	_message_            string
	_post_               string
	_post_metric_        string
	_post_by_comment_    string
	_post_by_media_      string
	_postAttachmentBill_ string
	_postCollection_     string
	_postContent_        string
	_postStar_           string
	_tag_                string
	_user_               string
	_userMetric_         string
	_walletRecharge_     string
	_walletStatement_    string
)

func initTableName() {
	m := conf.DatabaseSetting.TableNames()
	_anouncement_ = m[conf.TableAnouncement]
	_anouncementContent_ = m[conf.TableAnouncementContent]
	_attachment_ = m[conf.TableAttachment]
	_captcha_ = m[conf.TableCaptcha]
	_comment_ = m[conf.TableComment]
	_commentMetric_ = m[conf.TableCommentMetric]
	_commentContent_ = m[conf.TableCommentContent]
	_commentReply_ = m[conf.TableCommentReply]
	_following_ = m[conf.TableFollowing]
	_contact_ = m[conf.TableContact]
	_contactGroup_ = m[conf.TableContactGroup]
	_message_ = m[conf.TableMessage]
	_post_ = m[conf.TablePost]
	_post_metric_ = m[conf.TablePostMetric]
	_post_by_comment_ = m[conf.TablePostByComment]
	_post_by_media_ = m[conf.TablePostByMedia]
	_postAttachmentBill_ = m[conf.TablePostAttachmentBill]
	_postCollection_ = m[conf.TablePostCollection]
	_postContent_ = m[conf.TablePostContent]
	_postStar_ = m[conf.TablePostStar]
	_tag_ = m[conf.TableTag]
	_user_ = m[conf.TableUser]
	_userMetric_ = m[conf.TableUserMetric]
	_walletRecharge_ = m[conf.TableWalletRecharge]
	_walletStatement_ = m[conf.TableWalletStatement]
}
