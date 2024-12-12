package gosolutions

import (
	"C"
	"bufio"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strconv"
	"strings"
)

func checkDiff(list1, list2 []int) int {
    sort.Ints(list1)
    sort.Ints(list2)

    ans := 0
    for i := 0; i < len(list1); i++ {
        diff := list1[i] - list2[i]
        if diff < 0 {
            diff = -diff
        }
        ans += diff
    }

    return ans
}

func similarity(list1, list2 []int) int {
    freq := make(map[int]int)

    for _, num := range list2 {
        freq[num]++;
    }

    var ans int
    for _, num := range list1 {
        ans += num * freq[num]
    }

    return ans
}


func Day1() {
    // Get the path of the current Go source file
	_, currentFile, _, _ := runtime.Caller(0)

	// Get the directory of the Go source file
	currentDir := filepath.Dir(currentFile)
    parentDir := filepath.Dir(currentDir)
    file, err := os.Open(parentDir + "/testcases/1.txt")

    if err != nil {
        log.Fatal("Unable to open testcase file")
        return
    }

    defer file.Close()

    scanner := bufio.NewScanner(file)
    var list1, list2 []int
    for scanner.Scan(){
        fields := strings.Fields(scanner.Text())

        intFields := make([]int, len(fields))
        for i, v := range fields {
            num, err := strconv.Atoi(v)
            if err != nil {
                log.Fatal("Invalid input given")
                return
            }
            intFields[i] = num
        }

        list1 = append(list1, intFields[0])
        list2 = append(list2, intFields[1])
    }

    println("Part a: ", checkDiff(list1, list2))
    println("Part b: ", similarity(list1, list2))
}
