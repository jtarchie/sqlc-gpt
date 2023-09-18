package parser

import (
	"regexp"
	"strconv"
)

type ParsedQuery struct {
	Name     string
	Type     string
	SQL      string
	Filename string
	Line     int
}

func NewParsedQuery(name, type_, sql, filename string, line int) *ParsedQuery {
	return &ParsedQuery{name, type_, sql, filename, line}
}

func (pq *ParsedQuery) SQLWithBindings() *QueryWithBinding {
	namedToIndexed := make(map[string]string)
	index := 0

	// Replace named parameters with indexed placeholders
	regex := regexp.MustCompile(`@(\w+)`)
	modifiedSQL := regex.ReplaceAllStringFunc(pq.SQL, func(match string) string {
		// If the named parameter has not been seen before, increase the index and store the mapping
		name := match[1:]
		if _, exists := namedToIndexed[name]; !exists {
			index++
			namedToIndexed[name] = "$" + strconv.Itoa(index)
		}
		return namedToIndexed[name]
	})

	// Convert map to bindings
	var bindings []*Binding
	for name, placeholder := range namedToIndexed {
		bindings = append(bindings, NewBinding(name, placeholder))
	}

	return NewQueryWithBinding(modifiedSQL, bindings)
}

func (pq *ParsedQuery) Prepared() bool {
	regex := regexp.MustCompile(`IN \(@\w+\)`)
	return !regex.MatchString(pq.SQL)
}
