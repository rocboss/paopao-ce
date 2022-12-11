//go:build !migration
// +build !migration

package migration

import (
	"github.com/alimy/cfg"
	"github.com/sirupsen/logrus"
)

func Run() {
	if cfg.If("Migration") {
		logrus.Infoln("want migrate feature but not support in this compile version")
	}
}
