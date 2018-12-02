package main

import (
	"flag"
	"fmt"

	"github.com/Celebrian/Advent-of-Code-2018/Day1"
	"github.com/Celebrian/Advent-of-Code-2018/Day2"
)

func main() {
	//Flag for verbose or not
	verbose := flag.Bool("v", false, "Should the program be verbose or not(default not).")
	flag.Parse()

	fmt.Println("Press the Enter Key to start, and move to a new task")
	fmt.Scanln()

	//Day 1
	day1A := Day1.A("Day1/input.txt", *verbose)
	fmt.Printf("\nDay 1 Part 1 answer is %d.\n", day1A)

	fmt.Scanln()

	day1B := Day1.B("Day1/input.txt", *verbose)
	fmt.Printf("\nDay 1 Part 2 answer is %d.\n", day1B)

	fmt.Scanln()

	day2A := day2.A("Day2/input.txt", *verbose)
	fmt.Printf("\nDay 2 Part 1 answer is %d.\n", day2A)

	fmt.Scanln()

	day2B := day2.B("Day2/input.txt", *verbose)
	fmt.Printf("\nDay 2 Part 2 answer is %d.\n", day2A)

}
