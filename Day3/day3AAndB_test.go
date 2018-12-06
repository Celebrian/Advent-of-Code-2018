package day3

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

//Test day3A function
func TestA(t *testing.T) {
	//Define and declare test cases
	testCases := map[string][2]int{
		"day3AAndB1.txt": {4, 3},
	}

	//For each case, check to see if A returns the expected result.
	for input, expected := range testCases {
		actualA, actualB := AAndB(input, *verbose, 8)

		if actualA != expected[0] || actualB != expected[1] {
			t.Error("Failure: Expected ", expected[0], " and ", expected[1], " got ", actualA, " and ", actualB, "\n",
				"E.g. Should be: ", expected[0], "==", actualA, " and ", expected[1], "==", actualB, "\n")
		}
	}
}
