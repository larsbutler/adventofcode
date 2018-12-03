package solutions

import (
	"bytes"
	"fmt"
	"strconv"
)


func Day2Part1(input string) string {
	var lines []string = SplitLines(input)

	var contain3 int = 0
	var contain2 int = 0

	for _, line := range lines {
		var found3 bool = false
		var found2 bool = false
		var counts map[rune]int = UniqueLetterCount(line)
		for _, v := range counts {
			// For each line, the 3 and 2 counts only apply once
			if v == 3 && !found3 {
				contain3++
				found3 = true
			} else if v == 2 && !found2 {
				contain2++
				found2 = true
			}
		}
	}

	var checksum = contain3 * contain2
	return strconv.Itoa(checksum)
}

// Given a string, build a map keyed by single characters.
// Values indicate the frequency of occurrence of each character in the string.
func UniqueLetterCount(s string) map[rune]int {
	var m = make(map[rune]int)

	for _, char := range s {
		if _, ok := m[char]; ok {
			// The character is already in the cache
			m[char]++
		} else {
			m[char] = 1
		}
	}
	return m
}


func Day2Part2(input string) string {
	return Day2Part2QuadraticSolution(input)
}

// Solution is O(n^2 m) where:
// 	- n is the number of ids in the input file
//  - m is the number of characters in each id
//
// Example puzzle input: n=250, m=6 -> 1.625million operations
func Day2Part2QuadraticSolution(input string) string {
	var boxIds []string = SplitLines(input)

	var boxId1 string
	var boxId2 string
	var n int = len(boxIds)

	// loop through all boxIds, excluding the last one
	Search:
		for i := 0; i < n - 1; i++ {
			boxId1 = boxIds[i]
			// loop through all boxIds, starting after the ith box ID
			for j := i + 1; j < n; j++ {
				boxId2 = boxIds[j]
				var diff []int = PositionalDiff(boxId1, boxId2)
				var sumDiff int = SumInts(diff)
				if sumDiff == 1 {
					fmt.Printf("Target box IDs found: %s, %s, %v\n", boxId1, boxId2, diff)
					break Search
				}
			}
		}

	// Return letters common between the two correct box IDs
	return CommonChars(boxId1, boxId2)
}


// Given two strings of equal length, return an array of ints indicating if the
// strings have differing characters in each position. For example:
// ("abc", "abd") -> [0, 0, 1]  One char is different
// ("xyz", "abc") -> [1, 1, 1]  All chars are different
// ("abc", "cab") -> [1, 1, 1]  Strings contain the same chars, but the
//								positions are completely different
func PositionalDiff(aStr string, bStr string) []int {
	// Assume strings are of equal length

	// Convert to rune arrays so we can index each individual char
	var a []rune = []rune(aStr)
	var b []rune = []rune(bStr)

	var diff []int = make([]int, len(a))

	for i, aChar := range a {
		var bChar rune = b[i]
		if aChar == bChar {
			// chars are the same, distances are 0
			diff[i] = 0
		} else {
			// chars are different
			diff[i] = 1
		}
	}

	return diff
}

// Maintaining order, return the the characters that are common between both strings.
// For example:
// ("abd", "abc") -> "ab"
func CommonChars(aStr string, bStr string) string {
	var buf *bytes.Buffer = bytes.NewBufferString("")
	var a []rune = []rune(aStr)
	var b []rune = []rune(bStr)

	for i, aChar := range a {
		var bChar rune = b[i]
		if aChar == bChar {
			buf.WriteRune(aChar)
		}
	}
	return buf.String()
}
