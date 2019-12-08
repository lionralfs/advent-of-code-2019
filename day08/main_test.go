package main

import "testing"

func TestMinZeroDigits(t *testing.T) {
	tests := []struct {
		input    string
		width    int
		height   int
		min      int
		minIndex int
	}{
		// {
		// 	input:    "123456789012",
		// 	width:    3,
		// 	height:   2,
		// 	min:      0,
		// 	minIndex: 0,
		// },
		// {
		// 	input:    "100456789012",
		// 	width:    3,
		// 	height:   2,
		// 	min:      1,
		// 	minIndex: 1,
		// },
		{
			input:    "0222112220120000",
			width:    2,
			height:   2,
			min:      0,
			minIndex: 1,
		},
	}

	for i, test := range tests {
		min, minIndex, _ := fewestZeroDigits(test.input, test.width, test.height)

		if min != test.min {
			t.Errorf("Wrong min in test case %v; Got %v, expected, %v", i, min, test.min)
		}

		if minIndex != test.minIndex {
			t.Errorf("Wrong minIndex in test case %v; Got %v, expected, %v", i, minIndex, test.minIndex)
		}
	}
}
