package Day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

//Part2 of day 1
func Part2() {
	//Zero out the counter
	currentFrequency := 0
	//Open file
	file, err := os.Open("Day1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	//Close file on exit
	defer file.Close()

	//Create scanner for file
	scanner := bufio.NewScanner(file)

	//Bool for if a frequency has happened before, assumes not
	found := false

	//Map of previous values
	previousValues := make(map[int]bool)

	//Loop while not found and not at end of file
	for !found && scanner.Scan() {
		//Extract text on line
		line := scanner.Text()
		//Get operator
		operator := string(line[0])
		//Get value
		value, err := strconv.Atoi(line[1:])
		if err != nil {
			fmt.Println(err)
		}

		//Save current currentFrequency for output string:
		previousValue := currentFrequency

		//Add or subtract value depending on operator
		if operator == "+" {
			currentFrequency += value
		} else if operator == "-" {
			currentFrequency -= value
		}

		//Check if currentFrequency has happened before
		_, exists := previousValues[currentFrequency]
		//If it has, print that, and set found to true
		if exists == true {
			fmt.Printf("Current frequency  %d, change of %d; resulting frequency  %d. witch has already been seen.", previousValue, value, currentFrequency)
			found = true
			//If not add currentFrequency to map and print current values
		} else {
			previousValues[currentFrequency] = true
			//Print string
			fmt.Printf("Current frequency  %d, change of %d; resulting frequency  %d.\n", previousValue, value, currentFrequency)
		}
	}

}
