package Day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

//B is part 2 of Day 1
func B(fil string, verbose bool) int {

	//Bool for if a frequency has happened before, assumes not
	found := false

	//Map of previous values
	previousValues := map[int]bool{
		0: true,
	}

	//Variable for resulting frequency
	var resultingFrequency int

	//Zero out the counter
	currentFrequency := 0

	//Loop while not found
	for !found {
		//Open file
		file, err := os.Open(fil)
		if err != nil {
			log.Fatal(err)
		}

		//Create scanner for file
		scanner := bufio.NewScanner(file)

		//Scan each line
		for !found && scanner.Scan() {
			//Extract text on line
			line := scanner.Text()
			//Get operator
			operator := string(line[0])
			//Get value
			change, err := strconv.Atoi(line[1:])
			if err != nil {
				fmt.Println(err)
			}

			//Add or subtract value depending on operator
			if operator == "+" {
				resultingFrequency = currentFrequency + change
			} else {
				resultingFrequency = currentFrequency - change
			}

			//Check if currentFrequency/previousFrequency has happened before
			_, exists := previousValues[resultingFrequency]
			//If it has, print that, and set found to true
			if exists == true {
				if verbose {
					fmt.Printf("Current frequency  %d, change of %s%d; resulting frequency  %d, which has \nalready been seen.\n", currentFrequency, operator, change, resultingFrequency)
				}
				found = true
				//If not add currentFrequency to map and print current values
			} else {
				previousValues[resultingFrequency] = true
				if verbose {
					fmt.Printf("Current frequency  %d, change of %s%d; resulting frequency  %d.\n", currentFrequency, operator, change, resultingFrequency)
				}
			}
			currentFrequency = resultingFrequency
		}

		//Close file before reopening
		file.Close()
	}
	return resultingFrequency
}
