package parser_test

import (
	"github.com/jtarchie/sqlc-gpt/parser"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Binding", func() {
	Describe("ArgName", func() {
		It("should return the name without the first character", func() {
			binding := parser.NewBinding("name", "placeholderValue")
			Expect(binding.ArgName()).To(Equal("name"))
		})
	})

	Describe("ArgType", func() {
		Context("when arg name is 'name' or 'value'", func() {
			It("should return 'string'", func() {
				binding := parser.NewBinding("name", "placeholderValue")
				Expect(binding.ArgType()).To(Equal("string"))
			})

			It("should return 'string'", func() {
				binding := parser.NewBinding("value", "placeholderValue")
				Expect(binding.ArgType()).To(Equal("string"))
			})
		})

		Context("when arg name is 'names'", func() {
			It("should return '[]string'", func() {
				binding := parser.NewBinding("names", "placeholderValue")
				Expect(binding.ArgType()).To(Equal("[]string"))
			})
		})

		Context("when arg name is neither 'name', 'value', nor 'names'", func() {
			It("should return 'int64'", func() {
				binding := parser.NewBinding("id", "placeholderValue")
				Expect(binding.ArgType()).To(Equal("int64"))
			})
		})
	})
})
