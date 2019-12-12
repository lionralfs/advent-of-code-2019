package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"regexp"
	"strconv"
	"strings"
)

// Coordinate represents a point in a 3D coordinate system
type Coordinate struct {
	x, y, z int
}

// A Moon has a position and a velocity
type Moon struct {
	position, velocity *Coordinate
}

func (m Moon) String() string {
	return fmt.Sprintf("pos=<x=%v, y=%v, z=%v>, vel=<x=%v, y=%v, z=%v>", m.position.x, m.position.y, m.position.z, m.velocity.x, m.velocity.y, m.velocity.z)
}

func main() {
	partOne()
}

func partOne() {
	bytes, _ := ioutil.ReadFile("./input.txt")
	r := regexp.MustCompile(`^<x=(-?\d+), y=(-?\d+), z=(-?\d+)>$`)

	var moons []Moon

	for _, moon := range strings.Split(string(bytes), "\n") {
		matches := r.FindStringSubmatch(moon)
		x, _ := strconv.Atoi(matches[1])
		y, _ := strconv.Atoi(matches[2])
		z, _ := strconv.Atoi(matches[3])

		moons = append(moons, Moon{
			position: &Coordinate{
				x: x,
				y: y,
				z: z,
			},
			velocity: &Coordinate{
				x: 0,
				y: 0,
				z: 0,
			},
		})

	}

	// build pairs between moons
	var pairs [][]Moon

	for i, moonA := range moons {
		for j := i + 1; j < len(moons); j++ {
			pair := make([]Moon, 2)
			pair[0] = moonA
			pair[1] = moons[j]
			pairs = append(pairs, pair)
		}
	}

	steps := 1000

	// start the simulation
	for i := 0; i < steps; i++ {
		// for each pair of moons, apply gravity
		for _, pair := range pairs {
			moonA := pair[0]
			moonB := pair[1]

			// compare x positions
			if moonA.position.x > moonB.position.x {
				moonA.velocity.x--
				moonB.velocity.x++
			} else if moonA.position.x < moonB.position.x {
				moonA.velocity.x++
				moonB.velocity.x--
			}

			// compare y positions
			if moonA.position.y > moonB.position.y {
				moonA.velocity.y--
				moonB.velocity.y++
			} else if moonA.position.y < moonB.position.y {
				moonA.velocity.y++
				moonB.velocity.y--
			}

			// compare z positions
			if moonA.position.z > moonB.position.z {
				moonA.velocity.z--
				moonB.velocity.z++
			} else if moonA.position.z < moonB.position.z {
				moonA.velocity.z++
				moonB.velocity.z--
			}
		}

		// apply velocity by adding the velocity to the position
		for _, moon := range moons {
			moon.position.x += moon.velocity.x
			moon.position.y += moon.velocity.y
			moon.position.z += moon.velocity.z
		}
	}

	totalEnergy := 0
	for _, moon := range moons {
		pot := intAbs(moon.position.x) + intAbs(moon.position.y) + intAbs(moon.position.z)
		kin := intAbs(moon.velocity.x) + intAbs(moon.velocity.y) + intAbs(moon.velocity.z)
		total := pot * kin
		totalEnergy += total
	}

	fmt.Printf("[Part1] After %v steps, the total energy is: %v\n", steps, totalEnergy)
}

func intAbs(n int) int {
	return int(math.Abs(float64(n)))
}
