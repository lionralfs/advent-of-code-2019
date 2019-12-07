package main

import "fmt"

func main() {
	intcode := readInput()

	result1, configuration1 := largestOutput(intcode)
	fmt.Printf("[Part1] The largest output is: %v (using configuration %v)\n", result1, configuration1)

	result2, configuration2 := largestOutputFeedback(intcode)
	fmt.Printf("[Part2] The largest output is: %v (using configuration %v)\n", result2, configuration2)

}

func permutations(arr []int) [][]int {
	var helper func([]int, int)
	res := [][]int{}

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}

func largestOutput(originalProgram []int) (int, []int) {
	max := 0
	configuration := make([]int, 4)

	permutations := permutations([]int{0, 1, 2, 3, 4})

	for _, permutation := range permutations {

		arg2 := 0
		var latestProgram program

		for i := 0; i < 5; i++ {
			intcode := make([]int, len(originalProgram))
			copy(intcode, originalProgram)
			latestProgram = program{
				instructionPointer: 0,
				intcode:            intcode,
				outputs:            []int{},
				inputArgs:          []int{permutation[i], arg2},
			}
			latestProgram.run()
			arg2 = latestProgram.outputs[0]
		}

		if latestProgram.outputs[0] > max {
			max = latestProgram.outputs[0]
			configuration = permutation
		}

	}

	return max, configuration
}

func largestOutputFeedback(intcode []int) (int, []int) {
	permutations := permutations([]int{5, 6, 7, 8, 9})

	max := 0
	configuration := make([]int, 4)

	for _, permutation := range permutations {
		// stores the output of the last amplifier, starting with 0 (the input to the first one)
		previousOutput := 0
		amplifiers := make([]*Program, 5)

		// initialize amplifiers by passing a copy of the "source code"
		// and the phase they should run on
		for i := range amplifiers {
			codeCopy := make([]int, len(intcode))
			copy(codeCopy, intcode)
			amplifiers[i] = &Program{
				code:   codeCopy,
				inputs: []int{permutation[i]},
			}
		}

		last := amplifiers[4]

		// as long as the last one is not done
		for !last.done {
			// calculate the output and use it as input for the next amplifier for each amplifier
			for _, amplifier := range amplifiers {
				amplifier.inputs = append(amplifier.inputs, previousOutput)
				previousOutput = amplifier.Run()
			}
		}

		if last.output > max {
			max = last.output
			configuration = permutation
		}
	}

	return max, configuration
}
