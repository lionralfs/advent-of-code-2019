package main

import (
	"testing"
)

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

type want struct {
	finalProgramState []int
	outputs           []int
}

func TestExecuteIntcode(t *testing.T) {
	tests := []struct {
		program  program
		inputArg int
		want     want
	}{
		{
			inputArg: 2,
			program: program{
				instructionPointer: 0,
				intcode:            []int{3, 0, 4, 0, 99},
				outputs:            []int{},
			},
			want: want{
				finalProgramState: []int{2, 0, 4, 0, 99},
				outputs:           []int{2},
			},
		},
		{
			inputArg: 2, // doesn't really matter in this case
			program: program{
				instructionPointer: 0,
				intcode:            []int{1002, 4, 3, 4, 33},
				outputs:            []int{},
			},
			want: want{
				finalProgramState: []int{1002, 4, 3, 4, 99},
				outputs:           []int{},
			},
		},
		{
			inputArg: 2,
			program: program{
				instructionPointer: 0,
				intcode:            []int{1101, 100, -1, 4, 0},
				outputs:            []int{},
			},
			want: want{
				finalProgramState: []int{1101, 100, -1, 4, 99},
				outputs:           []int{},
			},
		},
		{
			inputArg: 1,
			program: program{
				instructionPointer: 0,
				intcode:            []int{3, 225, 1, 225, 6, 6, 1100, 1, 238, 225, 104, 0, 1102, 35, 92, 225, 1101, 25, 55, 225, 1102, 47, 36, 225, 1102, 17, 35, 225, 1, 165, 18, 224, 1001, 224, -106, 224, 4, 224, 102, 8, 223, 223, 1001, 224, 3, 224, 1, 223, 224, 223, 1101, 68, 23, 224, 101, -91, 224, 224, 4, 224, 102, 8, 223, 223, 101, 1, 224, 224, 1, 223, 224, 223, 2, 217, 13, 224, 1001, 224, -1890, 224, 4, 224, 102, 8, 223, 223, 1001, 224, 6, 224, 1, 224, 223, 223, 1102, 69, 77, 224, 1001, 224, -5313, 224, 4, 224, 1002, 223, 8, 223, 101, 2, 224, 224, 1, 224, 223, 223, 102, 50, 22, 224, 101, -1800, 224, 224, 4, 224, 1002, 223, 8, 223, 1001, 224, 5, 224, 1, 224, 223, 223, 1102, 89, 32, 225, 1001, 26, 60, 224, 1001, 224, -95, 224, 4, 224, 102, 8, 223, 223, 101, 2, 224, 224, 1, 223, 224, 223, 1102, 51, 79, 225, 1102, 65, 30, 225, 1002, 170, 86, 224, 101, -2580, 224, 224, 4, 224, 102, 8, 223, 223, 1001, 224, 6, 224, 1, 223, 224, 223, 101, 39, 139, 224, 1001, 224, -128, 224, 4, 224, 102, 8, 223, 223, 101, 3, 224, 224, 1, 223, 224, 223, 1102, 54, 93, 225, 4, 223, 99, 0, 0, 0, 677, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1105, 0, 99999, 1105, 227, 247, 1105, 1, 99999, 1005, 227, 99999, 1005, 0, 256, 1105, 1, 99999, 1106, 227, 99999, 1106, 0, 265, 1105, 1, 99999, 1006, 0, 99999, 1006, 227, 274, 1105, 1, 99999, 1105, 1, 280, 1105, 1, 99999, 1, 225, 225, 225, 1101, 294, 0, 0, 105, 1, 0, 1105, 1, 99999, 1106, 0, 300, 1105, 1, 99999, 1, 225, 225, 225, 1101, 314, 0, 0, 106, 0, 0, 1105, 1, 99999, 1008, 677, 677, 224, 1002, 223, 2, 223, 1005, 224, 329, 101, 1, 223, 223, 7, 677, 677, 224, 102, 2, 223, 223, 1006, 224, 344, 101, 1, 223, 223, 108, 677, 677, 224, 1002, 223, 2, 223, 1006, 224, 359, 1001, 223, 1, 223, 7, 677, 226, 224, 1002, 223, 2, 223, 1005, 224, 374, 1001, 223, 1, 223, 1107, 677, 226, 224, 1002, 223, 2, 223, 1005, 224, 389, 1001, 223, 1, 223, 107, 226, 677, 224, 102, 2, 223, 223, 1005, 224, 404, 1001, 223, 1, 223, 1108, 226, 677, 224, 1002, 223, 2, 223, 1006, 224, 419, 101, 1, 223, 223, 107, 226, 226, 224, 102, 2, 223, 223, 1005, 224, 434, 1001, 223, 1, 223, 108, 677, 226, 224, 1002, 223, 2, 223, 1006, 224, 449, 101, 1, 223, 223, 108, 226, 226, 224, 102, 2, 223, 223, 1006, 224, 464, 1001, 223, 1, 223, 1007, 226, 226, 224, 1002, 223, 2, 223, 1005, 224, 479, 101, 1, 223, 223, 8, 677, 226, 224, 1002, 223, 2, 223, 1006, 224, 494, 101, 1, 223, 223, 1007, 226, 677, 224, 102, 2, 223, 223, 1006, 224, 509, 101, 1, 223, 223, 7, 226, 677, 224, 1002, 223, 2, 223, 1005, 224, 524, 101, 1, 223, 223, 107, 677, 677, 224, 102, 2, 223, 223, 1005, 224, 539, 101, 1, 223, 223, 1008, 677, 226, 224, 1002, 223, 2, 223, 1005, 224, 554, 1001, 223, 1, 223, 1008, 226, 226, 224, 1002, 223, 2, 223, 1006, 224, 569, 1001, 223, 1, 223, 1108, 226, 226, 224, 102, 2, 223, 223, 1005, 224, 584, 101, 1, 223, 223, 1107, 226, 677, 224, 1002, 223, 2, 223, 1005, 224, 599, 1001, 223, 1, 223, 8, 226, 677, 224, 1002, 223, 2, 223, 1006, 224, 614, 1001, 223, 1, 223, 1108, 677, 226, 224, 102, 2, 223, 223, 1005, 224, 629, 1001, 223, 1, 223, 8, 226, 226, 224, 1002, 223, 2, 223, 1005, 224, 644, 1001, 223, 1, 223, 1107, 677, 677, 224, 1002, 223, 2, 223, 1005, 224, 659, 1001, 223, 1, 223, 1007, 677, 677, 224, 1002, 223, 2, 223, 1005, 224, 674, 101, 1, 223, 223, 4, 223, 99, 226},
				outputs:            []int{},
			},
			want: want{
				finalProgramState: []int{3, 225, 1, 225, 6, 6, 1101, 1, 238, 225, 104, 0, 1102, 35, 92, 225, 1101, 25, 55, 225, 1102, 47, 36, 225, 1102, 17, 35, 225, 1, 165, 18, 224, 1001, 224, -106, 224, 4, 224, 102, 8, 223, 223, 1001, 224, 3, 224, 1, 223, 224, 223, 1101, 68, 23, 224, 101, -91, 224, 224, 4, 224, 102, 8, 223, 223, 101, 1, 224, 224, 1, 223, 224, 223, 2, 217, 13, 224, 1001, 224, -1890, 224, 4, 224, 102, 8, 223, 223, 1001, 224, 6, 224, 1, 224, 223, 223, 1102, 69, 77, 224, 1001, 224, -5313, 224, 4, 224, 1002, 223, 8, 223, 101, 2, 224, 224, 1, 224, 223, 223, 102, 50, 22, 224, 101, -1800, 224, 224, 4, 224, 1002, 223, 8, 223, 1001, 224, 5, 224, 1, 224, 223, 223, 1102, 89, 32, 225, 1001, 26, 60, 224, 1001, 224, -95, 224, 4, 224, 102, 8, 223, 223, 101, 2, 224, 224, 1, 223, 224, 223, 1102, 51, 79, 225, 1102, 65, 30, 225, 1002, 170, 86, 224, 101, -2580, 224, 224, 4, 224, 102, 8, 223, 223, 1001, 224, 6, 224, 1, 223, 224, 223, 101, 39, 139, 224, 1001, 224, -128, 224, 4, 224, 102, 8, 223, 223, 101, 3, 224, 224, 1, 223, 224, 223, 1102, 54, 93, 225, 4, 223, 99, 6761139, 3, 5022, 677, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1105, 0, 99999, 1105, 227, 247, 1105, 1, 99999, 1005, 227, 99999, 1005, 0, 256, 1105, 1, 99999, 1106, 227, 99999, 1106, 0, 265, 1105, 1, 99999, 1006, 0, 99999, 1006, 227, 274, 1105, 1, 99999, 1105, 1, 280, 1105, 1, 99999, 1, 225, 225, 225, 1101, 294, 0, 0, 105, 1, 0, 1105, 1, 99999, 1106, 0, 300, 1105, 1, 99999, 1, 225, 225, 225, 1101, 314, 0, 0, 106, 0, 0, 1105, 1, 99999, 1008, 677, 677, 224, 1002, 223, 2, 223, 1005, 224, 329, 101, 1, 223, 223, 7, 677, 677, 224, 102, 2, 223, 223, 1006, 224, 344, 101, 1, 223, 223, 108, 677, 677, 224, 1002, 223, 2, 223, 1006, 224, 359, 1001, 223, 1, 223, 7, 677, 226, 224, 1002, 223, 2, 223, 1005, 224, 374, 1001, 223, 1, 223, 1107, 677, 226, 224, 1002, 223, 2, 223, 1005, 224, 389, 1001, 223, 1, 223, 107, 226, 677, 224, 102, 2, 223, 223, 1005, 224, 404, 1001, 223, 1, 223, 1108, 226, 677, 224, 1002, 223, 2, 223, 1006, 224, 419, 101, 1, 223, 223, 107, 226, 226, 224, 102, 2, 223, 223, 1005, 224, 434, 1001, 223, 1, 223, 108, 677, 226, 224, 1002, 223, 2, 223, 1006, 224, 449, 101, 1, 223, 223, 108, 226, 226, 224, 102, 2, 223, 223, 1006, 224, 464, 1001, 223, 1, 223, 1007, 226, 226, 224, 1002, 223, 2, 223, 1005, 224, 479, 101, 1, 223, 223, 8, 677, 226, 224, 1002, 223, 2, 223, 1006, 224, 494, 101, 1, 223, 223, 1007, 226, 677, 224, 102, 2, 223, 223, 1006, 224, 509, 101, 1, 223, 223, 7, 226, 677, 224, 1002, 223, 2, 223, 1005, 224, 524, 101, 1, 223, 223, 107, 677, 677, 224, 102, 2, 223, 223, 1005, 224, 539, 101, 1, 223, 223, 1008, 677, 226, 224, 1002, 223, 2, 223, 1005, 224, 554, 1001, 223, 1, 223, 1008, 226, 226, 224, 1002, 223, 2, 223, 1006, 224, 569, 1001, 223, 1, 223, 1108, 226, 226, 224, 102, 2, 223, 223, 1005, 224, 584, 101, 1, 223, 223, 1107, 226, 677, 224, 1002, 223, 2, 223, 1005, 224, 599, 1001, 223, 1, 223, 8, 226, 677, 224, 1002, 223, 2, 223, 1006, 224, 614, 1001, 223, 1, 223, 1108, 677, 226, 224, 102, 2, 223, 223, 1005, 224, 629, 1001, 223, 1, 223, 8, 226, 226, 224, 1002, 223, 2, 223, 1005, 224, 644, 1001, 223, 1, 223, 1107, 677, 677, 224, 1002, 223, 2, 223, 1005, 224, 659, 1001, 223, 1, 223, 1007, 677, 677, 224, 1002, 223, 2, 223, 1005, 224, 674, 101, 1, 223, 223, 4, 223, 99, 226},
				outputs:           []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 6761139},
			},
		},
		{
			inputArg: 2,
			program: program{
				instructionPointer: 0,
				intcode:            []int{1101, 100, -1, 4, 0},
				outputs:            []int{},
			},
			want: want{
				finalProgramState: []int{1101, 100, -1, 4, 99},
				outputs:           []int{},
			},
		},
		{
			inputArg: 8,
			program: program{
				instructionPointer: 0,
				intcode:            []int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8},
				outputs:            []int{},
			},
			want: want{
				finalProgramState: []int{},
				outputs:           []int{1},
			},
		},
		{
			inputArg: 9,
			program: program{
				instructionPointer: 0,
				intcode:            []int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8},
				outputs:            []int{},
			},
			want: want{
				finalProgramState: []int{},
				outputs:           []int{0},
			},
		},
		{
			inputArg: 7,
			program: program{
				instructionPointer: 0,
				intcode:            []int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8},
				outputs:            []int{},
			},
			want: want{
				finalProgramState: []int{},
				outputs:           []int{1},
			},
		},
		{
			inputArg: 9,
			program: program{
				instructionPointer: 0,
				intcode:            []int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8},
				outputs:            []int{},
			},
			want: want{
				finalProgramState: []int{},
				outputs:           []int{0},
			},
		},
		{
			inputArg: 7,
			program: program{
				instructionPointer: 0,
				intcode:            []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99},
				outputs:            []int{},
			},
			want: want{
				finalProgramState: []int{},
				outputs:           []int{999},
			},
		},
		{
			inputArg: 8,
			program: program{
				instructionPointer: 0,
				intcode:            []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99},
				outputs:            []int{},
			},
			want: want{
				finalProgramState: []int{},
				outputs:           []int{1000},
			},
		},
		{
			inputArg: 9,
			program: program{
				instructionPointer: 0,
				intcode:            []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99},
				outputs:            []int{},
			},
			want: want{
				finalProgramState: []int{},
				outputs:           []int{1001},
			},
		},
	}

	for i, test := range tests {
		program := test.program
		inputArg := test.inputArg
		want := test.want

		program.run(inputArg)

		if len(want.finalProgramState) > 0 && !equal(program.intcode, want.finalProgramState) {
			t.Errorf("Final program state is incorrect in test case %d\nGot %v, expected %v", i, program.intcode, want.finalProgramState)
		}

		if !equal(program.outputs, want.outputs) {
			t.Errorf("Output is incorrect in test case %d\nGot %v, expected %v", i, program.outputs, want.outputs)
		}
	}
}
