package main

import "testing"

func TestCalcFuel(t *testing.T) {
	tests := []struct {
		input int
		want  int
	}{
		{
			input: 12,
			want:  2,
		},
		{
			input: 14,
			want:  2,
		},
		{
			input: 1969,
			want:  654,
		},
		{
			input: 100756,
			want:  33583,
		},
	}

	for _, test := range tests {
		if output := calculateFuel(test.input); output != test.want {
			t.Errorf("calculateFuel(%d) = %v; want %d", test.input, output, test.want)
		}
	}
}

func TestCalcFuelRecursive(t *testing.T) {
	tests := []struct {
		input int
		want  int
	}{
		{
			input: 14,
			want:  2,
		},
		{
			input: 1969,
			want:  966,
		},
		{
			input: 100756,
			want:  50346,
		},
	}

	for _, test := range tests {
		if output := calculateFuelRec(test.input); output != test.want {
			t.Errorf("calculateFuelRec(%d) = %v; want %d", test.input, output, test.want)
		}
	}
}
