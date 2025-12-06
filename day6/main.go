package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	lines := [][]int{}
	rawLines := []string{}
	var operations []string
	for {
		scanner.Scan()
		line := scanner.Text()
		if line == "" {
			break
		}
		rawLines = append(rawLines, line)

		numbers := strings.Fields(line)
		intNumbers := []int{}

		for _, number := range numbers {
			i, err := strconv.Atoi(number)

			if err != nil {
				break
			}
			intNumbers = append(intNumbers, i)
		}
		if len(intNumbers) == 0 {
			operations = numbers
		} else {
			lines = append(lines, intNumbers)
		}

	}

	// Part 1

	total1 := 0

	for i, operation := range operations {
		column := []int{}

		for _, line := range lines {
			column = append(column, line[i])
		}

		partAnswer := 0
		for _, value := range column {
			if operation == "+" {
				partAnswer += value
			} else {
				if partAnswer == 0 {
					partAnswer = value
				} else {
					partAnswer *= value
				}
			}
		}
		total1 += partAnswer
	}

	fmt.Println("Part 1:", total1)

	// Part 2

	newWidth := len(rawLines)
	newHeight := len(rawLines[0])

	rotatedGrid := []string{}

	for y := range newHeight {
		newLine := ""
		for x := range newWidth {
			newLine += string(rawLines[x][newHeight-1-y])
		}
		rotatedGrid = append(rotatedGrid, newLine)
	}

	total2 := 0
	collectedNumbers := []int{}
	for _, line := range rotatedGrid {
		if line == "" {
			continue
		}
		var strNumber string
		if strings.Contains(line, "+") {
			strNumber = strings.ReplaceAll(line, "+", "")
		} else if strings.Contains(line, "*") {
			strNumber = strings.ReplaceAll(line, "*", "")
		} else {
			strNumber = line
		}

		intNumber, _ := strconv.Atoi(strings.ReplaceAll(strNumber, " ", ""))
		if intNumber != 0 {
			collectedNumbers = append(collectedNumbers, intNumber)
		}

		if strings.Contains(line, "+") {
			partSum := collectedNumbers[0]

			for _, num := range collectedNumbers[1:] {
				partSum += num
			}
			total2 += partSum
			collectedNumbers = []int{}
		} else if strings.Contains(line, "*") {
			partProduct := collectedNumbers[0]
			for _, num := range collectedNumbers[1:] {
				partProduct *= num
			}
			total2 += partProduct
			collectedNumbers = []int{}
		}

	}

	fmt.Println("Part 2:", total2)
}
