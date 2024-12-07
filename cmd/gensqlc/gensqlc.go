package main

import (
	"os"

	"github.com/sqlc-dev/sqlc/pkg/cli"
)

func main() {
	code := cli.Run([]string{
		"generate",
	})

	os.Exit(code)
}
