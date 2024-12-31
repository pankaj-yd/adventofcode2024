package day2

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

func increasing(report []int) bool {
    for i := 1; i < len(report); i++{
        diff := report[i] - report[i-1]
        if diff > 3 || diff < 1 {
            return false
        }
    }
    return true
}

func decreasing(report []int) bool {
    for i := 1; i < len(report); i++{
        diff := report[i-1] - report[i]
        if diff > 3 || diff < 1 {
            return false
        }
    }
    return true
}

func safeReports(reports [][]int) int{
    numSafeReports := 0
    for _, report := range reports {
        if len(report) <= 1 {
            numSafeReports++
        } else if increasing(report) || decreasing(report) {
            numSafeReports++
        }
    }

    return numSafeReports
}

func safeReportsTolerance(reports [][]int) int{
    numSafeReports := 0
    for _, report := range reports {
        if len(report) <= 1 {
            numSafeReports++
        } else if increasing(report) || decreasing(report) {
            numSafeReports++
        } else {
            // remove ith element and check if report becomes safe
            for i := 0; i < len(report); i++ {
                copyReport := make([]int, len(report))
                copy(copyReport, report)
                subReport := append(copyReport[:i], copyReport[i+1:]...)
                if increasing(subReport) || decreasing(subReport) {
                    numSafeReports++
                    break
                }
            }
        }
    }

    return numSafeReports
}

func Day2() {
    // Get the path of the current Go source file
	_, currentFile, _, _ := runtime.Caller(0)

	// Get the directory of the Go source file
	currentDir := filepath.Dir(currentFile)
    parentDir := filepath.Dir(filepath.Dir(currentDir))
    file, err := os.Open(parentDir + "/testcases/2.txt")

    if err != nil {
        log.Fatal("Unable to open testcase file")
        return
    }

    defer file.Close()

    scanner := bufio.NewScanner(file)
    var reports [][]int
    for scanner.Scan(){
        fields := strings.Fields(scanner.Text())

        report := make([]int, len(fields))
        for i, v := range fields {
            num, err := strconv.Atoi(v)
            if err != nil {
                log.Fatal("Invalid input given")
                return
            }
            report[i] = num
        }

        reports = append(reports, report)
    }
    println("Expected answers:")
    println("Part a: ", 2)
    println("Part b: ", 4)
    println()
    println("Your answers:")
    println("Part a: ", safeReports(reports))
    println("Part b: ", safeReportsTolerance(reports))
}
