package day2

import (
	"testing"
)

//Test day2A function
func TestB(t *testing.T) {
	//Define and declare test cases
	testCases := map[string]string{
		"day2B1.txt": "fgij",
	}

	//For each case, check to see if A returns the expected result.
	for input, expected := range testCases {
		actual := B(input, false)

		if actual != expected {
			t.Error("Failure: Expected ", expected, " got ", actual, "\n")
		}
	}
}
