package gw2api_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGw2api(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gw2api Suite")
}
