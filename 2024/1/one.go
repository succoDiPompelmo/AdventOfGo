package one

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Day1_Part1(dat []byte) int {
	lines := strings.Split(string(dat), "\n")

	left_list := []int{}
	right_list := []int{}

	for _, line := range lines {
		line = strings.TrimSpace(line)
		numbers := strings.Split(line, "   ")
		left, err := strconv.Atoi(numbers[0])
		check(err)

		right, err := strconv.Atoi(numbers[1])
		check(err)

		left_list = append(left_list, left)
		right_list = append(right_list, right)
	}

	// Sort the lists
	sort.Slice(left_list, func(i, j int) bool { return left_list[i] < left_list[j] })
	sort.Slice(right_list, func(i, j int) bool { return right_list[i] < right_list[j] })

	diff_list := []int{}
	for i := 0; i < len(left_list); i++ {
		diff_list = append(diff_list, right_list[i]-left_list[i])
	}

	// Find the sum of the differences, take the absolute value
	sum := 0
	for _, diff := range diff_list {
		sum += int(math.Abs(float64(diff)))
	}

	return sum
}

// Open file day_1.txt and parse the content. Each row has two numbers separated by three spaces.
func Day1_Part2(dat []byte) int {
	lines := strings.Split(string(dat), "\n")

	left_list := []int{}
	right_list := []int{}

	for _, line := range lines {
		line = strings.TrimSpace(line)
		numbers := strings.Split(line, "   ")
		left, err := strconv.Atoi(numbers[0])
		check(err)

		right, err := strconv.Atoi(numbers[1])
		check(err)

		left_list = append(left_list, left)
		right_list = append(right_list, right)
	}

	// Store the frequency of each number in the right list
	freq := make(map[int]int)
	for _, num := range right_list {
		freq[num]++
	}

	// Compute similarity score that is number on the left list multiplied by the frequency of the number in the right list
	score := 0
	for _, num := range left_list {
		score += num * freq[num]
	}

	return score
}
