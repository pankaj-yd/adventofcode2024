package day18

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"
)

var (
	rows = 6
	cols = 6
	fallenBytes = 12

	north = Dir{dx: -1, dy: 0}
	east = Dir{dx: 0, dy: 1}
	west = Dir{dx: 0, dy: -1}
	south = Dir{dx: 1, dy: 0}
	dirs = []Dir{north, east, south, west}

	path = make(map[Pos]struct{})
)

type Dir struct {
	dx, dy int
}

type Pos struct {
	x, y int
}

func (pos *Pos) inBound() bool {
	return pos.x >= 0 && pos.x <= rows && pos.y >= 0 && pos.y <= cols
}

func (pos *Pos) move(dir Dir) Pos {
	return Pos{x: pos.x + dir.dx, y: pos.y + dir.dy}
}

func bfs(walls map[Pos]struct{}, distance map[Pos]int) map[Pos]Pos{
	// fmt.Println(walls)
	startPos := Pos{0, 0}
	parents := make(map[Pos]Pos, 0)
	queue := make([]Pos, 0)
	queue = append(queue, startPos)
	distance[startPos] = 0

	for len(queue) != 0 {
		// fmt.Println(queue)
		newQueue := make([]Pos, 0)

		// relax edges from each position in all dirs
		for _, pos := range queue {
			for _, dir := range dirs {
				newPos := pos.move(dir)
				_, isWall := walls[newPos]

				// if invalid position, or a wall
				if !newPos.inBound() || isWall {
					continue
				}

				newDist := distance[pos] + 1

				dist, ok := distance[newPos]
				if !ok || dist > newDist {
					distance[newPos] = newDist
					newQueue = append(newQueue, newPos)
					parents[newPos] = pos
				}
			}
		}

		queue = newQueue
	}

	return parents
}

func solve1(input []Pos, currFallen int) int {
	walls := make(map[Pos]struct{}, 0)
	for i := 0; i < currFallen; i++ {
		walls[input[i]] = struct{}{}
	}

	distance := make(map[Pos]int, 0)
	parents := bfs(walls, distance)
	endPos := Pos{rows, cols}
	ans, ok := distance[endPos]
	if !ok {
		return -1
	}

	// update the path
	for {
		path[endPos] = struct{}{}
		parentPos, ok := parents[endPos]

		if !ok {
			break
		}
		endPos = parentPos
	}

	return ans
}

// part b
func solve2(input []Pos) string {
	for i := fallenBytes + 1; i < len(input); i++ {
		// if input[i-1] is a part of path
		// then re-compute the path
		_, ok := path[input[i-1]]

		// not a part, continue with next
		if !ok {
			continue
		}

		minDist := solve1(input, i)
		if minDist == -1 {
			ansPos := input[i-1]
			return fmt.Sprintf("%d,%d", ansPos.y, ansPos.x)
		}
		
	}
	return ""
}

func Day18() {
	// Get the path of the current Go source file
	_, currentFile, _, _ := runtime.Caller(0)

	// Get the directory of the Go source file
	currentDir := filepath.Dir(currentFile)
	parentDir := filepath.Dir(currentDir)
	testCasePath := filepath.Join(parentDir, "..", "testcases", "18.txt")
	file, err := os.Open(testCasePath)

	if err != nil {
		log.Fatal("Unable to open testcase file")
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	input := make([]Pos, 0)

	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), ",")
		var pos Pos
		pos.y, _ = strconv.Atoi(fields[0])
		pos.x, _ = strconv.Atoi(fields[1])
		input = append(input, pos)
	}

	// compute answers
	fmt.Println("Expected answers:")
	// different rows and columns from sample case
	fmt.Println("Part a: ", 22)
	fmt.Println("Part b: ", "6,1")
	fmt.Println()

	fmt.Println("Your answers:")
	startTime := time.Now()
	fmt.Println("Part a: ", solve1(input, fallenBytes))
	elapsed := time.Since(startTime)
	fmt.Println("Time taken:", elapsed.Milliseconds(), "ms")
	fmt.Println()
	startTime = time.Now()
	fmt.Println("Part b: ", solve2(input))
	elapsed = time.Since(startTime)
	fmt.Println("Time taken:", elapsed.Milliseconds(), "ms")
}
