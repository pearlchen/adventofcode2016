package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {

	// example routes for tests
	// TODO: Write tests
	// route := "R2, L3" // expect 5
	// route := "R2, R2, R2" // expect 2
	// route := "R5, L5, R5, R3" // expect 12

	// original directions from http://adventofcode.com/2016/day/1
	// final answer to expect = 241
	// TODO: Accept route as a value typed in somewhere
	route := "R1, R1, R3, R1, R1, L2, R5, L2, R5, R1, R4, L2, R3, L3, R4, L5, R4, R4, R1, L5, L4, R5, R3, L1, R4, R3, L2, L1, R3, L4, R3, L2, R5, R190, R3, R5, L5, L1, R54, L3, L4, L1, R4, R1, R3, L1, L1, R2, L2, R2, R5, L3, R4, R76, L3, R4, R191, R5, R5, L5, L4, L5, L3, R1, R3, R2, L2, L2, L4, L5, L4, R5, R4, R4, R2, R3, R4, L3, L2, R5, R3, L2, L1, R2, L3, R2, L1, L1, R1, L3, R5, L5, L1, L2, R5, R3, L3, R3, R5, R2, R5, R5, L5, L5, R2, L3, L5, L2, L1, R2, R2, L2, R2, L3, L2, R3, L5, R4, L4, L5, R3, L4, R1, R3, R2, R4, L2, L3, R2, L5, R5, R4, L2, R4, L1, L3, L1, L3, R1, R2, R1, L5, R5, R3, L3, L3, L2, R4, R2, L5, L1, L1, L5, L4, L1, L1, R1"

	// turn the `route` string into an array using split using both the comma and space as delimiter
	steps := strings.Split(route, ", ")
	fmt.Println("Number of steps in route:", len(steps))

	// use x,y cartesian-based coordinates to keep track of path walked
	var coords = map[string]int64{"x": 0, "y": 0}

	// define a struct for how to behave for a given direction
	type behaviour struct {
		direction  string // "N", "E", "S", "W"
		coord      string // reference to x or y
		multiplier int64  // based on cartesian plane, -1 is up, 1 is down, -1 is left, 1 is right
	}

	// calling this array `compass` since the positions of behaviours will "rotate around"
	compass := []behaviour{
		{direction: "N", coord: "y", multiplier: -1},
		{direction: "E", coord: "x", multiplier: 1},
		{direction: "S", coord: "y", multiplier: 1},
		{direction: "W", coord: "x", multiplier: -1},
	}

	// default start state is facing north
	var facing behaviour = compass[0]
	fmt.Println("Start facing: ", facing.direction)

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
		if directionShort == "R" {
			// update compass clockwise
			fmt.Println("   Turn RIGHT")
			firstSlice, compassTemp := compass[0], compass[1:] //shift (remove the first element)
			compass = append(compassTemp, firstSlice)          // push (re-add element at end)
		} else {
			// update compass counter-clockwise
			fmt.Println("   Turn LEFT")
			lastSlice, compassTemp := compass[len(compass)-1:], compass[:len(compass)-1] // pop (removes the last item)
			compass = append(lastSlice, compassTemp...)                                  //unshift (insert element at front)
		}

		// update to always face direction of first element
		facing = compass[0]
		fmt.Println("   Now facing", facing.direction)

		// update x or y coord
		fmt.Println("   Walked", blocks, "block(s)")
		coords[facing.coord] += blocks * facing.multiplier

		// debugging at end of step
		fmt.Println("   Now at x:", coords["x"], ", y:", coords["y"])
	}

	// FINAL ANSWER
	absoluteBlocks := math.Abs(float64(coords["x"])) + math.Abs(float64(coords["y"]))
	fmt.Println("Total blocks away:", absoluteBlocks)

}
