package parser_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/jtarchie/sqlc/parser"
)

var _ = Describe("QueryWithBinding", func() {
	Describe("ReturnedValues", func() {
		Context("when the SQL query contains a RETURNING clause", func() {
			It("should extract the RETURNING values from the ListRightPushUpsert query", func() {
				sql := `
INSERT INTO keys (name, value)
VALUES (@name, json_insert('[]', '$[#]', @value)) ON CONFLICT(name) DO
UPDATE
SET value = json_insert(
    value,
    '$[#]',
    json_extract(excluded.value, '$[0]')
  )
RETURNING CAST(json_valid(value) AS boolean) AS valid,
  CAST(json_array_length(value) AS INTEGER) AS length;
`
				query := parser.NewQueryWithBinding(sql, nil)
				returnedValues := query.ReturnedValues()
				Expect(len(returnedValues)).To(Equal(2))
				Expect(returnedValues[0].Statement).To(Equal("valid"))
				Expect(returnedValues[1].Statement).To(Equal("length"))
			})

			It("should extract the RETURNING values from the DeleteMany query", func() {
				sql := `
DELETE FROM keys
WHERE name IN (@names)
RETURNING value;
`
				query := parser.NewQueryWithBinding(sql, nil)
				returnedValues := query.ReturnedValues()
				Expect(len(returnedValues)).To(Equal(1))
				Expect(returnedValues[0].Statement).To(Equal("value"))
			})
		})

		Context("when the SQL query contains a SELECT clause", func() {
			It("should extract the SELECT values from the GetMany query", func() {
				sql := `
SELECT name,
  value
FROM keys
WHERE name IN (@names);
`
				query := parser.NewQueryWithBinding(sql, nil)
				returnedValues := query.ReturnedValues()
				Expect(len(returnedValues)).To(Equal(2))
				Expect(returnedValues[0].Statement).To(Equal("name"))
				Expect(returnedValues[1].Statement).To(Equal("value"))
			})
		})
	})
})
