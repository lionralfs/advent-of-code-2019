package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type node struct {
	name            string
	children        []*node
	parent          *node
	distanceToSanta int
}

func main() {
	bytes, _ := ioutil.ReadFile("./input.txt")
	output1 := calculateChecksum(string(bytes))
	fmt.Printf("[Part1] The total number of direct and indirect orbits is %d\n", output1)

	output2 := distanceToSanta(string(bytes))
	fmt.Printf("[Part2] The total distance to santa is %d\n", output2)
}

func readInput(orbitMap string) map[string]*node {
	relationships := strings.Split(orbitMap, "\n")
	known := map[string]*node{}
	for _, relationship := range relationships {
		elements := strings.Split(relationship, ")")
		el1 := elements[0]
		el2 := elements[1]
		if known[el1] == nil {
			known[el1] = &node{
				name: el1,
			}
		}
		if known[el2] == nil {
			known[el2] = &node{
				name: el2,
			}
		}

		// add el1 as parent of el2
		known[el2].parent = known[el1]

		// add el2 to children of el1
		children := known[el1].children
		known[el1].children = append(children, known[el2])
	}
	return known
}

func calculateChecksum(orbitMap string) int {
	known := readInput(orbitMap)
	return checksum(known["COM"], 0)
}

func checksum(node *node, depth int) int {
	if len(node.children) == 0 {
		return depth
	}
	sum := depth
	for _, child := range node.children {
		sum += checksum(child, depth+1)
	}
	return sum
}

func distanceToSanta(orbitMap string) int {
	known := readInput(orbitMap)
	current := known["SAN"]
	distance := -1

	for current != nil {
		current.distanceToSanta = distance

		distance++
		current = current.parent
	}

	current = known["YOU"]
	distance = -1

	for current != nil {
		if current.distanceToSanta > 0 {
			return distance + current.distanceToSanta
		}

		distance++
		current = current.parent
	}

	return 0
}
