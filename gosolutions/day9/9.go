package day9

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

func updateCheckSum(checkSum *int, fileId, fileSize, startIdx int) {
	// fmt.Println("fileId: ", fileId, " fileSize: ", fileSize, " startIdx: ", startIdx)
	n := startIdx + fileSize - 1
	*checkSum += ((n * (n + 1))/2 - (startIdx * (startIdx - 1))/2 ) * fileId
}

// part a
func solve1(inputString string) int {
	input := make([]int, 0)
	for _, ch := range inputString {
		input = append(input, int(ch - '0'))
	}

	checkSum := 0

	// move and add to checkSum from odd places
	n := len(input)
	j := n-1 // point to files from back of input
	if n & 1 == 0{
		// if n is even
		// move j pointer to file
		j--
	}
	i := 0 // pointer to free space
	expandedIdx := 0 // keeps track of what value does i maps to after expansion
	for i < j {
		// i is even, it's a file
		// add to checkSum, update expanded idx, i
		if i & 1 == 0 {
			fileId := i/2
			fileSize := input[i]
			updateCheckSum(&checkSum, fileId, fileSize, expandedIdx)
			expandedIdx += fileSize
			i++
			continue
		}

		// move file at idx j to freeSpace i
		freeSpace := input[i]
		fileSize := input[j]
		fitSize := min(freeSpace, fileSize)
		moveFileId := j/2
		updateCheckSum(&checkSum, moveFileId, fitSize, expandedIdx)
		expandedIdx += fitSize

		// update freesize and fileSize
		input[i] -= fitSize
		input[j] -= fitSize
		
		// if free space is filled, move to next idx
		if input[i] == 0 {
			i++
		}
		// if file is completely moved, move left more
		if input[j] == 0 {
			j -= 2
		}
	}

	// add files after i
	for i < n {
		fileId := i/2
		fileSize := input[i]
		updateCheckSum(&checkSum, fileId, fileSize, expandedIdx)
		expandedIdx += fileSize
		i += 2
	}


	return checkSum
}

// part b
type FreeSpace struct {
	idx int // index in input
	size int // size of free space
	eIdx int // expanded index
}

func getFreeSpaces(input []int) *[]FreeSpace{
	freeSpaces := make([]FreeSpace, 0)
	expIdx := 0
	for i := 1; i < len(input); i += 2 {
		expIdx += input[i-1]
		freeSpaces = append(freeSpaces, FreeSpace{idx: i, size: input[i], eIdx: expIdx})
		expIdx += input[i]
	}
	return &freeSpaces
}

func solve2(inputString string) int {
	input := make([]int, 0)
	for _, ch := range inputString {
		input = append(input, int(ch - '0'))
	}

	checkSum := 0

	freeSpaces := getFreeSpaces(input)
	// move and add to checkSum from odd places
	n := len(input)
	j := n-1 // point to files from back of input
	if n & 1 == 0{
		// if n is even
		// move j pointer to file
		j--
	}
	
	for j >= 0 {
		// move file at idx j to freeSpace i
		fileId := j/2
		fileSize := input[j]
		for i := 0; i < len(*freeSpaces); i++ {
			freeSpace := &(*freeSpaces)[i]
			
			if freeSpace.idx > j {
				break
			}

			if freeSpace.size >= fileSize {
				// move file to freeSpace
				updateCheckSum(&checkSum, fileId, fileSize, freeSpace.eIdx)

				// update the size and eIdx
				freeSpace.eIdx += fileSize
				freeSpace.size -= fileSize

				input[j] = -input[j]
				break
			}
		}
		j -= 2
		
	}
	// add files to check at even indices
	i := 0
	expandedIdx := 0
	for i < n {
		// update expandedIdx to include previous freeSpace
		if i != 0 {
			expandedIdx += input[i-1]
		}

		// moved file
		// just update expandedIdx and i
		if input[i] < 0 {
			expandedIdx += -input[i]
			i += 2;
			continue
		}

		fileId := i/2
		fileSize := input[i]
		updateCheckSum(&checkSum, fileId, fileSize, expandedIdx)
		expandedIdx += fileSize
		i += 2
	}


	return checkSum
}

func Day9() {
	// Get the path of the current Go source file
	_, currentFile, _, _ := runtime.Caller(0)

	// Get the directory of the Go source file
	currentDir := filepath.Dir(currentFile)
    parentDir := filepath.Dir(filepath.Dir(currentDir))
    file, err := os.Open(parentDir + "/testcases/9.txt")

    if err != nil {
        log.Fatal("Unable to open testcase file")
        return
    }

    defer file.Close()

    scanner := bufio.NewScanner(file)
	scanner.Scan()
	inputString := scanner.Text()

	// compute answers
	fmt.Println("Expected answers:")
    fmt.Println("Part a: ", 1928)
    fmt.Println("Part b: ", 2858)
    fmt.Println("Your answers:")
    fmt.Println("Part a: ", solve1(inputString))
    fmt.Println("Part b: ", solve2(inputString))
}