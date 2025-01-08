package day20

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

const (
	maxCheats = 20
	minSaving = 100
	cheatASize = 2
	cheatBSize = 20
)

var (
	rows, cols int
	startPos = Pos{}
	endPos   = Pos{}
	wall     = byte('#')
	minDist  = math.MaxInt64
)

type Dir struct {
	dx, dy int
}

type Pos struct {
	x, y int
}

type Cheat struct {
	start, end Pos
}

func (pos *Pos) inBound() bool {
	return pos.x >= 0 && pos.x < rows && pos.y >= 0 && pos.y < cols
}

func (pos *Pos) move(dir Dir) Pos {
	return Pos{x: pos.x + dir.dx, y: pos.y + dir.dy}
}

func (pos *Pos) distance(newPos Pos) int {
	return absInt(pos.x - newPos.x) + absInt(pos.y - newPos.y)
}

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

var (
	north = Dir{dx: -1, dy: 0}
	east  = Dir{dx: 0, dy: 1}
	west  = Dir{dx: 0, dy: -1}
	south = Dir{dx: 1, dy: 0}

	dirs = []Dir{north, east, south, west}

	savings = make(map[Cheat]int, 0)

	distanceFromStart = make(map[Pos]int, 0)
	distanceFromEnd   = make(map[Pos]int, 0)
)

func printGrid(grid [][]byte) {
	for _, row := range grid {
		fmt.Println(string(row))
	}
	fmt.Println()
}

func bfs(grid [][]byte, pos Pos, distance map[Pos]int) {
	queue := []Pos{pos}
	distance[pos] = 0

	for len(queue) != 0 {
		newQueue := make([]Pos, 0)
		for _, pos := range queue {
			// move in all dirs from state
			for _, dir := range dirs {
				newPos := pos.move(dir)
				newDist := distance[pos] + 1

				if !newPos.inBound() || grid[pos.x][pos.y] == wall {
					continue
				}

				currDist, ok := distance[newPos]
				if !ok || currDist > newDist {
					distance[newPos] = newDist
					newQueue = append(newQueue, newPos)
				}
			}
		}

		queue = newQueue
	}
}

func getAllPointsAtDistance(pos Pos, d int) []Pos {
	rowStart := max(0, pos.x - d)
	rowEnd := min(pos.x + d, rows - 1)
	colsStart := max(0, pos.y - d)
	colsEnd := min(pos.y + d, cols - 1)

	ansPos := make([]Pos, 0)
	for i := rowStart; i <= rowEnd; i++ {
		for j := colsStart; j <= colsEnd; j++ {
			adjPos := Pos{i, j}
			if adjPos.inBound() && pos.distance(adjPos) <= d {
				ansPos = append(ansPos, adjPos)
			}
		}
	}
	return ansPos
}

func solve1(inputGrid [][]byte, cheatSize int) int {
	if len(distanceFromStart) == 0 {
		rows = len(inputGrid)
		cols = len(inputGrid[0])
		bfs(inputGrid, startPos, distanceFromStart)
		bfs(inputGrid, endPos, distanceFromEnd)
		minDist = distanceFromStart[endPos]
	}

	savings = make(map[Cheat]int)

	for i := 0; i < len(inputGrid); i++ {
		for j := 0; j < len(inputGrid[0]); j++ {
			if inputGrid[i][j] == wall {
				continue
			}

			pos := Pos{i, j}
			adjPos := getAllPointsAtDistance(pos, cheatSize)
			for _, newPos := range adjPos {
				// invalid newPos
				if !newPos.inBound() || inputGrid[newPos.x][newPos.y] == wall {
					continue
				}
				
				// calc dist
				// start->pos + pos->newPos + newPos->end
				newSaving := minDist - (distanceFromStart[pos] + pos.distance(newPos) + distanceFromEnd[newPos])
				cheat := Cheat{start: pos, end: newPos}
				existingSavings, ok := savings[cheat];

				if newSaving > 0 && (!ok || existingSavings > newSaving) {
					savings[cheat] = newSaving
				}
			}
		}
	}

	ans := 0
	ansMap := make(map[int]int, 0)
	for _, v := range savings {
		if v >= minSaving {
			ans++
		}
		ansMap[v]++
	}
	fmt.Println("Map showing (Savings : Number of cheats)")
	fmt.Println(ansMap)
	return ans
}

func Day20() {
	// Get the path of the current Go source file
	_, currentFile, _, _ := runtime.Caller(0)

	// Get the directory of the Go source file
	currentDir := filepath.Dir(currentFile)
	parentDir := filepath.Dir(currentDir)
	testCasePath := filepath.Join(parentDir, "..", "testcases", "20.txt")
	file, err := os.Open(testCasePath)

	if err != nil {
		log.Fatal("Unable to open testcase file")
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	grid := make([][]byte, 0)

	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		grid = append(grid, []byte(scanner.Text()))
		sIdx := strings.Index(scanner.Text(), "S")
		eIdx := strings.Index(scanner.Text(), "E")

		if sIdx != -1 {
			startPos.x = len(grid) - 1
			startPos.y = sIdx
		}

		if eIdx != -1 {
			endPos.x = len(grid) - 1
			endPos.y = eIdx
		}
	}

	// compute answers
	fmt.Println("Expected answers:")
	// different rows and columns from sample case
	fmt.Println("Part a: ", 0)
	fmt.Println("Part b: ", 0)

	fmt.Println("Your answers:")
	startTime := time.Now()
	fmt.Println("Part a: ", solve1(grid, cheatASize))
	elapsed := time.Since(startTime)
	fmt.Println("Time taken:", elapsed.Milliseconds(), "ms")
	fmt.Println()
	startTime = time.Now()
	fmt.Println("Part b: ", solve1(grid, cheatBSize))
	elapsed = time.Since(startTime)
	fmt.Println("Time taken:", elapsed.Milliseconds(), "ms")
}
