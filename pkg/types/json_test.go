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

var _ = Describe("Json", Ordered, func() {
	type jsonCases []struct {
		j types.JsonType[json.RawMessage]
		b []byte
	}
	var samples jsonCases

	BeforeAll(func() {
		samples = jsonCases{
			{
				j: types.JsonType[json.RawMessage]{json.RawMessage(`null`)},
				b: []byte(`null`),
			},
			{
				j: types.JsonType[json.RawMessage]{json.RawMessage(`{}`)},
				b: []byte(`{}`),
			},
			{
				j: types.JsonType[json.RawMessage]{json.RawMessage(`[]`)},
				b: []byte(`[]`),
			},
			{
				j: types.JsonType[json.RawMessage]{json.RawMessage(`[{"b":true,"n":123},{"s":"foo","obj":{"f1":456,"f2":false}},[null]]`)},
				b: []byte(`[{"b":true,"n":123},{"s":"foo","obj":{"f1":456,"f2":false}},[null]]`),
			},
		}
	})

	It("driver valuer ", func() {
		for _, t := range samples {
			v, err := t.j.Value()
			Expect(err).To(BeNil())
			Expect(v).To(Equal(t.b))
		}
	})

	It("sql scan ", func() {
		for _, t := range samples {
			var jv types.JsonType[json.RawMessage]
			err := jv.Scan(t.b)
			Expect(err).To(BeNil())
			Expect(jv.Data).To(Equal(t.j.Data))
		}
	})
})
