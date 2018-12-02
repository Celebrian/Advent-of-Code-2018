package Day1

import (
	"testing"
)

//Test day1B function
func TestB(t *testing.T) {
	//Define and declare test cases
	testCases := map[string]int{
		"day1B1.txt": 0,
		"day1B2.txt": 10,
		"day1B3.txt": 5,
		"day1B4.txt": 14,
	}

	//For each case, check to see if Part1 returns the expected result.
	for input, expected := range testCases {
		actual := B(input, false)

		if actual != expected {
			t.Error("Failure: Expected ", expected, " got ", actual, "\n")
		}
	}
}
