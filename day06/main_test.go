package main

import "testing"

func TestChecksum(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{
			input: "COM)B\nB)C\nC)D\nD)E\nE)F\nB)G\nG)H\nD)I\nE)J\nJ)K\nK)L",
			want:  42,
		},
	}
	for i, test := range tests {
		if output := calculateChecksum(test.input); output != test.want {
			t.Errorf("Got %d, but expected %d in test case %d", output, test.want, i)
		}
	}
}
