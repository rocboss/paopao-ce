// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package core

import (
	"github.com/rocboss/paopao-ce/internal/dao/jinzhu/dbr"
)

const (
	UserStatusNormal = dbr.UserStatusNormal
	UserStatusClosed = dbr.UserStatusClosed
)

type (
	User                = dbr.User
	Post                = dbr.Post
	ConditionsT         = dbr.ConditionsT
	PostFormated        = dbr.PostFormated
	UserFormated        = dbr.UserFormated
	PostContentFormated = dbr.PostContentFormated
	Model               = dbr.Model
)
