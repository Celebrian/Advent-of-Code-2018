package day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

//B is part 2 of day 2
func B(fil string, verbose bool) string {

	//Open file
	file, err := os.Open(fil)
	if err != nil {
		log.Fatal(err)
	}
	//Close file on exit
	defer file.Close()

	//Create scanner for file
	scanner := bufio.NewScanner(file)

	//Create map with ID and a bool for if one character is different
	IDs := make(map[string]bool)

	//Create output string
	var output string

	//Loop while not at end of file
	for scanner.Scan() {

		//Extract text on line
		line := scanner.Text()

		//Loop through all IDs
		for id := range IDs {
			//Clear output string
			output = ""
			//For each id, loop through the current line
			for i := range line {
				//Compare the same index in id with line
				if id[i] == line[i] {
					//If they are the same, add to output
					output = fmt.Sprintf("%s%s", output, string(id[i]))
				} else {
					//If the letters are not the same check if this is the first time
					if IDs[id] {
						//If the letters have been the same before, set id to false and break to next id
						IDs[id] = false
						break
					} else {
						//If this is the first time, set id to true
						IDs[id] = true
					}
				}
			}
			//If, after comparing the new line with one from IDs, there was exactly one letter difference, we found it
			if IDs[id] {
				return output
			}
			if verbose {
				fmt.Printf("\n%s and %s differ by more than one letter", line, id)
			}
		}
		//After comparing line with all IDs, add this line to IDs
		IDs[line] = false
	}
	fmt.Println(IDs)
	return "Did not find exactly one"
}
