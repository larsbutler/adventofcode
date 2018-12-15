package main

import (
	"fmt"
	"flag"
	"io/ioutil"
	"os"
	"reflect"

	"./solutions"
)

var solvers = map[string]solutions.Solver {
	"1-1": solutions.Day1Part1,
	"1-2": solutions.Day1Part2,
	"2-1": solutions.Day2Part1,
	"2-2": solutions.Day2Part2,
	"3-1": solutions.Day3Part1And2,
	"3-2": solutions.Day3Part1And2,
	"4-1": solutions.Day4Part1,
}


func main() {
	var day *string = flag.String("day", "", "DAY")
	var inputFile *string = flag.String("input", "", "INPUT_FILE")
	var part *string = flag.String("part", "", "PART")
	flag.Parse()

	// Make a smart default for the input:
	if *inputFile == "" {
		*inputFile = fmt.Sprintf("input-day%s.txt", *day)
	}

	if *day == "" || *inputFile == "" || *part == "" {
		fmt.Println("Required flags:")
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Read the input file
	var input []byte
	var err error
	input, err = ioutil.ReadFile(*inputFile)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var solution string
	var solver solutions.Solver = solvers[fmt.Sprintf("%s-%s", *day, *part)]
	if solver == nil {
		fmt.Printf("No solution yet for day %s part %s. ", *day, *part)
		var keys []reflect.Value = reflect.ValueOf(solvers).MapKeys()
		fmt.Printf("Choose from: %s\n", keys)
		os.Exit(1)
	}

	// Convert byte array to string and solve:
	solution = solver(string(input))
	fmt.Printf("Solution: %s\n", solution)

	/*
	pkg, err := importer.Default().Import("solutions")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	for _, declName := range pkg.Scope().Names() {
		fmt.Println(declName)
	}
	*/
}
