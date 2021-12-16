package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const DEBUG = false			// for debugging everything

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func zalgo(text string, pattern string) []int {
	/* initializations */
	concat := pattern + "$" + text
	l := len(concat)
	p := len(pattern)

	/* build the Z array */
	Z := make([]int, l)
	L := 0
	R := 0
	k := 0

	for i := 1; i < l; i++ {		// loop thru length of concat
		if i > R {
			L, R = i, i
			for R < l && concat[R-L] == concat[R] {
				R++
			}
			Z[i] = R - L
			R--
		} else {
			k = i - L 				// idx which matches in [L,R] interval
			if Z[k] < R - i + 1 {
				Z[i] = Z[k]
			} else { 				// start from R and check manually
				L = i
				for R < l && concat[R-L] == concat[R] {
					R++
				}
				Z[i] = R - L
				R--
			}
		}
	}
	if DEBUG {
		fmt.Println(Z)
	}

	// loop thru the Z array
	matches := make([]string, 0)
	output := make([]int, 0)
	for i := 0; i < l; i++ {
		if Z[i] == p {
			foo := i - p - 1
			output = append(output, foo+1)
			matches = append(matches, text[foo:foo+p])
		}
	}

	// print matches for debugging
	if DEBUG {
		for i, v := range matches {
			fmt.Printf("match i=%d\t%s\n", i, v)
		}
	}

	return output
}

func main() {
	// read in the file & format the string
	dat, err := ioutil.ReadFile("rosalind_dna.txt")
	check(err)
	s := string(dat)
	strings.ToUpper(s)			// convert to uppercase, just in case
	input := strings.Split(s, "\r\n")

	text := input[0]			// grab the original string
	pattern := input[1] 		// grab the pattern

	// debug file input
	if DEBUG {
		fmt.Printf("Searching the text: %s\n", text)
		fmt.Printf("for the pattern: %s\n", pattern)
	}

	// run algorithm and output
	output := zalgo(text, pattern)
	for i := 0; i < len(output); i++ {
		fmt.Printf("%d ", output[i])
	}
}
