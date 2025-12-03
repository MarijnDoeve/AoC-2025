package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var lines []string

	for {
		scanner.Scan()
		line := scanner.Text()

		if len(line) == 0 {
			break
		}

		lines = append(lines, line)
	}

	sum2 := 0
	sum12 := 0
	for _, line := range lines {
		sum2 += findLargestOfSize(line, 2)
		sum12 += findLargestOfSize(line, 12)
	}
	fmt.Println("Part 1:", sum2)
	fmt.Println("Part 2:", sum12)
}

func findLargestOfSize(line string, size int) int {
	foundNumber := ""
	startPos := 0

	for i := range size {
		endPos := len(line) - size + i + 1

		largest := 0
		largestPos := 0

		for j, digit := range line[startPos:endPos] {
			intDigit := int(digit - '0')
			if intDigit > largest {
				largest = intDigit
				largestPos = j
			}
		}

		foundNumber += strconv.Itoa(largest)
		startPos += largestPos + 1
	}

	result, _ := strconv.Atoi(foundNumber)
	return result
}
