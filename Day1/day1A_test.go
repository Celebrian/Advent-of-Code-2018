package Day1

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

//Test part1 function
func TestA(t *testing.T) {
	//Define and declare test cases
	testCases := map[string]int{
		"day1A1.txt": 3,
		"day1A2.txt": 0,
		"day1A3.txt": -6,
	}

	//For each case, check to see if Part1 returns the expected result.
	for input, expected := range testCases {
		actual := A(input, *verbose)

		if actual != expected {
			t.Error("Failure: Expected ", expected, " got ", actual)
		}
	}
}
