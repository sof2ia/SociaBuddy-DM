package comment

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"testing"
)

func TestPostService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Post Service Suite")
}
