package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func toggler(c byte) byte {
	if c == 'A' {
		return 'T'
	} else if c == 'T' {
		return 'A'
	} else if c == 'G' {
		return 'C'
	} else if c == 'C' {
		return 'G'
	} else {
		return c
	}
}

func rev(str string) string {
	bytes := []rune(str)
	for i, j := 0, len(bytes) - 1; i < j; i,j = i+1, j-1 {
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}
	return string(bytes)
}

func main() {
	// read in the file & format the string
	dat, err := ioutil.ReadFile("dna.txt")
	check(err)
	s := string(dat)
	strings.ToUpper(s)
	buff := bytes.NewBufferString("")

	// build the opposing strand
	for _, letter := range dat {
		buff.WriteByte(toggler(letter))
	}

	// print the output
	// without the reverse, the DNA strand would be 3' to 5'
	// conventionally it is written 5' to 3'
	out := rev(buff.String())
	fmt.Printf("%s\n", out)
}
