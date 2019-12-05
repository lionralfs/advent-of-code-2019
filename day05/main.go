package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type program struct {
	instructionPointer int
	intcode            []int
	outputs            []int
}

func (p *program) getOperation() []int {
	result := make([]int, 4)
	instruction := p.intcode[p.instructionPointer]

	result[0] = instruction % 100
	result[1] = (instruction / 100) % 10
	result[2] = (instruction / 1000) % 10
	result[3] = (instruction / 10000) % 10

	return result
}

func (p *program) getArg(mode, value int) int {
	switch mode {
	case 0:
		return p.intcode[value]
	case 1:
		return value
	}

	panic(errors.New("Unknown mode: " + strconv.Itoa(mode)))
}

func (p *program) executeInstruction(inputArg int) int {
	operation := p.getOperation()

	switch operation[0] {
	case 1:
		arg1 := p.getArg(operation[1], p.intcode[p.instructionPointer+1])
		arg2 := p.getArg(operation[2], p.intcode[p.instructionPointer+2])
		writeAddress := p.intcode[p.instructionPointer+3]
		p.intcode[writeAddress] = arg1 + arg2

		return 4
	case 2:
		arg1 := p.getArg(operation[1], p.intcode[p.instructionPointer+1])
		arg2 := p.getArg(operation[2], p.intcode[p.instructionPointer+2])
		writeAddress := p.intcode[p.instructionPointer+3]
		p.intcode[writeAddress] = arg1 * arg2

		return 4
	case 3:
		writeAddress := p.intcode[p.instructionPointer+1]
		p.intcode[writeAddress] = inputArg
		return 2
	case 4:
		readAddress := p.getArg(operation[1], p.instructionPointer+1)
		p.outputs = append(p.outputs, p.intcode[readAddress])
		return 2
	case 99:
		return -1
	default:
		panic(errors.New("Unknown operation: " + strconv.Itoa(operation[0])))
	}
}

func (p *program) run(inputArg int) {
	for p.instructionPointer < len(p.intcode) {
		pointsToJump := p.executeInstruction(inputArg)

		if pointsToJump < 0 {
			return
		}

		p.instructionPointer += pointsToJump
	}

	panic(errors.New("Reached end of program without encountering opcode 99 (halt)"))
}

func main() {
	partOne := program{
		instructionPointer: 0,
		intcode:            readInput(),
		outputs:            []int{},
	}
	partOne.run(1)

	fmt.Printf("[Part1] Outputs: %v\n", partOne.outputs)
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
