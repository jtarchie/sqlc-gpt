// This is a generated file, do not touch
// hash: c905fd97dcc7700655e8d051ed927335f9f496f3

package code

import (
	"context"
	"database/sql"
)

func ListSet(ctx context.Context, db *sql.DB, name string, index int, value string) (*ListSetResult, error) {
	const query = `
		UPDATE keys
		SET value = json_replace(
			value,
			'$[' || IIF($3 >= 0, $3, '#' || $3) || ']',
			$4
		)
		WHERE name = $1
		RETURNING json_valid(value);
	`

	row := db.QueryRowContext(ctx, query, name, index, value)

	var result ListSetResult
	err := row.Scan(&result.Valid)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

type ListSetResult struct {
	Valid bool `json:"valid"`
}
