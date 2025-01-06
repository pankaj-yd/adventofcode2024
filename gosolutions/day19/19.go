package day19

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"slices"
	"strings"
	"time"
)

var (
	counters =  map[string]int{"":1}
)

func processPatterns(patterns map[string]struct{}) {
	allPatterns := make([]string, 0)
	for key := range patterns {
		allPatterns = append(allPatterns, key)
	}

	slices.SortFunc(allPatterns, func(a, b string) int{
		if len(a) < len(b) || len(a) == len(b) && a < b {
			return -1
		}
		return 1
	})

	for _, pattern := range allPatterns {
		calcWays(patterns, pattern)
		counters[pattern]++
	}
}

func calcWays(patterns map[string]struct{}, design string) int {
	ways, ok := counters[design]
	if ok {
		return ways
	}

	totalWays := 0
	for i := 1; i < len(design); i++ {
		_, ok := patterns[design[:i]]
		if !ok {
			continue
		}
		totalWays += calcWays(patterns, design[i:])
	}

	counters[design] = totalWays
	return totalWays
}

func solve1(patterns map[string]struct{}, designs []string) int {
	counters :=  make(map[string]int)
	counters[""] = 1

	// add all patterns in counters
	processPatterns(patterns)

	ans := 0
	for _, design := range designs {
		if calcWays(patterns, design) > 0 {
			ans++
		}
	}
	return ans
}

// part b
func solve2(patterns map[string]struct{}, designs []string) int {
	// add all patterns in counters
	if len(counters) == 1 {
		processPatterns(patterns)
	}

	ans := 0
	for _, design := range designs {
		ans += counters[design]
	}
	return ans
}

func Day19() {
	// Get the path of the current Go source file
	_, currentFile, _, _ := runtime.Caller(0)

	// Get the directory of the Go source file
	currentDir := filepath.Dir(currentFile)
	parentDir := filepath.Dir(currentDir)
	testCasePath := filepath.Join(parentDir, "..", "testcases", "19.txt")
	file, err := os.Open(testCasePath)

	if err != nil {
		log.Fatal("Unable to open testcase file")
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	patterns := make(map[string]struct{}, 0)

	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}

		fields := strings.Split(scanner.Text(), ",")
		for _, field := range fields {
			field = strings.TrimSpace(field)
			patterns[field] = struct{}{}
		}
	}

	designs := make([]string, 0)
	for scanner.Scan() {
		designs = append(designs, scanner.Text())
	}

	// compute answers
	fmt.Println("Expected answers:")
	// different rows and columns from sample case
	fmt.Println("Part a: ", 22)
	fmt.Println("Part b: ", "6,1")
	fmt.Println()

	fmt.Println("Your answers:")
	startTime := time.Now()
	fmt.Println("Part a: ", solve1(patterns, designs))
	elapsed := time.Since(startTime)
	fmt.Println("Time taken:", elapsed.Milliseconds(), "ms")
	fmt.Println()
	startTime = time.Now()
	fmt.Println("Part b: ", solve2(patterns, designs))
	elapsed = time.Since(startTime)
	fmt.Println("Time taken:", elapsed.Milliseconds(), "ms")
}
