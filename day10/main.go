package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

// A Coordinate represents a point on the field
type Coordinate struct {
	x int
	y int
}

// The Field is the map of astroids
type Field struct {
	data  [][]string
	sizeX int
	sizeY int
}

func main() {
	partOne()
}

func partOne() {
	bytes, _ := ioutil.ReadFile("./input.txt")
	field := NewField(string(bytes))

	coord, count := field.SeesMax()
	fmt.Printf("%v can see %v other astroids\n", coord, count)
}

// NewField creates a new field instance from a string
func NewField(input string) Field {
	rows := strings.Split(input, "\n")

	sizeX := len(rows[0])
	sizeY := len(rows)
	field := make([][]string, sizeX)
	for x := range field {
		field[x] = make([]string, sizeY)
	}

	for y, row := range rows {
		for x, value := range strings.Split(row, "") {
			field[x][y] = value
		}
	}

	return Field{
		data:  field,
		sizeX: sizeX,
		sizeY: sizeY,
	}
}

// CanSee takes two coordinates and determines if they can see each other
func (field Field) CanSee(a, b Coordinate) bool {
	xDifference := b.x - a.x
	yDifference := b.y - a.y

	var xDirection, yDirection int

	if xDifference == 0 {
		// they are on the same x axis, walk along the y axis
		xDirection = 0
		yDirection = 1
		if yDifference < 0 {
			yDirection = -1
		}
	} else if yDifference == 0 {
		// they are on the same y axis, walk along the x axis
		yDirection = 0
		xDirection = 1
		if xDifference < 0 {
			xDirection = -1
		}
	} else {
		// not in a straight line
		gcd := int(math.Abs(float64(gcd(xDifference, yDifference))))
		xDifference /= gcd
		yDifference /= gcd

		xDirection = xDifference
		yDirection = yDifference
	}

	// walk
	for x, y := a.x+xDirection, a.y+yDirection; x < field.sizeX && x >= 0 && y < field.sizeX && y >= 0; {
		if x == b.x && y == b.y {
			return true
		}
		if field.data[x][y] != "." {
			return false
		}

		x += xDirection
		y += yDirection
	}

	return true
}

// SeesMax finds the coordinate that can see the most astroids
// and returns it, along with the amount that it can see
func (field Field) SeesMax() (Coordinate, int) {
	astroids := make([]Coordinate, 0)
	for x, col := range field.data {
		for y, element := range col {
			if element == "#" {
				astroids = append(astroids, Coordinate{x, y})
			}
		}
	}

	maxCount := 0
	var astroid Coordinate
	for i, A := range astroids {
		count := 0
		for j, B := range astroids {
			if i == j {
				continue
			}
			if field.CanSee(A, B) {
				count++
			}
		}
		if count > maxCount {
			maxCount = count
			astroid = A
		}
	}

	return astroid, maxCount
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}
