package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
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

	possiblePasswords := countPasswords(min, max)
	fmt.Printf("[Part1] There are %d possible passwords\n", possiblePasswords)
}

func countPasswords(min, max int) int {
	result := 0
	current := min

	for current <= max {
		if validatePassword(current) {
			result++
		}
		current++
	}

	return result
}

func toDigits(n int) []int {
	result := make([]int, 6)
	asString := strconv.Itoa(n)

	for i, e := range asString {
		digit, _ := strconv.Atoi(string(e))
		result[i] = digit
	}

	return result
}

func validatePassword(password int) bool {
	// make sure it is 6-digits long
	if password < 100000 || password > 999999 {
		return false
	}

	digits := toDigits(password)
	hasAdjacentDigits := false
	lastSeen := -1

	for _, digit := range digits {
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
