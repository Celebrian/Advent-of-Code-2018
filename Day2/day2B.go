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
		//Start by setting all IDs as not exactly one character different
		for id := range IDs {
			IDs[id] = false
		}

		//Extract text on line
		line := scanner.Text()

		//Loop through all IDs
		for id, found := range IDs {
			//Clear output string
			output = ""
			//For each id, loop through the current line
			for i := range line {
				//Compare the same index in id with line
				if id[i] == line[i] {
					//If they are the same, add to output
					output = fmt.Sprintf("%s%s", output, string(id[i]))
					//If the letters are the same check if this has happened before
					if found {
						//If it has, change found to false and break
						IDs[id] = false
						break
					} else {
						//If not, it now has
						IDs[id] = true
					}
				}
			}
			//If, after comparing the new line with one from IDs, there was exactly one letter difference, break
			if IDs[id] {
				break
			}
		}
	}
	return output
}
