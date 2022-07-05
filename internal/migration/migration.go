//go:build !migration
// +build !migration

package migration

import (
	"github.com/rocboss/paopao-ce/internal/conf"
	"github.com/sirupsen/logrus"
)

func Run() {
	if conf.CfgIf("Migration") {
		logrus.Infoln("want migrate feature but not support in this compile version")
	}
}
