package Day1

import "testing"

//Test part1 function
func TestA(t *testing.T) {
	//Define and declare test cases
	testCases := map[string]int{
		"day1A1.test": 3,
		"day1A2.test": 0,
		"day1A3.test": -6,
	}

	//For each case, check to see if Part1 returns the expected result.
	for input, expected := range testCases {
		actual := A(input)

		if actual != expected {
			t.Error("Failure: Expected ", expected, " got ", actual)
		}
	}
}
