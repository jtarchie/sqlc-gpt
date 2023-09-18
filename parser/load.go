package parser

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
)

func Load(glob string) (ParsedQueries, error) {
	var queries []*ParsedQuery

	matches, err := filepath.Glob(glob)
	if err != nil {
		return nil, fmt.Errorf("could not load queries: %w", err)
	}

	if len(matches) == 0 {
		return nil, fmt.Errorf("could not find any files")
	}

	for _, filename := range matches {
		slog.Info("loading file", slog.String("filename", filename))
		contents, err := os.ReadFile(filename)
		if err != nil {
			return nil, fmt.Errorf("could not load file: %w", err)
		}

		segments := strings.Split(string(contents), "-- name:")
		line := 1
		for _, segment := range segments {
			segment = strings.TrimSpace(segment)
			if segment == "" {
				continue
			}

			parts := strings.SplitN(segment, "\n", 2)
			meta := strings.TrimSpace(parts[0])
			metaParts := strings.Split(meta, " :")
			name := strings.TrimSpace(metaParts[0])
			type_ := strings.TrimSpace(metaParts[1])
			sql := strings.TrimSpace(parts[1])

			queries = append(queries, NewParsedQuery(name, type_, sql, filename, line))
			line += strings.Count(sql, "\n") + 2
		}
	}

	return NewParsedQueries(queries), nil
}
