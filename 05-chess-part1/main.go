package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func main() {

	// test
	// doorId := "abc" // expect 18f47a30

	//real
	doorId := "ffykfhsq" // expect c6697b55

	// use the Door ID as the seed for the md5 hash
	// then increase the integer (i), looking for a combined hexademical string starting with "00000"
	// TODO: better way to do this than increasing the integer???
	i := 0
	password := make([]string, 0)

	for len(password) < 8 {
		h := md5.New()
		io.WriteString(h, doorId)
		i++
		io.WriteString(h, strconv.Itoa(i))
		hash := hex.EncodeToString(h.Sum(nil))
		if hash[:5] == "00000" {
			fmt.Println(hash, hash[:5], hash[5:6])
			password = append(password, hash[5:6])
		}
	}

	fmt.Println("PASSWORD:", strings.Join(password, ""))

}
