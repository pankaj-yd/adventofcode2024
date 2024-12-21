package main

import (
	"adventofcode/dec24/gosolutions/day1"
	"adventofcode/dec24/gosolutions/day2"
	"adventofcode/dec24/gosolutions/day3"
	"adventofcode/dec24/gosolutions/day4"
	"adventofcode/dec24/gosolutions/day5"
	"adventofcode/dec24/gosolutions/day6"
	"adventofcode/dec24/gosolutions/day7"
	"adventofcode/dec24/gosolutions/day8"
	"fmt"
	"log"
	"os"
	"strconv"
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
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		log.Fatal("Error: No arguments provided. Use -h or --help for more information.")
	}

	day := 0
	var err error
	for i := 0; i < len(args); i++ {
		switch args[i] {
		case "-h", "--help":
			fmt.Print(`
Usage: go run main.go [options]

Options:
  -d <day>    Specify the day to run (1-25)
			`)
			return
		case "-d": 
			if i+1 >= len(args) {
				log.Fatal("Error: -d flag requires a day number.")
			}

			day, err = strconv.Atoi(args[i+1])
			if err != nil {
				log.Fatalf("Error: Invalid day '%s': %v", args[i+1], err)
			}
			i++

			if day < 1 || day > 25 {
				log.Fatalf("Error: Day '%d' is out of range (1-25).", day)
			}
		default: 
			log.Fatalf("Error: Invalid argument '%s'. Use -h or --help for more information.", args[i])
		}
	}
    
	dayRuns[day]()
}
