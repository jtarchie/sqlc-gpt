package main

import (
	"fmt"
	"log/slog"
	"os"
	"text/template"

	"github.com/jtarchie/sqlc/parser"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stderr, nil))
	slog.SetDefault(logger)

	err := execute()
	if err != nil {
		slog.Error("could not execute", slog.String("error", err.Error()))
	}
}

func execute() error {
	queries, err := parser.Load("/Users/jtarchie/workspace/sqlettus/db/drivers/sqlite/*.sql")
	if err != nil {
		return fmt.Errorf("could not parse SQL files: %w", err)
	}

	err = queries.Validate()
	if err != nil {
		return fmt.Errorf("could not validate queries: %w", err)
	}

	data := map[string]interface{}{ // Fill in this data from the parsed SQL files
		"PackageName":   "blah",
		"ParsedQueries": queries,
	}

	t, err := template.New("queriesTemplate").Parse(tpl)
	if err != nil {
		return fmt.Errorf("could not parse template: %w", err)
	}

	file, err := os.Create("test/test.go")
	if err != nil {
		return fmt.Errorf("could create file: %w", err)
	}

	err = t.Execute(file, data)
	if err != nil {
		return fmt.Errorf("could execute template: %w", err)
	}

	return nil
}
