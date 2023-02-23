// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package util_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/rocboss/paopao-ce/pkg/util"
)

var _ = Describe("Ip", Ordered, func() {
	type iplocs []struct {
		ip       string
		location string
	}
	var samples iplocs

	BeforeAll(func() {
		samples = iplocs{
			{
				ip:       "",
				location: "",
			},
			{
				ip:       "103.197.70.244",
				location: "香港",
			},
		}
	})

	It("get ip location", func() {
		for _, t := range samples {
			Expect(util.GetIPLoc(t.ip)).To(Equal(t.location))
		}
	})
})
