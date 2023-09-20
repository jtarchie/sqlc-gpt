// This is a generated file, do not touch
// hash: 97c789afc18c219ea8895de438f70072513ea6fb

package code

import (
	"context"
	"database/sql"
	"fmt"
)

type FlushAllResult struct {
	RowsAffected int64
}

func FlushAll(ctx context.Context, db *sql.DB) (FlushAllResult, error) {
	const query = `
		DELETE FROM keys;
	`

	_, err := db.ExecContext(ctx, query)
	if err != nil {
		return FlushAllResult{}, fmt.Errorf("FlushAll: failed to execute query: %w", err)
	}

	return FlushAllResult{}, nil
}
