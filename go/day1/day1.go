package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

const halfPi float64 = math.Pi / 2.0


// Round float64 to the nearest int.
// Works for positive and negative numbers and zero.
// Examples:
// 	round(0.4) -> 0.0
//	round(0.5) -> 1.0
// 	round(-0.4) -> 0.0
//	round(-0.5) -> -1.0
//	round(0.0) -> 0.0
//	round(-0.0) -> 0.0
func round(f float64) int {
	return int(f + math.Copysign(0.5, f))
}

// Solution for http://adventofcode.com/2016/day/1.
// Solves any valid input without using "if" statements or maps.
func solve(input []string) int {
	var xLoc float64 = 0
	var yLoc float64 = 0

	// Start at 90* (or 0.5 * pi radians)
	// Angle is in radians
	var angle float64 = halfPi

	for _, each := range input {
		// Magnitude of vector
		var distance int
		// -1 or 1 (determines the direction of rotation)
		var sign float64

		// Unit vector components
		var x float64
		var y float64

		// Vector components (with non-unit length)
		var vx float64
		var vy float64

		// "L" or "R"
		direction := each[0]
		// The rest of the characters in the string are the distance (as an int)
		distance, _ = strconv.Atoi(string(each[1:]))
		// Magic happens here in calculating the sign.
		// Use the ascii code of L and R to determine the sign (direction of rotation)
		// 	* L yields +1 (positive rotation)
		//  * R yields -1 (negative rotation)
		sign = math.Pow(-1, 1.0 + float64((direction % 10 + direction / 10) % 2))
		// Add or subtract (0.5 * pi) radians (90 degrees) to make the rotation
		angle += sign * halfPi
		// Get the X component of the new angle (unit vector X component)
		x = math.Cos(angle)
		// Get the Y component of the new angle (unit vector Y component)
		y = math.Sin(angle)

		// Scale the unit vector by the distance
		vx = x * float64(distance)
		vy = y * float64(distance)

		// Update location by maintaining a sum of the vectors
		xLoc += vx
		yLoc += vy
		// fmt.Printf("Instruction: %s, vx: %d, vy: %d\n", each, vx, vy)
	}
	// fmt.Printf("Final location -> X: %d, Y: %d\n", xLoc, yLoc)

	// Get the final answer by calculating the "taxi cab distance"
	return int(round(math.Abs(xLoc)) + round(math.Abs(yLoc)))
}


func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s INPUT_FILE\n", os.Args[0])
		os.Exit(1)
	}

	filename := os.Args[1]
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	contents := string(bytes)

	lines := strings.Split(contents, "\n")
	input := strings.Split(lines[0], ", ")

	solution := solve(input)

	if lines[1] != "" {
		// There is an "expected" answer included in the input file.
		expected, err := strconv.Atoi(lines[1])
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		fmt.Printf("Solution: %d. Expected: %d.\n", solution, expected)
	} else {
		fmt.Printf("Solution: %d.\n", solution)
	}
}
