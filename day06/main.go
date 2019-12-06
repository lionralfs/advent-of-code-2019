package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type node struct {
	name     string
	children []*node
}

func main() {
	bytes, _ := ioutil.ReadFile("./input.txt")
	output := calculateChecksum(string(bytes))
	fmt.Printf("[Part1] The total number of direct and indirect orbits is %d\n", output)
}

func calculateChecksum(orbitMap string) int {
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
		children := known[el1].children
		known[el1].children = append(children, known[el2])
	}
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
