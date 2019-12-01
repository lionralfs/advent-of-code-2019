package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	total := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		asInt, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		total += (asInt / 3) - 2
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Printf("Total: %v\n", total)
}
