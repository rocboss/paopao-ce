// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package chain

import (
	"github.com/rocboss/paopao-ce/internal/conf"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/infra/metrics"
)

type OnlineUserMetric struct {
	metrics.BaseMetric
	ac     core.AppCache
	uid    int64
	expire int64
}

func OnUserOnlineMetric(ac core.AppCache, uid int64) {
	metrics.OnMeasure(&OnlineUserMetric{
		ac:     ac,
		uid:    uid,
		expire: conf.CacheSetting.OnlineUserExpire,
	})
}

func (m *OnlineUserMetric) Name() string {
	return "OnlineUserMetric"
}

func (m *OnlineUserMetric) Action() (err error) {
	// 暂时仅做标记，不存储其他相关信息
	m.ac.SetNx(conf.KeyOnlineUser.Get(m.uid), []byte{}, m.expire)
	return
}
