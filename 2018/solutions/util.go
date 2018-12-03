package solutions

import (
	"strings"
	"strconv"
)

// Take input (from a file) as a stream of bytes,
// spit out the solution as a string.
type Solver = func(input string) string

// Split a multi-line string (with \n for newline characters)
// and return an array of strings representing the text on each line.
func SplitLines(input string) []string {
	var lines []string = make([]string, 0)
	for _, s := range strings.Split(input, "\n") {
		s = strings.TrimSpace(s)
		// Drop empty lines:
		if s != "" {
			lines = append(lines, s)
		}
	}
	return lines
}

// Read an array of numbers in string form and convert to ints.
func AsInts(nums []string) []int {
	var output []int = make([]int, len(nums))
	for i, n := range nums {
		x, _ := strconv.Atoi(n)
		output[i] = x
	}
	return output
}

func SumInts(nums []int) int {
	var sum int = 0
	for _, x := range nums {
		sum += x
	}
	return sum
}
