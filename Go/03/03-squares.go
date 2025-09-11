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

func splitToInts(s string) []int {
	return lo.FilterMap(strings.Split(s, " "), func(s string, i int) (int, bool) {
		if s == "" {
			return 0, false
		}

		return ignoreError(strconv.Atoi(s)), true
	})
}

func getInput2() [][]int {
	readFile, _ := os.Open("03-squares-input.txt")

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	result := [][]int{}

	for fileScanner.Scan() {
		l1 := fileScanner.Text()
		fileScanner.Scan()
		l2 := fileScanner.Text()
		fileScanner.Scan()
		l3 := fileScanner.Text()

		is1 := splitToInts(l1)
		is2 := splitToInts(l2)
		is3 := splitToInts(l3)

		result = append(result, []int{is1[0], is2[0], is3[0]})
		result = append(result, []int{is1[1], is2[1], is3[1]})
		result = append(result, []int{is1[2], is2[2], is3[2]})
	}

	return result
}

func isPossible(tri []int) bool {
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

	fmt.Printf("Number of possible triangles: %d\n", result)

	// Part 2
	input = getInput2()

	ps = lo.Filter(input, func(tri []int, i int) bool {
		return isPossible(tri)
	})

	result = len(ps)

	fmt.Printf("Number of possible triangles: %d\n", result)
}
