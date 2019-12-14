package main

import "testing"

func TestInput1(t *testing.T) {
	want := 31
	output := calculateOreRequiredForOneFuel("./testinput1.txt")

	if output != want {
		t.Errorf("Got %v, expected %v", output, want)
	}
}

func TestInput2(t *testing.T) {
	want := 165
	output := calculateOreRequiredForOneFuel("./testinput2.txt")

	if output != want {
		t.Errorf("Got %v, expected %v", output, want)
	}
}

func TestInput3(t *testing.T) {
	want := 13312
	output := calculateOreRequiredForOneFuel("./testinput3.txt")

	if output != want {
		t.Errorf("Got %v, expected %v", output, want)
	}
}

func TestInput4(t *testing.T) {
	want := 180697
	output := calculateOreRequiredForOneFuel("./testinput4.txt")

	if output != want {
		t.Errorf("Got %v, expected %v", output, want)
	}
}

func TestInput5(t *testing.T) {
	want := 2210736
	output := calculateOreRequiredForOneFuel("./testinput5.txt")

	if output != want {
		t.Errorf("Got %v, expected %v", output, want)
	}
}

func TestRealInput(t *testing.T) {
	want := 168046
	output := calculateOreRequiredForOneFuel("./input.txt")

	if output != want {
		t.Errorf("Got %v, expected %v", output, want)
	}
}
