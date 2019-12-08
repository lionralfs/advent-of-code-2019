package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

func readInput() string {
	bytes, _ := ioutil.ReadFile("./input.txt")

	return string(bytes)
}

func main() {
	partOne()
	partTwo()
}

func partOne() {
	input := readInput()
	min, minIndex, layer := fewestZeroDigits(input, 25, 6)
	ones := countDigitOccurenceInLayer(&layer, 1)
	twos := countDigitOccurenceInLayer(&layer, 2)

	fmt.Printf("[Part1] The layer with the fewest 0's is at index %d (%d); The product of the count of 1's and count of 2's on that layer is %d\n", minIndex, min, ones*twos)
}

func partTwo() {
	input := readInput()
	width := 25
	height := 6
	layers := extractLayers(input, width, height)

	// start with a transparent layer
	temp := make([][]int, width)
	for x := range temp {
		temp[x] = make([]int, height)

		for y := range temp[x] {
			temp[x][y] = 2
		}
	}

	// go from the topmost layer back,
	// and for each layer, only override the
	// position if it was previously transparent
	for _, layer := range layers {
		for y, row := range layer {
			for x, digit := range row {
				if temp[x][y] == 2 {
					temp[x][y] = digit
				}
			}
		}
	}

	// encode temp as image
	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}
	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	// Set color for each pixel.
	for x := range temp {
		for y, value := range temp[x] {
			switch value {
			case 0:
				img.Set(x, y, color.Black)
			case 1:
				img.Set(x, y, color.White)
			}
		}
	}

	// Encode as PNG.
	f, _ := os.Create("output.png")
	png.Encode(f, img)

	fmt.Println("[Part2] The generated image is at outout.png")
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
		layer := make([][]int, 0)
		for j := i; j < i+width*height; j += width {
			row := digits[j : j+width]
			layer = append(layer, row)
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
