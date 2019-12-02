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

	b, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	// split the string of integers at each comma
	// which results in a list of strings
	list := strings.Split(string(b), ",")

	// parse the strings to integers
	opcodes := make([]int, len(list))

	for i, e := range list {
		code, err := strconv.Atoi(e)
		if err != nil {
			panic(err)
		}
		opcodes[i] = code
	}

	programCopy := make([]int, len(list))
	copy(programCopy, opcodes)

	partOne(programCopy)
	partTwo(opcodes)
}

func executeIntcode(program []int) []int {
	for i := 0; i < len(program); i += 4 {
		opcode := program[i]
		switch opcode {
		// addition
		case 1:
			program[program[i+3]] = program[program[i+1]] + program[program[i+2]]

		// multiplication
		case 2:
			program[program[i+3]] = program[program[i+1]] * program[program[i+2]]

		// halt
		case 99:
			return program
		}
	}

	panic(errors.New("Reached end of program without encountering opcode 99 (halt)"))
}

func partOne(program []int) {
	program[1] = 12
	program[2] = 2
	result := executeIntcode(program)

	fmt.Printf("[Part1] The value at position 0 after the running the program is: %d\n", result[0])
}

// just try every combination
func partTwo(program []int) {
	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			programCopy := make([]int, len(program))
			copy(programCopy, program)
			programCopy[1] = noun
			programCopy[2] = verb

			result := executeIntcode(programCopy)

			if result[0] == 19690720 {
				fmt.Printf("[Part2]\tFound the matching noun/verb pair; noun: %d, verb: %d\n\t100 * noun + verb = %d\n", noun, verb, 100*noun+verb)
				return
			}
		}
	}
}
