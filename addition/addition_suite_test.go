package addition_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestAddition(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Addition Suite")
}
