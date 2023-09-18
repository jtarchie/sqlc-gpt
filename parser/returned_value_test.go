package parser_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/jtarchie/sqlc/parser"
)

var _ = Describe("ReturnedValue", func() {

	Describe("ArgName", func() {
		It("should always return the string 'name'", func() {
			returnedValue := parser.NewReturnedValue("someStatement")
			Expect(returnedValue.ArgName()).To(Equal("someStatementRet"))
		})
	})

	Describe("ArgType", func() {
		It("should always return 'string' since ArgName always returns 'name'", func() {
			returnedValue := parser.NewReturnedValue("someStatement")
			Expect(returnedValue.ArgType()).To(Equal("int64"))
		})
	})
})
