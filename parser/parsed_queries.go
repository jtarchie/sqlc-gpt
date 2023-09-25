package parser

import (
	"fmt"
)

type ParsedQueries []*ParsedQuery

func NewParsedQueries(queries []*ParsedQuery) ParsedQueries {
	return ParsedQueries(queries)
}

// Range allows you to iterate over each of the parsed queries.
func (pqs ParsedQueries) Range(f func(pq *ParsedQuery)) {
	for _, pq := range pqs {
		f(pq)
	}
}

// Validate checks if any query name appears more than once and raises an error if it does.
func (pqs ParsedQueries) Validate() error {
	nameCounts := make(map[string]int)
	for _, pq := range pqs {
		nameCounts[pq.Name]++
	}

	for key, count := range nameCounts {
		if count > 1 {
			return fmt.Errorf("the query name %s appears %d times", key, count)
		}
	}

	return nil
}
