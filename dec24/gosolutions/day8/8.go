package day8

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)


// global variables
var (
	rows, cols int
)

var towers map[byte][]Pos

// interface
type Pos struct {
	x, y int
}

// interface methods
func (pos *Pos) valid() bool {
	return pos.x >= 0 && pos.x < rows && pos.y >= 0 && pos.y < cols;
}

func (pos *Pos) add(pos2 *Pos) Pos {
	return Pos{x: pos.x + pos2.x, y: pos.y + pos2.y}
}

func (pos *Pos) sub(pos2 *Pos) Pos {
	return Pos{x: pos.x - pos2.x, y: pos.y - pos2.y}
}

// pre-processing
func preProcess(input []string) {
	rows = len(input)
	cols = len(input[0])

	towers = make(map[byte][]Pos, 0)
	
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			ch := input[i][j]
			if 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || '0' <= ch && ch <= '9' {
				if _, ok := towers[ch]; !ok{
					towers[ch] = make([]Pos, 0)
				}
				towers[ch] = append(towers[ch], Pos{x: i, y: j})
			}
		} 
	}	
}

// part a
func addNode(nodes map[Pos]struct{}, pos1, pos2 *Pos) {
	delta := pos1.sub(pos2)
	nodePos := pos1.add(&delta)
	if nodePos.valid() {
		nodes[nodePos] = struct{}{}
	}
}

func solve1() int {
	nodes := make(map[Pos]struct{}, 0)
	for _, positions := range towers {
		for i := 0; i < len(positions); i++ {
			for j := i+1; j < len(positions); j++ {
				addNode(nodes, &positions[i], &positions[j])
				addNode(nodes, &positions[j], &positions[i])
			}
		}
	}
	
	return len(nodes)
}

// part b
func addAllNodes(nodes map[Pos]struct{}, pos1, pos2 *Pos) {
	nodePos := *pos1
	delta := pos1.sub(pos2)
	
	for {
		nodePos = nodePos.add(&delta)
		if !nodePos.valid() {
			break
		}

		nodes[nodePos] = struct{}{}
	}
}

func solve2() int {
	nodes := make(map[Pos]struct{}, 0)
	for _, positions := range towers {
		for i := 0; i < len(positions); i++ {
			for j := i+1; j < len(positions); j++ {
				nodes[positions[i]] = struct{}{}
				nodes[positions[j]] = struct{}{}

				addAllNodes(nodes, &positions[i], &positions[j])
				addAllNodes(nodes, &positions[j], &positions[i])
			}
		}
	}
	
	return len(nodes)
}

func Day8() {
	// Get the path of the current Go source file
	_, currentFile, _, _ := runtime.Caller(0)

	// Get the directory of the Go source file
	currentDir := filepath.Dir(currentFile)
    parentDir := filepath.Dir(filepath.Dir(currentDir))
    file, err := os.Open(parentDir + "/testcases/8.txt")

    if err != nil {
        log.Fatal("Unable to open testcase file")
        return
    }

    defer file.Close()

	// initialise map
	input := make([]string, 0)

    scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		input = append(input, scanner.Text())
	}

	// pre-prcoess
	preProcess(input)

	// compute answers
	fmt.Println("Expected answers:")
    fmt.Println("Part a: ", 332)
    fmt.Println("Part b: ", 1174)
    fmt.Println("Your answers:")
    fmt.Println("Part a: ", solve1())
    fmt.Println("Part b: ", solve2())
}