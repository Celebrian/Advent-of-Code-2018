package Day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

//A is part one of Day 1
func A(fil string, verbose bool) int {
	//Zero out the counter
	currentFrequency := 0
	//Open file
	file, err := os.Open(fil)
	if err != nil {
		log.Fatal(err)
	}
	//Close file on exit
	defer file.Close()

	//Create scanner for file
	scanner := bufio.NewScanner(file)

	//Loop while not at end of file
	for scanner.Scan() {
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

		//Print string if verbose
		if verbose {
			fmt.Printf("Current frequency  %d, change of %d; resulting frequency  %d.\n", previousValue, value, currentFrequency)
		}
	}
	return currentFrequency
}
