package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func main() {

	//tests
	// columns := parseInput(loadInput("test.txt")) // expect "advent"

	//real
	columns := parseInput(loadInput("input.txt")) // expect ???

	fmt.Println(columns)

	message := make([]string, len(columns))
	for i, column := range columns {
		occurances := countOccurances(column)
		sorted := sortMapByValue(occurances)
		message[i] = sorted[0].Key
	}
	fmt.Println("FINAL MESSAGE:", strings.Join(message, ""))
}

func loadInput(filename string) string {

	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading", filename, err)
		return ""
	}

	str := string(bs)

	return str
}

func parseInput(file string) []string {

	// split file at line break
	rows := strings.Split(file, "\n")

	// all rows have the same length so use first one to determine length of message
	messageLength := len(rows[0])
	columnLength := len(rows)
	fmt.Println(messageLength, columnLength)

	// flip rows into columns, sorted alphabetically
	columns := make([]string, messageLength)
	for i := 0; i < messageLength; i++ {
		column := make([]string, columnLength)
		for j := 0; j < columnLength; j++ {
			row := rows[j]
			column[j] = row[i : i+1]
		}
		// fmt.Println(i, column)
		columns[i] = strings.Join(column, "")
	}

	return columns
}

func countOccurances(letters string) map[string]int {

	// split into array so it can be sorted alphabetically
	sortedLetters := strings.Split(letters, "")
	sort.Strings(sortedLetters)

	// loop through and count occurances of letters
	occurances := make(map[string]int)
	occurances[sortedLetters[0]] = 1
	for i := 1; i < len(sortedLetters); i++ {
		letter := sortedLetters[i]
		if letter == sortedLetters[i-1] {
			occurances[letter]++
		} else {
			occurances[letter] = 1
		}
	}

	// fmt.Println(len(occurances), occurances)
	return occurances
}

/// -------------

// A data structure to hold a key/value pair.
type Pair struct {
	Key   string
	Value int
}

// A slice of Pairs that implements sort.Interface to sort by Value.
type PairList []Pair

func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }

// A function to turn a map into a PairList, then sort and return it.
func sortMapByValue(m map[string]int) PairList {
	p := make(PairList, len(m))
	i := 0
	for k, v := range m {
		p[i] = Pair{k, v}
		i++
	}
	sort.Sort(p)
	return p
}
