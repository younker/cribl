package main

import (
	"os"

	"github.com/younker/cribl/adapter"
)

func main() {
	sorter := adapter.NewScanner(os.Stdin, os.Stdout, os.Stderr)
	sorter.Sort()
}
