// This is a generated file, do not touch
// hash: 66971ae7e77c82d97b01b7d70902f3d0b6de4f50

package code

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
)

type GetManyResult struct {
	Name  []string
	Value []string
}

func GetMany(ctx context.Context, db *sql.DB, names []string) (GetManyResult, error) {
	query := `
        SELECT name, value
        FROM keys
        WHERE name IN (%s)`
	query = fmt.Sprintf(query, "'"+strings.Join(names, "','")+"'")

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return GetManyResult{}, fmt.Errorf("GetMany: failed to execute query: %w", err)
	}
	defer rows.Close()

	var result GetManyResult
	for rows.Next() {
		var name string
		var value string
		if err := rows.Scan(&name, &value); err != nil {
			return result, fmt.Errorf("GetMany: failed to scan row: %w", err)
		}

		result.Name = append(result.Name, name)
		result.Value = append(result.Value, value)
	}

	if rows.Err() != nil {
		return result, fmt.Errorf("GetMany: unexpected error while iterating rows: %w", rows.Err())
	}

	return result, nil
}
