package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	partOne()
	partTwo()
}

func partOne() {
	intcode := readInput()

	program := NewProgram(intcode)
	program.inputs = []int{1}

	var result int
	for {
		res, done := program.Run()
		if done {
			break
		}
		result = res
	}

	fmt.Printf("[Part1] The result is %v\n", result)
}

func partTwo() {
	intcode := readInput()

	program := NewProgram(intcode)
	program.inputs = []int{2}

	var result int
	for {
		res, done := program.Run()
		if done {
			break
		}
		result = res
	}

	fmt.Printf("[Part2] The result is %v\n", result)
}

func readInput() []int {
	bytes, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	// split the string of integers at each comma
	// which results in a list of strings
	list := strings.Split(string(bytes), ",")

	// parse the strings to integers
	opcodes := make([]int, len(list))

	for i, e := range list {
		code, err := strconv.Atoi(e)
		if err != nil {
			panic(err)
		}
		opcodes[i] = code
	}

	return opcodes
}
