package falamos_test

import (
	"context"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var ctx = context.Background()

func TestFalamos(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Falamos Suite")
}
