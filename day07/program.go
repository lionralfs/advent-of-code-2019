package main

import (
	"errors"
	"strconv"
)

func getOperation(instruction int) []int {
	result := make([]int, 4)

	result[0] = instruction % 100
	result[1] = (instruction / 100) % 10
	result[2] = (instruction / 1000) % 10
	result[3] = (instruction / 10000) % 10

	return result
}

// A Program can execute a set of instructions (intcode)
// on some inputs which produces some outputs
type Program struct {
	pointer int
	code    []int
	output  int
	inputs  []int
	done    bool
}

func (p *Program) getArg(mode, value int) int {
	switch mode {
	case 0:
		return p.code[value]
	case 1:
		return value
	}

	panic(errors.New("Unknown mode: " + strconv.Itoa(mode)))
}

// Run runs the program, returning the next result
func (p *Program) Run() int {
	for {
		position := p.pointer
		operation := getOperation(p.code[position])

		switch operation[0] {
		case 1: // addition
			arg1 := p.getArg(operation[1], p.code[position+1])
			arg2 := p.getArg(operation[2], p.code[position+2])
			writeAddress := p.code[position+3]
			p.code[writeAddress] = arg1 + arg2

			p.pointer += 4
		case 2: // multiplication
			arg1 := p.getArg(operation[1], p.code[position+1])
			arg2 := p.getArg(operation[2], p.code[position+2])
			writeAddress := p.code[position+3]
			p.code[writeAddress] = arg1 * arg2

			p.pointer += 4
		case 3: // input
			writeAddress := p.code[position+1]
			// take the first position off the input list and remove it
			input := p.inputs[0]
			p.inputs = p.inputs[1:]
			p.code[writeAddress] = input

			p.pointer += 2
		case 4: // output
			readAddress := p.getArg(operation[1], position+1)
			p.output = p.code[readAddress]

			p.pointer += 2

			return p.output
		case 5: // jump-if-true
			arg1 := p.getArg(operation[1], p.code[position+1])
			if arg1 != 0 {
				p.pointer = p.getArg(operation[2], p.code[position+2])
			} else {
				p.pointer += 3
			}

		case 6: // jump-if-false
			arg1 := p.getArg(operation[1], p.code[position+1])
			if arg1 == 0 {
				p.pointer = p.getArg(operation[2], p.code[position+2])
			} else {
				p.pointer += 3
			}
		case 7: // less than
			arg1 := p.getArg(operation[1], p.code[position+1])
			arg2 := p.getArg(operation[2], p.code[position+2])

			writeAddress := p.code[position+3]

			if arg1 < arg2 {
				p.code[writeAddress] = 1
			} else {
				p.code[writeAddress] = 0
			}

			if writeAddress == position {
				p.pointer = position
			} else {
				p.pointer += 4
			}

		case 8: // equals
			arg1 := p.getArg(operation[1], p.code[position+1])
			arg2 := p.getArg(operation[2], p.code[position+2])

			writeAddress := p.code[position+3]

			if arg1 == arg2 {
				p.code[writeAddress] = 1
			} else {
				p.code[writeAddress] = 0
			}

			p.pointer += 4
		case 99: // halt
			p.done = true
			return p.output
		default:
			panic(errors.New("Unknown operation: " + strconv.Itoa(operation[0])))
		}
	}
}
