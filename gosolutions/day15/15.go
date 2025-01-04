package day15

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

var (
	rowFactor = 100
)

type Dir struct {
	dx, dy int
}

type Pos struct {
	x, y int
}

type Box struct {
	left, right Pos
}

func (pos *Pos) inBound(grid [][]byte) bool {
	return pos.x >= 0 && pos.x < len(grid) && pos.y >= 0 && pos.y < len(grid[0])
}

func (pos *Pos) move(dir Dir) Pos{
	return Pos{x: pos.x + dir.dx, y: pos.y + dir.dy}
}

func (box *Box) move(dir Dir) Box {
	return Box{left: box.left.move(dir), right: box.right.move(dir)}
}

var (
	north = Dir{dx: -1, dy: 0}
	east = Dir{dx: 0, dy: 1}
	west = Dir{dx: 0, dy: -1}
	south = Dir{dx: 1, dy: 0}
	dir = map[byte]Dir{
		'>': east,
		'<': west,
		'^': north,
		'v': south,
	}
)

func moveInGridFast(grid [][]byte, pos *Pos, movDir Dir) {
	newPos := *pos

	for newPos.inBound(grid) && grid[newPos.x][newPos.y] != '.' && grid[newPos.x][newPos.y] != '#' {
		newPos = newPos.move(movDir)
	}

	if newPos.inBound(grid) && grid[newPos.x][newPos.y] == '.'{
		nextPos := pos.move(movDir)
		grid[newPos.x][newPos.y] = grid[nextPos.x][nextPos.y]
		grid[nextPos.x][nextPos.y] = grid[pos.x][pos.y]
		grid[pos.x][pos.y] = '.'

		pos.x = nextPos.x
		pos.y = nextPos.y
	}
}

func printGrid(grid [][]byte) {
	for _, row := range grid {
		fmt.Println(string(row))
	}
	fmt.Println()
}

func calcScore(grid [][]byte, box byte) int {
	score := 0
	for i, row := range grid {
		for j, ch := range row {
			if ch == box {
				score += rowFactor * i + j
			}
		}
	}
	return score
}

func solve1(inputGrid [][]byte, moves string) int {
	grid := make([][]byte, len(inputGrid))
    // Loop through each row (inner slice) and create a copy
    for i := range inputGrid {
        grid[i] = append([]byte{}, inputGrid[i]...) // Copy each inner slice
    }

	// find robot position
	pos := Pos{}
	for i := range grid {
		for j := range grid[i]{
			if grid[i][j] == '@' {
				pos.x = i
				pos.y = j
			}
		}
	}

	// move robot
	// fmt.Println(pos)
	for _, ch := range moves {
		moveInGridFast(grid, &pos, dir[byte(ch)])
		// fmt.Printf("Move %c:\n", ch)
		// printGrid(grid)
	}

	printGrid(grid)
	// calc score
	
	return calcScore(grid, 'O')
}


// part b
func getBoxAtPos(grid [][]byte, pos Pos) (Box, bool) {
	left := Pos{x: -1, y: -1}
	right := Pos{x: -1, y: -1}
	ok := false
	switch grid[pos.x][pos.y]{
	case '[':
		left = pos
		right = pos.move(east)
		ok = true
	case ']':
		left = pos.move(west)
		right = pos
		ok = true
	case '.':
		left = Pos{}
		right = Pos{}
	}
	return Box{left: left, right: right}, ok
}

func moveBox(grid [][]byte, box Box, movDir Dir){
	newBoxPos := box.move(movDir)
	grid[newBoxPos.left.x][newBoxPos.left.y] = '['
	grid[newBoxPos.right.x][newBoxPos.right.y] = ']'

	grid[box.left.x][box.left.y] = '.'
	grid[box.right.x][box.right.y] = '.'	
}

func moveInGridHori(grid [][]byte, pos *Pos, movDir Dir) {
	newPos := *pos

	for newPos.inBound(grid) && grid[newPos.x][newPos.y] != '.' && grid[newPos.x][newPos.y] != '#' {
		newPos = newPos.move(movDir)
		// fmt.Println(newPos)
	}

	if newPos.inBound(grid) && grid[newPos.x][newPos.y] == '.'{
		revDir := Dir{dx:-movDir.dx, dy:-movDir.dy}
		toMovePos := newPos.move(revDir)

		for newPos != *pos{
			grid[newPos.x][newPos.y] = grid[toMovePos.x][toMovePos.y]
			newPos = toMovePos
			toMovePos = toMovePos.move(revDir)
		}
		
		grid[pos.x][pos.y] = '.'

		robPos := pos.move(movDir)
		pos.x = robPos.x
		pos.y = robPos.y
	}
}

func moveInGridVert(grid [][]byte, pos *Pos, movDir Dir) {
	// fmt.Println("@:", pos)
	newPos := pos.move(movDir)
	switch grid[newPos.x][newPos.y] {
	case '.':
		grid[newPos.x][newPos.y] = grid[pos.x][pos.y]
		grid[pos.x][pos.y] = '.'

		pos.x = newPos.x
		pos.y = newPos.y
		return
	case '#':
		return
	}

	// there is a box above/below the robot
	// stores boxes at each level from robot
	levels := make(map[int]map[Box]struct{})
	box, _ := getBoxAtPos(grid, newPos)
	lvl := 0
	levels[lvl] = make(map[Box]struct{})
	levels[lvl][box] = struct{}{}
	
	canMove := true
	for {
		lvl++
		levels[lvl] = make(map[Box]struct{})
		for currBox := range levels[lvl-1] {
			// add boxes in dir to currBox
			boxAboveLeftPos, ok := getBoxAtPos(grid, currBox.left.move(movDir))
			if ok {
				levels[lvl][boxAboveLeftPos] = struct{}{}
			}
			// enocountered #
			if boxAboveLeftPos.left.x == -1 {
				// fmt.Println("Can not move")
				canMove = false
				break
			}
			
			boxAboveRightPos, ok := getBoxAtPos(grid, currBox.right.move(movDir))
			if ok {
				levels[lvl][boxAboveRightPos] = struct{}{}
			}
			// enocountered #
			if boxAboveRightPos.left.x == -1 {
				// fmt.Println("Can not move")
				canMove = false
				break
			}
		}

		// if no boxes are added, stop looping
		if len(levels[lvl]) == 0 {
			break
		}
	}

	if canMove {
		// move boxes at all lvls
		for i := lvl-1; i >= 0; i-- {
			for box := range levels[i]{
				moveBox(grid, box, movDir)
			}
		}

		// mov rob
		grid[newPos.x][newPos.y] = grid[pos.x][pos.y]
		grid[pos.x][pos.y] = '.'

		pos.x = newPos.x
		pos.y = newPos.y
	}

}

func solve2(grid [][]byte, moves string) int {
	// make new grid
	doubleGrid := make([][]byte, 0)
	pos := Pos{x: 1, y: 2}
	for i, row := range grid {
		doubleRow := make([]byte, 0)
		for j, ch := range row {
			switch ch {
			case '.':
				fallthrough
			case '#':
				doubleRow = append(doubleRow, ch, ch)
			case 'O':
				doubleRow = append(doubleRow, "[]"...)
			case '@':
				doubleRow = append(doubleRow, "@."...)
				pos.x *= i
				pos.y *= j
			}
		}
		doubleGrid = append(doubleGrid, doubleRow)
	}

	// move robot
	// fmt.Println(pos)
	// fmt.Println("Initial state:")
	// printGrid(doubleGrid)
	for _, ch := range moves {
		movDir := dir[byte(ch)]
		if movDir == east || movDir == west {
			moveInGridHori(doubleGrid, &pos, movDir)
		} else {
			moveInGridVert(doubleGrid, &pos, movDir)
		}
		// fmt.Printf("Move %c:\n", ch)
		// printGrid(doubleGrid)
	}
	printGrid(doubleGrid)

	// calc score
	return calcScore(doubleGrid, '[')
}


func Day15() {
	// Get the path of the current Go source file
	_, currentFile, _, _ := runtime.Caller(0)

	// Get the directory of the Go source file
	currentDir := filepath.Dir(currentFile)
	parentDir := filepath.Dir(currentDir)
	testCasePath := filepath.Join(parentDir, "..", "testcases", "15.txt")
	file, err := os.Open(testCasePath)

	if err != nil {
		log.Fatal("Unable to open testcase file")
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	grid := make([][]byte, 0)
	moves := ""

	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		grid = append(grid, []byte(scanner.Text()))

	}

	for scanner.Scan() {
		moves += scanner.Text()
	}

	// compute answers
	fmt.Println("Expected answers:")
	// different rows and columns from sample case
	fmt.Println("Part a: ", 10092)
	fmt.Println("Part b: ", 9021)
	fmt.Println("Your answers:")
	fmt.Println("Part a: ", solve1(grid, moves))
	fmt.Println("Part b: ", solve2(grid, moves))
}