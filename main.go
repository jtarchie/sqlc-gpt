package main

import (
	"bytes"
	"context"
	"crypto/sha1"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"text/template"

	"github.com/alecthomas/kong"
	"github.com/gosimple/slug"
	"github.com/jtarchie/sqlc/parser"
	"github.com/jtarchie/worker"
	"github.com/sashabaranov/go-openai"
)

type CLI struct {
	Glob              string `required:"" help:"glob to look for *.sql files"`
	OutputDir         string `required:"" help:"location to write the files to"`
	PackageName       string `required:"" default:"main" help:"the package name to the file"`
	OpenAIAccessToken string `help:"the API token for the OpenAI API" required:"" env:"OPENAI_ACCESS_TOKEN"`
	BaseURL           string `help:"url of the OpenAI HTTP domain" default:"https://api.openai.com/v1"`
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

func (c *CLI) executePrompt(query *parser.ParsedQuery, prompt string) (string, error) {
	config := openai.DefaultConfig(c.OpenAIAccessToken)
	config.BaseURL = c.BaseURL
	client := openai.NewClientWithConfig(config)

	slog.Info("openai start", slog.String("name", query.Name))

	response, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleAssistant,
					Content: prompt,
				},
			},
		},
	)

	slog.Info("open ai finished", slog.String("name", query.Name))
	if err != nil {
		return "", fmt.Errorf("could not translate: %w", err)
	}

	code := response.Choices[0].Message.Content
	code = strings.TrimSpace(code)
	code = strings.TrimPrefix(code, "```go")
	code = strings.TrimSuffix(code, "```")

	return code, nil
}

func logError(message string, err error, query *parser.ParsedQuery) {
	slog.Error(message, slog.String("error", err.Error()), slog.String("name", query.Name))
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

	err = os.MkdirAll(c.OutputDir, os.ModePerm)
	if err != nil {
		return fmt.Errorf("could not create output directory: %w", err)
	}

	var wg sync.WaitGroup

	workers := worker.New(2, 2, func(index int, query *parser.ParsedQuery) {
		defer wg.Done()

		filename := filepath.Join(c.OutputDir, fmt.Sprintf("%s.go", slug.Make(query.Name)))

		slog.Info("executing prompt", slog.String("name", query.Name))
		defer slog.Info("finished prompt", slog.String("name", query.Name))

		data := map[string]interface{}{
			"PackageName": c.PackageName,
			"Query":       query,
		}

		buffer := &bytes.Buffer{}

		t, err := template.New("promptTemplate").Parse(prompt)
		if err != nil {
			logError("could not parse template", err, query)
			return
		}

		err = t.Execute(buffer, data)
		if err != nil {
			logError("could execute template", err, query)
			return
		}

		promptHash := sha1.Sum(buffer.Bytes())

		file, err := os.OpenFile(filename, os.O_RDWR, os.ModePerm)
		if !errors.Is(err, os.ErrNotExist) {
			_ = file.Close()
			contents, err := os.ReadFile(filename)
			if err != nil {
				logError("could not read file", err, query)
				return
			}

			if bytes.Contains(contents, []byte(fmt.Sprintf("%x", promptHash))) {
				slog.Info("file is up to date", slog.String("filename", filename), slog.String("name", query.Name))
				return
			}
		}

		file, err = os.Create(filename)
		if err != nil {
			logError("could not create file", err, query)
			return
		}
		defer file.Close()

		result, err := c.executePrompt(query, strings.TrimSpace(buffer.String()))
		if err != nil {
			logError("could not execute prompt", err, query)
			return
		}

		_, _ = fmt.Fprint(file, "// This is a generated file, do not touch\n")
		_, _ = fmt.Fprintf(file, "// hash: %x\n\n", promptHash)
		_, err = fmt.Fprintf(file, "\n%s\n", result)
		if err != nil {
			logError("could not write prompt to file", err, query)
			return
		}
	})

	for _, query := range queries {
		wg.Add(1)
		workers.Enqueue(query)
	}

	wg.Wait()

	return nil
}
