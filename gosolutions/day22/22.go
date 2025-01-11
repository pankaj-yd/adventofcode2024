package day22

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"time"
)

const (
	mod = 16777216
	multBy64 = 6
	divBy32 = 5
	multBy2048 = 11
	partAIterations = 2000
)

func mix(currNum, toMix int) int {
	return currNum^toMix
}

func prune(currNum int) int {
	return currNum % mod
}

func generateNextNumber(currNum int) int {
	// mult by 64
	toMix := currNum << multBy64
	// mix
	currNum = mix(currNum, toMix)
	// prune
	currNum = prune(currNum)

	// div by 32
	toMix = currNum >> divBy32
	// mix
	currNum = mix(currNum, toMix)
	// prune
	currNum = prune(currNum)

	// mult by 2048
	toMix = currNum << multBy2048
	// mix
	currNum = mix(currNum, toMix)
	// prune
	currNum = prune(currNum)

	return currNum
}

func solve1(secrets []int) int {
	ans := 0
	for _, secret := range secrets {
		currNum := secret
		for i := 0; i < partAIterations; i++ {
			currNum = generateNextNumber(currNum)
		}
		ans += currNum
	}
	

	return ans
}

type Sequence struct {
	a, b, c, d int
}

// part b
func solve2(secrets []int) int {
	sequences := make([][]int, 0)
	for _, secret := range secrets {
		currNum := secret
		currSeq := []int{currNum%10}
		for i := 0; i < partAIterations; i++ {
			currNum = generateNextNumber(currNum)
			currSeq = append(currSeq, currNum%10)
		}
		sequences = append(sequences, currSeq)
	}
	// fmt.Println(sequences)
	
	diffSequences := make([][]int, 0)
	
	for _, sequence := range sequences {
		currSeq := make([]int, 0)
		for i := 1; i < len(sequence); i++ {
			currSeq = append(currSeq, sequence[i] - sequence[i-1])
		}

		diffSequences = append(diffSequences, currSeq)
	}

	uniqueSeq := make(map[Sequence]int, 0)
	ans := 0
	for i := range len(diffSequences) {
		currSeq := make(map[Sequence]struct{}, 0)
		for j := 0; j < len(diffSequences[i]) - 3; j++ {
			seq := Sequence{a: diffSequences[i][j], b: diffSequences[i][j+1], c: diffSequences[i][j+2], d: diffSequences[i][j+3]}

			_, ok := currSeq[seq]
			if !ok {
				currSeq[seq] = struct{}{}
				uniqueSeq[seq] += sequences[i][j+4]
				ans = max(ans, uniqueSeq[seq])
			}
		}
	}

	return ans
}

func Day22() {
	// Get the path of the current Go source file
	_, currentFile, _, _ := runtime.Caller(0)

	// Get the directory of the Go source file
	currentDir := filepath.Dir(currentFile)
	parentDir := filepath.Dir(currentDir)
	testCasePath := filepath.Join(parentDir, "..", "testcases", "22.txt")
	file, err := os.Open(testCasePath)

	if err != nil {
		log.Fatal("Unable to open testcase file")
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	secrets := make([]int, 0)

	for scanner.Scan() {
		intVal, _ := strconv.Atoi(scanner.Text())
		secrets = append(secrets, intVal)
	}

	// compute answers
	fmt.Println("Expected answers:")
	// different rows and columns from sample case
	fmt.Println("Part a: ", 0)
	fmt.Println("Part b: ", 0)

	fmt.Println("Your answers:")
	startTime := time.Now()
	fmt.Println("Part a: ", solve1(secrets))
	elapsed := time.Since(startTime)
	fmt.Println("Time taken:", elapsed.Milliseconds(), "ms")
	fmt.Println()
	startTime = time.Now()
	fmt.Println("Part b: ", solve2(secrets))
	elapsed = time.Since(startTime)
	fmt.Println("Time taken:", elapsed.Milliseconds(), "ms")
}
