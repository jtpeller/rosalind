// ============================================================================
// = main.go																  =
// = 	Description: The file from which to run any of the sequences		  =
// = 	Date: October 08, 2021												  =
// ============================================================================

package main

import (
	"errors"
	"flag"
	"fmt"
	"reflect"
	"time"

	prob "rosalind/problems"
	"rosalind/utils"
)


func main() {
	// program initialization (flags)
	id := flag.String("id", "", "which problem to run. Example: -id dna")
	comptime := flag.Bool("t", false, "True if you want approximate time-of-computation information printed. False otherwise")
	
	flag.Parse()		// remember to parse!

	_, exists := stubs[*id]

	if *id == "" {				// user must specify a problem to run
		utils.HandleError(errors.New("you need to specify a problem to run! "))
	} else if !exists {				// user must specify a sequence that exists
		utils.HandleError(errors.New("either this sequence has not been implemented yet, or your id is invalid! "))
	}

	start := time.Now()
	fmt.Println(call(*id))
	duration := time.Since(start)

	// output time if requested
	if *comptime {
		utils.PrintInfo("Computed Problem " + *id + " in " + duration.String())
	}
}

// calls the function by name using stub storage
func call(name string) string {
	f := reflect.ValueOf(stubs[name])

	// build result interface
	var res []reflect.Value = f.Call(nil)
	result := res[0].Interface().(string)
	return result
}

// the following is a (large) mapping from strings to the corresponding function
var stubs = map[string]interface{}{
	// bio_stronghold
	"dna": prob.DNA,
	"rna": prob.RNA,
	"revc": prob.REVC,
	"subs": prob.SUBS,
	"fib": prob.FIB,
	"gc": prob.GC,
	"hamm": prob.HAMM,
	"iprb": prob.IPRB,
}