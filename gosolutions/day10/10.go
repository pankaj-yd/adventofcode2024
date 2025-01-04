package day10

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

type Dir struct {
	dx, dy int
}

var (
	rows, cols int
	north = Dir{dx: -1, dy: 0}
	east  = Dir{dx: 0, dy: 1}
	south = Dir{dx: 1, dy: 0}
	west  = Dir{dx: 0, dy: -1}

	dirs = []Dir{north, east, south, west}
)

type Pos struct {
	x, y int
}

func (pos *Pos) valid() bool {
	return pos.x >= 0 && pos.x < rows && pos.y >= 0 && pos.y < cols;
}

func (pos *Pos) move(dir Dir) Pos {
	return Pos{x: pos.x + dir.dx, y: pos.y + dir.dy}
}

func dfs(input []string, pos Pos, currScore map[Pos]int, ) {
	if !pos.valid(){
		return
	}

	if input[pos.x][pos.y] == '9' {
		currScore[pos] += 1
		return
	}

	for _, dir := range dirs {
		newPos := pos.move(dir)
		if newPos.valid() && input[newPos.x][newPos.y] == input[pos.x][pos.y] + 1 {
			dfs(input, newPos, currScore)
		}
	}
}

// part a
func solve1(input []string) int {
	score := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if input[i][j] != '0' {
				continue
			}
			currScore := make(map[Pos]int, 0)
			dfs(input, Pos{x: i, y: j}, currScore)
			score += len(currScore)
		}
	}
	return score
}

// part b
func solve2(input []string) int {
	score := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if input[i][j] != '0' {
				continue
			}
			currScore := make(map[Pos]int, 0)
			dfs(input, Pos{x: i, y: j}, currScore)

			// add all scores to total score
			for _, val := range currScore {
				score += val
			}
		}
	}
	return score
}

func Day10() {
	// Get the path of the current Go source file
	_, currentFile, _, _ := runtime.Caller(0)

	// Get the directory of the Go source file
	currentDir := filepath.Dir(currentFile)
    parentDir := filepath.Dir(filepath.Dir(currentDir))
    file, err := os.Open(parentDir + "/testcases/10.txt")

    if err != nil {
        log.Fatal("Unable to open testcase file")
        return
    }

    defer file.Close()

    scanner := bufio.NewScanner(file)
	input := make([]string, 0)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	rows = len(input)
    cols = len(input[0])

	// compute answers
	fmt.Println("Expected answers:")
    fmt.Println("Part a: ", 36)
    fmt.Println("Part b: ", 81)
    fmt.Println("Your answers:")
    fmt.Println("Part a: ", solve1(input))
    fmt.Println("Part b: ", solve2(input))
}