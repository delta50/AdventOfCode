package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func parseDay1(input string) []int {
	var convertedInput = []int{}
	for _, i := range strings.Split(input, "\n") {
		if i == "" {
			break
		}
		convertedInt, err := strconv.Atoi(i)
		if err != nil {
			log.Fatal(err)
		}
		convertedInput = append(convertedInput, convertedInt)
	}
	return convertedInput
}

func (t *T) Day1Part1(input string) {
	convertedInput := parseDay1(input)
	for i, valueA := range convertedInput {
		for j, valueB := range convertedInput {
			if i == j {
				continue
			}
			if valueA+valueB == 2020 {
				fmt.Println(valueA * valueB)
				os.Exit(0)
			}
		}
	}
}

func (t *T) Day1Part2(input string) {
	convertedInput := parseDay1(input)
	for i, valueA := range convertedInput {
		for j, valueB := range convertedInput {
			for k, valueC := range convertedInput {
				if i == j || i == k || j == k {
					continue
				}
				if valueA+valueB+valueC == 2020 {
					fmt.Println(valueA * valueB * valueC)
					os.Exit(0)
				}
			}
		}
	}
}
