package day13

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
)

var p1 = regexp.MustCompile(`\d+`)
var errVal = 10000000000000
type Pos struct {
	x, y int
}

type Machine struct {
	A, B, P Pos
}

// part a
func absInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func calcSteps(machine *Machine) int{
	numerator := absInt(machine.P.x * machine.A.y - machine.P.y * machine.A.x)
	denominator := absInt(machine.A.y * machine.B.x - machine.A.x * machine.B.y)

	if numerator != 0 && denominator != 0 && numerator % denominator == 0{
		return numerator / denominator
	}
	
	return -1
}

func solve1(input []Machine) int {
	cost := 0
	for _, machine := range input {
		nB := calcSteps(&machine)
		nA := calcSteps(&Machine{A: machine.B, B: machine.A, P: machine.P})

		if nA == -1 || nB == -1 {
			continue
		}
		cost += nA * 3 + nB
	}
	return cost
}

// part b
func solve2(input []Machine) int {	
	cost := 0
	for _, machine := range input {
		machine.P.x += errVal
		machine.P.y += errVal

		nB := calcSteps(&machine)
		nA := calcSteps(&Machine{A: machine.B, B: machine.A, P: machine.P})
		
		if nA == -1 || nB == -1 {
			continue
		}

		cost += nA * 3 + nB
	}
	return cost
}

func parseInput(str string) Pos {
	ints := make([]int, 2)
	matches := p1.FindAllString(str, len(str))
	for i, match := range matches {
		val, _ := strconv.Atoi(match)
		ints[i] = val
	}

	pos := Pos{x: ints[0], y: ints[1]}
	return pos
}

func Day13() {
	// Get the path of the current Go source file
	_, currentFile, _, _ := runtime.Caller(0)

	// Get the directory of the Go source file
	currentDir := filepath.Dir(currentFile)
	parentDir := filepath.Dir(filepath.Dir(currentDir))
	file, err := os.Open(parentDir + "/testcases/13.txt")

	if err != nil {
		log.Fatal("Unable to open testcase file")
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	input := make([]Machine, 0)
	
	for scanner.Scan() {
		machine := Machine{}
		machine.A = parseInput(scanner.Text())

		scanner.Scan()
		machine.B = parseInput(scanner.Text())

		scanner.Scan()
		machine.P = parseInput(scanner.Text())

		scanner.Scan()

		input = append(input, machine)
	}

	// compute answers
	fmt.Println("Expected answers:")
	fmt.Println("Part a: ", 480)
	fmt.Println("Part b: ", 875318608908)
	fmt.Println("Your answers:")
	fmt.Println("Part a: ", solve1(input))
	fmt.Println("Part b: ", solve2(input))
}
