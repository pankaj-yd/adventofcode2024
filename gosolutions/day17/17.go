package day17

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

const (
	moduloAndVal  = 8 - 1
	iterationSize = 10e6
	maxRoutines   = 5
	registerA     = 4
	registerB     = 5
	registerC     = 6
)

func getPow(exp int) int {
	if exp == 0 {
		return 1
	}
	return 2 << (exp - 1)
}

func adv(comboOperand int, operandVals map[int]int) {
	numerator := operandVals[registerA]
	denominator := getPow(operandVals[comboOperand])

	operandVals[registerA] = numerator / denominator
}

func bxl(literalOperand int, operandVals map[int]int) {
	operandVals[registerB] = operandVals[registerB] ^ literalOperand
}

func bst(comboOperand int, operandVals map[int]int) {
	operandVals[registerB] = operandVals[comboOperand] & moduloAndVal
}

func jnz(literalOperand int, operandVals map[int]int, instructionPointer *int) bool {
	toJump := true
	if operandVals[registerA] != 0 {
		*instructionPointer = literalOperand
		toJump = false
	}
	return toJump
}

func bxc(_ int, operandVals map[int]int) {
	operandVals[registerB] = operandVals[registerB] ^ operandVals[registerC]
}

func bdv(comboOperand int, operandVals map[int]int) {
	numerator := operandVals[registerA]
	denominator := getPow(operandVals[comboOperand])

	operandVals[registerB] = numerator / denominator
}

func cdv(comboOperand int, operandVals map[int]int) {
	numerator := operandVals[registerA]
	denominator := getPow(operandVals[comboOperand])

	operandVals[registerC] = numerator / denominator
}

func out(comboOperand int, operandVals map[int]int) int {
	return operandVals[comboOperand] & moduloAndVal
}

func solve1(input []int, a, b, c int) []int {
	var (
		operandVals = map[int]int{
			0: 0,
			1: 1,
			2: 2,
			3: 3,
			4: a, // register A
			5: b, // register B
			6: c, // register C
			7: 7,
		}

		instructionPointer = 0
	)

	outVals := make([]int, 0)
	for instructionPointer < len(input) {
		opCode := input[instructionPointer]
		operand := input[instructionPointer+1]
		switch opCode {
		case 0:
			adv(operand, operandVals)
		case 1:
			bxl(operand, operandVals)
		case 2:
			bst(operand, operandVals)
		case 3:
			if !jnz(operand, operandVals, &instructionPointer) {
				instructionPointer -= 2
			}
		case 4:
			bxc(operand, operandVals)
		case 5:
			outVals = append(outVals, out(operand, operandVals))
		case 6:
			bdv(operand, operandVals)
		case 7:
			cdv(operand, operandVals)
		}
		instructionPointer += 2
	}

	return outVals
}

// part b
func checkVals(input []int, i int) int {
	// fmt.Println("Checking range: ", i * iterationSize, (i+1) * iterationSize)
	for j := i * iterationSize; j < (i+1) * iterationSize; j++ {
		outVals := solve1(input, j, 0, 0)
		if reflect.DeepEqual(input, outVals) {
			fmt.Println("Found Ans: ", j)
			return j
		}
	}
	return math.MaxInt64
}

func atomicallyGetAndUpdate(valRef *int64, increment int64) int64 {
	for {
		oldValue := atomic.LoadInt64(valRef)
		newValue := oldValue + increment

		if atomic.CompareAndSwapInt64(valRef, oldValue, newValue) {
			return int64(oldValue)
		}
	}
}

func atomicallyUpdate(valRef *int64, newVal int64) {
	for {
		oldValue := atomic.LoadInt64(valRef)
		if oldValue > newVal {
			if atomic.CompareAndSwapInt64(valRef, oldValue, newVal) {
				fmt.Println("Updated atmoically", atomic.LoadInt64(valRef))
				break
			}
		} else {
			break
		}
	}
}

func solve2(input []int) int64 {
	// to limit max number of goroutines
	sem := make(chan struct{}, maxRoutines)
	// signal := make(chan struct{}, 1)
	// wait group to wait for routines to finish
	var wg sync.WaitGroup

	var i int64
	var res int64 = math.MaxInt64

	for res == math.MaxInt64 {
		// fmt.Println(len(sem))
		if len(sem) == maxRoutines {
			time.Sleep(15 * time.Second)
			fmt.Println("Sleeping for 15 seconds")
			continue
		}

		// fmt.Println(res)
		// if res != math.MaxInt64 {
		// 	signal <- struct{}{}
		// 	break
		// }

		wg.Add(1)
		// fmt.Println("Launching new")
		go func() {
			defer wg.Done()

			// get the slot
			sem <- struct{}{}

			// try atomically updating the iteration counter
			iStart := atomicallyGetAndUpdate(&i, 1)
			
			ans := checkVals(input, int(iStart))
			atomicallyUpdate(&res, int64(ans))
			// release the slot
			<-sem
		}()
		time.Sleep(3 * time.Second)
	}

	fmt.Println("Exited Loop")
	wg.Wait()

	return res
}

func Day17() {
	// Get the path of the current Go source file
	_, currentFile, _, _ := runtime.Caller(0)

	// Get the directory of the Go source file
	currentDir := filepath.Dir(currentFile)
	parentDir := filepath.Dir(currentDir)
	testCasePath := filepath.Join(parentDir, "..", "testcases", "17.txt")
	file, err := os.Open(testCasePath)

	if err != nil {
		log.Fatal("Unable to open testcase file")
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	i := 0
	var a, b, c int
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}

		fields := strings.Split(scanner.Text(), " ")
		intVal, _ := strconv.Atoi(fields[len(fields)-1])

		switch i {
		case 0:
			a = intVal
		case 1:
			b = intVal
		case 2:
			c = intVal
		}
		i++
	}
	scanner.Scan()
	fields := strings.Split(scanner.Text(), " ")
	intFields := strings.Split(fields[1], ",")

	input := make([]int, 0)
	for _, intField := range intFields {
		val, _ := strconv.Atoi(intField)
		input = append(input, val)
	}

	// fmt.Println(registerA, registerB, registerC, input)

	// compute answers
	fmt.Println("Expected answers:")
	// different rows and columns from sample case
	fmt.Println("Part a: ", "6,7,5,2,1,3,5,1,7")
	fmt.Println("Part b: ", 45)
	fmt.Println("Your answers:")
	outVals := solve1(input, a, b, c)
	strs := make([]string, len(outVals))
	for i, num := range outVals {
		strs[i] = strconv.Itoa(num) // Convert integer to string
	}
	fmt.Println("Part a: ", strings.Join(strs, ","))
	fmt.Println("Part b: ", solve2(input))
}
