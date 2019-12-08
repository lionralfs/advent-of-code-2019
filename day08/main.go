package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func readInput() string {
	bytes, _ := ioutil.ReadFile("./input.txt")

	return string(bytes)
}

func main() {
	partOne()
}

func partOne() {
	input := readInput()
	min, minIndex, layer := fewestZeroDigits(input, 25, 6)
	ones := countDigitOccurenceInLayer(&layer, 1)
	twos := countDigitOccurenceInLayer(&layer, 2)

	fmt.Printf("[Part1] The layer with the fewest 0's is at index %d (%d); The product of the count of 1's and count of 2's on that layer is %d\n", minIndex, min, ones*twos)
}

func extractLayers(input string, width, height int) [][][]int {
	digits := make([]int, len(input))

	// parse the digits to ints
	for i, digit := range strings.Split(input, "") {
		asInt, _ := strconv.Atoi(digit)
		digits[i] = asInt
	}

	var layers [][][]int

	for i := 0; i < len(digits); i += width * height {
		layer := make([][]int, height)
		for j := i; j < i+width*height; j += width {
			layer[j%height] = digits[j : j+width]
		}
		layers = append(layers, layer)
	}

	return layers
}

func fewestZeroDigits(input string, width, height int) (int, int, [][]int) {
	layers := extractLayers(input, width, height)
	min := math.MaxInt64
	minIndex := -1
	var minLayer [][]int

	for i, layer := range layers {
		sum := countDigitOccurenceInLayer(&layer, 0)

		if sum < min {
			min = sum
			minIndex = i
			minLayer = layer
		}
	}

	return min, minIndex, minLayer
}

func countDigitOccurenceInLayer(layer *[][]int, digitToSearchFor int) int {
	sum := 0

	for _, row := range *layer {
		for _, digit := range row {
			if digit == digitToSearchFor {
				sum++
			}
		}
	}

	return sum
}
