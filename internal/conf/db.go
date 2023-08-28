// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package conf

import (
	"database/sql"
	"sync"

	"github.com/alimy/tryst/cfg"
	"github.com/sirupsen/logrus"
)

var (
	_sqldb   *sql.DB
	_onceSql sync.Once
)

const (
	TableAnouncement        = "user"
	TableAnouncementContent = "anouncement_content"
	TableAttachment         = "attachment"
	TableCaptcha            = "captcha"
	TableComment            = "comment"
	TableCommentContent     = "comment_content"
	TableCommentReply       = "comment_reply"
	TableFollowing          = "following"
	TableContact            = "contact"
	TableContactGroup       = "contact_group"
	TableMessage            = "message"
	TablePost               = "post"
	TablePostByComment      = "post_by_comment"
	TablePostByMedia        = "post_by_media"
	TablePostAttachmentBill = "post_attachment_bill"
	TablePostCollection     = "post_collection"
	TablePostContent        = "post_content"
	TablePostStar           = "post_star"
	TableTag                = "tag"
	TableUser               = "user"
	TableWalletRecharge     = "wallet_recharge"
	TableWalletStatement    = "wallet_statement"
)

type TableNameMap map[string]string

func MustSqlDB() *sql.DB {
	_onceSql.Do(func() {
		var err error
		if _, _sqldb, err = newSqlDB(); err != nil {
			logrus.Fatalf("new sql db failed: %s", err)
		}
	})
	return _sqldb
}

func newSqlDB() (driver string, db *sql.DB, err error) {
	if cfg.If("MySQL") {
		driver = "mysql"
		db, err = sql.Open(driver, MysqlSetting.Dsn())
		logrus.Infof("use MySQL as relation database")
	} else if cfg.Any("PostgreSQL", "PgSQL", "Postgres") {
		driver = "pgx"
		db, err = sql.Open(driver, PostgresSetting.Dsn())
		logrus.Infof("use PostgreSQL as relation database")
	} else if cfg.If("Sqlite3") {
		driver, db, err = OpenSqlite3()
		logrus.Infof("use Sqlite3 as relation database")
	} else {
		driver = "mysql"
		db, err = sql.Open(driver, MysqlSetting.Dsn())
		logrus.Infof("not set db fetuare so use MySQL as default relation database")
	}
	return
}
