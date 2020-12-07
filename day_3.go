package main

import (
	"fmt"
	"os"
	"strings"
)

func parseDay3(input string) [][]rune {
	convertedInput := [][]rune{}
	for _, value1 := range strings.Split(input, "\n") {
		if value1 == "" {
			break
		}

		row := []rune{}
		for _, value2 := range value1 {
			row = append(row, value2)
		}
		convertedInput = append(convertedInput, row)
	}
	return convertedInput
}

func (t *T) Day3Part1(input string) {
	x := 3
	y := 1
	convertedInput := parseDay3(input)
	width := len(convertedInput[0])

	trees := 0
	j := 0
	for i := 0; i < len(convertedInput); i += y {
		if convertedInput[i][j] == '#' {
			trees++
		}
		j += x
		if j >= width {
			j -= width
		}
	}
	fmt.Println(trees)
	os.Exit(0)
}

func (t *T) Day3Part2(input string) {
	x := [5]int{1, 3, 5, 7, 1}
	y := [5]int{1, 1, 1, 1, 2}

	convertedInput := parseDay3(input)
	width := len(convertedInput[0])

	total := 1
	for k := 0; k < 5; k++ {
		trees := 0
		j := 0
		for i := 0; i < len(convertedInput); i += y[k] {
			if convertedInput[i][j] == '#' {
				trees++
			}
			j += x[k]
			if j >= width {
				j -= width
			}
		}
		total *= trees
	}
	fmt.Println(total)
	os.Exit(0)
}
