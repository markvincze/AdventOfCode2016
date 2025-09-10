package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func abs(i int) int {
	return int(math.Abs(float64(i)))
}

func getInput() []string {
	readFile, _ := os.Open("02-bathroom-security-input.txt")

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	result := []string{}

	for fileScanner.Scan() {
		result = append(result, fileScanner.Text())
	}

	return result
}

func step(x int, y int, dir rune) (int, int) {
	switch dir {
	case 'U':
		if y > 0 {
			y = y - 1
		}
	case 'R':
		if x < 2 {
			x = x + 1
		}
	case 'D':
		if y < 2 {
			y = y + 1
		}
	case 'L':
		if x > 0 {
			x = x - 1
		}
	}

	return x, y
}

func step2(x int, y int, dir rune) (int, int) {
	switch dir {
	case 'U':
		if y > abs(x-2) {
			y = y - 1
		}
	case 'R':
		if x < 4-abs(y-2) {
			x = x + 1
		}
	case 'D':
		if y < 4-abs(x-2) {
			y = y + 1
		}
	case 'L':
		if x > abs(y-2) {
			x = x - 1
		}
	}

	return x, y
}

func getNumber(x int, y int) byte {
	switch x {
	case 0:
		switch y {
		case 0:
			return 1
		case 1:
			return 4
		case 2:
			return 7
		}
	case 1:
		switch y {
		case 0:
			return 2
		case 1:
			return 5
		case 2:
			return 8
		}
	case 2:
		switch y {
		case 0:
			return 3
		case 1:
			return 6
		case 2:
			return 9
		}
	}

	return 0
}

func getNumber2(x int, y int) rune {
	switch x {
	case 0:
		switch y {
		case 2:
			return '5'
		}
	case 1:
		switch y {
		case 1:
			return '2'
		case 2:
			return '6'
		case 3:
			return 'A'
		}
	case 2:
		switch y {
		case 0:
			return '1'
		case 1:
			return '3'
		case 2:
			return '7'
		case 3:
			return 'B'
		case 4:
			return 'D'
		}
	case 3:
		switch y {
		case 1:
			return '4'
		case 2:
			return '8'
		case 3:
			return 'C'
		}
	case 4:
		switch y {
		case 2:
			return '9'
		}
	}

	return '0'
}

func main() {
	input := getInput()

	x, y := 1, 1
	code := []rune{}

	for _, i := range input {
		for _, m := range i {
			// x, y = step(x, y, m)
			x, y = step2(x, y, m)
		}

		// code = append(code, getNumber(x, y))
		code = append(code, getNumber2(x, y))
	}

	for _, c := range code {
		fmt.Printf("%c", c)
	}
}
