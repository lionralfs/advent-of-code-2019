package main

import "testing"

func TestInput1(t *testing.T) {
	want := 31
	output := calculateOreRequiredForFuel("./testinput1.txt", 1)

	if output != want {
		t.Errorf("Got %v, expected %v", output, want)
	}
}

func TestInput2(t *testing.T) {
	want := 165
	output := calculateOreRequiredForFuel("./testinput2.txt", 1)

	if output != want {
		t.Errorf("Got %v, expected %v", output, want)
	}
}

func TestInput3(t *testing.T) {
	want := 13312
	output := calculateOreRequiredForFuel("./testinput3.txt", 1)

	if output != want {
		t.Errorf("Got %v, expected %v", output, want)
	}
}

func TestInput4(t *testing.T) {
	want := 180697
	output := calculateOreRequiredForFuel("./testinput4.txt", 1)

	if output != want {
		t.Errorf("Got %v, expected %v", output, want)
	}
}

func TestInput5(t *testing.T) {
	want := 2210736
	output := calculateOreRequiredForFuel("./testinput5.txt", 1)

	if output != want {
		t.Errorf("Got %v, expected %v", output, want)
	}
}

func TestRealInput(t *testing.T) {
	want := 168046
	output := calculateOreRequiredForFuel("./input.txt", 1)

	if output != want {
		t.Errorf("Got %v, expected %v", output, want)
	}
}
