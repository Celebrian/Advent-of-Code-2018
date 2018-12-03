package day3

import "testing"

//Test day3A function
func TestA(t *testing.T) {
	//Define and declare test cases
	testCases := map[string]int{
		"day3A1.txt": 4,
	}

	//For each case, check to see if A returns the expected result.
	for input, expected := range testCases {
		actual := A(input, false)

		if actual != expected {
			t.Error("Failure: Expected ", expected, " got ", actual, "\n")
		}
	}
}
