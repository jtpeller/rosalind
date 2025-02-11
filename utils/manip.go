package utils

import (
	"fmt"
	"strings"
)

func RemoveNewline(s string) string {
	s = strings.ReplaceAll(s, "\n", "")
	s = strings.ReplaceAll(s, "\r", "")
	return s
}

func Reverse(str string) string {
	bytes := []rune(str)
	for i, j := 0, len(bytes)-1; i < j; i, j = i+1, j-1 {
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}
	return string(bytes)
}

// toggles dna nucleotides
func ToggleDnaNucleotides(c byte) byte {
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

func Zalgo(text string, pattern string) []int {
	/* initializations */
	concat := pattern + "$" + text
	l := len(concat)
	p := len(pattern)

	/* build the Z array */
	Z := make([]int, l)
	L := 0
	R := 0
	k := 0

	for i := 1; i < l; i++ { // loop thru length of concat
		if i > R {
			L, R = i, i
			for R < l && concat[R-L] == concat[R] {
				R++
			}
			Z[i] = R - L
			R--
		} else {
			k = i - L // idx which matches in [L,R] interval
			if Z[k] < R-i+1 {
				Z[i] = Z[k]
			} else { // start from R and check manually
				L = i
				for R < l && concat[R-L] == concat[R] {
					R++
				}
				Z[i] = R - L
				R--
			}
		}
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
	if false {
		fmt.Println(matches)
	}

	return output
}
