// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package utils_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/rocboss/paopao-ce/pkg/utils"
)

var _ = Describe("Md5", Ordered, func() {
	type md5Cases []struct {
		value string
		md5   string
	}
	var samples md5Cases

	BeforeAll(func() {
		samples = md5Cases{
			{
				value: "123456",
				md5:   "e10adc3949ba59abbe56e057f20f883e",
			},
			{
				value: "",
				md5:   "d41d8cd98f00b204e9800998ecf8427e", // really odd, why?
			},
			{
				value: "paopaocestr",
				md5:   "8a5033dda1a8919224c66e68d846a289",
			},
		}
	})

	It("encode md5", func() {
		for _, t := range samples {
			Expect(utils.EncodeMD5(t.value)).To(Equal(t.md5))
		}
	})

})
