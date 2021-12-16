package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// read in the file & format the string
	dat, err := ioutil.ReadFile("rosalind_dna.txt")
	check(err)					// ensure no errors
	s := string(dat)
	strings.ToUpper(s)			// convert strings to upper

	// count up the number of each letter
	aCount := 0
	gCount := 0
	tCount := 0
	cCount := 0
	for _, letter := range dat {
		if letter == 'A' {
			aCount++
		} else if letter == 'G' {
			gCount++
		} else if letter == 'C' {
			cCount++
		} else if letter == 'T' {
			tCount++
		}
	}

	// print the output
	fmt.Printf("%d %d %d %d\n", aCount, cCount, gCount, tCount)
}
