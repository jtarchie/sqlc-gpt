// This is a generated file, do not touch
// hash: 5f5cfded02bc9a76348e39918a503bff27029c0f

package code

import (
	"context"
	"database/sql"
	"fmt"
)

type ListLengthResult struct {
	Value int `db:"value"`
}

func ListLength(ctx context.Context, db *sql.DB, name string) (ListLengthResult, error) {
	const query = `SELECT json_array_length(value) AS value
	FROM keys
	WHERE name = ?`

	row := db.QueryRowContext(ctx, query, name)

	var result ListLengthResult
	if err := row.Scan(&result.Value); err != nil {
		return ListLengthResult{}, fmt.Errorf("ListLength: %w", err)
	}

	return result, nil
}
