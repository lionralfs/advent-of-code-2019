package main

import (
	"fmt"
	"strconv"
)

const (
	black = 0
	white = 1
)

type panel struct {
	color int
	count int
}

func main() {
	partOne()
	partTwo()
}

func partOne() {
	intcode := readInput()

	program := NewProgram(intcode)
	program.inputs = []int{0}

	fields := make(map[string]panel)
	x, y := 0, 0
	currentDirection := 0
	panelsAtLeastPaintedOnce := 0

	for {
		// get the color to paint the panel with
		color, done := program.Run()
		if done {
			break
		}

		// paint the panel
		currentPanel := fields[strconv.Itoa(x)+"--"+strconv.Itoa(y)]

		if currentPanel.count == 0 {
			panelsAtLeastPaintedOnce++
		}

		switch color {
		case 0:
			currentPanel.color = black
		case 1:
			currentPanel.color = white
		default:
			panic("Unknown color: " + strconv.Itoa(color))
		}
		currentPanel.count++
		fields[strconv.Itoa(x)+"--"+strconv.Itoa(y)] = currentPanel

		// get the direction to turn to
		direction, _ := program.Run()

		// update the direction
		switch direction {
		case 0:
			currentDirection = (currentDirection + 90) % 360
		case 1:
			currentDirection = (currentDirection + 270) % 360
		default:
			panic("Unknown direction: " + strconv.Itoa(direction))
		}

		// walk 1 step into that direction
		switch currentDirection {
		case 0:
			y--
		case 90:
			x++
		case 180:
			y++
		case 270:
			x--
		}

		// set the color of the current panel as input
		program.inputs = append(program.inputs, fields[strconv.Itoa(x)+"--"+strconv.Itoa(y)].color)
	}

	fmt.Printf("[Part1] %v panels are painted at least once\n", panelsAtLeastPaintedOnce)
}

func partTwo() {
	intcode := readInput()

	program := NewProgram(intcode)
	program.inputs = []int{1}

	fields := make(map[string]panel)
	x, y := 50, 50
	currentDirection := 0
	fields[strconv.Itoa(x)+"--"+strconv.Itoa(y)] = panel{
		count: 0,
		color: white,
	}

	printableField := make([][]string, 100)
	for i := range printableField {
		printableField[i] = make([]string, 100)
	}

	for {
		// get the color to paint the panel with
		color, done := program.Run()
		if done {
			break
		}

		// paint the panel
		currentPanel := fields[strconv.Itoa(x)+"--"+strconv.Itoa(y)]

		switch color {
		case 0:
			currentPanel.color = black
			printableField[y][x] = "."
		case 1:
			currentPanel.color = white
			printableField[y][x] = "#"
		default:
			panic("Unknown color: " + strconv.Itoa(color))
		}
		currentPanel.count++
		fields[strconv.Itoa(x)+"--"+strconv.Itoa(y)] = currentPanel

		// get the direction to turn to
		direction, _ := program.Run()

		// update the direction
		switch direction {
		case 0:
			currentDirection = (currentDirection + 270) % 360
		case 1:
			currentDirection = (currentDirection + 90) % 360
		default:
			panic("Unknown direction: " + strconv.Itoa(direction))
		}

		// walk 1 step into that direction
		switch currentDirection {
		case 0:
			y--
		case 90:
			x++
		case 180:
			y++
		case 270:
			x--
		}

		// set the color of the current panel as input
		program.inputs = append(program.inputs, fields[strconv.Itoa(x)+"--"+strconv.Itoa(y)].color)
	}

	for y := range printableField {
		for _, val := range printableField[y] {
			if val == "" {
				val = "."
			}
			fmt.Printf("%v", val)
		}
		fmt.Println("")
	}
}
