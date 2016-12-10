package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	var rows [][]int

	// TODO: Write tests
	// rows = parseInput(loadInput("test.txt")) // expect 6
	rows = parseInput(loadInput("input.txt")) // expect 1577
	columns := flipRowsToColumns(rows)

	triangles := make([][]int, 0)
	for _, column := range columns {
		triangles = append(triangles, getPossibleTrianglesForColumn(column)...)
	}

	// Final answer
	fmt.Println("TOTAL VALID TRIANGLES", countValid(triangles), "out of", len(triangles), "possible")
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
			// fmt.Println(splitStrings, ">>>", splitInts)
			formatted = append(formatted, splitInts)
		}
	}

	return formatted
}

func flipRowsToColumns(input [][]int) [][]int {

	column0 := make([]int, 0)
	column1 := make([]int, 0)
	column2 := make([]int, 0)

	// flip columns into a rows
	for _, line := range input {
		column0 = append(column0, line[0])
		column1 = append(column1, line[1])
		column2 = append(column2, line[2])
	}

	columns := [][]int{column0, column1, column2}
	return columns
}

func getPossibleTrianglesForColumn(column []int) [][]int {

	triangles := make([][]int, 0)

	// lump possible triangles by groups of 3 in a column
	for i := 2; i < len(column); i = i + 3 {
		possible := []int{column[i], column[i-1], column[i-2]}
		triangles = append(triangles, possible)
	}

	return triangles
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
