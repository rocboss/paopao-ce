// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package web

import (
	"github.com/alimy/tryst/cfg"
	"github.com/robfig/cron/v3"
	"github.com/rocboss/paopao-ce/internal/conf"
	"github.com/rocboss/paopao-ce/internal/events"
	"github.com/sirupsen/logrus"
)

func onMaxOnlineJob() {
	spec := conf.JobManagerSetting.MaxOnlineInterval
	schedule, err := cron.ParseStandard(spec)
	if err != nil {
		panic(err)
	}
	events.OnTask(schedule, func() {
		onlineUserKeys, err := _wc.Keys(conf.PrefixOnlineUser + "*")
		if maxOnline := len(onlineUserKeys); err == nil && maxOnline > 0 {
			if _, err = _wc.PutHistoryMaxOnline(maxOnline); err != nil {
				logrus.Warnf("onMaxOnlineJob[2] occurs error: %s", err)
			}
		} else if err != nil {
			logrus.Warnf("onMaxOnlineJob[1] occurs error: %s", err)
		}
	})
}

func scheduleJobs() {
	cfg.Not("DisableJobManager", func() {
		lazyInitial()
		onMaxOnlineJob()
		logrus.Debug("schedule inner jobs complete")
	})
}
