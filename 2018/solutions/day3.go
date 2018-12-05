package solutions

import (
	"fmt"
	"regexp"
)

type Coord struct {
	x int
	y int
}

type Claim struct {
	id int
	left int
	top int
	width int
	height int
}

func (c *Claim) Coords() []Coord {
	var coords []Coord = make([]Coord, 0)

	for y:= c.top; y < c.top+c.height; y++ {  // rows
		for x:= c.left; x < c.left+c.width; x++ {  // columns
			var coord Coord = Coord{
				x: x,
				y: y,
			}
			coords = append(coords, coord)
		}
	}
	return coords
}


func getClaims(lines []string) *[]Claim {
	// each line: #123 @ 3,2: 5x4
	//
	//  - #123 = claimId
	//	- @ = separator (ignored)
	//  - 3,2 = 3 inches from left edge, 2 inches from top edge
	//	- 5x4 = 5 inches wide, 4 inches tall
	var claimMatch *regexp.Regexp = regexp.MustCompile("^#([0-9]+) @ ([0-9]+),([0-9]+): ([0-9]+)x([0-9]+)$")
	var claims []Claim = make([]Claim, 0)
	for _, line := range lines {
		var claimParts []string = claimMatch.FindStringSubmatch(line)
		var claim Claim = Claim{
			id: AsInt(claimParts[1]),
			left: AsInt(claimParts[2]),
			top: AsInt(claimParts[3]),
			width: AsInt(claimParts[4]),
			height: AsInt(claimParts[5]),
		}
		claims = append(claims, claim)
	}
	return &claims
}


func getMaxXY(claims *[]Claim) (int, int) {
	var maxX int = 0
	var maxY int = 0

	for _, claim := range *claims {
		var claimMaxX int = claim.left + claim.width
		var claimMaxY int = claim.top + claim.height
		if claimMaxX > maxX {
			maxX = claimMaxX
		}
		if claimMaxY > maxY {
			maxY = claimMaxY
		}
	}
	return maxX, maxY
}


func getMatrix(x int, y int) *[][]int {
	var matrix [][]int = make([][]int, y)  // y rows
	for i := range matrix {
		matrix[i] = make([]int, x)  // x columns
	}
	return &matrix
}


func printMatrix(m *[][]int) {
	fmt.Printf("[\n")
	for i := range *m {
		fmt.Printf("  %v,\n", (*m)[i])
	}
	fmt.Printf("]\n")
}


func matrixMemberCount(m *[][]int, target int) int {
	var count int = 0
	for i := range *m {
		var row []int = (*m)[i]
		for j := range row {
			if row[j] == target {
				count++
			}
		}
	}
	return count
}


func Day3Part1And2(input string) string {
	var lines []string = SplitLines(input)
	// Step1: collect all claims
	var claims *[]Claim = getClaims(lines)

	// Step2: get total fabric size (or at least the area covered by all claims
	var maxX int = 0
	var maxY int = 0
	maxX, maxY = getMaxXY(claims)

	// Step3: represent the fabric space as a matrix, initialized with all values to zero
	var fabric *[][]int = getMatrix(maxX, maxY)

	var overlappingClaimIds map[int]int = make(map[int]int)
	var nonOverlappingClaimIds map[int]int = make(map[int]int)

	// Step4: Start laying out the claims on the fabric, marking
	// the areas appropriately to indicate overlaps
	for _, claim := range *claims {
		var coords []Coord = claim.Coords()
		var overlaps bool = false
		for _, coord := range coords {
			// coord.x, coord.y
			var stakedClaim int = (*fabric)[coord.y][coord.x]
			if stakedClaim != 0 {
				// Another claim has already been laid here
				// Set the value to -1 to indicate an overlap
				(*fabric)[coord.y][coord.x] = -1
				// If the stored claim != -1, keep track of that claim id in the
				// overlapping claims map
				if stakedClaim != -1 {
					overlappingClaimIds[stakedClaim] = stakedClaim
					// Remove it from non overlapping map
					delete(nonOverlappingClaimIds, stakedClaim)
				}
				// No matter what, always store the current claim id in overlapping claims
				// map
				overlaps = true
			} else {
				// No other claim has yet been laid here
				// Set the value to the claim ID
				(*fabric)[coord.y][coord.x] = claim.id
			}
		}
		if overlaps {
			overlappingClaimIds[claim.id] = claim.id
		} else {
			nonOverlappingClaimIds[claim.id] = claim.id
		}
	}

	// printMatrix(fabric)

	var part1Solution string = AsStr(matrixMemberCount(fabric, -1))
	var part2Solution string = ""
	if len(nonOverlappingClaimIds) == 1 {
		// Part 2 solution found
		var k int
		// Get the first (and only key)
		for k = range nonOverlappingClaimIds {
			break
		}
		part2Solution = AsStr(k)
	}
	return fmt.Sprintf("Part1=%s, Part2=%s\n", part1Solution, part2Solution)
}
