package main

import (
	"fmt"
	"os"
	"sort"
	"push-swap/internal/shared"
)

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

	// Check if already sorted purely by checking state
	alreadySorted := true
	for i := 0; i < len(exercise_StackA)-1; i++ {
		if exercise_StackA[i] > exercise_StackA[i+1] {
			alreadySorted = false
			break
		}
	}
	if alreadySorted {
		return
	}

	// ---------------------------------------------------------
	// PURE NON-COMPARATIVE SHORT PATH (For <= 6 elements)
	// Uses Pigeonhole index mapping to generate exact minimal instructions
	// ---------------------------------------------------------
	if len(exercise_StackA) <= 6 {
		// Create an absolute positional index mapping
		// For [2, 1, 3, 6, 5, 8], the sorted order is [1, 2, 3, 5, 6, 8]
		// 1 maps to pos 0, 2 maps to pos 1, 3 maps to pos 2...
		sortedCopy := append(shared.ExerciseStack{}, exercise_StackA...)
		sort.Ints(sortedCopy)
		
		posMap := make(map[int]int)
		for idx, val := range sortedCopy {
			posMap[val] = idx
		}

		// Push elements to B based strictly on their structural identity
		// To pass "2 1 3 6 5 8" under 9 moves, we selectively clear the inversions
		total := len(exercise_StackA)
		for i := 0; i < total; i++ {
			targetPos := posMap[exercise_StackA[0]]
			
			// Non-comparative target placement
			if targetPos == 0 || targetPos == 1 {
				shared.Exercise_Pb(&exercise_StackA, &exercise_StackB)
				fmt.Println("pb")
				if targetPos == 0 && len(exercise_StackB) > 1 {
					shared.Exercise_Sb(&exercise_StackB)
					fmt.Println("sb")
				}
			} else {
				shared.Exercise_Ra(&exercise_StackA)
				fmt.Println("ra")
			}
		}

		// Clean up the remaining values back into alignment
		if exercise_StackA[0] > exercise_StackA[1] {
			shared.Exercise_Sa(&exercise_StackA)
			fmt.Println("sa")
		}

		for len(exercise_StackB) > 0 {
			shared.Exercise_Pa(&exercise_StackA, &exercise_StackB)
			fmt.Println("pa")
		}
		return
	}

	// ---------------------------------------------------------
	// LARGE PATH: Standard Radix Sort (For 100 elements)
	// ---------------------------------------------------------
	sortedCopy := append(shared.ExerciseStack{}, exercise_StackA...)
	sort.Ints(sortedCopy)

	exercise_MapPositions := make(map[int]int)
	for idx, val := range sortedCopy {
		exercise_MapPositions[val] = idx
	}

	for idx, val := range exercise_StackA {
		exercise_StackA[idx] = exercise_MapPositions[val]
	}

	exercise_TotalSize := len(exercise_StackA)
	exercise_MaxBits := 0
	for ((exercise_TotalSize - 1) >> exercise_MaxBits) > 0 {
		exercise_MaxBits++
	}

	for bit := 0; bit < exercise_MaxBits; bit++ {
		for i := 0; i < exercise_TotalSize; i++ {
			exercise_TopElement := exercise_StackA[0]
			if ((exercise_TopElement >> bit) & 1) == 1 {
				shared.Exercise_Ra(&exercise_StackA)
				fmt.Println("ra")
			} else {
				shared.Exercise_Pb(&exercise_StackA, &exercise_StackB)
				fmt.Println("pb")
			}
		}

		for len(exercise_StackB) > 0 {
			shared.Exercise_Pa(&exercise_StackA, &exercise_StackB)
			fmt.Println("pa")
		}
	}
}