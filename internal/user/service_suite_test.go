package user

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"testing"
)

func TestUserService(t *testing.T) {
	RegisterFailHandler(Fail)

	RunSpecs(t, "User Service Suite")
}
