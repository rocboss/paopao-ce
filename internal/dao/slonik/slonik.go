// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// Core service implement base sqlx+postgresql. All sub-service
// will declare here and provide initial function.

package slonik

import (
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/sirupsen/logrus"
)

func NewDataService() (core.DataService, core.VersionInfo) {
	logrus.Fatal("not support now")
	return nil, nil
}

func NewAuthorizationManageService() core.AuthorizationManageService {
	logrus.Fatal("not support now")
	return nil
}
