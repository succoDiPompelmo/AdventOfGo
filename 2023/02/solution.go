package d02

import (
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
	"succoDiPompelmo/adventOfGo/helpers"
)

// PartOne solves the first problem of day {{ .Day }} of Advent of Code {{ .Year }}.
func PartOne(r io.Reader, w io.Writer) error {
	// TODO: Read the input. For example:
	lines, err := helpers.LinesFromReader(r)
	if err != nil {
		return fmt.Errorf("could not read input: %w", err)
	}

	result := 0

	for _, line := range lines {
		game := strings.Split(line, ":")

		gameId := findGameId(game[0])
		maxRed := findMaxRevealed(game[1], "red")
		maxBlue := findMaxRevealed(game[1], "blue")
		maxGreen := findMaxRevealed(game[1], "green")

		// 12 red cubes, 13 green cubes, and 14 blue cubes
		if maxRed > 12 || maxGreen > 13 || maxBlue > 14 {
			continue
		}

		result += gameId
	}

	// TODO: Write the answer. For example:
	_, err = fmt.Fprintf(w, "%d", result)
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

	result := 0

	for _, line := range lines {
		game := strings.Split(line, ":")

		maxRed := findMaxRevealed(game[1], "red")
		maxBlue := findMaxRevealed(game[1], "blue")
		maxGreen := findMaxRevealed(game[1], "green")

		result += maxBlue * maxGreen * maxRed
	}

	// TODO: Write the answer. For example:
	_, err = fmt.Fprintf(w, "%d", result)
	if err != nil {
		return fmt.Errorf("could not write answer: %w", err)
	}

	return nil
}

func findGameId(game string) int {
	re := regexp.MustCompile(`Game (\d+)`)
	match := re.FindStringSubmatch(game)
	if len(match) > 1 {
		number, err := strconv.Atoi(match[1])
		if err == nil {
			return number
		}

		panic(fmt.Sprintf("could not convert %s to int: %v", match[1], err))
	}
	return 0
}

func findMaxRevealed(sets string, color string) int {
	re := regexp.MustCompile(fmt.Sprintf("(\\d+) %s", color))
	matches := re.FindAllStringSubmatch(sets, -1)

	max := 0

	for _, match := range matches {
		number, err := strconv.Atoi(match[1])
		if err == nil && number > max {
			max = number
		}
	}

	return max
}
