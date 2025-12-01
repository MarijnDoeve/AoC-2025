package main

import (
	"bufio"
	"fmt"
	"log"
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

	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}

	dial := 50
	count := 0
	count2 := 0

	for _, l := range lines {
		if len(l) > 0 {
			firstChar := string(l[0])
			rest, err := strconv.Atoi(l[1:])
			if err != nil {
				log.Fatal(err)
			}

			if firstChar == "R" {
				dial += rest
			} else {
				dial -= rest
			}

			for dial > 99 {
				dial -= 100
				count2++
			}

			for dial < 0 {
				dial += 100
				count2++
			}

			if dial == 0 {
				count++
			}
		}

	}

	fmt.Println("Count1:", count)
	fmt.Println("Count2:", count2)
}
