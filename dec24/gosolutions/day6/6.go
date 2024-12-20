package day6

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

var rows, cols int
var startPos Pos
var startDir Dir

type Pos struct {
	x, y int
}

func (pos *Pos) valid() bool {
	return pos.x >= 0 && pos.x < rows && pos.y >= 0 && pos.y < cols
}

func (pos *Pos) move(dir Dir) Pos {
	return Pos{x : pos.x + dir.dx, y: pos.y + dir.dy}
}

type Dir struct {
	dx, dy int
}

func (dir *Dir) turn() Dir {
	return Dir{dx: dir.dy, dy: -dir.dx}
}

func uniqueVisits(input [][]byte) int {
	visited := make(map[Pos]struct{})
	dir := startDir
	pos := startPos
	for {
		if !pos.valid(){
			break;
		}
		
		visited[pos] = struct{}{}
		nextPos := pos.move(dir)
		if nextPos.valid() && input[nextPos.x][nextPos.y] == byte('#') {
			dir = dir.turn()
		} else {
			pos = nextPos
		}

	}
	return len(visited)
}

type PosVector struct {
	pos Pos
	dir Dir
}

func cycle(input [][]byte) bool {
	visited := make(map[PosVector]struct{})
	dir := startDir
	pos := startPos
	for {
		if !pos.valid(){
			break;
		}
		currVector := PosVector{pos: pos, dir: dir} 
		if _, exist := visited[currVector]; exist {
			return true
		}
		visited[currVector] = struct{}{}

		nextPos := pos.move(dir)
		if nextPos.valid() && input[nextPos.x][nextPos.y] == '#' {
			dir = dir.turn()
		} else {
			pos = nextPos
		}

	}
	return false
}

func obstacles(input [][]byte) int {
	// store the path
	visited := make(map[Pos]struct{})
	dir := startDir
	pos := startPos
	for {
		if !pos.valid(){
			break;
		}
		visited[pos] = struct{}{}
		nextPos := pos.move(dir)
		if nextPos.valid() && input[nextPos.x][nextPos.y] == '#' {
			dir = dir.turn()
		} else {
			pos = nextPos
		}
	}

	// place obstacle at every visited location
	// except spawn
	res := 0
	for currPos := range visited{
		if currPos == startPos {
			continue
		}

		input[currPos.x][currPos.y] = '#'
		if cycle(input){
			res++
		}
		input[currPos.x][currPos.y] = '.'
	}

	return res
}

func Day6() {
	// Get the path of the current Go source file
	_, currentFile, _, _ := runtime.Caller(0)

	// Get the directory of the Go source file
	currentDir := filepath.Dir(currentFile)
    parentDir := filepath.Dir(filepath.Dir(currentDir))
    file, err := os.Open(parentDir + "/testcases/6.txt")

    if err != nil {
        log.Fatal("Unable to open testcase file")
        return
    }

    defer file.Close()

	// initialise map
	input := make([][]byte, 0)

    scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		if(scanner.Text() == ""){
			break
		}
		row := scanner.Text()
		input = append(input, []byte(row))

		idx := bytes.Index(scanner.Bytes(), []byte("^"))
		if idx != -1 {
			startPos.x = len(input) - 1
			startPos.y = idx
			startDir.dx = -1
			startDir.dy = 0
		}
	}
	rows = len(input)
	cols = len(input[0])


	fmt.Println("Expected answers:")
    fmt.Println("Part a: ", 5129)
    fmt.Println("Part b: ", 1888)
    fmt.Println("Your answers:")
    fmt.Println("Part a: ", uniqueVisits(input))
    fmt.Println("Part b: ", obstacles(input))
}