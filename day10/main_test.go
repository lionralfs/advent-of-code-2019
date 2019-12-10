package main

import (
	"io/ioutil"
	"testing"
)

func TestSameXCoord(t *testing.T) {
	bytes, _ := ioutil.ReadFile("./testinput.txt")
	field := NewField(string(bytes))

	a := Coordinate{x: 0, y: 0}
	b := Coordinate{x: 0, y: 6}

	cansee := field.CanSee(a, b)

	if !cansee {
		t.Errorf("Expected %v to be able to see %v", a, b)
	}

	cansee2 := field.CanSee(b, a)
	if !cansee2 {
		t.Errorf("Expected %v to be able to see %v", b, a)
	}
}

func TestSameYCoord(t *testing.T) {
	bytes, _ := ioutil.ReadFile("./testinput.txt")
	field := NewField(string(bytes))

	a := Coordinate{x: 6, y: 0}
	b := Coordinate{x: 0, y: 0}

	cansee := field.CanSee(a, b)

	if !cansee {
		t.Errorf("Expected %v to be able to see %v", a, b)
	}

	cansee2 := field.CanSee(b, a)
	if !cansee2 {
		t.Errorf("Expected %v to be able to see %v", b, a)
	}
}

func TestDiagonal(t *testing.T) {
	bytes, _ := ioutil.ReadFile("./testinput.txt")
	field := NewField(string(bytes))

	a := Coordinate{x: 0, y: 0}
	b := Coordinate{x: 6, y: 2}

	cansee := field.CanSee(a, b)

	if cansee {
		t.Errorf("Expected %v to be unable to see %v", a, b)
	}

	cansee2 := field.CanSee(b, a)
	if cansee2 {
		t.Errorf("Expected %v to be unable to see %v", b, a)
	}
}

func TestCount1(t *testing.T) {
	bytes, _ := ioutil.ReadFile("./testinput1.txt")
	field := NewField(string(bytes))

	coord, count := field.SeesMax()

	wantCount := 33
	wantCoord := Coordinate{x: 5, y: 8}

	if count != wantCount {
		t.Errorf("Expected max count %v, got %v", wantCount, count)
	}
	if coord.x != wantCoord.x || coord.y != wantCoord.y {
		t.Errorf("Expected coordinate %v, got %v", wantCoord, coord)
	}
}

func TestCount2(t *testing.T) {
	bytes, _ := ioutil.ReadFile("./testinput2.txt")
	field := NewField(string(bytes))

	coord, count := field.SeesMax()

	wantCount := 35
	wantCoord := Coordinate{x: 1, y: 2}

	if count != wantCount {
		t.Errorf("Expected max count %v, got %v", wantCount, count)
	}
	if coord.x != wantCoord.x || coord.y != wantCoord.y {
		t.Errorf("Expected coordinate %v, got %v", wantCoord, coord)
	}
}

func TestCount3(t *testing.T) {
	bytes, _ := ioutil.ReadFile("./testinput3.txt")
	field := NewField(string(bytes))

	coord, count := field.SeesMax()

	wantCount := 41
	wantCoord := Coordinate{x: 6, y: 3}

	if count != wantCount {
		t.Errorf("Expected max count %v, got %v", wantCount, count)
	}
	if coord.x != wantCoord.x || coord.y != wantCoord.y {
		t.Errorf("Expected coordinate %v, got %v", wantCoord, coord)
	}
}

func TestCount4(t *testing.T) {
	bytes, _ := ioutil.ReadFile("./testinput4.txt")
	field := NewField(string(bytes))

	coord, count := field.SeesMax()

	wantCount := 210
	wantCoord := Coordinate{x: 11, y: 13}

	if count != wantCount {
		t.Errorf("Expected max count %v, got %v", wantCount, count)
	}
	if coord.x != wantCoord.x || coord.y != wantCoord.y {
		t.Errorf("Expected coordinate %v, got %v", wantCoord, coord)
	}
}
