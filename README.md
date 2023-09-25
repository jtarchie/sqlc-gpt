# SQL to Go Code Generator

## Overview

This tool is designed to transform SQL queries into Go code. The transformation
is achieved by using OpenAI's ChatGPT API. By automating this translation, we
aim to reduce boilerplate and improve developer productivity when writing SQL
queries in Go.

## Features

- Automated translation of SQL queries to Go code.
- Checks for up-to-date translations to avoid redundant API calls.
- Multithreaded processing for better performance with larger sets of SQL
  queries.
- Generates Go files with comments to indicate they are auto-generated.

## Setup

### Prerequisites

- Go installed on your machine.
- An active OpenAI API subscription.

### Installation

1. Clone the repository.
2. Navigate to the repository directory.
3. Run `go build` to compile the tool.

## Examples

```bash
# generates files based off of examples/sql files
go run github.com/jtarchie/sqlc-gpt --glob='./examples/sql/*.sql' --output-dir=./examples/code/ --package-name=code
# fix go imports, because for some reason ChatGPT just can't get it
goimports -w examples/code/
```
