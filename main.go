package main

import (
	"os"

	"github.com/younker/thoughtful_ai/adapter"
)

func main() {
	sorter := adapter.NewScanner(os.Stdin, os.Stdout, os.Stderr)
	sorter.Sort()
}
