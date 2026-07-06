package parser_test

import (
	"github.com/divya/resume-analyser/internal/parser/mock"
	. "github.com/onsi/ginkgo/v2"
)

var _ = Describe("Registry", func() {

	Context("", func() {
		It("should return registered parser", func() {

			mock := &mock.MockParser{}

			Register("abc", ".abc", mock)

			p, err := Get(".abc")

			Expect(err).To(BeNil())

			Expect(p).To(Equal(mock))
		})
	})
})
