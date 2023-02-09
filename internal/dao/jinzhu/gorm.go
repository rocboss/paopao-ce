// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package jinzhu

import (
	"github.com/rocboss/paopao-ce/internal/conf"
)

var (
	_anouncement        string
	_anouncementContent string
	_attachment         string
	_captcha            string
	_comment            string
	_commentContent     string
	_commentReply       string
	_contact            string
	_contactGroup       string
	_message            string
	_post               string
	_postAttachmentBill string
	_postCollection     string
	_postContent        string
	_postStar           string
	_tag                string
	_user               string
	_walletRecharge     string
	_walletStatement    string
)

func initTableName() {
	m := conf.DatabaseSetting.TableNames()
	_anouncement = m[conf.TableAnouncement]
	_anouncementContent = m[conf.TableAnouncementContent]
	_attachment = m[conf.TableAttachment]
	_captcha = m[conf.TableCaptcha]
	_comment = m[conf.TableComment]
	_commentContent = m[conf.TableCommentContent]
	_commentReply = m[conf.TableCommentReply]
	_contact = m[conf.TableContact]
	_contactGroup = m[conf.TableContactGroup]
	_message = m[conf.TableMessage]
	_post = m[conf.TablePost]
	_postAttachmentBill = m[conf.TablePostAttachmentBill]
	_postCollection = m[conf.TablePostCollection]
	_postContent = m[conf.TablePostContent]
	_postStar = m[conf.TablePostStar]
	_tag = m[conf.TableTag]
	_user = m[conf.TableUser]
	_walletRecharge = m[conf.TableWalletRecharge]
	_walletStatement = m[conf.TableWalletStatement]
}
