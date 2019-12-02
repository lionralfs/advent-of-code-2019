package main

import (
	"fmt"
	"testing"
)

func TestProgramExecution(t *testing.T) {
	tests := []struct {
		program []int
		want    []int
	}{
		{
			program: []int{1, 0, 0, 0, 99},
			want:    []int{2, 0, 0, 0, 99},
		},
		{
			program: []int{1, 0, 0, 3, 99},
			want:    []int{1, 0, 0, 2, 99},
		},
		{
			program: []int{2, 3, 0, 3, 99},
			want:    []int{2, 3, 0, 6, 99},
		},
		{
			program: []int{2, 4, 4, 5, 99, 0},
			want:    []int{2, 4, 4, 5, 99, 9801},
		},
		{
			program: []int{1, 1, 1, 4, 99, 5, 6, 0, 99},
			want:    []int{30, 1, 1, 4, 2, 5, 6, 0, 99},
		},
	}

	for testCase, test := range tests {
		program, want := test.program, test.want
		output := executeIntcode(program)
		fmt.Println(output)

		if len(output) != len(want) {
			t.Errorf("Program length doesn't match output length")
			return
		}

		for i, e := range want {
			if output[i] != e {
				t.Errorf("Test case %d failed, got %d; want %d at index %d", testCase, output[i], want[i], i)
			}
		}
	}
}
