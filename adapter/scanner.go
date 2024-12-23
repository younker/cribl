package adapter

import (
	"bufio"
	"fmt"
	"io"
)

type Scanner struct{}

func NewScanner() Scanner {
	return Scanner{}
}

func (s Scanner) Sort() {
	scanner := bufio.NewScanner(io.Reader)
	for scanner.Scan() {
		input := scanner.Text()
		fmt.Printf("input: %s\n", input)
		// parse input -> normalize to model objects
		// validate input
		// send to dispatch
		// print stack
	}
}
