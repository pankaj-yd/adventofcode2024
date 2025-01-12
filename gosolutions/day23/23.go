package day23

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
	cycles = make(map[string]struct{})
)

func insertCycle(a, b, c string){
	elems := []string{a, b, c}
	slices.Sort(elems)

	cycle := strings.Join(elems, ",")
	cycles[cycle] = struct{}{}
}

func addThreeCyclesForVertex(graph map[string]map[string]struct{}, u string) {
	queue := make([]string, 0)
	for v := range graph[u] {
		queue = append(queue, v)
	}

	// now for all the neighbors of u, check which pairs of neighbors have edge between them
	for i := 0; i < len(queue); i++ {
		for j := i+1; j < len(queue); j++ {
			v := queue[i]
			w := queue[j]
			_, ok := graph[v][w]
			if ok {
				insertCycle(u, v, w)
			}
		}
	}

}

func solve1(graph map[string]map[string]struct{}) int {
	ans := 0
	vertices := make([]string, 0)
	for key := range graph {
		vertices = append(vertices, key)
	}

	// for every vertex try to find all 3 cycles
	for _, vertex := range vertices {
		addThreeCyclesForVertex(graph, vertex)
	}

	for cycle := range cycles {
		computers := strings.Split(cycle, ",")
		startsWithT := false
		for _, computer := range computers {
			if computer[0] == 't' {
				startsWithT = true
				break
			}
		}
		if startsWithT {
			ans++
		}
	}

	return ans
}

// part b
var (
	password = ""
)

func setPassword(ansSubset []string){
	copySubset := make([]string, 0)
	copySubset = append(copySubset, ansSubset...)

	slices.Sort(copySubset)
	ansStr := strings.Join(copySubset, ",")

	if len(password) < len(ansStr){
		password = ansStr
	}
}

func canAdd(vertex string, graph map[string]map[string]struct{}, ansSubset []string) bool {
	for _, u := range ansSubset {
		_, ok := graph[vertex][u]
		if !ok {
			return false
		}
	}
	return true
}

func makeGroup(vertices []string, idx int, graph map[string]map[string]struct{}, ansGroup []string) {
	if idx == len(vertices) {
		setPassword(ansGroup)
		return
	}

	// add idx to current group if possible
	if canAdd(vertices[idx], graph, ansGroup){
		// add
		ansGroup = append(ansGroup, vertices[idx])

		// process it further
		makeGroup(vertices, idx + 1, graph, ansGroup)

		// remove
		ansGroup = ansGroup[:len(ansGroup) - 1]
	}

	// do not add current vertex
	makeGroup(vertices, idx + 1, graph, ansGroup)
}

func solve2(graph map[string]map[string]struct{}) string {
	vertices := make([]string, 0)
	for key := range graph {
		vertices = append(vertices, key)
	}
	ansGroup := make([]string, 0)
	makeGroup(vertices, 0, graph, ansGroup)

	return password
}

func Day23() {
	// Get the path of the current Go source file
	_, currentFile, _, _ := runtime.Caller(0)

	// Get the directory of the Go source file
	currentDir := filepath.Dir(currentFile)
	parentDir := filepath.Dir(currentDir)
	testCasePath := filepath.Join(parentDir, "..", "testcases", "23.txt")
	file, err := os.Open(testCasePath)

	if err != nil {
		log.Fatal("Unable to open testcase file")
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	graph := make(map[string]map[string]struct{})

	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), "-")
		_, ok := graph[fields[0]]
		if !ok {
			graph[fields[0]] = make(map[string]struct{})
		}
		_, ok = graph[fields[1]]
		if !ok {
			graph[fields[1]] = make(map[string]struct{})
		}

		graph[fields[0]][fields[1]] = struct{}{}
		graph[fields[1]][fields[0]] = struct{}{}
	}

	// compute answers
	fmt.Println("Expected answers:")
	// different rows and columns from sample case
	fmt.Println("Part a: ", 7)
	fmt.Println("Part b: ", "co,de,ka,ta")

	fmt.Println("Your answers:")
	startTime := time.Now()
	fmt.Println("Part a: ", solve1(graph))
	elapsed := time.Since(startTime)
	fmt.Println("Time taken:", elapsed.Milliseconds(), "ms")
	fmt.Println()
	startTime = time.Now()
	fmt.Println("Part b: ", solve2(graph))
	elapsed = time.Since(startTime)
	fmt.Println("Time taken:", elapsed.Milliseconds(), "ms")
}
