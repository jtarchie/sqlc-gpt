// This is a generated file, do not touch
// hash: b4b7e156d60cad0f46523771f1774fca0b28a9c4

package code

import (
	"context"
	"database/sql"
	"fmt"
)

// KeysAppendValueResult represents the result of the KeysAppendValue function.
type KeysAppendValueResult struct {
	ValueLength int
}

// KeysAppendValue inserts a new row into the "keys" table with the provided name
// and value. If a row with the same name already exists, it updates the value
// by concatenating the new value with the existing one. It returns the length
// of the updated value or an error if any occurred.
func KeysAppendValue(ctx context.Context, db *sql.DB, name, value string) (KeysAppendValueResult, error) {
	const query = `INSERT INTO keys (name, value)
	VALUES (?, ?)
	ON CONFLICT(name) DO
	UPDATE
	SET value = value || excluded.value
	RETURNING length(value)`

	result := KeysAppendValueResult{}

	err := db.QueryRowContext(ctx, query, name, value).Scan(&result.ValueLength)
	if err != nil {
		return result, fmt.Errorf("KeysAppendValue: %w", err)
	}

	return result, nil
}
