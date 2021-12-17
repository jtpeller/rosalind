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
	newlabel := false
	for i := 0; i < len(b); i++ {
		v := b[i]
		if !newlabel && strings.HasPrefix(v, ">") {
			labels = append(labels, v[1:])	// get rid of the ">"
			newlabel = true
		} else if newlabel {		// loop thru everything until a >
			temp := b[i]
			nadded := 0
			for i+nadded+1 < len(b) && !strings.HasPrefix(b[i+nadded+1], ">") {
				temp += b[i+nadded+1]
				nadded++
			}
			dnas = append(dnas, temp)
			i += nadded - 1
			newlabel = false
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

/*
Given: Two DNA strings s and t of equal length (not exceeding 1 kbp).
Return: The Hamming distance dH(s,t)
*/
func HAMM() string {
	// parse input
	s := utils.GetInput(bsdata + "rosalind_hamm.txt")
	t := strings.Split(s, "\r\n")
	s1 := t[0]
	s2 := t[1]
	
	// compute
	dist := 0
	if len(s1) != len(s2) {
		panic("strings are of unequal length")
	}
	for i := range s1 {
		if s1[i] != s2[i] {
			dist++
		}
	}

	// assemble
	out := fmt.Sprintf("%d", dist)
	return out
}

/*
Given: Three positive integers k, m, and n, representing a population
containing k+m+n organisms: k individuals are homozygous dominant for
a factor, m are heterozygous, and n are homozygous recessive.
Return: The probability that two randomly selected mating organisms
will produce an individual possessing a dominant allele (and thus
displaying the dominant phenotype). Assume that any two organisms
can mate.
*/
func IPRB() string {
	s := utils.GetInput(bsdata + "rosalind_iprb.txt")
	s = utils.RemoveNewline(s)
	t := strings.Split(s, " ")
	if len(t) != 3 {
		panic("something bad happened")
	}
	k, _ := strconv.ParseInt(t[0], 10, 64)	// homozygous dominant
	m, _ := strconv.ParseInt(t[1], 10, 64)	// heterozygous
	n, _ := strconv.ParseInt(t[2], 10, 64)	// homozygous recessive

	total := k+m+n		// total population

	// convert to floats
	kf := float64(k)
	mf := float64(m)
	nf := float64(n)
	tf := float64(total)

	// homozygous dominant
	popk := kf/tf

	// heterozygous
	popm := mf/tf
	h1d2 := popm * (kf/(tf-1))
	h1h2 := popm * ((mf-1)/(tf-1))*0.75
	h1r2 := popm * (nf/(tf-1))*0.5
	h1 := h1d2 + h1h2 + h1r2

	// homozygous
	popn := nf/tf
	r1d2 := popn * (kf/(tf-1))
	r1h2 := popn * (mf/(tf-1))*0.5
	r1 := r1d2 + r1h2

	// build output 
	prob := popk + h1 + r1
	out := fmt.Sprintf("%.5f", prob)
	return out
}