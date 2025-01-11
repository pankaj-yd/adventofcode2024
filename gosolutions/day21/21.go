package day21

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
	"strings"
	"time"
)

const (
	partADirKeyboards = 3
	partBDirKeyboards = 25
	invalidChar       = byte('.')
	pressButtonChar   = byte('A')
	startChar         = byte('A')
)

var (
	north = Dir{dx: -1, dy: 0}
	east  = Dir{dx: 0, dy: 1}
	west  = Dir{dx: 0, dy: -1}
	south = Dir{dx: 1, dy: 0}

	dirs = []Dir{north, east, south, west}

	getDir = map[byte]Dir{
		'^': north,
		'>': east,
		'v': south,
		'<': west,
	}

	getSymbol = map[Dir]byte{
		north: '^',
		east:  '>',
		south: 'v',
		west:  '<',
	}
)

var (
	numpadPos = Pos{x: 3, y: 2}
	numPad    = []string{"789", "456", "123", ".0A"}

	dirPadPos = Pos{x: 0, y: 2}
	dirPad    = []string{".^A", "<v>"}
)

type Dir struct {
	dx, dy int
}

type Pos struct {
	x, y int
}

type Move struct {
	start, end Pos
}

func (p *Pos) move(dir Dir) Pos {
	return Pos{x: p.x + dir.dx, y: p.y + dir.dy}
}

func (p *Pos) diff(pos Pos) Dir {
	return Dir{dx: p.x - pos.x, dy: p.y - pos.y}
}

type Pad struct {
	pad   []string
	paths map[Move][]string
}

func (p *Pad) valid(pos Pos) bool {
	if pos.x < 0 || pos.x >= len(p.pad) || pos.y < 0 || pos.y >= len(p.pad[0]) || p.pad[pos.x][pos.y] == invalidChar {
		return false
	}
	return true
}

func (p *Pad) getPos(ch byte) Pos {
	for i, keys := range p.pad {
		idx := strings.Index(keys, string(ch))
		if idx != -1 {
			return Pos{i, idx}
		}
	}
	return Pos{-1, -1}
}

func (p *Pad) printPaths(paths [][]Pos) {
	for _, path := range paths {
		for _, pos := range path {
			fmt.Printf("%c->", p.pad[pos.x][pos.y])
		}
		fmt.Println()
	}
}

func (p *Pad) addPaths(move Move, paths [][]Pos) {
	_, ok := p.paths[move]
	if !ok {
		p.paths[move] = make([]string, 0)
	}

	for _, path := range paths {
		symbols := make([]byte, 0)
		for i := 0; i < len(path)-1; i++ {
			dir := path[i+1].diff(path[i])
			symbols = append(symbols, getSymbol[dir])
		}
		symbols = append(symbols, pressButtonChar)
		p.paths[move] = append(p.paths[move], string(symbols))
	}

}

func reverse[T any](slice []T) {
	left, right := 0, len(slice)-1
	for left < right {
		slice[left], slice[right] = slice[right], slice[left]
		left++
		right--
	}
}

func (pad *Pad) getPaths(move Move) []string {
	_, ok := pad.paths[move]
	if ok {
		return pad.paths[move]
	}

	queue := []Pos{move.start}
	distance := make(map[Pos]int)
	parents := make(map[Pos]map[Pos]struct{})

	distance[move.start] = 0
	reachedEnd := false
	for len(queue) != 0 || !reachedEnd {
		newQueue := make([]Pos, 0)
		for _, pos := range queue {
			for _, dir := range dirs {
				nextPos := pos.move(dir)

				if !pad.valid(nextPos) {
					continue
				}

				nextDist := distance[pos] + 1
				currDist, ok := distance[nextPos]

				// update new shortest distance
				// update parents
				if !ok || currDist > nextDist {
					distance[nextPos] = nextDist
					newQueue = append(newQueue, nextPos)
					// reset parents if new distance is smaller
					parents[nextPos] = make(map[Pos]struct{})
				}

				// add parent
				if !ok || currDist >= nextDist {
					parents[nextPos][pos] = struct{}{}
				}

				if nextPos == move.end {
					reachedEnd = true
				}

			}
		}

		queue = newQueue
	}

	// fmt.Println(parents)
	allPathsInPos := make([][]Pos, 0)
	currPath := make([]Pos, 0)

	var dfs func(pos Pos)
	dfs = func(currPos Pos) {
		// add currposition to the path
		currPath = append(currPath, currPos)

		// if we reached end of path then add it to all paths
		_, ok := parents[currPos]
		if !ok {
			copyPath := make([]Pos, 0)
			copyPath = append(copyPath, currPath...)
			reverse(copyPath)
			allPathsInPos = append(allPathsInPos, copyPath)
		}

		for parent := range parents[currPos] {
			dfs(parent)
		}
		currPath = currPath[:len(currPath)-1]
	}

	dfs(move.end)
	pad.addPaths(move, allPathsInPos)

	return pad.paths[move]
}

func translateCodeToDir(code string, pad Pad) []string {
	ans := []string{}

	startPos := pad.getPos(startChar)

	for i := range code {
		move := Move{start: startPos, end: pad.getPos(code[i])}
		allPaths := pad.getPaths(move)
		newAns := make([]string, 0)

		if len(ans) == 0 {
			newAns = allPaths
		} else {
			for _, currPath := range ans {
				for _, newPath := range allPaths {
					newAns = append(newAns, currPath+newPath)
				}
			}
		}

		ans = newAns
		startPos = move.end
	}

	return ans
}

var (
	Numpad, Dirpad Pad
)

func printMinLens(numToDirs []string) {
	minLen := math.MaxInt32
	for _, str := range numToDirs {
		minLen = min(minLen, len(str))
	}
	fmt.Println(minLen)
}

func getMinLenTranslations(numToDirs []string) []string {
	minLen := math.MaxInt64
	ansTranslations := make([]string, 0)
	for _, currCode := range numToDirs {
		res := translateCodeToDir(currCode, Dirpad)
		newMinLen := math.MaxInt64
		for i := range res {
			if len(res[i]) < newMinLen {
				newMinLen = len(res[i])
			}
		}

		if minLen > newMinLen {
			minLen = newMinLen
			ansTranslations = make([]string, 0)
		}

		for i := range res {
			if len(res[i]) == minLen {
				ansTranslations = append(ansTranslations, res[i])
			}
		}

	}
	return ansTranslations
}

func getIntPart(code string) int {
	re := regexp.MustCompile(`\d+`)

	// Find the first match in the input string
	firstMatch := re.FindString(code)

	ans, _ := strconv.Atoi(firstMatch)
	return ans
}

func solve1(codes []string, dirBoards int) int {
	Numpad = Pad{pad: numPad, paths: make(map[Move][]string)}
	Dirpad = Pad{pad: dirPad, paths: make(map[Move][]string)}

	ans := 0
	for _, code := range codes {
		numToDirs := translateCodeToDir(code, Numpad)
		// one translated above
		// translate remaining
		for i := 0; i < dirBoards-1; i++ {
			numToDirs = getMinLenTranslations(numToDirs)
			// fmt.Println("i: ", i, " translations: ", len(numToDirs), "each len: ", len(numToDirs[0]))
		}

		ans += len(numToDirs[0]) * getIntPart(code)
	}

	return ans
}

func Day21() {
	// Get the path of the current Go source file
	_, currentFile, _, _ := runtime.Caller(0)

	// Get the directory of the Go source file
	currentDir := filepath.Dir(currentFile)
	parentDir := filepath.Dir(currentDir)
	testCasePath := filepath.Join(parentDir, "..", "testcases", "21.txt")
	file, err := os.Open(testCasePath)

	if err != nil {
		log.Fatal("Unable to open testcase file")
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	codes := make([]string, 0)

	for scanner.Scan() {
		codes = append(codes, scanner.Text())
	}

	// compute answers
	fmt.Println("Expected answers:")
	// different rows and columns from sample case
	fmt.Println("Part a: ", 126384)
	fmt.Println("Part b: ", "")

	fmt.Println("Your answers:")
	startTime := time.Now()
	fmt.Println("Part a: ", solve1(codes, partADirKeyboards))
	elapsed := time.Since(startTime)
	fmt.Println("Time taken:", elapsed.Milliseconds(), "ms")
	fmt.Println()
	startTime = time.Now()
	// ansB := solve1(codes, partBDirKeyboards)
	fmt.Println("Part b: ", "Need to Do")
	elapsed = time.Since(startTime)
	fmt.Println("Time taken:", elapsed.Milliseconds(), "ms")
}
