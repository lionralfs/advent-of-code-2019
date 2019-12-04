package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	bytes, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	passwordRange := strings.Split(string(bytes), "-")
	if len(passwordRange) != 2 {
		panic(errors.New("Expected range to be of length 2, it has length " + string(len(passwordRange))))
	}
	min, err := strconv.Atoi(passwordRange[0])
	if err != nil {
		panic(err)
	}
	max, err := strconv.Atoi(passwordRange[1])
	if err != nil {
		panic(err)
	}

	if min > max {
		panic(errors.New("min is larger than max"))
	}

	possiblePasswordsPart1 := countPasswords(min, max, validatePassword)
	fmt.Printf("[Part1] There are %d possible passwords\n", possiblePasswordsPart1)

	possiblePasswordsPart2 := countPasswords(min, max, validatePasswordWithGroups)
	fmt.Printf("[Part2] There are %d possible passwords\n", possiblePasswordsPart2)
}

func countPasswords(min, max int, validator func(password int) bool) int {
	result := 0
	current := min

	for current <= max {
		if validator(current) {
			result++
		}
		current++
	}

	return result
}

func validatePassword(password int) bool {
	// make sure it is 6-digits long
	if password < 100000 || password > 999999 {
		return false
	}

	hasAdjacentDigits := false
	var lastSeen rune = -1

	for _, digit := range strconv.Itoa(password) {
		if lastSeen > digit {
			return false
		}

		if lastSeen == digit {
			hasAdjacentDigits = true
		}

		lastSeen = digit
	}

	return hasAdjacentDigits
}

func validatePasswordWithGroups(password int) bool {
	// make sure it is 6-digits long
	if password < 100000 || password > 999999 {
		return false
	}

	hasTwoAdjacentDigits := false
	repeatingDigits := 0
	var lastSeen rune = -1

	for _, digit := range strconv.Itoa(password) {
		if lastSeen > digit {
			return false
		}

		if lastSeen == digit {
			repeatingDigits++
		} else {
			if repeatingDigits == 1 {
				hasTwoAdjacentDigits = true
			}
			repeatingDigits = 0
		}
		lastSeen = digit
	}

	if repeatingDigits == 1 {
		hasTwoAdjacentDigits = true
	}

	return hasTwoAdjacentDigits
}
