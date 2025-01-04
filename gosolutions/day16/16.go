package day16

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
	"runtime"
)

const (
	turnCost   = 1000
	normalCost = 1
	start      = byte('S')
	end        = byte('E')
	wall       = byte('#')
)

var (
	startPos = Pos{}
	endPos   = Pos{}
	minDistance = 0

	north = Dir{dx: -1, dy: 0}
	east = Dir{dx: 0, dy: 1}
	west = Dir{dx: 0, dy: -1}
	south = Dir{dx: 1, dy: 0}
	dirs = []Dir{north, east, south, west}

	// stores distnace from startPos, startDir to particular pos, dir
	distance = make(map[State]int, 0)

	// stores minDistance parents in a set for a pos,dir
	parents = make(map[State]map[State]struct{})
)

type Dir struct {
	dx, dy int
}

type Pos struct {
	x, y int
}

type State struct {
	pos Pos
	dir Dir
}

func (dir *Dir) rotateCW() Dir {
	// -1, 0 north
	// 0, 1 east
	return Dir{dx: dir.dy, dy: -dir.dx}
}

func (dir *Dir) rotateACW() Dir {
	// 0, 1 east
	// -1, 0 north
	return Dir{dx: -dir.dy, dy: dir.dx}
}

func (pos *Pos) inBound(grid [][]byte) bool {
	return pos.x >= 0 && pos.x < len(grid) && pos.y >= 0 && pos.y < len(grid[0])
}

func (pos *Pos) move(dir Dir) Pos {
	return Pos{x: pos.x + dir.dx, y: pos.y + dir.dy}
}

func printGrid(grid [][]byte) {
	for _, row := range grid {
		fmt.Println(string(row))
	}
	fmt.Println()
}

func (child State) resetParents() {
	parents[child] = make(map[State]struct{})
}

func (child State) addParent(parent State) {
	_, ok := parents[child]
	if !ok {
		parents[child] = make(map[State]struct{})
	}

	parents[child][parent] = struct{}{}
}

func solve1(grid [][]byte) int {
	for i, row := range grid {
		for j, ch := range row {
			switch ch {
			case start:
				startPos.x = i
				startPos.y = j
			case end:
				endPos.x = i
				endPos.y = j
			}
		}
	}

	// for every i,j there are 4 ways to reach
	queue := make([]State, 0)
	initialState := State{pos: startPos, dir: east}
	queue = append(queue, initialState)
	distance[initialState] = 0

	// Floyd Warshall Algo
	shouldContinue := true
	ans := math.MaxInt64
	for shouldContinue {
		if len(queue) == 0 {
			break
		}

		// relax all outgoing edges from every element in queue
		newQueue := make([]State, 0)
		for _, state := range queue {
			for _, dir := range []Dir{state.dir, state.dir.rotateCW(), state.dir.rotateACW()} {
				// initialise new state as current state, with same cost
				newState := state
				newScore := distance[state]

				if dir == state.dir { // same dir, move one step and add score
					newState.pos = state.pos.move(dir)
					newScore += normalCost
				} else { // if different direction, the rotate at same position, update dir, score
					newState.dir = dir
					newScore += turnCost
				}

				// check if new state is invalid
				if !newState.pos.inBound(grid) || grid[newState.pos.x][newState.pos.y] == wall {
					continue
				}

				
				score, ok := distance[newState]
				if !ok || score > newScore {
					distance[newState] = newScore
					newQueue = append(newQueue, newState)

					// reset parents and add the state as new parent
					newState.resetParents()
					newState.addParent(state)
				}

				// add parent if the score is same
				if ok && score == newScore {
					newState.addParent(state)
				}

				if grid[newState.pos.x][newState.pos.y] == end {
					ans = min(ans, newScore)
				}
			}
		}
		queue = newQueue
	}
	minDistance = ans
	return ans
}

// part b
func solve2(grid [][]byte) int {
	// add all end states which have minDistance
	queue := make([]State, 0)

	for _, dir := range dirs {
		endState := State{pos: endPos, dir: dir}
		endDistance, ok := distance[endState]
		if ok && endDistance == minDistance {
			queue = append(queue, endState)
		}
	}

	goodSpots := make(map[Pos]struct{})

	for len(queue) != 0 {

		newQueue := make([]State, 0)
		for _, state := range queue {
			// add the good spots to the queue
			goodSpots[state.pos] = struct{}{}

			for parent := range parents[state] {
				newQueue = append(newQueue, parent)
			}
		}
		queue = newQueue
	}
	// fmt.Println(goodSpots)
	return len(goodSpots)
}

func Day16() {
	// Get the path of the current Go source file
	_, currentFile, _, _ := runtime.Caller(0)

	// Get the directory of the Go source file
	currentDir := filepath.Dir(currentFile)
	parentDir := filepath.Dir(currentDir)
	testCasePath := filepath.Join(parentDir, "..", "testcases", "16.txt")
	file, err := os.Open(testCasePath)

	if err != nil {
		log.Fatal("Unable to open testcase file")
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	grid := make([][]byte, 0)

	for scanner.Scan() {
		grid = append(grid, []byte(scanner.Text()))

	}

	// compute answers
	fmt.Println("Expected answers:")
	// different rows and columns from sample case
	fmt.Println("Part a: ", 7036)
	fmt.Println("Part b: ", 45)
	fmt.Println("Your answers:")
	fmt.Println("Part a: ", solve1(grid))
	fmt.Println("Part b: ", solve2(grid))
}


/*
// Faulty DFS, need to debug

func dfs(grid [][]byte, pos Pos, currDir Dir, prevDir Dir, cache map[State]int) int {

	if !pos.inBound(grid) || grid[pos.x][pos.y] == '#' {
		// fmt.Print(pos, currDir, prevDir, "-->")
		// fmt.Println("go back")
		return math.MaxInt64
	}
	//   fmt.Print("[", pos.x,  ",", pos.y, "],")
	if grid[pos.x][pos.y] == end {
		return 0
	}

	currState := State{pos: pos, dir: currDir}
	if score, ok := cache[currState]; ok {
		// fmt.Println("using cached result")
		return score
	}

	score := math.MaxInt64
	ch := grid[pos.x][pos.y]
	grid[pos.x][pos.y] = '#'

	nextPos := pos.move(currDir)

	// calc new score
	// fmt.Print(pos, currDir, prevDir, "-->")
	// fmt.Println(nextPos, currDir, prevDir)
	newScore := dfs(grid, nextPos, currDir, prevDir, cache)
	if newScore != math.MaxInt64 {
		newScore += normalCost
	}
	score = min(score, newScore)
	grid[pos.x][pos.y] = ch

	// rotate CW
	if !oppositeDirs(currDir.rotateCW(), prevDir) || currDir.rotateCW() != prevDir {
		// fmt.Print(pos, currDir, prevDir, "-->")
		// fmt.Println(pos, currDir.rotateCW(), currDir)
		newScore = dfs(grid, pos, currDir.rotateCW(), currDir, cache)
		if newScore != math.MaxInt64 {
			newScore += turnCost
		}
		score = min(score, newScore)
	}
	// rotate ACW
	if !oppositeDirs(currDir.rotateACW(), prevDir) || currDir.rotateACW() != prevDir {
		// fmt.Print(pos, currDir, prevDir, "-->")
		// fmt.Println(pos, currDir.rotateACW(), currDir)
		newScore = dfs(grid, pos, currDir.rotateACW(), currDir, cache)
		if newScore != math.MaxInt64 {
			newScore += turnCost
		}
		score = min(score, newScore)
	}

	cache[currState] = score

	return score
}

func moveCost(dir1, dir2 Dir) int {
	if dir1 != dir2 {
		return turnCost
	}

	return normalCost
}

func solve1(grid [][]byte) int {
	pos := Pos{x: -1, y: -1}
	for i, row := range grid {
		for j, ch := range row {
			if ch == start {
				pos.x = i
				pos.y = j
			}
		}
	}
	cache := make(map[State]int, 0)
	score := dfs(grid, pos, east, west, cache)

	// printGrid(grid)
	return score
}

func oppositeDirs(dir1, dir2 Dir) bool {
	if dir1.dx == -dir2.dx && dir1.dy == -dir2.dy {
		return true
	}
	return false
}
*/
