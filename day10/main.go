package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
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

// An Astroid with some info associated with them (relative to a certain position)
type Astroid struct {
	coordinates Coordinate
	// the manhattan distance to the
	distance  int
	angle     float64
	destroyed bool
}

func main() {
	partOne()
	partTwo()
}

func partOne() {
	bytes, _ := ioutil.ReadFile("./input.txt")
	field := NewField(string(bytes))

	coord, count := field.SeesMax()
	fmt.Printf("[Part1] %v can see %v other astroids\n", coord, count)
}

func partTwo() {
	bytes, _ := ioutil.ReadFile("./input.txt")
	field := NewField(string(bytes))

	coord, _ := field.SeesMax()

	astroids := make([]*Astroid, 0)
	for x, col := range field.data {
		for y, element := range col {
			if element == "#" {
				xDifference := x - coord.x
				yDifference := coord.y - y

				angle := math.Atan2(float64(xDifference), float64(yDifference)) * (180 / math.Pi)

				if angle < 0 {
					angle = 360 + angle
				}

				astroid := &Astroid{
					coordinates: Coordinate{x, y},
					distance:    int(math.Abs(float64(xDifference)) + math.Abs(float64(yDifference))),
					angle:       angle, // TODO
				}
				astroids = append(astroids, astroid)
			}
		}
	}

	// start at a negative number so we can be sure to hit the 0 angles first
	var currentAngle float64 = -1.0
	for i := range astroids {
		bestTarget := NextTarget(astroids, currentAngle)
		currentAngle = bestTarget.angle
		bestTarget.destroyed = true

		// update angles that we've already passed
		for _, e := range astroids {
			if e.angle <= currentAngle {
				e.angle += 360
			}
		}

		field.data[bestTarget.coordinates.x][bestTarget.coordinates.y] = strconv.Itoa(i + 1)

		if i+1 == 200 {
			x := bestTarget.coordinates.x
			y := bestTarget.coordinates.y
			fmt.Printf("[Part2] The 200th destroyed astroid is at (x, y): (%v, %v), so the solution is %v\n", x, y, x*100+4)
			return
		}

	}
}

// NextTarget selects the next target to shoot
func NextTarget(list []*Astroid, oldAngle float64) *Astroid {

	var bestTarget *Astroid = nil
	for _, astroid := range list {
		if astroid.destroyed {
			continue
		}

		angle := astroid.angle
		if angle <= oldAngle {
			// skip angles that we've already passed
			continue
		}

		if bestTarget == nil {
			bestTarget = astroid
			continue
		}

		// if we find a smaller angle than the one we already assumed to be the smallest, save it
		if angle < bestTarget.angle {
			bestTarget = astroid
		} else if angle == bestTarget.angle {
			// in case we found one with the same angle, use the one with the smaller distance
			if astroid.distance < bestTarget.distance {
				bestTarget = astroid
			}
		}

	}

	return bestTarget
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

// Print prints the field in a human-readable format
func (field Field) Print() {
	for y := 0; y < field.sizeY; y++ {
		for x := 0; x < field.sizeX; x++ {
			fmt.Printf("%v ", field.data[x][y])
		}
		fmt.Println("")
	}
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}
