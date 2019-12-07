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

func TestExecuteIntcode(t *testing.T) {
	tests := []struct {
		intcode           []int
		want              int
		wantConfiguration []int
	}{
		{
			intcode:           []int{3, 15, 3, 16, 1002, 16, 10, 16, 1, 16, 15, 15, 4, 15, 99, 0, 0},
			want:              43210,
			wantConfiguration: []int{4, 3, 2, 1, 0},
		},
		{
			intcode:           []int{3, 23, 3, 24, 1002, 24, 10, 24, 1002, 23, -1, 23, 101, 5, 23, 23, 1, 24, 23, 23, 4, 23, 99, 0, 0},
			want:              54321,
			wantConfiguration: []int{0, 1, 2, 3, 4},
		},
		{
			intcode:           []int{3, 31, 3, 32, 1002, 32, 10, 32, 1001, 31, -2, 31, 1007, 31, 0, 33, 1002, 33, 7, 33, 1, 33, 31, 31, 1, 32, 31, 31, 4, 31, 99, 0, 0, 0},
			want:              65210,
			wantConfiguration: []int{1, 0, 4, 3, 2},
		},
	}

	for i, test := range tests {
		want := test.want

		result, configuration := largestOutput(test.intcode)

		if result != want {
			t.Errorf("largestOutput is incorrect in test case %d\nGot %v, expected %v", i, result, want)
		}

		if !equal(test.wantConfiguration, configuration) {
			t.Errorf("Configuration is incorrect in test case %d\nGot %v, expected %v", i, configuration, test.wantConfiguration)
		}
	}
}
