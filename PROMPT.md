You are an AI assistant providing code for a code generation tool. Only output
the requested code, without additional explanations, introductions, conclusions,
or prose.

Generate a Go function for the "Queries" struct that executes the provided SQL
query using standard `database/sql` calls. The generated function should follow
these criteria:

1. The function should be named according to the query's name provided in the
   comment above the SQL. For example, if the name is "DeleteMany", the function
   should be `DeleteMany`.
2. The function should accept parameters: a `context.Context` and the necessary
   parameters for the SQL query, like a `struct` or individual variables, as
   required.
3. Embed the SQL query in the function as a constant string. Properly format the
   SQL within the Go code.
4. Execute the SQL query and return the results as specified in the SQL. If
   there's a `RETURNING` clause, return the results in the appropriate format.
5. Wrap any errors with context about the failure and the function's name.
   Return these as the second return value.
6. Do not use a prepared statement if the SQL has dynamic elements such as an
   `IN` clause.
7. Implement the function without adding any external functions or dependencies
   beyond the already provided standard libraries.
8. Only output the function, not extra external declarations required.
9. When a result has to be returned with more than two values a result struct
   needs to be created, ensure it has a prefix of the query name on it. For
   example, if the query name is `ListAll` the result struct would be
   `ListAllResult`. Output the definition of the struct type only if it will be used.
10. Inline all code, don't create external functions.
11. `sqlx` does not exist and function calls to `In` and `Rebind` should never
    be used!
12. Please keep number place holders in the queries.

Please provide the SQL query you want to convert into a Go function, and ensure
to include any necessary context and parameter details.

Given the SQL schema for SQLite:

```sql
CREATE TABLE IF NOT EXISTS keys (
  name TEXT NOT NULL PRIMARY KEY,
  value TEXT NOT NULL
);
```

Given this SQL query for SQLite:

```sql
-- name: ListRange :many
SELECT json_each.value
FROM keys,
  json_each(keys.value)
WHERE keys.name = @name
  AND json_each.key >= IIF(@start >= 0, @start, json_array_length(keys.value) + @start)
  AND json_each.key <= IIF(@end >= 0, @end, json_array_length(keys.value) + @end);
```

And this existing Go code:

```go
import (
  "database/sql"
  "fmt"
  "context"
)

type Queries struct {
    db *sql.DB
}
```

Generate the function according to these criteria.
