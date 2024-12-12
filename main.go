package main

import (
	"adventofcode/dec24/gosolutions"
	"fmt"
	"log"
	"os"
	"strconv"
)

var dayRuns = map[int]func(){
	1: gosolutions.Day1,
	2: gosolutions.Day2,
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
