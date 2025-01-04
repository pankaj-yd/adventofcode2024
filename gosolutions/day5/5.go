package day5

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"slices"
	"strconv"
	"strings"
)

// key, vals
// all vals which must come before key
var order map[int]map[int]struct{}

func isValid(update []int) bool {
	for i := 0; i < len(update); i++ {
		for j:= i+1; j < len(update); j++{
			if valsBefore, ok := order[update[i]]; ok{
				if _, ok := valsBefore[update[j]]; ok{
					return false
				}
			}
		}
	}
	return true
}

func solve1(updates [][]int) int{
	res := 0
	for _, update := range updates {
		if isValid(update){
			res += update[len(update)/2]
		}
	}
	return res
}

func solve2(updates [][]int) int{
	res := 0
	for _, update := range updates {
		if isValid(update){
			continue
		}
		slices.SortFunc(update, func(a, b int) int {
			// if a < b return -1
			// if a > b return  1
			if valsBeforeA, ok := order[a]; ok{
				if _, ok := valsBeforeA[b]; ok{
					// a > b
					return 1
				}
			}
			// a < b
			return -1
		})
		res += update[len(update)/2]
	}
	return res
}


func Day5() {
	// Get the path of the current Go source file
	_, currentFile, _, _ := runtime.Caller(0)

	// Get the directory of the Go source file
	currentDir := filepath.Dir(currentFile)
    parentDir := filepath.Dir(filepath.Dir(currentDir))
    file, err := os.Open(parentDir + "/testcases/5.txt")

    if err != nil {
        log.Fatal("Unable to open testcase file")
        return
    }

    defer file.Close()

	// initialise map
	order = make(map[int]map[int]struct{})
    scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		if(scanner.Text() == ""){
			break
		}
		fields := strings.Split(scanner.Text(), "|")
		x, _ := strconv.Atoi(fields[0])
		y, _ := strconv.Atoi(fields[1])
		if _, ok := order[y]; !ok{
			order[y] = make(map[int]struct{})
		}
		order[y][x] = struct{}{}
	}

	// get updates
	updates := [][]int{}
	for scanner.Scan(){
		fields := strings.Split(scanner.Text(), ",")
		update := []int{}
		for _, field := range fields{
			x, _ := strconv.Atoi(field)
			update = append(update, x)
		}
		updates = append(updates, update)
	}

	fmt.Println("Expected answers:")
    fmt.Println("Part a: ", 143)
    fmt.Println("Part b: ", 123)
    fmt.Println("Your answers:")
    fmt.Println("Part a: ", solve1(updates))
    fmt.Println("Part b: ", solve2(updates))
}