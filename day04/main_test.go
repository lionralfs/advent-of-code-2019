package main

import "testing"

func TestPasswordValidator(t *testing.T) {
	tests := []struct {
		input int
		want  bool
	}{
		{
			input: 111111,
			want:  true,
		},
		{
			input: 223450,
			want:  false,
		},
		{
			input: 123789,
			want:  false,
		},
	}

	for i, test := range tests {
		if valid := validatePassword(test.input); valid != test.want {
			t.Errorf("Got %v, but want %v in test case %d", valid, test.want, i)
		}
	}
}
