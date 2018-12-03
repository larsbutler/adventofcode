package solutions

import (
	"fmt"
	"strconv"
)

func Day1Part1(input string) string {
	return strconv.Itoa(SumInts(AsInts(SplitLines(input))))
}

func Day1Part2(input string) string {
	// Frequency starts at 0
	var frequency int = 0
	// Process each number, add it to the frequency
	// Keep track of how many times we encounter each frequency
	// Return the first one that we encounter twice
	var freqs = make(map[int]int)
	var nums []int = AsInts(SplitLines(input))
	// This is the frequency we're looking for
	var target int
	// We start the frequency at 0, so we have technically encountered it once:
	freqs[0] = 1

	Search:
		for i := 0; ; i++ {
			for _, num := range nums {
				frequency += num
				if val, ok := freqs[frequency]; ok {
					// The cache already contains the frequency
					// Increment the cache counter
					freqs[frequency] = val + 1
					if freqs[frequency] == 2 {
						fmt.Printf(
							"Target frequency found: target=%d, loop iteration=%d\n",
							frequency,
							i,
						)
						target = frequency
						break Search
					}
				} else {
					// This the first time we've seen this frequency
					freqs[frequency] = 1
				}
			}
		}
	return strconv.Itoa(target)
}
