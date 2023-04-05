package sakila_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestSakila(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Sakila Suite")
}
