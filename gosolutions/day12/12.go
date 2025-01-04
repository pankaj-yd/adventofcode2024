package day12

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"slices"
)

var (
	rows, cols int
	north = Dir{-1, 0}
	south = Dir{1, 0}
	east = Dir{0, 1}
	west = Dir{0, -1}
	
	dirs = []Dir{north, east, south, west}
)

type Pos struct {
	x, y int
}

type Dir struct {
	dx, dy int
}

func (pos *Pos) inBounds() bool {
	return pos.x >= 0 && pos.x < rows && pos.y >= 0 && pos.y < cols;
}

func (pos *Pos) move(dir Dir) Pos {
	return Pos{x: pos.x + dir.dx, y: pos.y + dir.dy}
}

var plots = make(map[Pos]map[Pos]struct{}, 0)

func dfs(input []string, startPos, currPos Pos, visited map[Pos]struct{}) {
	if !currPos.inBounds() || input[startPos.x][startPos.y] != input[currPos.x][currPos.y] {
		return
	}

	if _, ok := plots[startPos]; !ok{
		plots[startPos] = make(map[Pos]struct{}, 0)
	}
	plots[startPos][currPos] = struct{}{}
	visited[currPos] = struct{}{}

	for _, dir := range dirs {
		newPos := currPos.move(dir)
		if _, ok := visited[newPos]; !ok {
			dfs(input, startPos, currPos.move(dir), visited)
		}
	}
}

func getArea(samePlots map[Pos]struct{}) int {
	return len(samePlots)
}

func getPerimeter(samePlots map[Pos]struct{}) int {
	perimeter := 0
	for pos := range samePlots {
		for _, dir := range dirs {
			if _, ok := samePlots[pos.move(dir)]; !ok {
				perimeter++;
			}
		}
	}
	return perimeter
}

// part a
func solve1(input []string) int {
	visited := make(map[Pos]struct{})

	// make plots
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			startPos := Pos{i, j}
			// if pos is not visited
			if _, ok := visited[startPos]; !ok {
				dfs(input, startPos, startPos, visited)
			} 
		}
	}
	
	ans := 0
	for _, samePlots := range plots {
		ans += getArea(samePlots) * getPerimeter(samePlots)
	}

	return ans
}

// part b
// generate all points which lie on the perimeter of polygon
// whether those lie on interior or exterior perimeter
type LineSeg struct {
	first, second Pos
}

func (lineseg *LineSeg) move(dir Dir) LineSeg {
	return LineSeg{
		first: Pos{x: lineseg.first.x + dir.dx, y: lineseg.first.y + dir.dy}, 
		second: Pos{x: lineseg.second.x + dir.dx, y: lineseg.second.y + dir.dy}}
}

func generateFigure(samePlots map[Pos]struct{}) map[LineSeg]bool {
	points := make(map[LineSeg]bool)
	allPos := make([]Pos, 0)
	for pos := range samePlots {
		allPos = append(allPos, pos)
	}

	slices.SortFunc(allPos, func(a, b Pos) int {
		if a.x < b.x || a.x == b.x && a.y < b.y {
			return -1
		}
		return 1
	})

	for _, pos := range allPos {
		for _, dir := range dirs {
			if _, ok := samePlots[pos.move(dir)]; !ok {
				// scale the co-ordinates
				newX := pos.x * 2
				newY := pos.y * 2
				
				// in our graph, right is +y, down is +x
				// add points in clockwise order
				topLeft := Pos{x: newX - 1, y: newY - 1}
				topRight := Pos{x: newX - 1, y: newY + 1}
				bottomRight := Pos{x: newX + 1, y: newY + 1}
				bottomLeft := Pos{x: newX + 1, y: newY - 1}
				// fmt.Println("Position: ", pos, " Direction: ", dir)
				// fmt.Println(topLeft, topRight, bottomRight, bottomLeft)
				switch dir {
				case north:
					points[LineSeg{first: topLeft, second: topRight}] = false
					// fmt.Println(topLeft, "-->", topRight)
				case south:
					points[LineSeg{first: bottomRight, second: bottomLeft}] = false
				case east:
					points[LineSeg{first: topRight, second: bottomRight}] = false
				case west:
					points[LineSeg{first: bottomLeft, second: topLeft}] = false
				}
			}
		}
	}
	return points
}

func isOrthogonal(firstDir, secondDir Dir) bool {
	return firstDir.dx * secondDir.dx + firstDir.dy * secondDir.dy == 0
}

func getDir(pos1, pos2 Pos) Dir {
	return Dir{dx: pos2.x - pos1.x, dy: pos2.y - pos1.y}
}

func orthoLineSeg(pos Pos, orthoDir Dir, lineSegs map[LineSeg]bool) bool {
	orthoPos := pos.move(orthoDir)
	orthoSeg := LineSeg{first: orthoPos, second: pos}
	if _, ok := lineSegs[orthoSeg]; ok {
		return true
	}
	orthoSeg = LineSeg{first: pos, second: orthoPos}
	if _, ok := lineSegs[orthoSeg]; ok {
		return true
	}

	return false
}

func orthoLineSegPresent(lineSeg LineSeg, orthoDir Dir, lineSegs map[LineSeg]bool) bool {
	return orthoLineSeg(lineSeg.first, orthoDir, lineSegs) || orthoLineSeg(lineSeg.second, orthoDir, lineSegs)
}

func markLineSegVisited(lineSeg LineSeg, dir Dir , lineSegs map[LineSeg]bool){
	stopVisit := false
	for {
		// if line seg does not exist, stop visiting
		if _, ok := lineSegs[lineSeg]; !ok {
			break
		}

		lineSegs[lineSeg] = true

		// if there are orthogonal linesegs present, stop visiting
		for _, d := range dirs {
			if isOrthogonal(dir, d) && orthoLineSegPresent(lineSeg, d, lineSegs){
				stopVisit = true
				break
			}
		}

		if stopVisit {
			break
		}

		lineSeg = lineSeg.move(dir)
	}

}

func markSameSideLineSegVisited(lineseg LineSeg, lineSegs map[LineSeg]bool) {
	firstDir := getDir(lineseg.second, lineseg.first)
	markLineSegVisited(lineseg, firstDir, lineSegs)

	secondDir := getDir(lineseg.first, lineseg.second)
	markLineSegVisited(lineseg, secondDir, lineSegs)
}

func getSides(samePlots map[Pos]struct{}) int {
	sides := 0
	lineSegs := generateFigure(samePlots)
	// points := make(map[Pos]struct{})
	// fmt.Print("[")
	// for key, _ := range lineSegs {
	// 	points[key.first] = struct{}{}
	// 	points[key.second] = struct{}{}
	// 	fmt.Print("[(", key.first.x, ",", key.first.y, "), ")
	// 	fmt.Print("(", key.second.x, ",", key.second.y, ")], ")
	// }
	// fmt.Println("]")

	// get sides
	for lineseg, val := range lineSegs {
		if val {
			continue
		}
		
		// mark curr lineseg as done
		lineSegs[lineseg] = true

		// all other lineSegs which are part of the side to which this
		// line seg belongs also as visited
		markSameSideLineSegVisited(lineseg, lineSegs)

		sides++;
	}
	return sides
}

func solve2() int {	
	var ans int
	allKeys := make([]Pos, 0)
	for keys := range plots {
		allKeys = append(allKeys, keys)
	}

	slices.SortFunc(allKeys, func(a, b Pos) int {
		if a.x < b.x || a.x == b.x && a.y < b.y {
			return -1
		}
		return 1
	})

	for _, key := range allKeys {
		// fmt.Println("Calculating for: ", key)
		area := getArea(plots[key])
		sides := getSides(plots[key])
		// fmt.Println(sides)
		ans +=  area * sides
	}

	return ans
}

func Day12() {
	// Get the path of the current Go source file
	_, currentFile, _, _ := runtime.Caller(0)

	// Get the directory of the Go source file
	currentDir := filepath.Dir(currentFile)
	parentDir := filepath.Dir(filepath.Dir(currentDir))
	file, err := os.Open(parentDir + "/testcases/12.txt")

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
	fmt.Println("Part a: ", 1930)
	fmt.Println("Part b: ", 1206)
	fmt.Println("Your answers:")
	fmt.Println("Part a: ", solve1(input))
	fmt.Println("Part b: ", solve2())
}
