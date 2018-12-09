package main

import (
	"flag"
	"fmt"

	"github.com/Celebrian/Advent-of-Code-2018/Day1"
	"github.com/Celebrian/Advent-of-Code-2018/Day2"

	"github.com/Celebrian/Advent-of-Code-2018/Day3"
)

func main() {
	//Flag for verbose or not
	verbose := flag.Bool("v", false, "Should the program be verbose or not(default not).")
	flag.Parse()

	fmt.Println("Press the Enter Key to start, and move to a new task")
	_, _ = fmt.Scanln()

	//Day 1
	day1A := Day1.A("Day1/input.txt", *verbose)
	fmt.Printf("\nDay 1 Part 1 answer is %d.\n", day1A)
	_, _ = fmt.Scanln()
	day1B := Day1.B("Day1/input.txt", *verbose)
	fmt.Printf("\nDay 1 Part 2 answer is %d.\n", day1B)

	_, _ = fmt.Scanln()

	//Day 2
	day2A := day2.A("Day2/input.txt", *verbose)
	fmt.Printf("\nDay 2 Part 1 answer is %d.\n", day2A)
	_, _ = fmt.Scanln()
	day2B := day2.B("Day2/input.txt", *verbose)
	fmt.Printf("\nDay 2 Part 2 answer is %s.\n", day2B)

	_, _ = fmt.Scanln()

	//Day 3
	day3A, day3B := day3.AAndB("Day3/input.txt", *verbose, 1000)
	fmt.Printf("\nDay 3 Part 1 answer is %d.\n\nDay 3 Part 2 answer is %d.\n", day3A, day3B)
}
