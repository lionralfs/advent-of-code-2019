package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	closest := closestIntersection(string(bytes))
	fmt.Printf("[Part1] The closest manhattan distance is %d\n", closest)

	closestSteps := closestIntersectionBySteps(string(bytes))
	fmt.Printf("[Part2] The closest distance by steps is %d\n", closestSteps)
}

func intAbs(n int) int {
	if n < 0 {
		return -n
	}

	return n
}

// calculates the manhattan distance to the (0, 0) point
func manhattanDistance(x, y int) int {
	return intAbs(x) + intAbs(y)
}

type walkCallback func(x, y int)

func walk(walkInstructionsRaw string, wc walkCallback) {
	walkInstructions := strings.Split(walkInstructionsRaw, ",")

	x, y := 0, 0

	for _, instruction := range walkInstructions {
		direction := string(instruction[0])
		steps, parseErr := strconv.Atoi(instruction[1:])

		if parseErr != nil {
			panic(parseErr)
		}

		switch direction {
		case "U":
			for i := y + 1; i <= y+steps; i++ {
				wc(x, i)
			}
			y += steps
		case "D":
			for i := y - 1; i >= y-steps; i-- {
				wc(x, i)
			}
			y -= steps
		case "L":
			for i := x - 1; i >= x-steps; i-- {
				wc(i, y)
			}
			x -= steps
		case "R":
			for i := x + 1; i <= x+steps; i++ {
				wc(i, y)
			}
			x += steps
		}
	}
}

func closestIntersection(input string) int {
	cables := strings.Split(input, "\n")

	if count := len(cables); count != 2 {
		panic(errors.New("Expected 2 cables, but got " + strconv.Itoa(count)))
	}

	minDistance := math.MaxInt64
	visited := make(map[string]bool)

	walk(cables[0], func(x, y int) {
		visited[strconv.Itoa(x)+"--"+strconv.Itoa(y)] = true
	})

	walk(cables[1], func(x, y int) {
		if visited[strconv.Itoa(x)+"--"+strconv.Itoa(y)] {
			if distance := manhattanDistance(x, y); distance < minDistance {
				minDistance = distance
			}
		}
	})

	return minDistance
}

func closestIntersectionBySteps(input string) int {
	cables := strings.Split(input, "\n")

	if count := len(cables); count != 2 {
		panic(errors.New("Expected 2 cables, but got " + strconv.Itoa(count)))
	}

	minDistance := math.MaxInt64
	cable1Steps := make(map[string]int)
	cable1StepCount := 0

	walk(cables[0], func(x, y int) {
		cable1StepCount++
		cable1Steps[strconv.Itoa(x)+"--"+strconv.Itoa(y)] = cable1StepCount
	})

	cable2StepCount := 0
	walk(cables[1], func(x, y int) {
		cable2StepCount++
		cable1StepsUntilHere := cable1Steps[strconv.Itoa(x)+"--"+strconv.Itoa(y)]

		if cable1StepsUntilHere != 0 {
			if sumOfSteps := cable1StepsUntilHere + cable2StepCount; sumOfSteps < minDistance {
				minDistance = sumOfSteps
			}
		}
	})

	return minDistance
}
