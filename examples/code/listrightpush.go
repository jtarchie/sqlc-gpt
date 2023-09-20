// This is a generated file, do not touch
// hash: b45b6682dd39c5b0231f26760ac4b65e975e8ffe

package code

import (
	"context"
	"database/sql"
)

type ListRightPushResult struct {
	Valid  bool
	Length int
}

func ListRightPush(ctx context.Context, db *sql.DB, name, value string) ([]ListRightPushResult, error) {
	const query = `
		UPDATE keys
		SET value = json_insert(value, '$[#]', ?)
		WHERE name = ?
		RETURNING json_valid(value) AS valid,
			json_array_length(value) AS length
	`
	rows, err := db.QueryContext(ctx, query, value, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []ListRightPushResult
	for rows.Next() {
		var result ListRightPushResult
		err := rows.Scan(&result.Valid, &result.Length)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}

	return results, nil
}
