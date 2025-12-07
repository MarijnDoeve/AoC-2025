package main

import (
	"bufio"
	"fmt"
	"os"
)

const EMPTY = '.'
const SPLITTER = '^'
const START = 'S'

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	diagram := [][]rune{}

	for {
		scanner.Scan()
		line := scanner.Text()
		if line == "" {
			break
		}

		splitLine := []rune{}
		for _, i := range line {
			splitLine = append(splitLine, i)
		}
		diagram = append(diagram, splitLine)
	}

	nBeams := make([][]int, len(diagram))
	for i := range nBeams {
		nBeams[i] = make([]int, len(diagram[i]))
	}

	nSplits := 0

	for nRow, row := range diagram[:len(diagram)-1] {
		for nPos, pos := range row {
			nTimelinesInPos := nBeams[nRow][nPos]

			switch pos {
			case EMPTY:
				if nTimelinesInPos == 0 {
					continue
				}
				nBeams[nRow+1][nPos] += nTimelinesInPos

			case SPLITTER:
				if nTimelinesInPos == 0 {
					continue
				}
				nSplits++
				nBeams[nRow+1][nPos-1] += nTimelinesInPos
				nBeams[nRow+1][nPos+1] += nTimelinesInPos

			case START:
				nBeams[nRow+1][nPos] = 1
			}
		}
	}

	nTimelines := 0
	for _, n := range nBeams[len(nBeams)-1] {
		nTimelines += n
	}

	fmt.Println("Part 1:", nSplits)
	fmt.Println("Part 2:", nTimelines)
}
