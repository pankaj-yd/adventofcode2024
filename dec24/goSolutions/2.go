package gosolutions

import (
	"C"
	"bufio"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)




func Day2() {
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
    return
}
