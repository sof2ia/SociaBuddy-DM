package comment

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"testing"
)

func TestComService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Comment Service Suite")
}
