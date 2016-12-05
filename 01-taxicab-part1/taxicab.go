package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {

	// example routes for tests
	// route := "R2, L3"         // expect 5
	// route := "R2, R2, R2"     // expect 2
	// route := "R5, L5, R5, R3" // expect 12

	// original directions from http://adventofcode.com/2016/day/1
	// final answer to expect = 241
	route := "R1, R1, R3, R1, R1, L2, R5, L2, R5, R1, R4, L2, R3, L3, R4, L5, R4, R4, R1, L5, L4, R5, R3, L1, R4, R3, L2, L1, R3, L4, R3, L2, R5, R190, R3, R5, L5, L1, R54, L3, L4, L1, R4, R1, R3, L1, L1, R2, L2, R2, R5, L3, R4, R76, L3, R4, R191, R5, R5, L5, L4, L5, L3, R1, R3, R2, L2, L2, L4, L5, L4, R5, R4, R4, R2, R3, R4, L3, L2, R5, R3, L2, L1, R2, L3, R2, L1, L1, R1, L3, R5, L5, L1, L2, R5, R3, L3, R3, R5, R2, R5, R5, L5, L5, R2, L3, L5, L2, L1, R2, R2, L2, R2, L3, L2, R3, L5, R4, L4, L5, R3, L4, R1, R3, R2, R4, L2, L3, R2, L5, R5, R4, L2, R4, L1, L3, L1, L3, R1, R2, R1, L5, R5, R3, L3, L3, L2, R4, R2, L5, L1, L1, L5, L4, L1, L1, R1"

	// turn the `route` string into an array using split using both the comma and space as delimiter
	steps := strings.Split(route, ", ")
	fmt.Println("Number of steps in route:", len(steps))

	// use x,y cartesian-based coordinates to keep track of path walked
	var coords = map[string]int64{"x": 0, "y": 0}

	// define a struct for how to behave for a given direction
	// TODO: can reduce complexity of right/left if turn to face direction first
	type behaviour struct {
		facing string // "N", "E", "S", "W"
		coord  string // reference to x or y
		right  int64  // if going right, increment (1) or decrement (-1)
		left   int64  // if going left, increment (1) or decrement (-1)
	}

	// calling this array `compass` since the positions of behaviours will "rotate around"
	compass := []behaviour{
		{facing: "N", coord: "x", right: 1, left: -1},
		{facing: "E", coord: "y", right: 1, left: -1},
		{facing: "S", coord: "x", right: -1, left: 1},
		{facing: "W", coord: "y", right: -1, left: 1},
	}

	// default start state is facing north
	var facing behaviour = compass[0]
	fmt.Println("Start facing: ", facing.facing)

	// start "walking" by looping through the steps
	for i := 0; i < len(steps); i++ {

		var step = steps[i] //e.g. "R33"
		fmt.Println("(", i, ")", step, ":")

		// parse `step` into something useable:
		var directionShort string = step[0:1]                //e.g. "R"
		var blocksString = step[1:len(step)]                 //e.g. "33" (second letter to end of string)
		blocks, err := strconv.ParseInt(blocksString, 0, 64) //convert string number to a real number e.g. 33
		if err != nil {                                      // output error if conversion goes wrong
			fmt.Println(err)
		}

		// update coords based on whether instructions say
		// left/right (`directionShort`) and how far (`blocks`)
		// TODO: can reduce complexity of right/left if turn to face direction first
		fmt.Println("   Currently facing: ", facing.facing)
		if directionShort == "R" {
			fmt.Println("   Turn RIGHT and walk", blocks, "blocks")
			coords[facing.coord] += blocks * facing.right

			// update compass clockwise
			firstSlice, compassTemp := compass[0], compass[1:] //shift (remove the first element)
			compass = append(compassTemp, firstSlice)          // push (re-add element at end)
		} else {
			fmt.Println("   Turn LEFT and walk", blocks, "blocks")
			coords[facing.coord] += blocks * facing.left

			// update compass counter-clockwise
			lastSlice, compassTemp := compass[len(compass)-1:], compass[:len(compass)-1] // pop (removes the last item)
			compass = append(lastSlice, compassTemp...)                                  //unshift (insert element at front)
		}

		// update to always face direction of first element
		facing = compass[0]

		// debugging
		fmt.Println("   Now at ", coords["x"], coords["y"])
	}

	// FINAL ANSWER
	absoluteBlocks := math.Abs(float64(coords["x"])) + math.Abs(float64(coords["y"]))
	fmt.Println("Total blocks away:", absoluteBlocks)

}
