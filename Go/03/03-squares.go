package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

func ignoreError[T any](val T, _ error) T {
	return val
}

func getInput() [][]int {
	readFile, _ := os.Open("03-squares-input.txt")

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	result := [][]int{}

	for fileScanner.Scan() {
		tri := lo.FilterMap(strings.Split(fileScanner.Text(), " "), func(s string, i int) (int, bool) {
			if s == "" {
				return 0, false
			}

			return ignoreError(strconv.Atoi(s)), true
		})

		result = append(result, tri)
	}

	return result
}

func isPossible(tri []int) bool {
	fmt.Printf("Tri: %d %d %d\n", tri[0], tri[1], tri[2])
	return tri[0]+tri[1] > tri[2] &&
		tri[0]+tri[2] > tri[1] &&
		tri[1]+tri[2] > tri[0]
}

func main() {
	input := getInput()

	ps := lo.Filter(input, func(tri []int, i int) bool {
		return isPossible(tri)
	})

	result := len(ps)

	fmt.Printf("Number of possible triangles: %d", result)
}
