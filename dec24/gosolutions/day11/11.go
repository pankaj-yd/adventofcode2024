package day11

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

const (
	numBlinksA = 25
	numBlinksB = 75
)

type Entry struct {
	val, blinks int
}

// cache
var cache = make(map[Entry]int, 0)

// part a
func blink(entry Entry) int {
	if entry.blinks == 0 {
		return 1
	}
	
	if val, ok := cache[entry]; ok {
		return val
	}

	// stone value is 0
	if entry.val == 0 {
		cache[entry] = blink(Entry{1, entry.blinks - 1})
		return cache[entry]
	}

	// stone value has even digits
	if len(strconv.Itoa(entry.val)) & 1 == 0 {
		// split the stone and blink
		valString := strconv.Itoa(entry.val)
		leftString := valString[0:len(valString)/2]
		rightString := valString[len(valString)/2:]

		leftVal, _ := strconv.Atoi(leftString)
		rightVal, _ := strconv.Atoi(rightString)

		cache[entry] = blink(Entry{leftVal, entry.blinks - 1}) + blink(Entry{rightVal, entry.blinks - 1})
		
		return cache[entry]
	}

	// none of the rules apply, apply last rule
	val := entry.val * 2024

	cache[entry] = blink(Entry{val, entry.blinks - 1})
	return cache[entry]
}

func solve1(input []int) int {
	res := 0
	for _, stoneVal := range input {
		res += blink(Entry{stoneVal, numBlinksA})
	}
	return res
}

// part b
func solve2(input []int) int {
	res := 0
	for _, stoneVal := range input {
		res += blink(Entry{stoneVal, numBlinksB})
	}
	return res
}

func Day11() {
	// Get the path of the current Go source file
	_, currentFile, _, _ := runtime.Caller(0)

	// Get the directory of the Go source file
	currentDir := filepath.Dir(currentFile)
    parentDir := filepath.Dir(filepath.Dir(currentDir))
    file, err := os.Open(parentDir + "/testcases/11.txt")

    if err != nil {
        log.Fatal("Unable to open testcase file")
        return
    }

    defer file.Close()

    scanner := bufio.NewScanner(file)
	scanner.Scan()
	inputString := scanner.Text()
	fields := strings.Split(inputString, " ")

	input := make([]int, 0)
	for _, inField := range fields {
		val, _ := strconv.Atoi(inField)
		input = append(input, val)
	}
	
	

	// compute answers
	fmt.Println("Expected answers:")
    fmt.Println("Part a: ", 217812)
    fmt.Println("Part b: ", 259112729857522)
    fmt.Println("Your answers:")
    fmt.Println("Part a: ", solve1(input))
    fmt.Println("Part b: ", solve2(input))
}