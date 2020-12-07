package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func parseDay4(input string) []string {
	return strings.Split(input, "\n\n")
}

func (t *T) Day4Part1(input string) {
	fields := []string{"byr:", "iyr:", "eyr:", "hgt:", "hcl:", "ecl:", "pid:"}
	convertedInput := parseDay4(input)
	validPassports := 0
	for _, passport := range convertedInput {
		isValid := true
		for _, field := range fields {
			if !strings.Contains(passport, field) {
				isValid = false
			}
		}
		if isValid == true {
			validPassports++
		}
	}
	fmt.Println(validPassports)
	os.Exit(0)
}

func (t *T) Day4Part2(input string) {
	convertedInput := parseDay4(input)
	validPassports := 0
	for _, passport := range convertedInput {
		isValid :=
			Day4ByrIsValid(passport) &&
			Day4IyrIsValid(passport) &&
			Day4EyrIsValid(passport) &&
			Day4HgtIsValid(passport) &&
			Day4HclIsValid(passport) &&
			Day4EclIsValid(passport) &&
			Day4PidIsValid(passport)
		if isValid == true {
			validPassports++
		}
	}
	fmt.Println(validPassports)
	os.Exit(0)
}

func Day4ByrIsValid(passport string) bool{
	byr := regexp.MustCompile(`byr:(\d{4})(\s|$)`).FindStringSubmatch(passport)
	if byr == nil {
		return false
	}
	year, err := strconv.Atoi(byr[1])
	if err != nil {
		log.Fatal(err)
	}
	if year < 1920 || year > 2002 {
		return false
	}
	return true
}

func Day4IyrIsValid(passport string) bool{
	iyr := regexp.MustCompile(`iyr:(\d{4})(\s|$)`).FindStringSubmatch(passport)
	if iyr == nil {
		return false
	}
	year, err := strconv.Atoi(iyr[1])
	if err != nil {
		log.Fatal(err)
	}
	if year < 2010 || year > 2020 {
		return false
	}
	return true
}

func Day4EyrIsValid(passport string) bool{
	eyr := regexp.MustCompile(`eyr:(\d{4})(\s|$)`).FindStringSubmatch(passport)
	if eyr == nil {
		return false
	}
	year, err := strconv.Atoi(eyr[1])
	if err != nil {
		log.Fatal(err)
	}
	if year < 2020 || year > 2030 {
		return false
	}
	return true
}

func Day4HgtIsValid(passport string) bool{
	hgt := regexp.MustCompile(`hgt:(\d+)(cm|in)(\s|$)`).FindStringSubmatch(passport)
	if hgt == nil {
		return false
	}
	height, err := strconv.Atoi(hgt[1])
	if err != nil {
		log.Fatal(err)
	}
	if hgt[2] == "cm" && height >= 150 && height <= 193 {
		return true
	}
	if hgt[2] == "in" && height >= 59 && height <= 76 {
		return true
	}
	return false
}

func Day4HclIsValid(passport string) bool{
	return regexp.MustCompile(`hcl:#[0-9|a-f]{6}(\s|$)`).FindString(passport) != ""
}

func Day4EclIsValid(passport string) bool{
	return regexp.MustCompile(`ecl:(amb|blu|brn|gry|grn|hzl|oth)(\s|$)`).FindString(passport) != ""
}

func Day4PidIsValid(passport string) bool{
	return regexp.MustCompile(`pid:\d{9}(\s|$)`).FindString(passport) != ""
}
