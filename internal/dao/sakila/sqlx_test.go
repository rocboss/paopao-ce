package sakila_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/rocboss/paopao-ce/internal/dao/sakila"
)

var _ = Describe("Sqlx", Ordered, func() {
	type queries []struct {
		originQuery string
		fixedQuery  string
	}
	var samples queries

	BeforeAll(func() {
		samples = queries{
			{
				originQuery: `SELECT * FROM @user WHERE username=?@_`,
				fixedQuery:  `SELECT * FROM user WHERE username=?_`,
			},
			{
				originQuery: `SELECT * FROM @user WHERE username=?`,
				fixedQuery:  `SELECT * FROM user WHERE username=?`,
			},
			{
				originQuery: `SELECT * FROM @@user WHERE 用户名=?`,
				fixedQuery:  `SELECT * FROM @@user WHERE 用户名=?`,
			},
			{
				originQuery: `SELECT * FROM @@user, @@@@contact WHERE 用户名=?`,
				fixedQuery:  `SELECT * FROM @@user, @@@@contact WHERE 用户名=?`,
			},
			{
				originQuery: `SELECT @@name, @@@@@id FROM @@user, @@@@contact WHERE 用户名=?`,
				fixedQuery:  `SELECT @@name, @@@@@id FROM @@user, @@@@contact WHERE 用户名=?`,
			},
			{
				originQuery: `SELECT @name, @id FROM @user, @contact WHERE 用户名=?`,
				fixedQuery:  `SELECT name, id FROM user, contact WHERE 用户名=?`,
			},
		}
	})

	It("test internal t func", func() {
		for _, t := range samples {
			Expect(sakila.FnTest_t(t.originQuery)).To(Equal(t.fixedQuery))
		}
	})
})
