package Day1

import "testing"

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
		actual := A(input, false)

		if actual != expected {
			t.Error("Failure: Expected ", expected, " got ", actual)
		}
	}
}
