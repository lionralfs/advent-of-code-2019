package main

import "testing"

func TestGetClosestIntersection(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{
			input: "R75,D30,R83,U83,L12,D49,R71,U7,L72\nU62,R66,U55,R34,D71,R55,D58,R83",
			want:  159,
		},
		{
			input: "R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51\nU98,R91,D20,R16,D67,R40,U7,R15,U6,R7",
			want:  135,
		},
	}

	for i, test := range tests {
		if d := closestIntersection(test.input); d != test.want {
			t.Errorf("Got %d, but want %d in test case %d", d, test.want, i)
		}
	}
}

func TestGetClosestIntersectionBySteps(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{
			input: "R75,D30,R83,U83,L12,D49,R71,U7,L72\nU62,R66,U55,R34,D71,R55,D58,R83",
			want:  610,
		},
		{
			input: "R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51\nU98,R91,D20,R16,D67,R40,U7,R15,U6,R7",
			want:  410,
		},
	}

	for i, test := range tests {
		if d := closestIntersectionBySteps(test.input); d != test.want {
			t.Errorf("Got %d, but want %d in test case %d", d, test.want, i)
		}
	}
}
