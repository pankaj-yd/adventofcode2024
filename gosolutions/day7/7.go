package day7

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

type Equation struct {
	result int64
	operands []int64
}

type Operator int
const (
	Plus Operator = iota
	Mult
	Concat
)

func (eq Equation) canCaliberate(currIdx int, currVal int64, operators []Operator) bool {
	if currIdx == len(eq.operands) {
		return currVal == eq.result
	}

	for _, operator := range operators {
		switch operator {
		case Plus:
			if eq.canCaliberate(currIdx + 1, currVal + eq.operands[currIdx], operators) {
				return true
			}
		case Mult:
			if eq.canCaliberate(currIdx + 1, currVal * eq.operands[currIdx], operators) {
				return true
			}
		case Concat:
			concatInt, _ := strconv.ParseInt(strconv.FormatInt(currVal, 10) + strconv.FormatInt(eq.operands[currIdx], 10), 10, 64)
			if eq.canCaliberate(currIdx + 1, concatInt, operators) {
				return true
			}
		
		}
	}
	return false
}

func solve1(input []Equation) int64 {
	var res int64
	for _, equation := range input {
		if equation.canCaliberate(1, equation.operands[0], []Operator{Plus, Mult}){
			res += equation.result
		}
	}
	return res
}

func solve2(input []Equation) int64 {
	var res int64
	for _, equation := range input {
		if equation.canCaliberate(1, equation.operands[0], []Operator{Plus, Mult, Concat}){
			res += equation.result
		}
	}
	return res
}

func Day7() {
	// Get the path of the current Go source file
	_, currentFile, _, _ := runtime.Caller(0)

	// Get the directory of the Go source file
	currentDir := filepath.Dir(currentFile)
    parentDir := filepath.Dir(filepath.Dir(currentDir))
    file, err := os.Open(parentDir + "/testcases/7.txt")

    if err != nil {
        log.Fatal("Unable to open testcase file")
        return
    }

    defer file.Close()

	// initialise map
	input := make([]Equation, 0)

    scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		fields := strings.Split(scanner.Text(), ":")
		
		res, _ := strconv.ParseInt(fields[0], 10, 64)
		operands := make([]int64, 0)
		fields[1] = strings.TrimSpace(fields[1])
		for _, val := range strings.Split(fields[1], " "){
			intVal, _ := strconv.ParseInt(val, 10, 64)
			operands = append(operands, intVal)
		}
		input = append(input, Equation{result: res, operands: operands})
	}
	fmt.Println("Expected answers:")
    fmt.Println("Part a: ", 3749)
    fmt.Println("Part b: ", 11387)
    fmt.Println("Your answers:")
    fmt.Println("Part a: ", solve1(input))
    fmt.Println("Part b: ", solve2(input))
}