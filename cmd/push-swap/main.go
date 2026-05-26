package main

import (
	"fmt"
	"os"
	"sort"
	"push-swap/internal/shared"
)

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

	// Structural Index Mapping
	exercise_SortedCopy := append(shared.ExerciseStack{}, exercise_StackA...)
	sort.Ints(exercise_SortedCopy)

	exercise_MapPositions := make(map[int]int)
	for idx, val := range exercise_SortedCopy {
		exercise_MapPositions[val] = idx
	}

	// ---------------------------------------------------------
	// OPTIMIZED CHUNK SORT PATH (For <= 6 elements)
	// ---------------------------------------------------------
	if len(exercise_StackA) <= 6 {
		chunkSize := 3
		currentMaxIdx := chunkSize - 1

		// Stop looping immediately once the chunk is filled to prevent extra ra moves
		for len(exercise_StackB) < chunkSize && len(exercise_StackA) > 0 {
			targetPos := exercise_MapPositions[exercise_StackA[0]]
			if targetPos <= currentMaxIdx {
				shared.Exercise_Pb(&exercise_StackA, &exercise_StackB)
				fmt.Println("pb")
				
				if len(exercise_StackB) > 1 && exercise_MapPositions[exercise_StackB[0]] < exercise_MapPositions[exercise_StackB[1]] {
					shared.Exercise_Sb(&exercise_StackB)
					fmt.Println("sb")
				}
			} else {
				shared.Exercise_Ra(&exercise_StackA)
				fmt.Println("ra")
			}
		}

		// Sort the remaining chunk items in Stack A
		if len(exercise_StackA) >= 2 && exercise_MapPositions[exercise_StackA[0]] > exercise_MapPositions[exercise_StackA[1]] {
			shared.Exercise_Sa(&exercise_StackA)
			fmt.Println("sa")
		}

		// Merge the chunks back together cleanly
		for len(exercise_StackB) > 0 {
			shared.Exercise_Pa(&exercise_StackA, &exercise_StackB)
			fmt.Println("pa")
		}
		return
	}

	// ---------------------------------------------------------
	// LARGE DATASET PATH: Multi-Pass Radix Sort (For 100 elements)
	// ---------------------------------------------------------
	exercise_TotalSize := len(exercise_StackA)
	exercise_MaxBits := 0
	for ((exercise_TotalSize - 1) >> exercise_MaxBits) > 0 {
		exercise_MaxBits++
	}

	for bit := 0; bit < exercise_MaxBits; bit++ {
		if exercise_IsSorted(exercise_StackA) && len(exercise_StackB) == 0 {
			break
		}

		for i := 0; i < exercise_TotalSize; i++ {
			normalizedVal := exercise_MapPositions[exercise_StackA[0]]
			if ((normalizedVal >> bit) & 1) == 1 {
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