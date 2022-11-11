package localoss

import (
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/rocboss/paopao-ce/internal/conf"
	"github.com/sirupsen/logrus"
)

type localossSrv struct {
	// TODO
}

// RouteLocalOSS register LocalOSS route if needed
func RouteLocalOSS(e *gin.Engine) {
	if !conf.CfgIf("LocalOSS") {
		return
	}

	savePath, err := filepath.Abs(conf.LocalOSSSetting.SavePath)
	if err != nil {
		logrus.Fatalf("get localOSS save path err: %v", err)
	}
	e.Static("/oss", savePath)

	logrus.Infof("register LocalOSS route in /oss on save path: %s", savePath)
}
