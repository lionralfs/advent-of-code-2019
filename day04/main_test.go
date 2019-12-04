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

func TestPasswordValidatorWithGroups(t *testing.T) {
	tests := []struct {
		input int
		want  bool
	}{
		{
			input: 112233,
			want:  true,
		},
		{
			input: 123444,
			want:  false,
		},
		{
			input: 111122,
			want:  true,
		},
		{
			input: 123455,
			want:  true,
		},
		{
			input: 223334,
			want:  true,
		},
	}

	for i, test := range tests {
		if valid := validatePasswordWithGroups(test.input); valid != test.want {
			t.Errorf("Got %v, but want %v in test case %d (password=%d)", valid, test.want, i, test.input)
		}
	}
}

func TestSolutionToPart1(t *testing.T) {
	want := 1767
	result := countPasswords(145852, 616942, validatePassword)

	if result != want {
		t.Errorf("Got %d, but want %d", result, want)
	}
}

func TestSolutionToPart2(t *testing.T) {
	want := 1192
	result := countPasswords(145852, 616942, validatePasswordWithGroups)

	if result != want {
		t.Errorf("Got %d, but want %d", result, want)
	}
}
