package main

import (
	"fmt"
	"log/slog"
	"os"
	"text/template"

	"github.com/jtarchie/sqlc/parser"
	"github.com/alecthomas/kong"
)

type CLI struct {
	Glob        string `required:"" help:"glob to look for *.sql files"`
	Filename    string `required:"" help:"location to write the file"`
	PackageName string `required:"" default:"main" help:"the package name to the file"`
}

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stderr, nil))
	slog.SetDefault(logger)

	cli := &CLI{}
	ctx := kong.Parse(cli)

	err := ctx.Run()
	if err != nil {
		slog.Error("could not execute", slog.String("error", err.Error()))
	}
}

func (c *CLI) Run() error {
	queries, err := parser.Load(c.Glob)
	if err != nil {
		return fmt.Errorf("could not parse SQL files: %w", err)
	}

	err = queries.Validate()
	if err != nil {
		return fmt.Errorf("could not validate queries: %w", err)
	}

	data := map[string]interface{}{
		"PackageName":   c.PackageName,
		"ParsedQueries": queries,
	}

	t, err := template.New("queriesTemplate").Parse(tpl)
	if err != nil {
		return fmt.Errorf("could not parse template: %w", err)
	}

	file, err := os.Create(c.Filename)
	if err != nil {
		return fmt.Errorf("could create file: %w", err)
	}

	err = t.Execute(file, data)
	if err != nil {
		return fmt.Errorf("could execute template: %w", err)
	}

	return nil
}
