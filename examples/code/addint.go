// This is a generated file, do not touch
// hash: 5de70026f10210f6b084e5a73428b6f197501329

package code

import (
	"context"
	"database/sql"
	"fmt"
)

func AddInt(ctx context.Context, name string, value string, db *sql.DB) (int, error) {
	const query = `
		INSERT INTO keys (name, value)
		VALUES (?, ?)
		ON CONFLICT(name) DO
		UPDATE
		SET value = CAST(value AS INTEGER) + CAST(excluded.value AS INTEGER)
		WHERE printf("%d", value) = value
		RETURNING CAST(value AS INTEGER)
	`

	var v int
	err := db.QueryRowContext(ctx, query, name, value).Scan(&v)
	if err != nil {
		return 0, fmt.Errorf("code.AddInt: %v", err)
	}

	return v, nil
}
