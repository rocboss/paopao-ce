// Copyright 2023 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package types_test

import (
	"encoding/json"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/rocboss/paopao-ce/pkg/types"
)

var _ = Describe("JsonBox", Ordered, func() {
	type jsonCases []struct {
		j *types.JsonBox[json.RawMessage]
		b []byte
	}
	var samples jsonCases

	BeforeAll(func() {
		samples = jsonCases{
			{
				j: types.NewJsonBox(json.RawMessage(`null`)),
				b: []byte(`null`),
			},
			{
				j: types.NewJsonBox(json.RawMessage(`{}`)),
				b: []byte(`{}`),
			},
			{
				j: types.NewJsonBox(json.RawMessage(`[]`)),
				b: []byte(`[]`),
			},
			{
				j: types.NewJsonBox(json.RawMessage(`[{"b":true,"n":123},{"s":"foo","obj":{"f1":456,"f2":false}},[null]]`)),
				b: []byte(`[{"b":true,"n":123},{"s":"foo","obj":{"f1":456,"f2":false}},[null]]`),
			},
		}
	})

	It("boxes Box and Unbox", func() {
		for _, t := range samples {
			jv := types.NewJsonBox[json.RawMessage]()
			jv.Box(json.RawMessage(t.b))
			Expect(jv.Unbox()).To(Equal(t.j.Unbox()))
		}
	})

	It("json marshaler", func() {
		for _, t := range samples {
			mv, err := t.j.MarshalJSON()
			Expect(err).To(BeNil())
			Expect(mv).To(Equal(t.b))
		}
	})

	It("json unmarshaler", func() {
		for _, t := range samples {
			jv := types.NewJsonBox[json.RawMessage]()
			err := jv.UnmarshalJSON(t.b)
			Expect(err).To(BeNil())
			Expect(t.j.Unbox()).To(Equal(jv.Unbox()))
		}
	})

	It("driver valuer", func() {
		for _, t := range samples {
			v, err := t.j.Value()
			Expect(err).To(BeNil())
			Expect(v).To(Equal(t.b))
		}
	})

	It("sql scaner", func() {
		for _, t := range samples {
			jv := types.NewJsonBox[json.RawMessage]()
			err := jv.Scan(t.b)
			Expect(err).To(BeNil())
			Expect(jv.Unbox()).To(Equal(t.j.Unbox()))
		}
	})
})
