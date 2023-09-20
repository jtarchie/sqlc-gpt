// This is a generated file, do not touch
// hash: 25c058322044f34b97935558743cc7e29bd75e51

package code

import (
	"context"
	"database/sql"
	"fmt"
)

func Substr(ctx context.Context, db *sql.DB, name string, start, end int) (string, error) {
	const query = `
	SELECT SUBSTR(
		value,
		IIF(? < 0, ?, ? + 1),
		IIF(
			? < 0,
			LENGTH(value) - ?,
			? + ? + 1
		)
	)
	FROM keys
	WHERE name = ?`
	var result string
	err := db.QueryRowContext(ctx, query, start, start, start, end, end, start, end, name).Scan(&result)
	if err != nil {
		return "", fmt.Errorf("Substr: %w", err)
	}
	return result, nil
}
