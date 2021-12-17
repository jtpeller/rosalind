// ============================================================================
// = bs.go
// = 	Description: implementation for problems in the Bioinformatics
// = 		Stronghold section of Rosalind.
// = 	Date: December 16, 2021
// ============================================================================

package problems

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"rosalind/utils"
	"strconv"
	"strings"
)

var bsdata = "./problems/bsdata/"

type fasta struct {
	label string
	dna string
}

/*
Given: A DNA string s of length at most 1000 nt.
Return: Four integers (separated by spaces) counting the respective #
of times that the symbols 'A', 'C', 'G', and 'T' occur in s
*/
func DNA() string {
	s := utils.GetInput(bsdata + "rosalind_dna.txt")

	// count up the number of each letter
	aCount := 0
	gCount := 0
	tCount := 0
	cCount := 0
	for _, letter := range s {
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

	// assemble output
	str := fmt.Sprintf("%d %d %d %d", aCount, cCount, gCount, tCount)
	return str
}


/*
Given: A DNA string t having length at most 1000 nt.
Return: The transcribed RNA string of t.
*/
func RNA() string {
	s := utils.GetInput(bsdata + "rosalind_rna.txt")
	rna := strings.Replace(s, "T", "U", -1)
	return rna
}

/*
Given: A DNA string s of length at most 1000 bp.
Return: The reverse complement s^c of s.
*/
func REVC() string {
	// read in the file & format the string
	dat, err := ioutil.ReadFile(bsdata + "rosalind_revc.txt")
	utils.HandleError(err)
	buff := bytes.NewBufferString("")
	dat = dat[:len(dat)-1]

	// build the opposing strand
	for _, letter := range dat {
		buff.WriteByte(utils.ToggleDnaNucleotides(letter))
	}

	// assemble the output
	out := utils.Reverse(buff.String())
	return out
}

/*
 Given: Two DNA strings s and t (each of length at most 1 kbp).
 Return: All locations of t as a substring of s.
*/
func SUBS() string {
	// read in the file & format the string
	dat, err := ioutil.ReadFile(bsdata + "rosalind_subs.txt")
	utils.HandleError(err)
	s := string(dat)
	input := strings.Split(s, "\r\n")

	text := input[0]			// grab the original string
	pattern := input[1] 		// grab the pattern

	// run algorithm and output
	z := utils.Zalgo(text, pattern)
	out := ""
	for i := 0; i < len(z); i++ {
		out += fmt.Sprintf("%d ", z[i])
	}
	return out
}

/*
Given: Positive integers n <= 40 and k <= 5.
Return: The total number of rabbit pairs that will be present after n
months, if we begin with 1 pair and in each generation, every pair of
reproduction-age rabbits produces a litter of k rabbit pairs (instead of only 1 pair).
*/
func FIB() string {
	// parse input
	s := utils.GetInput(bsdata + "rosalind_fib.txt")
	s = utils.RemoveNewline(s)
	sv := strings.Split(s, " ")
	months, _ := strconv.ParseInt(sv[0], 10, 64)
	offspring, _ := strconv.ParseInt(sv[1], 10, 64)
	out := rabbits(months, offspring)
	return strconv.FormatInt(out, 10)
}

// helper for FIB()
func rabbits(mo, off int64) int64 {
	if mo == 1 {
		return 1
	} else if mo == 2 {
		return off
	}
	gen1 := rabbits(mo-1, off)
	gen2 := rabbits(mo-2, off)

	if mo <= 4 {
		return gen1 + gen2
	}
	return gen1 + gen2 * off
}

/*
Given: At most 10 DNA strings in FASTA format (of length at most 1 kbp each).
Return: The ID of the string having the highest GC-content, followed by the 
	GC-content of that string. Rosalind allows for a default error of 0.001 in 
	all decimal answers unless otherwise stated; please see the note on absolute
	error below.
*/
func GC() string {
	// parse the input
	s := utils.GetInput(bsdata + "rosalind_gc.txt")
	b := strings.Split(s, "\r\n")
	labels := make([]string, 0)
	dnas := make([]string, 0)
	for _, v := range b {
		if strings.HasPrefix(v, ">") {
			labels = append(labels, v[1:])	// get rid of the ">"
		} else {
			dnas = append(dnas, v)
		}
	}

	// compute which has the highest gc content
	maxlabel := labels[0]
	maxgc := 0.0
	for i, dna := range dnas {
		gc := gccontent(dna)
		if gc > maxgc {
			maxgc = gc
			maxlabel = labels[i]
		}
	}

	// assemble output
	out := fmt.Sprintf("%s\n%.6f", maxlabel, maxgc*100)
	return out
}

func gccontent(dna string) float64 {
	total := float64(len(dna))
	gc := float64(0)
	for _, nt := range dna {
		if nt == 'G' || nt == 'C' {
			gc++
		}
	}
	return gc/total
}