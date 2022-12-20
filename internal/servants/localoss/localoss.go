// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package localoss

import (
	"path/filepath"

	"github.com/gin-gonic/gin"
	api "github.com/rocboss/paopao-ce/auto/api/s/v1"
	"github.com/rocboss/paopao-ce/internal/conf"
	"github.com/sirupsen/logrus"
)

// RouteLocalOSS register LocalOSS route if needed
func RouteLocalOSS(e *gin.Engine) {
	savePath, err := filepath.Abs(conf.LocalOSSSetting.SavePath)
	if err != nil {
		logrus.Fatalf("get localOSS save path err: %v", err)
	}
	e.Static("/oss", savePath)

	logrus.Infof("register LocalOSS route in /oss on save path: %s", savePath)
}

// RouteLocaloss register LocalOSS route if needed
func RouteLocaloss(e *gin.Engine) {
	api.RegisterUserServant(e, newUserSrv(), newUserBinding(), newUserRender())
}
