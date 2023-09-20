// This is a generated file, do not touch
// hash: 3a9eb2a6c62e1fd5fd8d309468b80432022bb62d

package code

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
)

func DeleteMany(ctx context.Context, db *sql.DB, names []string) ([]string, error) {
	const query = `
		DELETE FROM keys
		WHERE name IN (%s)
		RETURNING value`

	// Generate the placeholders for the IN clause
	placeholders := make([]string, len(names))
	for i := range names {
		placeholders[i] = "?"
	}
	inClause := fmt.Sprintf("(%s)", strings.Join(placeholders, ", "))

	sqlQuery := fmt.Sprintf(query, inClause)

	rows, err := db.QueryContext(ctx, sqlQuery, names)
	if err != nil {
		return nil, fmt.Errorf("DeleteMany: %v", err)
	}
	defer rows.Close()

	var values []string
	for rows.Next() {
		var value string
		if err := rows.Scan(&value); err != nil {
			return nil, fmt.Errorf("DeleteMany: %v", err)
		}
		values = append(values, value)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("DeleteMany: %v", err)
	}

	return values, nil
}
