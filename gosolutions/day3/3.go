package day3

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"
)

var (
	p1 = regexp.MustCompile(`mul\(([0-9]?[0-9]?[0-9]),([0-9]?[0-9]?[0-9])\)`)
	p2 = regexp.MustCompile(`mul\(([0-9]?[0-9]?[0-9]),([0-9]?[0-9]?[0-9])\)|don't\(\).*?do\(\)|don't\(\)`)
)

func solveMul(input *string) int{
	matches := p1.FindAllStringSubmatch(*input, len(*input))
	res := 0
	for _, match := range matches {
		n1, _ := strconv.Atoi(match[1])
		n2, _ := strconv.Atoi(match[2])
		res += n1 * n2
	}

	return res
}

func solveMulConditional(input *string) int{
	matches := p2.FindAllStringSubmatch(*input, len(*input))
	res := 0
	for _, match := range matches {
		if strings.HasPrefix(match[0], "mul") {
			n1, _ := strconv.Atoi(match[1])
			n2, _ := strconv.Atoi(match[2])
			res += n1 * n2 
		} else if match[0] == "don't()" {
			break
		}
	}

	return res
}

func Day3() {
	// Get the path of the current Go source file
	_, currentFile, _, _ := runtime.Caller(0)

	// Get the directory of the Go source file
	currentDir := filepath.Dir(currentFile)
    parentDir := filepath.Dir(filepath.Dir(currentDir))
    file, err := os.Open(parentDir + "/testcases/3.txt")

    if err != nil {
        log.Fatal("Unable to open testcase file")
        return
    }

    defer file.Close()

    scanner := bufio.NewScanner(file)
	input := ""
	for scanner.Scan(){
		input += scanner.Text()
	}

	println("Expected answers:")
    println("Part a: ", 161)
    println("Part b: ", 48)
    println()
    println("Your answers:")
    println("Part a: ", solveMul(&input))
    println("Part b: ", solveMulConditional(&input))
}