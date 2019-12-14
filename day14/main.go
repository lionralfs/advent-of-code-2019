package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

// Unit represents the type of resource
type Unit string

// A Resource has a unit and an amount associated with them
type Resource struct {
	amount int
	unit   Unit
}

// A Reaction has a list of inputs and a single output resource
type Reaction struct {
	inputs []Resource
	output Resource
}

func main() {
	part1Result := calculateOreRequiredForOneFuel("./input.txt")

	fmt.Printf("[Part1] %v ORE is required to make 1 FUEL\n", part1Result)
}

func calculateOreRequiredForOneFuel(inputfile string) int {
	reactions := readInput(inputfile)
	available := make(map[Unit]int)

	oreUsed := 0

	var produce func(amount int, unit Unit)
	produce = func(amount int, unit Unit) {
		if unit == "ORE" {
			oreUsed += amount
			return
		}

		if available[unit] >= amount {
			available[unit] -= amount
			return
		}

		// find a reaction that produces the required resource
		for _, reaction := range reactions {
			if reaction.output.unit != unit {
				continue
			}

			factor := 1
			for factor*reaction.output.amount+available[unit] < amount {
				factor++
			}

			// produce all required inputs
			for _, input := range reaction.inputs {
				produce(factor*input.amount, input.unit)
			}
			available[reaction.output.unit] += factor*reaction.output.amount - amount
		}
	}

	produce(1, "FUEL")

	return oreUsed
}

func readInput(file string) []Reaction {
	bytes, _ := ioutil.ReadFile(file)

	var result []Reaction

	for _, s := range strings.Split(string(bytes), "\n") {
		parts := strings.Split(s, " => ")
		inputs := strings.Split(parts[0], ", ")

		reactionOutput := parseResource(parts[1])

		reaction := Reaction{
			inputs: []Resource{},
			output: reactionOutput,
		}

		for _, input := range inputs {
			inputResource := parseResource(input)
			reaction.inputs = append(reaction.inputs, inputResource)
		}

		result = append(result, reaction)
	}

	return result
}

func parseResource(s string) Resource {
	r := regexp.MustCompile(`^(\d+) ([A-Z]+)$`)
	matches := r.FindStringSubmatch(s)

	amount, _ := strconv.Atoi(matches[1])

	return Resource{
		amount: amount,
		unit:   Unit(matches[2]),
	}
}
