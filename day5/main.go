package main

import (
	"bufio"
	"cmp"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	ingredients := [][2]int{}

	for {
		scanner.Scan()

		line := scanner.Text()
		if line == "" {
			break
		}
		numbers := strings.Split(line, "-")

		if len(numbers) != 2 {
			panic("no two numbers")
		}
		a, _ := strconv.Atoi(numbers[0])
		b, _ := strconv.Atoi(numbers[1])

		ingredients = append(ingredients, [2]int{a, b})
	}

	// Part 2
	slices.SortFunc(ingredients, func(a, b [2]int) int { return cmp.Compare(a[0], b[0]) })
	merged := [][2]int{ingredients[0]}

	for _, current := range ingredients[1:] {
		last := merged[len(merged)-1]

		if current[0] <= last[1] {
			merged[len(merged)-1][1] = max(last[1], current[1])
		} else {
			merged = append(merged, current)
		}
	}

	sum1 := 0

	for {
		scanner.Scan()
		line := scanner.Text()

		if line == "" {
			break
		}

		number, _ := strconv.Atoi(line)
		for _, x := range merged {
			if number >= x[0] && number <= x[1] {
				sum1++
				break
			}
		}
	}

	fmt.Println("Part 1: ", sum1)

	sum2 := 0
	for _, numRange := range merged {
		sum2 += numRange[1] - numRange[0] + 1
	}
	fmt.Println("Part 2: ", sum2)
}
