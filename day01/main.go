package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type fuelCalcFn func(int) int

func calculateFuel(mass int) int {
	return (mass / 3) - 2
}

func calculateFuelRec(mass int) int {
	fuel := calculateFuel(mass)
	if fuel <= 0 {
		return 0
	}
	return fuel + calculateFuelRec(fuel)
}

func totalFuelCalculator(file *os.File, fn fuelCalcFn) int {
	scanner := bufio.NewScanner(file)
	total := 0
	for scanner.Scan() {
		mass, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		total += fn(mass)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return total
}

func partOne() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	total := totalFuelCalculator(file, calculateFuel)

	fmt.Printf("[Part1] Total: %v\n", total)
}

func partTwo() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	total := totalFuelCalculator(file, calculateFuelRec)

	fmt.Printf("[Part2] Total: %v\n", total)
}

func main() {
	partOne()
	partTwo()
}
