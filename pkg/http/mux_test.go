// Copyright 2023 Michael Li <alimy@gility.net>. All rights reserved.
// Use of this source code is governed by Apache License 2.0 that
// can be found in the LICENSE file.

package http

import (
	g "github.com/onsi/ginkgo/v2"
	m "github.com/onsi/gomega"
)

var _ = g.Describe("Mux", g.Ordered, func() {
	var smm muxMap[int]
	var pmm muxMap[int]

	g.BeforeAll(func() {
		smm = make(simpleMuxMap[int])
		pmm = &prefixMuxMap[int]{
			m:      make(map[string]int),
			prefix: "/connect",
		}
	})

	g.It("simple mux map", func() {
		ok := smm.set("/core.v1.AuthenticateService/", 1)
		m.Expect(ok).To(m.BeTrue())

		ok = smm.set("/core.v1.AuthenticateService/", 2)
		m.Expect(ok).To(m.BeFalse())

		smm.set("/greet.v1.GreetService/", 2)
		val, exist := smm.get("/greet.v1.GreetService/")
		m.Expect(val).To(m.Equal(2))
		m.Expect(exist).To(m.BeTrue())

		_, exist = smm.get("/greet.v1.OtherService/")
		m.Expect(exist).To(m.BeFalse())

		val, exist = smm.match("/core.v1.AuthenticateService/login")
		m.Expect(val).To(m.Equal(1))
		m.Expect(exist).To(m.BeTrue())

		val, exist = smm.match("/core.v1.AuthenticateService/logout")
		m.Expect(val).To(m.Equal(1))
		m.Expect(exist).To(m.BeTrue())
	})

	g.It("prefix mux map", func() {
		ok := pmm.set("/core.v1.AuthenticateService/", 1)
		m.Expect(ok).To(m.BeTrue())

		ok = pmm.set("/core.v1.AuthenticateService/", 2)
		m.Expect(ok).To(m.BeFalse())

		pmm.set("/greet.v1.GreetService/", 2)
		val, exist := pmm.get("/greet.v1.GreetService/")
		m.Expect(val).To(m.Equal(2))
		m.Expect(exist).To(m.BeTrue())

		_, exist = pmm.get("/greet.v1.OtherService/")
		m.Expect(exist).To(m.BeFalse())

		val, exist = pmm.match("/connect/core.v1.AuthenticateService/login")
		m.Expect(val).To(m.Equal(1))
		m.Expect(exist).To(m.BeTrue())

		val, exist = pmm.match("/connect/core.v1.AuthenticateService/logout")
		m.Expect(val).To(m.Equal(1))
		m.Expect(exist).To(m.BeTrue())
	})
})
