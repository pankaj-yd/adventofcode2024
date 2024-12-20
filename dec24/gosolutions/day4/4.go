package day4
import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)
var rows, cols int

// Part A
var dirs = [8][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}, {1, 1}, {-1, 1}, {1, -1}, {-1, -1}}

func validIdx(i, j int) bool {
	return i >= 0 && i < rows && j >= 0 && j < cols;
}

func checkXmas(s []string, r, c int) int{
	xmas := 0
	for _, dir := range dirs{
		i := r
		j := c
		found := true
		for _, ch := range []byte("MAS"){
			i += dir[0]
			j += dir[1]
			if !validIdx(i, j) || s[i][j] != ch {
				found = false
				break
			}
		}
		if found {
			xmas++
		}
	}

	return xmas
}

func countXmas(s []string) int{
	res := 0
	for i := 0; i < rows; i++{
		for j := 0; j < cols; j++{
			if s[i][j] != 'X'{
				continue
			}
			res += checkXmas(s, i, j)
		}
	}

	return res
}

// Part B
func checkShape(s []string, r, c int) int{
	xmas := 0
    /*
    M.M
    .A.
    S.S
    */
   if r+2 < rows && c+2 < cols && s[r][c] == 'M' && s[r][c+2] == 'M' && s[r+1][c+1] == 'A' && s[r+2][c] == 'S' && s[r+2][c+2] == 'S'{
    xmas++;
   }

    /*
    M.S
    .A.
    M.S
    */
   if r+2 < rows && c+2 < cols && s[r][c] == 'M' && s[r][c+2] == 'S' && s[r+1][c+1] == 'A' && s[r+2][c] == 'M' && s[r+2][c+2] == 'S'{
    xmas++;
   }

   /*
    S.M
    .A.
    S.M
    */
   if r+2 < rows && c+2 < cols && s[r][c] == 'S' && s[r][c+2] == 'M' && s[r+1][c+1] == 'A' && s[r+2][c] == 'S' && s[r+2][c+2] == 'M'{
    xmas++;
   }

   /*
    S.S
    .A.
    M.M
    */
   if r+2 < rows && c+2 < cols && s[r][c] == 'S' && s[r][c+2] == 'S' && s[r+1][c+1] == 'A' && s[r+2][c] == 'M' && s[r+2][c+2] == 'M'{
    xmas++;
   }

	return xmas
}

func countShape(s []string) int{
	xmas := 0
	for i := 0; i < rows; i++{
		for j := 0; j < cols; j++{
			xmas += checkShape(s, i, j)
		}
	}
	return xmas
}


func Day4() {
	// Get the path of the current Go source file
	_, currentFile, _, _ := runtime.Caller(0)

	// Get the directory of the Go source file
	currentDir := filepath.Dir(currentFile)
    parentDir := filepath.Dir(filepath.Dir(currentDir))
    file, err := os.Open(parentDir + "/testcases/4.txt")

    if err != nil {
        log.Fatal("Unable to open testcase file")
        return
    }

    defer file.Close()

    scanner := bufio.NewScanner(file)
	var input []string
	for scanner.Scan(){
		input = append(input, scanner.Text())
	}

	rows = len(input)
	cols = len(input[0])

	fmt.Println("Expected answers:")
    fmt.Println("Part a: ", 2493)
    fmt.Println("Part b: ", 1890)
    fmt.Println("Your answers:")
    fmt.Println("Part a: ", countXmas(input))
    fmt.Println("Part b: ", countShape(input))
}