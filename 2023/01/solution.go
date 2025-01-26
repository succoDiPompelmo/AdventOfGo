package d01

import (
	"fmt"
	"io"
	"strings"
	"unicode"

	"succoDiPompelmo/adventOfGo/helpers"
)

// PartOne solves the first problem of day {{ .Day }} of Advent of Code {{ .Year }}.
func PartOne(r io.Reader, w io.Writer) error {
	// TODO: Read the input. For example:
	lines, err := helpers.LinesFromReader(r)
	if err != nil {
		return fmt.Errorf("could not read input: %w", err)
	}

	answer := 0

	for _, line := range lines {
		numbers := extractNumbers(line, false)

		if len(numbers) == 0 {
			panic("no numbers found")
		}

		answer += numbers[0] * 10
		answer += numbers[len(numbers)-1]
	}

	// TODO: Write the answer. For example:
	_, err = fmt.Fprintf(w, "%d", answer)
	if err != nil {
		return fmt.Errorf("could not write answer: %w", err)
	}

	return nil
}

// PartTwo solves the second problem of day {{ .Day }} of Advent of Code {{ .Year }}.
func PartTwo(r io.Reader, w io.Writer) error {
	// TODO: Read the input. For example:
	lines, err := helpers.LinesFromReader(r)
	if err != nil {
		return fmt.Errorf("could not read input: %w", err)
	}

	answer := 0

	for _, line := range lines {
		numbers := extractNumbers(line, true)

		if len(numbers) == 0 {
			panic("no numbers found")
		}

		answer += numbers[0] * 10
		answer += numbers[len(numbers)-1]
	}

	// TODO: Write the answer. For example:
	_, err = fmt.Fprintf(w, "%d", answer)
	if err != nil {
		return fmt.Errorf("could not write answer: %w", err)
	}

	return nil
}

func extractNumbers(s string, checkSpelledDigit bool) []int {
	digits := []int{}
	for i, r := range s {
		if unicode.IsDigit(r) {
			digit := int(r - '0')
			digits = append(digits, digit)
		} else if checkSpelledDigit && unicode.IsLetter(r) {
			spelledOutDigit, err := findSpelledOutDigit(s[i:])
			if err == nil {
				digits = append(digits, spelledOutDigit)
			}
		}
	}

	return digits
}

var wordsToDigits = map[string]int{
	"zero":  0,
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func findSpelledOutDigit(s string) (int, error) {
	for word, digit := range wordsToDigits {
		if strings.HasPrefix(s, word) {
			return digit, nil
		}
	}

	return 0, fmt.Errorf("could not find a digit")
}
