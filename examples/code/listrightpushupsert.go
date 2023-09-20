// This is a generated file, do not touch
// hash: 6e0d340ecd9528099b43328c8ca87f5f59c8e9f2

package code

import (
	"context"
	"database/sql"
)

type ListRightPushUpsertResult struct {
	Valid  bool `db:"valid"`
	Length int  `db:"length"`
}

func ListRightPushUpsert(ctx context.Context, db *sql.DB, name string, value string) (ListRightPushUpsertResult, error) {
	const query = `
	INSERT INTO keys (name, value)
	VALUES (?, json_insert('[]', '$[#]', ?))
	ON CONFLICT(name) DO UPDATE
	SET value = json_insert(
		value,
		'$[#]',
		json_extract(excluded.value, '$[0]')
	)
	RETURNING CAST(json_valid(value) AS boolean) AS valid,
		CAST(json_array_length(value) AS INTEGER) AS length;
	`

	var result ListRightPushUpsertResult
	row := db.QueryRowContext(ctx, query, name, value)
	err := row.Scan(&result.Valid, &result.Length)
	if err != nil {
		return ListRightPushUpsertResult{}, err
	}

	return result, nil
}
