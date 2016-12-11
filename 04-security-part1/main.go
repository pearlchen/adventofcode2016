package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func main() {

	// tests
	// encryptedRooms := loadInput("test.txt")

	// real
	encryptedRooms := loadInput("input.txt")

	// loop through input to get sum of sector IDs for valid rooms
	totalSum := 0
	for _, room := range encryptedRooms {
		name, sectorId, hash := parseEncryptedRoom(room)
		valid := validate(name, hash)
		if valid {
			totalSum += sectorId
		}
		fmt.Println(name, sectorId, hash, valid)
	}
	fmt.Println("TOTAL SUM:", totalSum)

}

func loadInput(filename string) []string {

	// use ioutil to read a file into a byte string
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("error", err)
		return make([]string, 0)
	}

	// convert to string
	str := string(bs)

	// split file contents at end of line
	lines := strings.Split(str, "\n")

	return lines
}

// Each room consists of an encrypted name (lowercase letters separated by dashes)
// followed by a dash, a sector ID, and a checksum in square brackets.
func parseEncryptedRoom(encrypted string) (string, int, string) {
	var name, hash string
	var sectorId int

	// hash
	i := strings.Index(encrypted, "[")       // find the start of the hash
	hash = encrypted[i+1 : len(encrypted)-1] // extract string between square brackets

	// sectorId
	j := strings.LastIndex(encrypted, "-")
	sectorId, err := strconv.Atoi(encrypted[j+1 : i])
	if err != nil {
		fmt.Println("   ", encrypted[j+1:i], "could not be converted")
	}

	// name
	name = encrypted[:j]

	return name, sectorId, hash
}

// A room is real if the checksum is the five most common letters in
// the encrypted name, in order, with ties broken by alphabetization.
func validate(name, hash string) bool {

	// split name into segments and rejoin to get rid of dash
	lettersOnly := strings.Join(strings.Split(name, "-"), "")

	// split into array so it can be sorted alphabetically
	letters := strings.Split(lettersOnly, "")
	sort.Strings(letters)

	// loop through and count occurances of letters
	occurances := make(map[string]int)
	occurances[letters[0]] = 1
	for i := 1; i < len(letters); i++ {
		letter := letters[i]
		if letter == letters[i-1] {
			occurances[letter]++
		} else {
			occurances[letter] = 1
		}
	}
	// fmt.Println(len(occurances), occurances)

	// order map by occurances
	sorted := sortMapByValue(occurances)
	// fmt.Println(len(sorted), sorted)

	// create comparision hash
	var comparisionHash = make([]string, 0)
	var group = make([]string, 0)

	// TODO: This loop is garbage
	for j := 0; j < len(sorted)-1; j++ {
		if sorted[j].Value == sorted[j+1].Value {
			// fmt.Println("will need to add", sorted[j].Key, "to a group for", sorted[j].Value)
			group = append(group, sorted[j].Key)
			// fmt.Println("   ", group)
			if j == len(sorted)-2 {
				group = append(group, sorted[j+1].Key)
				// fmt.Println("       ", group)
			}
		} else {
			if len(group) > 0 {
				group = append(group, sorted[j].Key)
				// fmt.Println("sort", group, "alphabetically then add to comparisionHash")
				sort.Strings(group)
				comparisionHash = append(comparisionHash, group...)
				group = group[:0]
			} else {
				// fmt.Println("okay to add", sorted[j].Key)
				comparisionHash = append(comparisionHash, sorted[j].Key)
				group = group[:0]
			}
		}
	}
	if len(group) > 0 {
		// fmt.Println("!!!sort", group, "alphabetically then add to comparisionHash")
		sort.Strings(group)
		comparisionHash = append(comparisionHash, group...)
	}
	// fmt.Println("  comparisionHash", comparisionHash)

	// compare to hash
	comparision := strings.Join(comparisionHash[:len(hash)], "")
	fmt.Println(comparision, "==", hash, "?")
	if comparision != hash {
		return false
	}

	return true
}

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
	sort.Sort(sort.Reverse(p))
	return p
}
