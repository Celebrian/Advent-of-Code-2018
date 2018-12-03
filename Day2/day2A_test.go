package day2

import (
	"flag"
	"testing"
)

var verbose *bool

//Function to parse verbose flag
func init() {
	verbose = flag.Bool("v", false, "Should the program be verbose or not(default not).")
	flag.Parse()
}

//Test day2A function
func TestA(t *testing.T) {
	//Define and declare test cases
	testCases := map[string]int{
		"day2A1.txt": 12,
	}

	//For each case, check to see if A returns the expected result.
	for input, expected := range testCases {
		actual := A(input, *verbose)

		if actual != expected {
			t.Error("Failure: Expected ", expected, " got ", actual, "\n")
		}
	}
}
