package main

import (
	"adventofcode/gosolutions/day1"
	"adventofcode/gosolutions/day10"
	"adventofcode/gosolutions/day11"
	"adventofcode/gosolutions/day12"
	"adventofcode/gosolutions/day13"
	"adventofcode/gosolutions/day14"
	"adventofcode/gosolutions/day15"
	"adventofcode/gosolutions/day16"
	"adventofcode/gosolutions/day2"
	"adventofcode/gosolutions/day3"
	"adventofcode/gosolutions/day4"
	"adventofcode/gosolutions/day5"
	"adventofcode/gosolutions/day6"
	"adventofcode/gosolutions/day7"
	"adventofcode/gosolutions/day8"
	"adventofcode/gosolutions/day9"
	"flag"
	"fmt"
	"log"
	"slices"
)

var dayRuns = map[int]func(){
	1: day1.Day1,
	2: day2.Day2,
	3: day3.Day3,
	4: day4.Day4,
	5: day5.Day5,
	6: day6.Day6,
	7: day7.Day7,
	8: day8.Day8,
	9: day9.Day9,
	10: day10.Day10,
	11: day11.Day11,
	12: day12.Day12,
	13: day13.Day13,
	14: day14.Day14,
	15: day15.Day15,
	16: day16.Day16,
}

func runAll(){
	days := make([]int, 0)
	for day := range dayRuns {
		days = append(days, day)
	}

	slices.Sort(days)

	for _, day := range days {
		fmt.Println("Day: ", day)
		dayRuns[day]()
		fmt.Println("-----------------------")
	}
}

func main() {
	dayPtr := flag.Int("d", 0, "Specify the day to run (1-25)")
	runAllPtr := flag.Bool("a", false, "Run all days")
	helpPtr := flag.Bool("h", false, "Display usage information")

	flag.Parse()

	if *helpPtr {
		printUsage()
		return
	}

	if !*runAllPtr && *dayPtr == 0 {
		log.Fatalf("Error: No arguments provided. Use -h or --help for more information.")
	}

	if *runAllPtr {
		runAll()
	}

	if *dayPtr != 0{
		dayFunc, ok := dayRuns[*dayPtr]
		if !ok {
			log.Fatalf("Error: Day '%d' is not added.", *dayPtr)
		}
		dayFunc()
	}
}

func printUsage() {
	usage := `Usage: go run main.go [options]

Options:
  -d   Specify the day to run (1-25)
  -a   Run all days
  -h   Display usage information
`
	fmt.Print(usage)
}
