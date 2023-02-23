// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package convert_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/rocboss/paopao-ce/pkg/convert"
)

var _ = Describe("Convert", func() {
	It("str to string", func() {
		for _, t := range []struct {
			str  string
			want string
		}{
			{
				str:  "test",
				want: "test",
			},
		} {
			Expect(convert.StrTo(t.str).String()).To(Equal(t.want))
		}
	})

	It("str to int", func() {
		var s convert.StrTo = "123"
		i, err := s.Int()
		Expect(err).To(BeNil())
		Expect(i).To(Equal(123))
	})

	It("str must int", func() {
		var s convert.StrTo = "123"
		Expect(s.MustInt()).To(Equal(int(123)))
	})

	It("str to uint32", func() {
		s := convert.StrTo("123")
		i, err := s.UInt32()
		Expect(err).To(BeNil())
		Expect(i).To(Equal(uint32(123)))
	})

	It("str to uint32 negative", func() {
		s := convert.StrTo("-123")
		i, err := s.UInt32()
		Expect(err).To(BeNil())
		Expect(i).NotTo(Equal(uint32(123)))
	})

	It("str to int64", func() {
		var s convert.StrTo = "123"
		i, err := s.Int64()
		Expect(err).To(BeNil())
		Expect(i).To(Equal(int64(123)))
	})

	It("str must int64", func() {
		var s convert.StrTo = "123"
		Expect(s.MustInt64()).To(Equal(int64(123)))
	})

	It("str to float64", func() {
		var s convert.StrTo = "123.456"
		f, err := s.Float64()
		Expect(err).To(BeNil())
		Expect(f).To(Equal(float64(123.456)))
	})

	It("str must float64", func() {
		var s convert.StrTo = "123.456"
		Expect(s.Float64()).To(Equal(float64(123.456)))
	})
})
