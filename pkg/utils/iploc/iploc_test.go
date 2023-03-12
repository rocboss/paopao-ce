// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package iploc_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/rocboss/paopao-ce/pkg/utils/iploc"
)

var _ = Describe("Iploc", Ordered, func() {
	type iplocCases []struct {
		ip      string
		country string
		city    string
	}
	var samples iplocCases

	BeforeAll(func() {
		samples = iplocCases{
			{
				ip:      "127.0.0.1",
				country: "本机地址",
				city:    " CZ88.NET",
			},
			{
				ip:      "180.89.94.9",
				country: "北京市",
				city:    "鹏博士宽带",
			},
		}
	})

	It("find country and city by ip", func() {
		for _, t := range samples {
			country, city := iploc.Find(t.ip)
			Expect(country).To(Equal(t.country))
			Expect(city).To(Equal(t.city))
		}
	})
})
