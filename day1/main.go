package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func mod(a, b int) int {
	return (a%b + b) % b
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}

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

	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}

	dial := 50
	count1 := 0
	count2 := 0

	for _, l := range lines {

		direction := string(l[0])
		num, err := strconv.Atoi(l[1:])
		if err != nil {
			log.Fatal(err)
		}

		if direction == "R" {
			dial += num
			count2 += dial / 100
			dial = mod(dial, 100)
		} else {
			if dial-num < 0 {
				count2 += abs((100 + dial - num) / 100)
				if dial != 0 {
					count2++
				}
			}
			dial = mod(dial-num, 100)
			if dial == 0 {
				count2++
			}
		}

		if dial == 0 {
			count1++
		}

	}

	fmt.Println("Count1:", count1)
	fmt.Println("Count2:", count2)
}
