// This is a generated file, do not touch
// hash: ecac154d8e6329148d2265c7b4301b9008900fb6

package code

import (
	"context"
	"database/sql"
	"fmt"
)

func Set(ctx context.Context, db *sql.DB, name, value string) error {
	const query = `
INSERT INTO keys (name, value)
VALUES (?, ?)
ON CONFLICT(name) DO UPDATE
SET value = excluded.value;`

	_, err := db.ExecContext(ctx, query, name, value)
	if err != nil {
		return fmt.Errorf("Set: %v", err)
	}

	return nil
}
