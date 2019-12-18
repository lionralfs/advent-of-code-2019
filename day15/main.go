package main

import (
	"fmt"
	"strconv"

	"github.com/lionralfs/advent-of-code-2019/intcode"
)

// A Node has (possibly negative) coordinates, an instruction that took the robot there (from the previous node),
// and an output that was generated when the robot moved here
type Node struct {
	instructionToGetHere int
	output               int
	x                    int
	y                    int
}

// A Path is a list of nodes
type Path struct {
	nodes []Node
}

func main() {
	partOne()
	partTwo()
}

func partOne() {
	code := intcode.ReadInput("./input.txt")
	walkPath := createPathWalker(code)

	startPath := Path{
		nodes: []Node{
			Node{
				instructionToGetHere: -1,
				x:                    0,
				y:                    0,
			},
		},
	}

	found := make(map[string]bool)
	found["0--0"] = true

	frontier := []Path{startPath}

	// as long as we have paths on the queue
	for len(frontier) > 0 {
		// dequeue first path
		path := frontier[0]
		frontier = frontier[1:]
		currentNode := path.nodes[len(path.nodes)-1]

		// check if we've reached the goal
		output := currentNode.output

		if output == 2 {
			fmt.Printf("[Part1] Found oxygen system, it took %v steps\n", len(path.nodes)-1)
			break
		}

		// otherwise, add all neighbors to the queue
		neighbors := []Node{}

		up := Node{instructionToGetHere: 1, x: currentNode.x, y: currentNode.y - 1}
		down := Node{instructionToGetHere: 2, x: currentNode.x, y: currentNode.y + 1}
		left := Node{instructionToGetHere: 3, x: currentNode.x - 1, y: currentNode.y}
		right := Node{instructionToGetHere: 4, x: currentNode.x + 1, y: currentNode.y}

		neighbors = append(neighbors, up, down, left, right)

		for _, neighbor := range neighbors {
			// skip neighbors that we've already seens
			if found[strconv.Itoa(neighbor.x)+"--"+strconv.Itoa(neighbor.y)] {
				continue
			}

			found[strconv.Itoa(neighbor.x)+"--"+strconv.Itoa(neighbor.y)] = true

			// make a new path with the neighbor node attached at the end
			newPathNodes := make([]Node, len(path.nodes)+1)
			copy(newPathNodes, path.nodes)
			newPathNodes[len(newPathNodes)-1] = neighbor
			newPath := Path{nodes: newPathNodes}

			// attempt to walk there, to see whether it is a wall or not
			output := walkPath(newPath, false)
			newPath.nodes[len(newPath.nodes)-1].output = output

			// if it's not a wall, put the path back on the queue
			if output != 0 {
				frontier = append(frontier, newPath)
			}
		}
	}
}

func partTwo() {
	code := intcode.ReadInput("./input.txt")
	walkPath := createPathWalker(code)

	startPath := Path{
		nodes: []Node{
			Node{
				instructionToGetHere: -1,
				x:                    0,
				y:                    0,
			},
		},
	}

	found := make(map[string]string)
	found["0--0"] = "."

	frontier := []Path{startPath}

	minX, maxX, minY, maxY := 0, 0, 0, 0

	// as long as we have paths on the queue
	for len(frontier) > 0 {
		// dequeue first path
		path := frontier[0]
		frontier = frontier[1:]
		currentNode := path.nodes[len(path.nodes)-1]

		// add all neighbors to the queue
		neighbors := []Node{}

		up := Node{instructionToGetHere: 1, x: currentNode.x, y: currentNode.y - 1}
		down := Node{instructionToGetHere: 2, x: currentNode.x, y: currentNode.y + 1}
		left := Node{instructionToGetHere: 3, x: currentNode.x - 1, y: currentNode.y}
		right := Node{instructionToGetHere: 4, x: currentNode.x + 1, y: currentNode.y}

		neighbors = append(neighbors, up, down, left, right)

		for _, neighbor := range neighbors {
			// skip neighbors that we've already seen
			if found[strconv.Itoa(neighbor.x)+"--"+strconv.Itoa(neighbor.y)] != "" {
				continue
			}

			// set new min/max coordinates
			if neighbor.x < minX {
				minX = neighbor.x
			}
			if neighbor.x > maxX {
				maxX = neighbor.x
			}
			if neighbor.y < minY {
				minY = neighbor.y
			}
			if neighbor.y > maxY {
				maxY = neighbor.y
			}

			// make a new path with the neighbor node attached at the end
			newPathNodes := make([]Node, len(path.nodes)+1)
			copy(newPathNodes, path.nodes)
			newPathNodes[len(newPathNodes)-1] = neighbor
			newPath := Path{nodes: newPathNodes}

			// attempt to walk there, to see whether it is a wall or not
			output := walkPath(newPath, false)
			newPath.nodes[len(newPath.nodes)-1].output = output

			symbol := ""
			if output == 0 {
				symbol = "#"
			}
			if output == 1 {
				symbol = "."
			}
			if output == 2 {
				symbol = "O"
			}
			found[strconv.Itoa(neighbor.x)+"--"+strconv.Itoa(neighbor.y)] = symbol

			// if it's not a wall, put the path back on the queue
			if output != 0 {
				frontier = append(frontier, newPath)
			}
		}
	}

	field := makeField(found, minX, maxX, minY, maxY)
	steps := 0

	for ; ; steps++ {
		// un-comment the following few lines to see the visualization (x and y axis are swapped though)

		// for _, col := range field {
		// 	fmt.Println(col)
		// }
		// time.Sleep(200 * time.Millisecond)

		nextOxigenatedCoordinates := [][2]int{}

		for x := range field {
			for y := range field[x] {
				if field[x][y] == "." {
					// check up
					if field[x-1][y] == "O" {
						nextOxigenatedCoordinates = append(nextOxigenatedCoordinates, [2]int{x, y})
						continue
					}

					// check down
					if field[x+1][y] == "O" {
						nextOxigenatedCoordinates = append(nextOxigenatedCoordinates, [2]int{x, y})
						continue
					}

					// check left
					if field[x][y-1] == "O" {
						nextOxigenatedCoordinates = append(nextOxigenatedCoordinates, [2]int{x, y})
						continue
					}

					// check right
					if field[x][y+1] == "O" {
						nextOxigenatedCoordinates = append(nextOxigenatedCoordinates, [2]int{x, y})
						continue
					}
				}
			}
		}

		if len(nextOxigenatedCoordinates) == 0 {
			break
		}

		for _, coordinates := range nextOxigenatedCoordinates {
			field[coordinates[0]][coordinates[1]] = "O"
		}
	}

	fmt.Printf("[Part2] It takes %v steps to completely fill everything with oxygen\n", steps)
}

func makeField(found map[string]string, minX, maxX, minY, maxY int) [][]string {
	sizeX := maxX - minX + 1
	sizeY := maxY - minY + 1
	result := make([][]string, sizeX)

	for x := range result {
		result[x] = make([]string, sizeY)
		for y := range result[x] {
			symbol := found[strconv.Itoa(x+minX)+"--"+strconv.Itoa(y+minY)]
			if symbol == "" {
				symbol = "#"
			}
			result[x][y] = symbol
		}
	}

	return result
}

func createPathWalker(code []int) func(path Path, throwOnError bool) int {
	return func(path Path, throwOnError bool) int {
		// make a copy of the intcode
		intcodeCopy := make([]int, len(code))
		copy(intcodeCopy, code)
		program := intcode.NewProgram(intcodeCopy)

		result := 0
		// ignore the first one, since it is the start node
		for _, node := range path.nodes[1:] {
			program.AddInput(node.instructionToGetHere)

			result, _ = program.Run()

			if throwOnError && result == 0 {
				panic("Unable to walk path")
			}
		}

		return result
	}
}
