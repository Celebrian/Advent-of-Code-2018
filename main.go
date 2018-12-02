package main

import (
	"fmt"

	"github.com/Celebrian/Advent-of-Code-2018/Day1"
)

func main() {
	fmt.Println("Press the Enter Key to move to a new task")
	fmt.Scanln()

	//Day 1
	fmt.Println("\nDay1 Part1")
	day1A := Day1.A("Day1/input.txt")
	fmt.Printf("\nDay 1 Part 1 answer is %d.\n", day1A)

	fmt.Scanln()

	fmt.Println("\nDay1 Part2")
	day1B := Day1.B("Day1/input.txt")
	fmt.Printf("\nDay 1 Part 2 answer is %d.\n", day1B)

	fmt.Scanln()
}
