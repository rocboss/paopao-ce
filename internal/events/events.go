// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package events

import (
	"sync"

	"github.com/alimy/tryst/cfg"
	"github.com/sirupsen/logrus"
)

var (
	_onceInitial sync.Once
)

func Initial() {
	_onceInitial.Do(func() {
		initEventManager()
		if cfg.If("UseJobManager") {
			initJobManager()
			logrus.Debugln("initial JobManager")
		}
	})
}
