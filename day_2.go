package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Day2Password struct {
	min      int
	max      int
	letter   string
	password string
}

func parseDay2(input string) []Day2Password {
	var convertedInput = []Day2Password{}
	for _, i := range strings.Split(input, "\n") {
		if i == "" {
			continue
		}
		substrings := regexp.MustCompile(`(\d+)-(\d+) (\D): (\D+)`).FindStringSubmatch(i)

		min, err := strconv.Atoi(substrings[1])
		if err != nil {
			log.Fatal(err)
		}

		max, err := strconv.Atoi(substrings[2])
		if err != nil {
			log.Fatal(err)
		}

		convertedInput = append(convertedInput, Day2Password{min, max, substrings[3], substrings[4]})
	}
	log.Print(len(convertedInput))
	return convertedInput
}

func (t *T) Day2Part1(input string) {
	convertedInput := parseDay2(input)
	validPassword := 0
	for _, pass := range convertedInput {
		occurances := strings.Count(pass.password, pass.letter)
		if occurances >= pass.min && occurances <= pass.max {
			validPassword++
		}
	}
	fmt.Println(validPassword)
	os.Exit(0)
}

func (t *T) Day2Part2(input string) {
	convertedInput := parseDay2(input)
	validPassword := 0
	for _, pass := range convertedInput {
		firstMatch := string(pass.password[pass.min-1]) == pass.letter
		secondMatch := string(pass.password[pass.max-1]) == pass.letter
		if (firstMatch || secondMatch) && firstMatch != secondMatch {
			validPassword++
		}
	}
	fmt.Println(validPassword)
	os.Exit(0)
}
