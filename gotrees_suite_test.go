package gotrees_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGotrees(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gotrees Suite")
}
