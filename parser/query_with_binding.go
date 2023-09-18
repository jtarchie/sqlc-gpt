package parser

import (
	"regexp"
	"sort"
	"strings"
)

// QueryWithBinding represents the Ruby QueryWithBinding struct.
type QueryWithBinding struct {
	SQL      string
	Bindings []*Binding
}

// NewQueryWithBinding is a constructor function for QueryWithBinding.
func NewQueryWithBinding(sql string, bindings []*Binding) *QueryWithBinding {
	return &QueryWithBinding{SQL: sql, Bindings: bindings}
}

// SortedBindings sorts the Bindings slice based on the Name field.
func (qwb *QueryWithBinding) SortedBindings() []*Binding {
	sort.Slice(qwb.Bindings, func(i, j int) bool {
		return qwb.Bindings[i].Name < qwb.Bindings[j].Name
	})

	return qwb.Bindings
}

// ReturnedValues uses a regex to find all RETURNING or SELECT clauses and
// create ReturnedValue instances based on them.
func (qwb *QueryWithBinding) ReturnedValues() []*ReturnedValue {
	var results []*ReturnedValue

	partRegex := regexp.MustCompile(`(?ms).*SELECT(.*)FROM|.*RETURNING(.*;)`)
	matches := partRegex.FindAllStringSubmatch(qwb.SQL, -1)

	if len(matches) == 0 {
		return nil
	}

	statement := matches[0][1]
	if statement == "" {
		statement = matches[0][2]
	}
	statement = strings.TrimSpace(statement)

	fieldRegex := regexp.MustCompile(`(?ms)\s+AS\s+(\w+)\s*[,;]|^\s*(\w+)[,;]?$`)

	matches = fieldRegex.FindAllStringSubmatch(statement, -1)
	for _, field := range matches {
		value := field[1]
		if value == "" {
			value = field[2]
		}

		results = append(results, NewReturnedValue(value))
	}

	return results
}
