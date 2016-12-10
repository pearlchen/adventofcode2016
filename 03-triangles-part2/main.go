package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	var triangles [][]int

	// TODO: Write tests
	// triangles = parseInput(loadInput("test.txt")) // expect 0
	triangles = parseInput(loadInput("input.txt")) // expect 862
	fmt.Println("Num triangles:", len(triangles))

	// Final answer
	fmt.Println("TOTAL VALID TRIANGLES", countValid(triangles))
}

func loadInput(filename string) string {

	// use ioutil to read a file into a byte string
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("error", err)
		return ""
	}

	// convert to string
	str := string(bs)

	return str
}

func parseInput(file string) [][]int {

	// split file contents at end of line
	lines := strings.Split(file, "\n")
	formatted := make([][]int, 0)

	// split each line into an array of the 3 numbers
	// use strings.Fields() to automatically trim whitespace
	// then convert to int so comparisions can be made with numbers not strings
	// TODO: Is there a shorter way to do this?
	for _, line := range lines {
		splitStrings := strings.Fields(line)
		if len(splitStrings) != 0 {
			splitInts := make([]int, 3)
			for j := 0; j < 3; j++ {
				numInt, err := strconv.Atoi(splitStrings[j])
				if err != nil {
					fmt.Println("Error converting", splitStrings[j], err)
				}
				splitInts[j] = numInt
			}
			fmt.Println(splitStrings, ">>>", splitInts)
			formatted = append(formatted, splitInts)
		}
	}

	return formatted
}

// In a valid triangle,
// the sum of any two sides must be larger than the remaining side.
func validate(a, b, c int) bool {
	if a+b <= c {
		return false
	}
	if b+c <= a {
		return false
	}
	if c+a <= b {
		return false
	}
	return true
}

func countValid(triangles [][]int) int {

	// count how many are valid triangles
	numValid := 0
	for i, triangle := range triangles {
		isTriangle := validate(triangle[0], triangle[1], triangle[2])
		fmt.Println(i, triangle)
		if isTriangle {
			numValid++
			fmt.Println("   ", isTriangle)
		}
	}

	return numValid
}
