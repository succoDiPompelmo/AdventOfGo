package d02

import (
	"log"
	"os"
	"succoDiPompelmo/adventOfGo/helpers"
	"testing"
)

func ExamplePartOne() {
	file, err := os.Open("testdata/full_input.txt")
	if err != nil {
		log.Fatalf("could not open input file: %v", err)
	}
	defer file.Close()

	if err := PartOne(file, os.Stdout); err != nil {
		log.Fatalf("could not solve: %v", err)
	}
	// Output: 2237
}

func ExamplePartTwo() {
	file, err := os.Open("testdata/full_input.txt")
	if err != nil {
		log.Fatalf("could open input file: %v", err)
	}
	defer file.Close()

	if err := PartTwo(file, os.Stdout); err != nil {
		log.Fatalf("could not solve: %v", err)
	}
	// Output: 66681
}

func Benchmark(b *testing.B) {
	testCases := map[string]struct {
		solution  helpers.Solution
		inputFile string
	}{
		"PartOne": {
			solution:  helpers.SolutionFunc(PartOne),
			inputFile: "testdata/full_input.txt",
		},

		"PartTwo": {
			solution:  helpers.SolutionFunc(PartTwo),
			inputFile: "testdata/full_input.txt",
		},
	}

	for name, test := range testCases {
		b.Run(name, func(b *testing.B) {
			helpers.BenchmarkSolution(b, test.solution, test.inputFile)
		})
	}
}
