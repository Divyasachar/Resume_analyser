package ats_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestAts(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Ats Suite")
}
