package adapter

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/younker/cribl/model"
)

type Scanner struct {
	input  io.Reader
	output io.Writer
	outerr io.Writer
}

func NewScanner(in io.Reader, out io.Writer, oerr io.Writer) Scanner {
	return Scanner{
		input:  in,
		output: out,
		outerr: oerr,
	}
}

func (s Scanner) Sort() {
	scanner := bufio.NewScanner(s.input)
	for scanner.Scan() {
		input := scanner.Text()
		pkg, err := csvInputToPackage(input)
		if err != nil {
			fmt.Println(s.outerr, fmt.Errorf("unable to parse input '%s' due to: %w", input, err))
			continue
		}

		fmt.Fprintln(s.output, model.Dispatch(pkg))
	}
}

func csvInputToPackage(line string) (model.Package, error) {
	// The specific traits of this package are provided as a CSV input where the
	// values are ordered as follows:
	//   mass,height,width,depth
	traits := strings.Split(line, ",")
	if len(traits) != 4 {
		return model.Package{}, fmt.Errorf("inputs should contain 4 values 'mass,height,width,depth' in that order. We received invalid input '%s'", line)
	}

	mass, err := strconv.ParseFloat(traits[0], 32)
	if err != nil {
		return model.Package{}, err
	}

	height, err := strconv.Atoi(traits[1])
	if err != nil {
		return model.Package{}, err
	}

	width, err := strconv.Atoi(traits[2])
	if err != nil {
		return model.Package{}, err
	}

	depth, err := strconv.Atoi(traits[3])
	if err != nil {
		return model.Package{}, err
	}

	return model.Package{
		Mass: model.Mass(mass),
		Dimensions: model.Dimensions{
			Height: height,
			Width:  width,
			Depth:  depth,
		},
	}, nil
}
