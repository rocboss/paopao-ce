// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/rocboss/paopao-ce/internal/conf"
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/sirupsen/logrus"
)

type metrics struct {
	siteInfo *prometheus.GaugeVec
	ds       core.DataService
	wc       core.WebCache
}

func (m *metrics) updateSiteInfo() {
	if onlineUserKeys, err := m.wc.Keys(conf.PrefixOnlineUser + "*"); err == nil {
		maxOnline := len(onlineUserKeys)
		m.siteInfo.With(prometheus.Labels{"name": "max_online"}).Set(float64(maxOnline))
	} else {
		logrus.Warnf("update promethues metrics[site_info_max_online] occurs error: %s", err)
	}
	if registerUserCount, err := m.ds.GetRegisterUserCount(); err == nil {
		m.siteInfo.With(prometheus.Labels{"name": "register_user_count"}).Set(float64(registerUserCount))
	} else {
		logrus.Warnf("update promethues metrics[site_info_register_user_count] occurs error: %s", err)
	}
}

func (m *metrics) onUpdate() {
	logrus.Debugf("update promethues metrics job running")
	m.updateSiteInfo()
}

func newMetrics(reg prometheus.Registerer, ds core.DataService, wc core.WebCache) *metrics {
	m := &metrics{
		siteInfo: prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Namespace: "paopao",
				Subsystem: "site",
				Name:      "simple_info",
				Help:      "paopao-ce site simple information.",
			},
			[]string{
				// metric name
				"name",
			}),
		ds: ds,
		wc: wc,
	}
	reg.MustRegister(m.siteInfo)
	return m
}
