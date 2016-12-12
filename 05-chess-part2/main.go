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
	// doorId := "abc" // expect 05ace8e3

	//real
	doorId := "ffykfhsq" // expect ???

	// use the Door ID as the seed for the md5 hash
	// then increase the integer (i), looking for a combined hexademical string starting with "00000"
	// TODO: better way to do this than increasing the integer???
	i := 0
	password := make([]string, 8)
	passwordLength := 0

	for passwordLength < 8 {
		h := md5.New()
		io.WriteString(h, doorId)
		i++
		io.WriteString(h, strconv.Itoa(i))
		hash := hex.EncodeToString(h.Sum(nil))
		if hash[:5] == "00000" {
			hi, err := strconv.Atoi(hash[5:6])
			if err != nil {
				fmt.Println("error converting", hash[5:6])
				continue
			}
			if hi <= 7 {
				if password[hi] == "" {
					fmt.Println(hash, hash[:5], hi, hash[6:7])
					password[hi] = hash[6:7]
					fmt.Println("    ", password, strings.Join(password, ""))
					passwordLength = len(strings.Join(password, ""))
				}
			}
		}
	}

	fmt.Println("PASSWORD:", passwordLength, ".", strings.Join(password, ""))

}
