// This is a generated file, do not touch
// hash: 907b15794921d49e22ec965bdc091346f66f308f

package code

import (
	"context"
	"database/sql"
	"fmt"
)

func ListRange(ctx context.Context, db *sql.DB, name string, start int, end int) ([]string, error) {
	const query = `SELECT json_each.value
		FROM keys,
		  json_each(keys.value)
		WHERE keys.name = @name
		  AND json_each.key >= IIF(@start >= 0, @start, json_array_length(keys.value) + @start)
		  AND json_each.key <= IIF(@end >= 0, @end, json_array_length(keys.value) + @end)`
	rows, err := db.QueryContext(ctx, query, name, start, end)
	if err != nil {
		return nil, fmt.Errorf("ListRange: %w", err)
	}
	defer rows.Close()

	var result []string
	for rows.Next() {
		var value string
		err := rows.Scan(&value)
		if err != nil {
			return nil, fmt.Errorf("ListRange: %w", err)
		}
		result = append(result, value)
	}

	return result, nil
}
