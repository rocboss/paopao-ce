// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

//go:build !pyroscope
// +build !pyroscope

package debug

import (
	"github.com/alimy/cfg"
	"github.com/sirupsen/logrus"
)

func StartPyroscope() {
	if cfg.If("Pyroscope") {
		logrus.Infoln("want Pyroscope feature but not support in this compile version")
	}
}
