// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package jinzhu

import (
	"github.com/rocboss/paopao-ce/internal/conf"
)

var (
	tableAnouncement        string
	tableAnouncementContent string
	tableAttachment         string
	tableCaptcha            string
	tableComment            string
	tableCommentContent     string
	tableCommentReply       string
	tableContact            string
	tableContactGroup       string
	tableMessage            string
	tablePost               string
	tablePostAttachmentBill string
	tablePostCollection     string
	tablePostContent        string
	tablePostStar           string
	tableTag                string
	tableUser               string
	tableWalletRecharge     string
	tableWalletStatement    string
)

func initTableName() {
	m := conf.DatabaseSetting.TableNames()
	tableAnouncement = m[conf.TableAnouncement]
	tableAnouncementContent = m[conf.TableAnouncementContent]
	tableAttachment = m[conf.TableAttachment]
	tableCaptcha = m[conf.TableCaptcha]
	tableComment = m[conf.TableComment]
	tableCommentContent = m[conf.TableCommentContent]
	tableCommentReply = m[conf.TableCommentReply]
	tableContact = m[conf.TableContact]
	tableContactGroup = m[conf.TableContactGroup]
	tableMessage = m[conf.TableMessage]
	tablePost = m[conf.TablePost]
	tablePostAttachmentBill = m[conf.TablePostAttachmentBill]
	tablePostCollection = m[conf.TablePostCollection]
	tablePostContent = m[conf.TablePostContent]
	tablePostStar = m[conf.TablePostStar]
	tableTag = m[conf.TableTag]
	tableUser = m[conf.TableUser]
	tableWalletRecharge = m[conf.TableWalletRecharge]
	tableWalletStatement = m[conf.TableWalletStatement]
}
