package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {

	// example routes for tests
	// route := "R2, L3" (expect 5)
	// route := "R2, R2, R2" (expect 2)
	// route := "R5, L5, R5, R3" (expect 12)

	// original directions from http://adventofcode.com/2016/day/1 (expect 241)
	route := "R1, R1, R3, R1, R1, L2, R5, L2, R5, R1, R4, L2, R3, L3, R4, L5, R4, R4, R1, L5, L4, R5, R3, L1, R4, R3, L2, L1, R3, L4, R3, L2, R5, R190, R3, R5, L5, L1, R54, L3, L4, L1, R4, R1, R3, L1, L1, R2, L2, R2, R5, L3, R4, R76, L3, R4, R191, R5, R5, L5, L4, L5, L3, R1, R3, R2, L2, L2, L4, L5, L4, R5, R4, R4, R2, R3, R4, L3, L2, R5, R3, L2, L1, R2, L3, R2, L1, L1, R1, L3, R5, L5, L1, L2, R5, R3, L3, R3, R5, R2, R5, R5, L5, L5, R2, L3, L5, L2, L1, R2, R2, L2, R2, L3, L2, R3, L5, R4, L4, L5, R3, L4, R1, R3, R2, R4, L2, L3, R2, L5, R5, R4, L2, R4, L1, L3, L1, L3, R1, R2, R1, L5, R5, R3, L3, L3, L2, R4, R2, L5, L1, L1, L5, L4, L1, L1, R1"

	// turn the big string into an array using split using the comma and space as delimiter
	steps := strings.Split(route, ", ")
	// fmt.Println("Number of steps:", len(steps))

	// use x,y cartesian-based coordinates to keep track of path walked
	var x int64 = 0
	var y int64 = 0

	// default start state is facing north
	facing := "N"
	fmt.Println("Start facing: ", facing)

	for i := 0; i < len(steps); i++ {

		// parse `step` into something useable:
		var step = steps[i]                                  //e.g. "R33"
		var direction string = step[0:1]                     //e.g. "R"
		var blocksString = step[1:len(step)]                 //e.g. "33" (second letter to end of string)
		blocks, err := strconv.ParseInt(blocksString, 0, 64) //convert string number to a real number e.g. 33
		if err != nil {                                      // output error if conversion goes wrong
			fmt.Println(err)
		}

		// update coords based on whether instructions say
		// left/right (direction) and how far ('blocks')
		// TODO: could be optimized
		if facing == "N" {
			if direction == "R" {
				facing = "E"
				x += blocks
			} else {
				facing = "W"
				x -= blocks
			}
		} else if facing == "E" {
			if direction == "R" {
				facing = "S"
				y += blocks
			} else {
				facing = "N"
				y -= blocks
			}
		} else if facing == "S" {
			if direction == "R" {
				facing = "W"
				x -= blocks
			} else {
				facing = "E"
				x += blocks
			}
		} else if facing == "W" {
			if direction == "R" {
				facing = "N"
				x -= blocks
			} else {
				facing = "S"
				x += blocks
			}
		}

		// debugging
		fmt.Println(i, ": ", step, " ", facing, ">>", blocks, " = ", x, ",", y)
	}

	// FINAL ANSWER
	absoluteBlocks := math.Abs(float64(x)) + math.Abs(float64(y))
	fmt.Println("Total blocks away:", absoluteBlocks)

}
