package main

import (
	"bufio"
	"cmp"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type junctionBox struct {
	x int32
	y int32
	z int32
}

func newJunctionBox(x, y, z string) junctionBox {
	xI, ex := strconv.Atoi(x)
	yI, ey := strconv.Atoi(y)
	zI, ez := strconv.Atoi(z)

	if ex != nil || ey != nil || ez != nil {
		fmt.Println(ex, ey, ez)
		panic("Wrong int")
	}

	return junctionBox{int32(xI), int32(yI), int32(zI)}
}

func junctionBoxDistance(a, b junctionBox) float64 {
	return math.Sqrt(math.Pow(float64(a.x)-float64(b.x), 2) + math.Pow(float64(a.y)-float64(b.y), 2) + math.Pow(float64(a.z)-float64(b.z), 2))
}

type junctionBoxPair struct {
	a        junctionBox
	b        junctionBox
	distance float64
}

func newJunctionBoxPair(a, b junctionBox) junctionBoxPair {
	distance := junctionBoxDistance(a, b)
	return junctionBoxPair{a, b, distance}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	junctionBoxes := []junctionBox{}

	for {
		scanner.Scan()

		line := scanner.Text()

		if line == "" {
			break
		}

		parts := strings.Split(line, ",")
		if len(parts) != 3 {
			panic("Not 3")
		}

		junctionBoxes = append(junctionBoxes, newJunctionBox(
			parts[0],
			parts[1],
			parts[2],
		))
	}

	var nConnections int
	if nBoxes := len(junctionBoxes); nBoxes == 20 {
		nConnections = 10
	} else if nBoxes == 1000 {
		nConnections = 1000
	} else {
		panic("Unknown number of boxes")
	}

	pairs := []junctionBoxPair{}

	for i := 0; i < len(junctionBoxes); i++ {
		for j := i + 1; j < len(junctionBoxes); j++ {
			pairs = append(pairs, newJunctionBoxPair(junctionBoxes[i], junctionBoxes[j]))
		}
	}

	slices.SortFunc(pairs, func(a, b junctionBoxPair) int { return cmp.Compare(a.distance, b.distance) })

	circuits := [][]junctionBox{}

	for _, box := range junctionBoxes {
		circuits = append(circuits, []junctionBox{box})
	}

	for i, current := range pairs {

		indexes := []int{}
		for j, circuit := range circuits {
			if slices.Contains(circuit, current.a) {
				indexes = append(indexes, j)
			}
			if slices.Contains(circuit, current.b) {
				indexes = append(indexes, j)
			}
		}

		if len(indexes) != 2 {
			panic("Did not find both boxes")
		}

		if indexes[0] == indexes[1] {
			continue
		}

		newCircuit := slices.Concat(circuits[indexes[0]], circuits[indexes[1]])
		slices.Sort(indexes)
		circuits = slices.Delete(circuits, indexes[1], indexes[1]+1)
		circuits = slices.Delete(circuits, indexes[0], indexes[0]+1)
		circuits = append(circuits, newCircuit)

		if i == nConnections-1 {
			slices.SortFunc(circuits, func(a, b []junctionBox) int { return cmp.Compare(len(b), len(a)) })

			sum := len(circuits[0]) * len(circuits[1]) * len(circuits[2])
			fmt.Println("Part 1:", sum)

		}

		if len(circuits) == 1 {
			answer := int64(current.a.x) * int64(current.b.x)
			fmt.Println("Part 2:", answer)
			os.Exit(0)
		}
	}

}
