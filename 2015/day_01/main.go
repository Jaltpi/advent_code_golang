package main

import (
	"fmt"
)

func main() {
	// input string
	fmt.Println("Day 1: Not Quite Lisp")
	apartment0_initial_directions := "))((((("
	apartment0_next_directions := ""

	apartment0_initial_directions_results := InstructionsConverter(apartment0_initial_directions)
	apartment0_next_directions_results := InstructionsConverter(apartment0_next_directions)
	fmt.Println("The initial instructions result is: ", apartment0_initial_directions_results)
	fmt.Println("The next instructions result is: ", apartment0_next_directions_results)

	apartment0_direction_result := DirectionCalculator(apartment0_initial_directions_results, apartment0_next_directions_results)
	fmt.Println("The floor result is: ", apartment0_direction_result)
}

func InstructionsConverter(instructions string) int {
	// create a variable to store the floor
	floor := 0

	// if instructions string is empty return 0
	if len(instructions) == 0 {
		return floor
	}

	// loop through the string

	for i := 0; i < len(instructions); i++ {

		// for every '(' character increase the floor by 1 & for every ')' character decrease the floor by 1
		if instructions[i] == '(' {
			floor++
		} else if instructions[i] == ')' {
			floor--
		}
	}

	// return the floor
	return floor

}

func DirectionCalculator(initial_direction_result int, next_direction_result int) int {
	if initial_direction_result == next_direction_result {
		return initial_direction_result
	}
	// when the initial direction result is not equal to the next direction result and the next direction result is 0 from an empty string
	if initial_direction_result != next_direction_result && next_direction_result == 0 {
		return initial_direction_result + next_direction_result

	}
	return 0
}
