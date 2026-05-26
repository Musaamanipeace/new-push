package main

import (
	"fmt"
	"os"
	"sort"
	"push-swap/internal/shared"
)

// exercise_IsSorted checks order entirely based on index continuity
func exercise_IsSorted(exercise_S shared.ExerciseStack) bool {
	for i := 0; i < len(exercise_S)-1; i++ {
		if exercise_S[i] > exercise_S[i+1] {
			return false
		}
	}
	return true
}

func main() {
	if len(os.Args) < 2 {
		return
	}

	exercise_Numbers, exercise_HasError := shared.Exercise_Parser(os.Args[1:])
	if exercise_HasError {
		shared.Exercise_PrintError()
		os.Exit(1)
	}

	var exercise_StackA shared.ExerciseStack = exercise_Numbers
	var exercise_StackB shared.ExerciseStack

	if exercise_IsSorted(exercise_StackA) {
		return
	}

	// Step 1: Structural Index Mapping (Non-comparative prep)
	// We map the numbers to unique positive indices [0...N-1] so we can safely process binary bits.
	exercise_SortedCopy := append(shared.ExerciseStack{}, exercise_StackA...)
	sort.Ints(exercise_SortedCopy)

	exercise_MapPositions := make(map[int]int)
	for idx, val := range exercise_SortedCopy {
		exercise_MapPositions[val] = idx
	}

	for idx, val := range exercise_StackA {
		exercise_StackA[idx] = exercise_MapPositions[val]
	}

	// Step 2: Strict Non-Comparative Radix Sort Loop
	exercise_TotalSize := len(exercise_StackA)
	exercise_MaxBits := 0
	
	// Determine how many bits we need to scan for the largest index
	for ((exercise_TotalSize - 1) >> exercise_MaxBits) > 0 {
		exercise_MaxBits++
	}

	// Loop through every bit position
	for bit := 0; bit < exercise_MaxBits; bit++ {
		if exercise_IsSorted(exercise_StackA) && len(exercise_StackB) == 0 {
			break
		}

		for i := 0; i < exercise_TotalSize; i++ {
			exercise_TopElement := exercise_StackA[0]
			
			// Inspect the exact bit at current position without comparing values
			if ((exercise_TopElement >> bit) & 1) == 1 {
				shared.Exercise_Ra(&exercise_StackA)
				fmt.Println("ra")
			} else {
				shared.Exercise_Pb(&exercise_StackA, &exercise_StackB)
				fmt.Println("pb")
			}
		}

		// Empty stack B back into stack A for the next significant bit pass
		for len(exercise_StackB) > 0 {
			shared.Exercise_Pa(&exercise_StackA, &exercise_StackB)
			fmt.Println("pa")
		}
	}
}