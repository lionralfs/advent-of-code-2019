package main

import (
	"fmt"
	"strconv"
)

func main() {
	partOne()
	partTwo()
}

func partOne() {
	intcode := readInput()
	program := NewProgram(intcode)
	screen := make(map[string]int)

	blockTiles := 0
	for {
		x, done := program.Run()
		if done {
			break
		}

		y, _ := program.Run()
		tileID, _ := program.Run()

		if tileID == 2 {
			blockTiles++
		}

		screen[strconv.Itoa(x)+"--"+strconv.Itoa(y)] = tileID
	}

	fmt.Printf("[Part1] There are %v block tiles\n", blockTiles)
}

func partTwo() {
	intcode := readInput()
	intcode[0] = 2
	program := NewProgram(intcode)

	// build the board first
	maxX, maxY, ballX, pX := 0, 0, 0, 0
	for {
		x, _ := program.Run()
		if x == -1 {
			program.Run()
			break
		}

		y, _ := program.Run()
		tileID, _ := program.Run()

		if x > maxX {
			maxX = x
		}

		if y > maxY {
			maxY = y
		}

		if tileID == 3 {
			pX = x
		}

		if tileID == 4 {
			ballX = x
		}
	}

	score, _ := program.Run()
	fmt.Printf("The score is %v!\n", score)

	// start playing
	if ballX > pX {
		program.inputs = []int{1}
	} else if ballX < pX {
		program.inputs = []int{-1}
	} else {
		program.inputs = []int{0}
	}

	for {
		x, done := program.Run()

		if done {
			break
		}

		y, _ := program.Run()
		z, _ := program.Run()

		if x == -1 && y == 0 {
			fmt.Printf("The score is %v!\n", z)
			continue
		}

		if z == 3 {
			pX = x
		}

		if z == 4 {
			ballX = x

			if ballX > pX {
				program.inputs = append(program.inputs, 1)
			} else if ballX < pX {
				program.inputs = append(program.inputs, -1)
			} else {
				program.inputs = append(program.inputs, 0)
			}
		}
	}
}
