package day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

//A is part 1 of day 2
func A(fil string, verbose bool) int {

	//Open file
	file, err := os.Open(fil)
	if err != nil {
		log.Fatal(err)
	}
	//Close file on exit
	defer file.Close()

	//Create scanner for file
	scanner := bufio.NewScanner(file)

	//Store total occurrences of each letter
	totalTwo := 0
	totalThree := 0

	//Loop while not at end of file
	for scanner.Scan() {

		//Extract text on line
		line := scanner.Text()

		//Were there two or three occurrences of a letter in this line
		two := false
		three := false

		//Create a map that maps each letter to a count of how many times it occurs
		charNr := make(map[string]int)

		//Loop over the line, and for each letter count up the occurrence for that letter
		for _, char := range line {
			charNr[string(char)]++
		}

		//Loop over all letters, if some happened twice or three times
		for i := range charNr {
			//Extract amount of occurrences
			nr := charNr[i]

			//If that is 2 or 3, count it
			if nr == 2 {
				two = true
			} else if nr == 3 {
				three = true
			}
		}
		//Start to create output
		output := fmt.Sprint(line, " contains")

		if two {
			totalTwo++
			output = fmt.Sprintf("%s at least one letter that appears exactly two times", output)
		}
		if two && three {
			output = fmt.Sprintf("%s and", output)
		}
		if three {
			totalThree++
			output = fmt.Sprintf("%s at least one letter that appears exactly three times", output)
		}
		if !two && !three {
			output = fmt.Sprintf("%s no letters that appear exactly two or three times.", output)
		}
		if verbose {
			fmt.Println(output)
		}
	}
	return totalTwo * totalThree
}
