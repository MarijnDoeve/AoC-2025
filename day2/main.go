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
	scanner.Scan()

	line := scanner.Text()

	ranges := strings.Split(line, ",")

	sum1 := 0
	sum2 := 0

	for _, a_range := range ranges {
		if len(a_range) == 0 {
			continue
		}

		parts := strings.Split(a_range, "-")

		a, _ := strconv.Atoi(parts[0])
		b, _ := strconv.Atoi(parts[1])

		for i := a; i <= b; i++ {
			str_i := strconv.Itoa(i)

			half_len := len(str_i) / 2

			// PART 2
			for j := half_len; j >= 0; j-- {
				part_to_find := str_i[:j]

				result := strings.ReplaceAll(str_i, part_to_find, "")

				if result == "" {
					sum2 += i
					break
				}
			}

			// PART 1
			if len(str_i)%2 != 0 {
				continue
			}

			if str_i[:half_len] == str_i[half_len:] {
				sum1 += i
			}
		}
	}

	fmt.Println("Part 1:", sum1)
	fmt.Println("Part 2:", sum2)
}
