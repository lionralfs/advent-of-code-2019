package main

import (
	"errors"
	"io/ioutil"
	"strconv"
	"strings"
)

type program struct {
	instructionPointer int
	intcode            []int
	outputs            []int
	inputArgs          []int
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

func (p *program) executeInstruction() int {
	operation := p.getOperation()
	position := p.instructionPointer

	switch operation[0] {
	case 1: // addition
		arg1 := p.getArg(operation[1], p.intcode[position+1])
		arg2 := p.getArg(operation[2], p.intcode[position+2])
		writeAddress := p.intcode[position+3]
		p.intcode[writeAddress] = arg1 + arg2

		return position + 4
	case 2: // multiplication
		arg1 := p.getArg(operation[1], p.intcode[position+1])
		arg2 := p.getArg(operation[2], p.intcode[position+2])
		writeAddress := p.intcode[position+3]
		p.intcode[writeAddress] = arg1 * arg2

		return position + 4
	case 3: // input
		writeAddress := p.intcode[position+1]
		p.intcode[writeAddress] = p.inputArgs[0]
		p.inputArgs = p.inputArgs[1:]
		return position + 2
	case 4: // output
		readAddress := p.getArg(operation[1], position+1)
		p.outputs = append(p.outputs, p.intcode[readAddress])
		return position + 2
	case 5: // jump-if-true
		arg1 := p.getArg(operation[1], p.intcode[position+1])
		if arg1 != 0 {
			return p.getArg(operation[2], p.intcode[position+2])
		}
		return position + 3
	case 6: // jump-if-false
		arg1 := p.getArg(operation[1], p.intcode[position+1])
		if arg1 == 0 {
			return p.getArg(operation[2], p.intcode[position+2])
		}
		return position + 3
	case 7: // less than
		arg1 := p.getArg(operation[1], p.intcode[position+1])
		arg2 := p.getArg(operation[2], p.intcode[position+2])

		writeAddress := p.intcode[position+3]

		if arg1 < arg2 {
			p.intcode[writeAddress] = 1
		} else {
			p.intcode[writeAddress] = 0
		}
		if writeAddress == position {
			return position
		}
		return position + 4
	case 8: // equals
		arg1 := p.getArg(operation[1], p.intcode[position+1])
		arg2 := p.getArg(operation[2], p.intcode[position+2])

		writeAddress := p.intcode[position+3]

		if arg1 == arg2 {
			p.intcode[writeAddress] = 1
		} else {
			p.intcode[writeAddress] = 0
		}

		return position + 4
	case 99: // halt
		return -1
	default:
		panic(errors.New("Unknown operation: " + strconv.Itoa(operation[0])))
	}
}

func (p *program) run() {
	for p.instructionPointer < len(p.intcode) {
		jumpTo := p.executeInstruction()

		if jumpTo < 0 {
			return
		}

		p.instructionPointer = jumpTo
	}

	panic(errors.New("Reached end of program without encountering opcode 99 (halt)"))
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
