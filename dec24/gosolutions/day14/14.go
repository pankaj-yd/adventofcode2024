package day14

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
)

var p1 = regexp.MustCompile(`[-]?\d+`)

var (
	rows = 103
	cols = 101
	time = 100
)

type Pos struct {
	x, y int
}

type Dir struct {
	dx, dy int
}

type Robot struct {
	pos Pos
	vel Dir
}

type State struct {
	robots []Robot
	time int
	score int
}

func (robot *Robot) move(steps int) {
	robot.pos.x += steps * robot.vel.dx
	robot.pos.x = (robot.pos.x % cols + cols) % cols

	robot.pos.y += steps * robot.vel.dy
	robot.pos.y = (robot.pos.y % rows + rows) % rows
}

func getQuadrant(robPos *Pos) int {
	// 1 2
	// 4 3
	if robPos.x < cols/2 && robPos.y < rows/2{
		return 1
	}

	if robPos.x > cols/2 && robPos.y < rows/2{
		return 2
	}

	if robPos.x > cols/2 && robPos.y > rows/2{
		return 3
	}

	if robPos.x < cols/2 && robPos.y > rows/2{
		return 4
	}

	return 0
}

func getSafetyScore(robots []Robot) int {
	score := 1
	quadrants := make([]int, 5)
	for _, robot := range robots {
		quadrants[getQuadrant(&robot.pos)]++
	}
	for i, val := range quadrants {
		if i == 0 {
			continue
		} 
		score *= val
	}
	return score
}

func solve1(robots []Robot) int {
	for i := range robots {
		robots[i].move(time)
	}
	// printGrid(input)

	return getSafetyScore(robots)
}

func printGrid(robots []Robot){
	grid := make([][]byte, 0)
	for i := 0; i < rows; i++ {
		row := make([]byte, cols)
		for j := range row {
			row[j] = '.'
		}
		grid = append(grid, row)
	}

	for _, robot := range robots {
		grid[robot.pos.y][robot.pos.x] = '*'
	}

	for _, row := range grid {
		fmt.Println(string(row))
	}
}

// part b
func solve2(robots []Robot) int {
	minscoreState := State{score: math.MaxInt64}
	for i := 0; i < rows*cols; i++ {
		copyRobots := make([]Robot, len(robots))
		copy(copyRobots, robots)
		for j := range copyRobots {
			copyRobots[j].move(i)
		}
		currState := State{robots: copyRobots, time: i, score: getSafetyScore(copyRobots)}
		if minscoreState.score > currState.score {
			minscoreState = currState
		}
	}

	printGrid(minscoreState.robots)

	return minscoreState.time
}

// utils
func parseLine(str string) *Robot {
	matches := p1.FindAllString(str, len(str))
	
	intFields := make([]int, 4)
	for i, match := range matches {
		val, _ := strconv.Atoi(match)
		intFields[i] = val
	}

	robPos := Pos{x: intFields[0], y: intFields[1]}
	robVel := Dir{dx: intFields[2], dy: intFields[3]}

	return &Robot{pos: robPos, vel: robVel}
}

func copySlice(robots []*Robot) []Robot {
	copy := make([]Robot, 0)

	for _, robot := range robots {
		copy = append(copy, *robot)
	}
	return copy
}

func Day14() {
	// Get the path of the current Go source file
	_, currentFile, _, _ := runtime.Caller(0)

	// Get the directory of the Go source file
	currentDir := filepath.Dir(currentFile)
	parentDir := filepath.Dir(filepath.Dir(currentDir))
	file, err := os.Open(parentDir + "/testcases/14.txt")

	if err != nil {
		log.Fatal("Unable to open testcase file")
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	input := make([]*Robot, 0)
	
	for scanner.Scan() {
		robot := parseLine(scanner.Text())
		input = append(input, robot)
	}
	// compute answers
	fmt.Println("Expected answers:")
	// different rows and columns from sample case
	fmt.Println("Part a: ", 21)
	fmt.Println("Part b: ", 0)
	fmt.Println("Your answers:")
	fmt.Println("Part a: ", solve1(copySlice(input)))
	fmt.Println("Part b: ", solve2(copySlice(input)))
}
