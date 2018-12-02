package Day1

import "testing"

//Test part1 function
func Test_B(t *testing.T) {
	//Define and declare test cases
	testCases := map[string]int{
		"Day1/day1B1.test": 0,
		"Day1/day1B2.test": 10,
		"Day1/day1B3.test": 5,
		"day1/day1B4.test": 14,
	}

	//For each case, check to see if Part1 returns the expected result.
	for input, expected := range testCases {
		actual := B(input)

		if actual != expected {
			t.Error("Failure: Expected ", expected, " got ", actual)
		}
	}
}
