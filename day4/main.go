package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	grid := [][]bool{}

	for {
		scanner.Scan()
		line := scanner.Text()

		if line == "" {
			break
		}

		lineSlice := []bool{}

		for _, value := range line {
			lineSlice = append(lineSlice, value == '@')
		}

		grid = append(grid, lineSlice)
	}

	maxX := len(grid[0]) - 1
	maxY := len(grid) - 1

	total := 0

	for {
		removed := false
		for y, row := range grid {
			for x, value := range row {
				if !value {
					continue
				}

				adjecentRolls := 0

				for i := max(0, y-1); i <= min(maxY, y+1); i++ {
					for j := max(0, x-1); j <= min(maxX, x+1); j++ {
						if i == y && j == x {
							continue
						}

						if grid[i][j] {
							adjecentRolls++
						}
					}
				}

				if adjecentRolls < 4 {
					total++
					grid[y][x] = false
					removed = true
				}
			}
		}

		if !removed {
			break
		}
	}
	fmt.Println("Part 2:", total)
}
