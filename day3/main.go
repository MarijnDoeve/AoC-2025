package main

import (
	"bufio"
	"fmt"
	"os"
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
	sum := 0
	for _, line := range lines {
		largest1 := 0
		largest2 := 0

		largest1_pos := 0

		for i, digit := range line[:len(line)-1] {
			int_digit := int(digit - '0')
			if int_digit > largest1 {
				largest1 = int_digit
				largest1_pos = i
			}
		}

		for _, digit := range line[largest1_pos+1:] {
			int_digit := int(digit - '0')
			if int_digit > largest2 {
				largest2 = int_digit
			}
		}
		sum += largest1*10 + largest2
	}
	fmt.Println(sum)
}
