// This is a generated file, do not touch
// hash: dc615b3f281344f543ae97b19eebca75656053d9

package code

import (
	"context"
	"database/sql"
	"fmt"
)

func AddFloat(ctx context.Context, db *sql.DB, name, value string) (float64, error) {
	const query = `
		INSERT INTO keys (name, value)
		VALUES (?, ?)
		ON CONFLICT(name) DO UPDATE
			SET value = CAST(value AS REAL) + CAST(excluded.value AS REAL)
		WHERE printf("%.17f", value) GLOB SUBSTRING(value, 1, 1) || '*'
		RETURNING CAST(value AS REAL) AS value
	`

	var result float64
	err := db.QueryRowContext(ctx, query, name, value).Scan(&result)
	if err != nil {
		err = fmt.Errorf("failed to execute AddFloat: %w", err)
		return 0, err
	}

	return result, nil
}
