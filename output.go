package main

const tpl = `
// Code generated. DO NOT EDIT.
package {{.PackageName}}
import (
  "database/sql"
  "fmt"
  "context"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

type Queries struct{
  db      DBTX
	tx      *sql.Tx
  prepare map[string]*sql.Stmt
  queries map[string]string
}

func New(db DBTX) (*Queries, error) {
  queries := &Queries{
    db: db,
  }

  setupQueries(queries)

  return queries, nil
}

func setupQueries(q *Queries) {
  q.queries = map[string]string{}
  {{range .ParsedQueries}}
    // from {{.Filename}}:{{.Line}}
    q.queries[{{printf "%q" .Name}}] = ` + "`" + `
  {{.SQLWithBindings.SQL}}` + "`" + `
  {{end}}
}

func (q *Queries) PrepareWithContext(ctx context.Context) error {
  var err error
  {{range .ParsedQueries}}
    {{if .Prepared}}
      // from {{.Filename}}:{{.Line}}
      if q.prepare[{{printf "%q" .Name}}], err = q.db.PrepareContext(ctx, q.queries[{{printf "%q" .Name}}]); err != nil {
        return fmt.Errorf("error preparing query {{.Name}}: %w", err)
      }
    {{end}}
  {{end}}

  return nil
}

func (q *Queries) Close() error {
  for name, query := range q.prepare {
    if err := query.Close(); err != nil {
      return fmt.Errorf("error closing %q: %w", name, err)
    }
  }

  return nil
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
  return &Queries{
    db:                      tx,
    tx:                      tx,
    queries: q.queries,
    prepare: q.prepare,
  }
}

{{range .ParsedQueries}}
// from {{.Filename}}:{{.Line}}
func (q *Queries) {{.Name}}(ctx context.Context,
  {{- range .SQLWithBindings.SortedBindings -}}
    {{.ArgName}} {{.ArgType}},
  {{- end -}}
) (
  {{- range .SQLWithBindings.ReturnedValues -}}
    {{.ArgType}},
  {{- end -}}
  error) {
  {{if eq .Type "exec"}}
    _, err := q.exec(
      ctx,
      q.prepare[{{printf "%q" .Name}}],
      q.queries[{{printf "%q" .Name}}],
      {{range .SQLWithBindings.SortedBindings}}
        {{.ArgName}},
      {{end}}
    )
    if err != nil {
      return fmt.Errorf("could not execute {{.Name}}: %w", err)
    }
    return nil
  {{else if eq .Type "one"}}
    row := q.queryRow(
      ctx,
      q.prepare[{{printf "%q" .Name}}],
      q.queries[{{printf "%q" .Name}}],
      {{range .SQLWithBindings.SortedBindings}}
        {{.ArgName}},
      {{end}}
    )

    {{ range .SQLWithBindings.ReturnedValues }}
      var {{.ArgName}} {{.ArgType}}
    {{ end }}

    err := row.Scan(
      {{ range .SQLWithBindings.ReturnedValues }}
      &{{.ArgName}},
      {{ end }}
    )

    if err != nil {
      return {{range .SQLWithBindings.ReturnedValues}}{{.DefaultValue}},{{end}}fmt.Errorf("could not query row {{.Name}}: %w", err)
    }

    return {{range .SQLWithBindings.ReturnedValues}}{{.ArgName}},{{end}}nil
  {{else if eq .Type "many"}}

  {{else}}
    return {{range .SQLWithBindings.ReturnedValues}}{{.DefaultValue}},{{end}}nil
  {{end}}
}
{{end}}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
  switch {
  case stmt != nil && q.tx != nil:
    return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
  case stmt != nil:
    return stmt.ExecContext(ctx, args...)
  default:
    return q.db.ExecContext(ctx, query, args...)
  }
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
  switch {
  case stmt != nil && q.tx != nil:
    return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
  case stmt != nil:
    return stmt.QueryContext(ctx, args...)
  default:
    return q.db.QueryContext(ctx, query, args...)
  }
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
  switch {
  case stmt != nil && q.tx != nil:
    return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
  case stmt != nil:
    return stmt.QueryRowContext(ctx, args...)
  default:
    return q.db.QueryRowContext(ctx, query, args...)
  }
}
`
