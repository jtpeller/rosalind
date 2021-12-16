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
	dat, err := ioutil.ReadFile("rosalind_rna.txt")
	check(err) // ensure no errors
	s := string(dat)
	strings.ToUpper(s) // convert strings to upper

	rna := strings.Replace(s, "T", "U", -1)
	fmt.Println(rna)
}
