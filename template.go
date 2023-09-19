package main

const tpl = `
// Code generated. DO NOT EDIT.
package {{.PackageName}}
import (
  "database/sql"
  "fmt"
  "context"
  "io"

  "github.com/jmoiron/sqlx"
)

type Query interface {
  NamedExecContext(context.Context, string, interface{}) (sql.Result, error)
  GetContext(context.Context, interface{}, string, ...interface{}) error
  SelectContext(context.Context, interface{}, string, ...interface{}) error
  Rebind(string) string
}

type Queries struct{
  db      Query
}

func NewQueries(db *sql.DB) (*Queries) {
  sqlxDB := sqlx.NewDb(db, "sqlite")

  queries := &Queries{
    db: sqlxDB,
  }

  return queries
}

func (q *Queries) Close() error {
  closer, ok := q.db.(io.Closer)
  if !ok {
    return nil
  }

  err := closer.Close()
  if err != nil {
    return fmt.Errorf("could not close sqlx db: %w", err)
  }

  return nil
}

func (q *Queries) WithTx(tx *sqlx.Tx) (Query) {
  return &Queries{
    db: tx,
  }
}

{{range .ParsedQueries}}
// from {{.Filename}}:{{.Line}}
func (q *Queries) {{.Name}}(ctx context.Context,
  {{- range .Bindings -}}
    {{.ArgName}} {{.ArgType}},
  {{- end -}}
) (
  {{- if eq .Type "one" -}}
    map[string]interface{},
  {{- else if eq .Type "batchone" -}}
    []map[string]interface{},
  {{- else if eq .Type "batchmany" -}}
    []map[string]interface{},
  {{- else if eq .Type "many" -}}
    []map[string]interface{},
  {{- end -}}
  error) {
  {{- if eq .Type "exec" -}}
    _, err := q.db.NamedExecContext(ctx, {{printf "%q" .SQL}}, map[string]interface{}{
      {{- range .Bindings -}}
        {{printf "%q" .ArgName}}: {{.ArgName}},
      {{- end -}}
    })
    if err != nil {
      return fmt.Errorf("could not execute {{.Name}}: %w", err)
    }
    
    return nil
  {{else if eq .Type "one"}}
    payload := map[string]interface{}{}
    err := q.db.GetContext(ctx, &payload, {{printf "%q" .SQL}}, map[string]interface{}{
      {{- range .Bindings -}}
        {{printf "%q" .ArgName}}: {{.ArgName}},
      {{- end -}}
    })
    if err != nil {
      return nil, fmt.Errorf("could not execute {{.Name}}: %w", err)
    }
    return payload, nil
  {{else if eq .Type "many"}}
    payload := []map[string]interface{}{}
    err := q.db.SelectContext(ctx, &payload, {{printf "%q" .SQL}}, map[string]interface{}{
      {{- range .Bindings -}}
        {{printf "%q" .ArgName}}: {{.ArgName}},
      {{- end -}}
    })
    if err != nil {
      return nil, fmt.Errorf("could not execute {{.Name}}: %w", err)
    }
    return payload, nil
  {{- else if eq .Type "batchone" -}}
    payload := []map[string]interface{}{}

    query, _, err := sqlx.In({{printf "%q" .SQL}}, {{ (index .Bindings 0).ArgName }})
    if err != nil {
      return nil, fmt.Errorf("could not bind IN {{.Name}}: %w", err)
    }

    query = q.db.Rebind(query)
    
    err = q.db.SelectContext(ctx, &payload, query, map[string]interface{}{
      {{- range .Bindings -}}
        {{printf "%q" .ArgName}}: {{.ArgName}},
      {{- end -}}
    })
    if err != nil {
      return nil, fmt.Errorf("could not execute {{.Name}}: %w", err)
    }

    return payload, nil
  {{- else if eq .Type "batchmany" -}}
    payload := []map[string]interface{}{}

    query, _, err := sqlx.In({{printf "%q" .SQL}}, {{ (index .Bindings 0).ArgName }})
    if err != nil {
      return nil, fmt.Errorf("could not bind IN {{.Name}}: %w", err)
    }

    query = q.db.Rebind(query)
    
    err = q.db.SelectContext(ctx, &payload, query, map[string]interface{}{
      {{- range .Bindings -}}
        {{printf "%q" .ArgName}}: {{.ArgName}},
      {{- end -}}
    })
    if err != nil {
      return nil, fmt.Errorf("could not execute {{.Name}}: %w", err)
    }

    return payload, nil
  {{- else -}}
    return nil
  {{- end -}}
}
{{end}}
`
