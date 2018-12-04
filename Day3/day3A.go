package day3

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

//A is part 1 of day 3
//nolint: gocyclo
func A(fil string, verbose bool, fabricSize int) int {

	//Open file
	file, err := os.Open(fil)
	if err != nil {
		log.Fatal(err)
	}
	//Close file on exit
	defer file.Close()

	//Create fabric in memory
	var fabric = make([][]string, fabricSize)
	for i := range fabric {
		fabric[i] = make([]string, fabricSize)
	}

	//Fill fabric with "  .  " to signify blank
	blank := "  .  "
	for y := fabricSize - 1; y >= 0; y-- {
		for x := 0; x < fabricSize; x++ {
			fabric[x][y] = blank
		}
	}

	//Count overlap
	overlap := 0
	//Create scanner for file
	scanner := bufio.NewScanner(file)

	//Loop while not at end of file
	for scanner.Scan() {

		//Extract text on line
		line := scanner.Text()

		//Split line on space, creating 4 parts: #<id>	@	<fromLeft>,<fromTop>:	<width>x<height>
		parts := strings.SplitN(line, " ", 4)

		//Extract id and remove leading # (it is still string)
		textID := parts[0][1:]

		//Reset default fabricID
		fabricID := "  .  "
		//Check digits and make correct input
		if len(textID) == 1 {
			fabricID = strings.Replace(fabricID, ".", textID, 1)
		} else if len(textID) == 2 {
			fabricID = strings.Replace(fabricID, ". ", textID, 1)
		} else if len(textID) == 3 {
			fabricID = strings.Replace(fabricID, " . ", textID, 1)
		} else if len(textID) == 4 {
			fabricID = strings.Replace(fabricID, " .  ", textID, 1)
		} else {
			log.Fatal("textID has either zero or over three digits, textID has: ", len(textID), " digits")
		}

		//Extract fromLeft and fromTop
		from := strings.Split(parts[2], ",")
		fromLeft, err := strconv.Atoi(from[0])
		if err != nil {
			log.Fatal(err)
		}
		fromTop, err := strconv.Atoi(strings.TrimSuffix(from[1], ":"))
		if err != nil {
			log.Fatal(err)
		}

		//Extract width and height
		rectangle := strings.Split(parts[3], "x")
		width, err := strconv.Atoi(rectangle[0])
		if err != nil {
			log.Fatal(err)
		}
		height, err := strconv.Atoi(rectangle[1])
		if err != nil {
			log.Fatal(err)
		}

		//Traverse the fabric, if string is blank, insert id
		for j := fromTop; j < fromTop+height; j++ {
			for i := fromLeft; i < fromLeft+width; i++ {
				if fabric[j][i] == blank {
					fabric[j][i] = fabricID
				} else if fabric[j][i] != "  X  " {
					fabric[j][i] = "  X  "
					overlap++
				}
			}
		}
	}
	//If verbose, print fabric
	if verbose {
		for _, row := range fabric {
			fmt.Println(row)
		}
	}
	return overlap
}
