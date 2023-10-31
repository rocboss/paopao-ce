// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package types_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/rocboss/paopao-ce/pkg/types"
)

var _ = Describe("Bitmap", Ordered, func() {
	It("sql scaner for Bitmap", func() {
		bitmap := types.NewBitmap()
		bitmap.Data.Map.Add(128)
		value, err := bitmap.Value()
		Expect(err).To(BeNil())

		var bm types.Bitmap
		bm.Scan(value)
		Expect(bm.Data.Map.Contains(128)).To(BeTrue())
	})

	It("sql scaner for Bitmap64", func() {
		bitmap := types.NewBitmap64()
		bitmap.Data.Map.Add(128)
		value, err := bitmap.Value()
		Expect(err).To(BeNil())

		var bm types.Bitmap64
		bm.Scan(value)
		Expect(bm.Data.Map.Contains(128)).To(BeTrue())
	})
})
