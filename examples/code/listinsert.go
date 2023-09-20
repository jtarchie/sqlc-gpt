// This is a generated file, do not touch
// hash: 32b536521555421f025ae0117a8038a69d3807a3

package code

import (
	"context"
	"database/sql"
	"fmt"
)

type ListInsertResult struct {
	Value int `json:"value"`
}

func ListInsert(ctx context.Context, db *sql.DB, pivot, offset int, value, name string) (ListInsertResult, error) {
	sql := `
		UPDATE keys
		SET value = json_array_insert(value, ?, ?, ?)
		WHERE name = ?
		RETURNING json_array_length(value) AS value
	`

	rows, err := db.QueryContext(ctx, sql, pivot, offset, value, name)
	if err != nil {
		return ListInsertResult{}, fmt.Errorf("ListInsert: %v", err)
	}
	defer rows.Close()

	var result ListInsertResult
	if rows.Next() {
		err := rows.Scan(&result.Value)
		if err != nil {
			return ListInsertResult{}, fmt.Errorf("ListInsert: %v", err)
		}
	}

	return result, nil
}
