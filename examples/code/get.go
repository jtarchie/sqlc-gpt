// This is a generated file, do not touch
// hash: a10e7133552c779a462419230731102ae9611875

package code

import (
	"context"
	"database/sql"
	"fmt"
)

func GetOne(ctx context.Context, db *sql.DB, name string) (string, error) {
	const query = `SELECT value
	FROM keys
	WHERE name = ?`

	row := db.QueryRowContext(ctx, query, name)
	var value string
	err := row.Scan(&value)
	if err != nil {
		return "", fmt.Errorf("GetOne: failed to get value: %w", err)
	}

	return value, nil
}
